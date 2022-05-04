package MiddleWare

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthFunc(ctx context.Context) (context.Context, error) {
	// 获取传输的用户名和密码
	md, ok := metadata.FromIncomingContext(ctx)
	if ok != true {
		return ctx, fmt.Errorf("missing key word")
	}
	var account string
	var passwd string
	if val, ok := md["admin"]; ok {
		account = val[0]
	}
	if val, ok := md["passwd"]; ok {
		passwd = val[0]
	}
	admin := viper.GetStringMapString("server")
	if admin == nil {
		return ctx, status.Errorf(codes.InvalidArgument, "配置读取失败")
	}
	if account != admin["admin"] || passwd != admin["passwd"] {
		return ctx, status.Errorf(codes.Unauthenticated, "Token不合法")
	}

	return ctx, nil
}
