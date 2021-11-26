package data

import (
	"context"
	uClient "web-api/api/user"
	"web-api/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserServiceClient, NewWebRepo)

// Data .
type Data struct {
	//	db *gorm.DB
	// TODO wrapped database client
	//logger     *log.Logger
	userClient uClient.UserClient
}

// NewData .
func NewData(logger log.Logger, client uClient.UserClient) (*Data, func(), error) {
	logs := log.NewHelper(logger)
	//	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	// if err != nil {
	// 	return nil, nil, err
	// }
	d := &Data{
		//db:         db.Debug(),
		userClient: client,
		//	logger:     log.Logger(logger),
		// logger:   ,
	}
	return d, func() {
		logs.Info("message:", "链接数据库")
	}, nil
}

func NewUserServiceClient(r registry.Discovery) uClient.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///T.user.srv"), //链接远程服务
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			metadata.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := uClient.NewUserClient(conn)

	return c
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
