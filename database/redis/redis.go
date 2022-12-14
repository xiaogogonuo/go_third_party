package main

// https://github.com/go-redis/redis

// If you are using Redis 6, install go-redis/v8:
// go get github.com/go-redis/redis/v8

// If you are using Redis 7, install go-redis/v9:
// go get github.com/go-redis/redis/v9

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			fmt.Println(111)
			return nil
		},
		Username:     "",
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		PoolFIFO:     true,
		PoolSize:     10,
	})
	return client
}

// NewFailOverClient Redis Sentinel模式
func NewFailOverClient() *redis.Client {
	failOverClient := redis.NewFailoverClient(&redis.FailoverOptions{})
	return failOverClient
}

// NewClusterClient Redis Cluster模式
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

func Set(ctx context.Context) {

}

func main() {

}
