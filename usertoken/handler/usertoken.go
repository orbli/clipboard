package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/micro/go-micro/client"

	pbToken "gitlab.com/orbli/clipboard/token/proto"
	pbUser "gitlab.com/orbli/clipboard/user/proto"
	pb "gitlab.com/orbli/clipboard/usertoken/proto"
)

type (
	UsertokenService struct {
		Client client.Client
	}
)

var (
	_ pb.UsertokenServiceHandler = UsertokenService{}
)

func (uts UsertokenService) Create(ctx context.Context, req *pb.Usertoken, res *pb.Usertoken) error {
	userService := pbUser.NewUserService("orbli.micro.user", uts.Client)
	_, err := userService.Read(ctx, &pbUser.User{Id: req.UserId})
	if err != nil {
		return err
	}
	tokenService := pbToken.NewTokenService("orbli.micro.token", uts.Client)
	parent := fmt.Sprintf("%d@%s", req.UserId, "User")
	tokenRes, err := tokenService.Create(ctx, &pbToken.Token{Parent: parent})
	if err != nil {
		return err
	}
	*res = pb.Usertoken{
		UserId: req.UserId,
		Token:  tokenRes.Token,
	}
	return nil
}

func (uts UsertokenService) Read(ctx context.Context, req *pb.Usertoken, res *pb.Usertoken) error {
	tokenService := pbToken.NewTokenService("orbli.micro.token", uts.Client)
	tokenRes, err := tokenService.Read(ctx, &pbToken.Token{Token: req.Token})
	if err != nil {
		return err
	}
	x := strings.Split(tokenRes.Parent, "@")
	if x[1] != "User" {
		return errors.New("incorrect token")
	}
	userId, _ := strconv.ParseUint(x[0], 10, 64)
	// Does not check userid since assume it exists (i.e. no system internal logic error)
	*res = pb.Usertoken{
		UserId: userId,
		Token:  req.Token,
	}
	return nil
}

func (uts UsertokenService) Delete(ctx context.Context, req *pb.Usertoken, res *pb.Usertoken) error {
	if err := uts.Read(ctx, req, res); err != nil {
		return err
	}
	if req.UserId != res.UserId {
		return errors.New("incorrect token")
	}
	tokenService := pbToken.NewTokenService("orbli.micro.token", uts.Client)
	_, err := tokenService.Delete(ctx, &pbToken.Token{Token: req.Token})
	if err != nil {
		return err
	}
	*res = pb.Usertoken{}
	return nil
}
