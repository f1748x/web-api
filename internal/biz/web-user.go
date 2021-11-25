package biz

import (
	"context"
	"fmt"
	v1 "web-api/api/user"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	NickName  string
	Pwd       string
	AwatarUrl string
	Country   string
	Province  string
	City      string
	Uname     string
}
type UserRepo interface {
	Registry(ctx context.Context, user *User) error
	Login(ctx context.Context, name, pwd string) (*v1.GetUserRes, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Add(ctx context.Context, user *User) error {

	return uc.repo.Registry(ctx, user)

}
func (uc *UserUseCase) Login(ctx context.Context, name, pwd string) (*v1.GetUserRes, error) {
	fmt.Println("web-service--------biz.login-")

	return uc.repo.Login(ctx, name, pwd)
}
