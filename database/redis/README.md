# 拉取镜像
```shell
docker pull redis:7.0.6
```

# create directories and files that need to be mounted on the macOS or linux host
```shell
mkdir -p ~/Documents/opt/redis/stand-alone/data
mkdir -p ~/Documents/opt/redis/stand-alone/conf
touch ~/Documents/opt/redis/stand-alone/conf/redis.conf
```

# start a redis instance with persistent storage
```shell
docker run -p 6379:6379 --name redis0 -v ~/Documents/opt/redis/stand-alone/data:/data \
-v ~/Documents/opt/redis/stand-alone/conf/redis.conf:/etc/redis/redis.conf \
-d redis:7.0.6 redis-server /etc/redis/redis.conf \
--save 60 1 --loglevel warning \
--requirepass rust
```

# 单机模式
## 1、在linux主机创建需要挂载的目录和文件
```shell
mkdir -p ~/Documents/volume/redis/stand-alone/data
mkdir -p ~/Documents/volume/redis/stand-alone/conf
touch ~/Documents/volume/redis/stand-alone/conf/redis.conf
```
## 2、创建容器并在后台运行
```shell
docker run \
--publish 6379:6379 \
--name redis-stand-alone \
-v ~/Documents/volume/redis/stand-alone/data:/data \
-v ~/Documents/volume/redis/stand-alone/conf/redis.conf:/etc/redis/redis.conf \
-d redis:7.0.6 redis-server /etc/redis/redis.conf \
--port 6379 \
--appendonly yes \
--appendfilename appendonly.aof \
--save 60 1 \
--loglevel warning \
--requirepass rust \
--masterauth rust
```
## 3、进入容器并连接redis客户端
```shell
docker exec -it redis-stand-alone redis-cli

127.0.0.1:6379> auth rust
OK
127.0.0.1:6379> info replication
# Replication
role:master
connected_slaves:0
master_failover_state:no-failover
master_replid:beff8708128dc84d407181dea448946425db6e4d
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
```

# 主从模式
## 1、在linux主机创建需要挂载的目录和文件
```shell
# master
mkdir -p ~/Documents/volume/redis/master-slave/master/data
mkdir -p ~/Documents/volume/redis/master-slave/master/conf
touch ~/Documents/volume/redis/master-slave/master/conf/redis.conf

# slave
mkdir -p ~/Documents/volume/redis/master-slave/slave/data
mkdir -p ~/Documents/volume/redis/master-slave/slave/conf
touch ~/Documents/volume/redis/master-slave/slave/conf/redis.conf
```
## 2、创建容器并在后台运行
```shell
# master
docker run \
--publish 6380:6379 \
--name redis-master \
-v ~/Documents/volume/redis/master-slave/master/data:/data \
-v ~/Documents/volume/redis/master-slave/master/conf/redis.conf:/etc/redis/redis.conf \
-d redis:7.0.6 redis-server /etc/redis/redis.conf \
--port 6379 \
--appendonly yes \
--appendfilename appendonly.aof \
--save 60 1 \
--loglevel warning \
--requirepass rust \
--masterauth rust

# slave
docker run \
--publish 6381:6379 \
--name redis-slave \
-v ~/Documents/volume/redis/master-slave/slave/data:/data \
-v ~/Documents/volume/redis/master-slave/slave/conf/redis.conf:/etc/redis/redis.conf \
-d redis:7.0.6 redis-server /etc/redis/redis.conf \
--port 6379 \
--appendonly yes \
--appendfilename appendonly.aof \
--save 60 1 \
--loglevel warning \
--requirepass go \
--slaveof 172.17.0.3 6379 \
--masterauth rust
```
## 3、进入容器并连接redis客户端
```shell
# master
docker exec -it redis-master redis-cli

127.0.0.1:6379> auth rust
OK
127.0.0.1:6379> info replication
role:master
connected_slaves:1
slave0:ip=172.17.0.4,port=6379,state=online,offset=224,lag=1
master_failover_state:no-failover
master_replid:3624f9851a93e3a496a4c2aa520a3374fb13eb8e
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:224
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:224

# slave
docker exec -it redis-slave redis-cli

127.0.0.1:6379> auth go
OK
127.0.0.1:6379> info replication
role:slave
master_host:172.17.0.3
master_port:6379
master_link_status:up
master_last_io_seconds_ago:4
master_sync_in_progress:0
slave_read_repl_offset:140
slave_repl_offset:140
slave_priority:100
slave_read_only:1
replica_announced:1
connected_slaves:0
master_failover_state:no-failover
master_replid:3624f9851a93e3a496a4c2aa520a3374fb13eb8e
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:140
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:57
repl_backlog_histlen:84
```

# 哨兵模式

# 集群模式

# connecting via redis-cli
```shell
# method1
docker exec -it <container_name or container_id> bash
redis-cli

# method2: 
docker exec -it <container_name or container_id> redis-cli

# if password set, auth password first
127.0.0.1:6379> auth your_redis_server_password
OK
127.0.0.1:6379> info replication
# Replication
role:master
connected_slaves:0
master_failover_state:no-failover
master_replid:0737e08ebd40d7e049283fe936aabc3b43592a1c
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
```

# docker builds redis cluster
## create directories and files that need to be mounted on the macOS or linux host
```shell
mkdir -p ~/Documents/opt/redis/cluster/node1/data
mkdir -p ~/Documents/opt/redis/cluster/node2/data
mkdir -p ~/Documents/opt/redis/cluster/node3/data
```
## create cluster
```shell
docker create --name redis-1 -v ~/Documents/opt/redis/cluster/node1/data:/data  \
-v ~/Documents/opt/redis/cluster/node1/conf/redis.conf:/etc/redis/redis.conf \
-p 6380:6379 redis:7.0.6 --cluster-enabled yes    \
--cluster-config-file redis.conf


docker run -p 6380:6379 --name redis-1 -v ~/Documents/opt/redis/cluster/node1/data:/data \
-v ~/Documents/opt/redis/cluster/node1/conf/redis.conf:/etc/redis/redis.conf \
-d redis:7.0.6 redis-server /etc/redis/redis.conf \
--save 60 1 --loglevel warning

docker run -p 6380:6379 --name redis-2 -v ~/Documents/opt/redis/cluster/node2/data:/data \
-v ~/Documents/opt/redis/cluster/node2/conf/redis.conf:/etc/redis/redis.conf \
-d redis:7.0.6 redis-server /etc/redis/redis.conf \
--save 60 1 --loglevel warning

docker run -p 6382:6379 --name redis-3 -v ~/Documents/opt/redis/cluster/node3/data:/data \
-v ~/Documents/opt/redis/cluster/node3/conf/redis.conf:/etc/redis/redis.conf \
-d redis:7.0.6 redis-server /etc/redis/redis.conf \
--save 60 1 --loglevel warning
```

* [Docker command line](https://docs.docker.com/engine/reference/commandline/run/)

* [Docker安装redis（单机、主从、哨兵、集群）](https://blog.csdn.net/u010148813/article/details/128099031)
* [docker安装redis并搭建集群](https://blog.csdn.net/qq_38327769/article/details/124063769)
* [Docker搭建redis集群](https://blog.csdn.net/emgexgb_sef/article/details/126327579)

* [Redis后台运行配置](https://www.cnblogs.com/s1mmons/p/16626299.html)
* [redis主要启动主要参数与配置文件说明](https://article.itxueyuan.com/pW0AO)

* [golang连接redis集群遇见的坑](https://blog.csdn.net/weixin_42854904/article/details/124990524)
* [使用 go-redis 连接 Amazon ElastiCache for Redis 集群](https://aws.amazon.com/cn/blogs/china/all-roads-lead-to-rome-use-go-redis-to-connect-amazon-elasticache-for-redis-cluster/)