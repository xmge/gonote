# 谈谈ok那些事儿

ok即断言，可以断言某一个变量是特定类型，也可以断言map中存在某一个key值

## 1.类型转换断言

```go
var v interface{} = 1

// 类型断言成功
v1, ok := v.(int)
fmt.Println(v1, ok)   // 1 true

// 类型断言失败 有ok 则不会panic 并返回断言的类型的零值，在这里是float64的零值：0
v2, ok := v.(float64)
fmt.Println(v2, ok)   // 0 false

// 类型断言失败 无ok 则直接panic
v3 := v.(string)      // 将直接panic
```

## 2.map中key值存在断言

```go
// key存在的断言成功 有ok
m := map[string]interface{}{"hello": "xmge"}
v1, ok := m["hello"]
fmt.Printf("v1:%v,ok:%t \n", v1, ok)    // v1:xmge,ok:true

// key存在的断言失败 有ok
v2, ok := m["key"]
fmt.Printf("v2:%v,ok:%t \n", v2, ok)   // v2:<nil>,ok:false

// key存在的断言失败 无ok
v3 := m["key"]
fmt.Printf("v3:%v \n", v3)             // v3:<nil> （不会像类型转换断言一样直接panic）
```

## 3.channel有值断言

ok不是断言channel是否关闭，而是断言是否能取出值。
随便说一下，向关闭的通道放值则panic，从关闭的chanel取值将不再阻塞

```go
c := make(chan bool,2)
c <- true
c <- true

// 通道未关闭 有值(无值不关闭则会阻塞)
b1,ok := <-c
fmt.Printf("b1:%t,ok:%t \n",b1,ok)
// b1:true,ok:true

// 通道已关闭 通道中还有值
close(c)
b2,ok := <-c
fmt.Printf("b2:%t,ok:%t \n",b2,ok)
// b2:true,ok:true

// 通道已关闭 通道中无值
b3,ok := <-c
fmt.Printf("b3:%t,ok:%t \n",b3,ok)
// b3:false,ok:false
```