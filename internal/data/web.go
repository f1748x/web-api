package data

import (
	"context"
	"fmt"
	v1 "web-api/api/user"
	"web-api/internal/biz"

	// "github.com/coreos/etcd/error"

	"github.com/go-kratos/kratos/v2/log"
)

type webDataRepo struct {
	log  *log.Helper
	data *Data
}

func NewWebRepo(logger log.Logger, data *Data) biz.UserRepo {
	return webDataRepo{
		log:  log.NewHelper(logger),
		data: data,
	}
}

func (c webDataRepo) Registry(ctx context.Context, user *biz.User) error {
	// c.Data.userClient
	//c.data.userClient
	// c.data.userClient.CreateUser()
	_, err := c.data.userClient.CreateUser(ctx, &v1.CreateUserReq{
		Nickname:  user.NickName,
		Pwd:       user.Pwd,
		AvatarUrl: user.AwatarUrl,
		Country:   user.Country,
		Province:  user.Province,
		City:      user.City,
		Uname:     user.Uname,
	})
	if err != nil {
		//用户创建失败！
		return err
	}
	//用户创建成功!
	return nil
}

//用户登陆
func (c webDataRepo) Login(ctx context.Context, name, pwd string) (*v1.GetUserRes, error) {
	user, err := c.data.userClient.GetUser(ctx, &v1.GetUserReq{
		Nickname: name,
		Pwd:      pwd,
	})
	if err != nil {
		fmt.Println("client--------------")
		fmt.Println(err.Error())
		return nil, err
	}
	//user.UserDetail
	return user, nil
}
