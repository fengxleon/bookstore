# bookstore
Demo for go-zero

在Add服务里添加了Update rpc


启服配置：

1、安装mysql、etcd、redis

2、修改api\etc\bookstore-api.yaml rpc\add\etc\add.yaml rpc\check\etc\check.yaml 配置

3、在rpc\add 下运行 go run add.go，在rpc\check 下运行 go run check.go。启服后在api目录下运行go run bookstore.go

4、测试连接 curl -i "http://localhost:8888/add?book=Little_Prince&price=52"
