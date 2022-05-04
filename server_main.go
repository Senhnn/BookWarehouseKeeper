package main

import (
	"BWKV1/DB"
	"BWKV1/MiddleWare"
	"BWKV1/RpcService"
	"BWKV1/Services"
	"BWKV1/config"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
)

func Init() {
	// 配置初始化
	config.Init()
	// log初始化
	MiddleWare.LoggerInit()
	// redis初始化
	DB.RedisInit()
	// mysql初始化
	DB.MysqlInit()
}

func main() {
	// 初始化
	Init()

	//tlsConfig := InitAdminServiceCert()

	serverCfg := viper.GetStringMapString("server")
	ip := serverCfg["ip"]
	port := serverCfg["port"]

	// 开启服务器监听
	listener, err := net.Listen("tcp", ip+":"+port)
	fmt.Println("listen on:", ip, ":", port)
	if err != nil {
		panic(any(err))
	}
	defer listener.Close()

	// 创建GRPC server
	server := grpc.NewServer(
		//grpc.Creds(credentials.NewTLS(tlsConfig)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(MiddleWare.Logger),
			//grpc_auth.UnaryServerInterceptor(MiddleWare.AuthFunc),
			//grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	// 注册grpc server
	Services.RegisterAdminServiceServer(server, &RpcService.AdminServer{})
	Services.RegisterWorkerServiceServer(server, &RpcService.WorkerServer{})

	// 服务启动
	server.Serve(listener)
}

func InitAdminServiceCert() *tls.Config {
	// 加载服务器私钥和证书
	cert, err := tls.LoadX509KeyPair("cert/server/server.pem", "cert/server/server.key")
	if err != nil {
		panic(any(err))
	}

	// 生成证书池，将根证书加入证书池
	certPool := x509.NewCertPool()
	rootBuf, err := ioutil.ReadFile("cert/server/ca.pem")
	if err != nil {
		panic(any(err))
	}
	if !certPool.AppendCertsFromPEM(rootBuf) {
		panic(any("Fail to add ca"))
	}

	// 初始化TLSConfig
	// ClientAuth有5中类型，如果要进行双向认证必须是RequireAndVerifyClientCert
	tlsConfig := &tls.Config{
		//ServerName:   "Syhan",
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	}

	return tlsConfig
}
