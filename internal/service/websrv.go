package service

import (
	"context"
	"fmt"

	pb "web-api/api/web"
	"web-api/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type WebSrvService struct {
	// pb.UnimplementedwebsrvServer
	uc *biz.UserUseCase
	pb.UnimplementedWebsrvServer
	log *log.Helper
}

func NewWebsrvService(uc *biz.UserUseCase, logger log.Logger) *WebSrvService {
	return &WebSrvService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

//用户注册
func (s *WebSrvService) Registory(ctx context.Context, req *pb.RegistoryReq) (*pb.RegistoryReply, error) {
	m := &biz.User{

		NickName:  req.Nickname,
		AwatarUrl: req.AvatarUrl,
		Uname:     req.Uname,
		Pwd:       req.Pwd,
		Country:   req.Country,
		City:      req.City,
		Province:  req.Province,
	}
	err := s.uc.Add(ctx, m)
	if err != nil {
		return &pb.RegistoryReply{
			Ok:  "false",
			Msg: "注册失败！",
		}, nil
	}
	return &pb.RegistoryReply{
		Ok:  "true",
		Msg: "注册成功！",
	}, nil
}
func (s *WebSrvService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	fmt.Println("web-service-----1----")
	u, err := s.uc.Login(ctx, req.Nickname, req.Pwd)
	fmt.Println("web-service--------2-")

	if err != nil {
		fmt.Println("web-service-----err---1-")
		fmt.Println(err.Error())
		return &pb.LoginReply{UserDetail: nil, Ok: "true", Msg: "登陆失败!"}, err
	}
	ud := &pb.User{
		Id:        u.UserDetail.Id,
		Nickname:  u.UserDetail.Nickname,
		AvatarUrl: u.UserDetail.AvatarUrl,
		Country:   u.UserDetail.Country,
		Province:  u.UserDetail.Province,
		City:      u.UserDetail.City,
		Uname:     u.UserDetail.Uname,
		Status:    u.UserDetail.Status,
	}
	fmt.Println("web-service-----err---2-")
	return &pb.LoginReply{
		UserDetail: ud,
		Msg:        "恭喜登陆成功了！",
		Ok:         "true",
	}, nil
}
