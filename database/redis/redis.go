package main

// https://github.com/go-redis/redis

// If you are using Redis 6, install go-redis/v8:
// go get github.com/go-redis/redis/v8

// If you are using Redis 7, install go-redis/v9:
// go get github.com/go-redis/redis/v9

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx = context.Background()

// NewClient 单机模式
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6380",
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			fmt.Println("connection is established")
			// do stuff
			return nil
		},
		Username:     "",
		Password:     "rust",
		DB:           0,
		MaxRetries:   3,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		PoolFIFO:     true,
		PoolSize:     10,
	})
	return client
}

// NewMasterSlave 主从模式
// 一主一从，从机作为主机的备份，实时复制主机数据，只能读不能写。
// 主机故障后，可以手动切换客户端的连接地址到从机继续使用。
//一般用于防止数据丢失，或简单的读写分离。
func NewMasterSlave() {

}

// NewFailOverClient 哨兵模式
func NewFailOverClient() *redis.Client {
	failOverClient := redis.NewFailoverClient(&redis.FailoverOptions{})
	return failOverClient
}

// NewClusterClient 集群模式
func NewClusterClient() *redis.ClusterClient {
	// 连接redis集群
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{ // 填写master主机
			"192.168.21.22:30001",
			"192.168.21.22:30002",
			"192.168.21.22:30003",
		},
		Password:     "123456",              // 设置密码
		DialTimeout:  50 * time.Microsecond, // 设置连接超时
		ReadTimeout:  50 * time.Microsecond, // 设置读取超时
		WriteTimeout: 50 * time.Microsecond, // 设置写入超时
	})
	// 发送一个ping命令,测试是否通
	s := clusterClient.Do(ctx, "ping").String()
	fmt.Println(s)
	return clusterClient
}

func main() {
	rdb := NewClient()
	rdb.Set(ctx, "city", "china", 60*time.Second)
	val, err := rdb.Get(ctx, "city").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
