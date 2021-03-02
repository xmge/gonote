package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	"log"
)

// 操作数据库有问题时，通过备份库恢复数据

type Record struct {
	WithdrawId string `json:"withdraw_id"`
	Status    string  `json:"status"`
}

func main() {
	//srcConn := "postgres://root:DA*449VBty7R@appserver-colorphone-1214-cluster-cluster.cluster-cv0oenksk2p5.rds.cn-northwest-1.amazonaws.com.cn:5432/appserver_mickey?sslmode=disable;"
	//dstConn := "postgres://root:DA*449VBty7R@appserver-colorphone-cluster.cluster-cv0oenksk2p5.rds.cn-northwest-1.amazonaws.com.cn:5432/appserver_mickey?sslmode=disable;"

	srcConn := "postgresql://bytepower_root:rDqe0NklUtVGBNsg@bytepower-develop-cluster.cluster-cowgwfadfddw.rds.cn-northwest-1.amazonaws.com.cn:5432/appserver_amber?sslmode=disable;"
	dstConn := "postgresql://bytepower_root:rDqe0NklUtVGBNsg@bytepower-develop-cluster.cluster-cowgwfadfddw.rds.cn-northwest-1.amazonaws.com.cn:5432/appserver_coffee?sslmode=disable;"
	srcDb, err := xorm.NewEngine("postgres", srcConn)
	if err != nil {
		log.Panic(err)
	}

	dstDb,err := xorm.NewEngine("postgres", dstConn)
	if err != nil {
		log.Panic(err)
	}
	offset := 0
	step := 10000
	loop := 1
	for {
		rs := make([]Record,0)
		err = srcDb.SQL("select withdraw_id,status from withdraw_money_record order by created desc limit ? OFFSET ? ",step,offset).Find(&rs)
		if err != nil {
			log.Panic(err)
		}

		for _, record := range rs {
			fmt.Println(record)
			_,err := dstDb.SQL("update withdraw_money_record set status = ? where withdraw_id = ?",record.Status,record.WithdrawId).Execute()
			if err != nil {
				log.Println(err)
				return
			}
		}

		if len(rs) < step {
			break
		}
		fmt.Println("loop:",loop)
		fmt.Println("offset:",offset)
		loop++
		offset +=step
	}

	fmt.Println("mission success")
}

