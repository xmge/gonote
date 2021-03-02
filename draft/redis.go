package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

func init()  {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "68.79.27.61:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_,err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	Set()
}

func Set() error {
	names := []string{"小张","小刘","小马","小武","小张","小马","小武","小张"}
	key := "names:set"
	_,err := rdb.SAdd(ctx,key,names).Result()
	if err != nil {
		return err
	}
	expireTime := time.Hour * 24
	_,err = rdb.Expire(ctx,key,expireTime).Result()
	if err != nil {
		return err
	}

	// 推荐 set 的元素个数小于 8192（此数字来自阿里的使用限制）
	setNames,err := rdb.SMembers(ctx,key).Result()
	if err != nil {
		return err
	}
	fmt.Println(setNames)
	return err
}

func ZSet() error {


	// 生成排行榜
	type User struct {
		Name string `json:"name"`
		Score int `json:"score"`
	}
	users:= []User{
		{"小黑",12},
		{"小白",2},
		{"小红",1},
		{"小蓝",213},
	}

	key := "amber:leaderboard"
	memebers := make([]*redis.Z,0)
	expireTime := time.Hour * 24
	for _, user := range users{
		//如果排行榜比较大，这里需要分批上传，以防阻塞 redis
		memebers = append(memebers, &redis.Z{float64(user.Score),user.Name})
	}
	_,err := rdb.ZAdd(ctx,key,memebers...).Result()
	if err != nil {
		return err
	}

	_,err = rdb.Expire(ctx,key,expireTime).Result()
	if err != nil {
		return err
	}

	count,err := rdb.SCard(ctx,key).Result()
	if err != nil {
		return err
	}
	_ = count
	// 如果 count > 8192,则取数据时要分批取。

	// 获取排行榜,排名从高到低
	zs,err := rdb.ZRevRangeWithScores(ctx,key,0,-1).Result()
	if err != nil {
		return err
	}
	for _, z := range zs {
		fmt.Println(z)
	}
	return nil
}

func Lock() (bool,error) {
	key := "key"
	expireTime := time.Hour * 12
	return rdb.SetNX(ctx,key,1,expireTime).Result()
}


func HotKey() (string,error) {
	key := "key"
	v, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil { // 如果 key 不存在
		//1、从数据库查询要缓存的值
		value := "value"
		expireTime := time.Hour * 24

		//2、 把热点数据放入缓存中，并制定过期时间
		_,err := rdb.Set(ctx,key,value, expireTime).Result()
		if err != nil {
			return "",err
		}
	}
	if err != nil { 	// 查询 redis 出错
		return "",err
	}
	return v,nil
}

