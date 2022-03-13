package main

import (
	"BWKV1/DB"
	"BWKV1/config"
	"crypto/tls"
	"crypto/x509"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net"
)

func Init() {
	// 配置初始化
	config.Init()
	// redis初始化
	DB.RedisInit()
	// mysql初始化
	DB.MysqlInit()
}

func main() {
	// 初始化
	Init()

	tlsConfig := InitAdminServiceCert()

	serverCfg := viper.GetStringMapString("server")
	ip := serverCfg["ip"]
	port := serverCfg["port"]

	// 开启服务器监听
	listener, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		panic(any(err))
	}
	defer listener.Close()

	// 创建GRPC server
	server := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))

	// 注册grpc server
	//Services.RegisterAdminServiceServer(server, &RpcHandler.AdminServices{})

	// 服务启动
	server.Serve(listener)
}

func InitAdminServiceCert() *tls.Config {
	// 加载服务器私钥和证书
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		panic(any(err))
	}

	// 生成证书池，将根证书加入证书池
	certPool := x509.NewCertPool()
	rootBuf, err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		panic(any(err))
	}
	if !certPool.AppendCertsFromPEM(rootBuf) {
		panic(any("Fail to add ca"))
	}

	// 初始化TLSConfig
	// ClientAuth有5中类型，如果要进行双向认证必须是RequireAndVerifyClientCert
	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	}

	return tlsConfig
}
