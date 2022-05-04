package main

import (
	"BWKV1/Services"
	"context"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
)

type Auth struct {
	User   string
	Passwd string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"admin": a.User, "passwd": a.Passwd}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}

var auth = &Auth{
	User:   "admin",
	Passwd: "1234",
}

func main() {
	// 加载客户端私钥和证书
	//cert, err := tls.LoadX509KeyPair("cert/client/client.pem", "cert/client/client.key")
	//if err != nil {
	//	panic(any(err))
	//}

	// 将根证书加入证书池
	certPool := x509.NewCertPool()
	rootBuf, err := ioutil.ReadFile("cert/client/ca.pem")
	if err != nil {
		panic(any(err))
	}
	if !certPool.AppendCertsFromPEM(rootBuf) {
		panic(any("Fail to append ca"))
	}

	// 新建凭证
	// 注意ServerName需要与服务器证书内的Common Name一致
	// 客户端是根据根证书和ServerName对服务端进行验证的
	//creds := credentials.NewTLS(&tls.Config{
	//	ServerName:   "Syhan",
	//	Certificates: []tls.Certificate{cert},
	//	RootCAs:      certPool,
	//})

	// 不使用认证建立连接
	conn, err := grpc.Dial("127.0.0.1:5000",
		//grpc.WithTransportCredentials(creds),
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(auth),
	)
	if err != nil {
		panic(any(err))
	}
	defer conn.Close()

	afd := Services.NewAdminServiceClient(conn)
	wfd := Services.NewWorkerServiceClient(conn)

	res, err := afd.AddNewBook(context.Background(), &Services.AddNewBookRequest{
		BookInfo: &Services.BookInfo{
			Isbn:     "677",
			Author:   "Nietzsche",
			Nation:   "Germany",
			BookName: "The birth of the tragic",
			BookType: "philosophy",
		},
		Num: 200,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Ret)
	fmt.Println(res.ErrorInfo)

	newRes, err := wfd.GetAllBook(context.Background(), &Services.GetAllBookRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(newRes.BookList))
	fmt.Println(newRes.BookList)
}
