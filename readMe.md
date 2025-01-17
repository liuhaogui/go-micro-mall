# readMe

## 功能介绍
- 核心功能组件：
    - api gateway
    - 注册中心 （consul）
    - 配置中心
    - 熔断
    - 链路追踪
    - jwt
    - metrics
    - 集成日志推送 日志系统
    - [x] 健康检查
    - [x] Swagger

### 相关组件运行
- jaeger
```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.17
```
- jaeger web Ul: http://localhost:16686

- consul
```
consul agent -dev
```
- consul web Ul: http://localhost:8500/ui/dc1/services



### consul config
 - 根路径：`micro/config`

- jwt-key
```
{
	"key": "asdf1saf233asdfas3df"
}
```
- database
```
{
    "user":{
        "address":"127.0.0.1",
        "port":3306,
        "user_name":"root",
        "user_password":"",
        "db_name":"micro_book_mall"
    }
}
```

- 生成proto
```
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. user.proto
```