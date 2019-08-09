package handler

import (
	"context"
	"errors"
	"strconv"

	"gitlab.com/orbli/clipboard/user/model"
	pb "gitlab.com/orbli/clipboard/user/proto"
	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	UserService struct{}
)

var (
	_ pb.UserServiceHandler = UserService{}
)

func (UserService) Create(ctx context.Context, req *pb.User, res *pb.User) error {
	message, err := pbToInternal(req)
	if err != nil {
		return err
	}

	if err = (UserService{}.Read(ctx, req, res)); err == nil {
		return errors.New("User exist!")
	}

	if err := storage.Set(message); err != nil {
		return err
	}

	pbv, err := internalToPb(message)
	if err != nil {
		return err
	}
	return UserService{}.Read(ctx, pbv, res)
}

func (UserService) Read(ctx context.Context, req *pb.User, res *pb.User) error {
	v, err := storage.Get(strconv.FormatUint(req.Id, 10))
	if err != nil {
		return err
	}
	pbv, err := internalToPb(v.(model.User))
	if err != nil {
		return err
	}
	*res = *pbv
	return nil
}

func (UserService) Update(ctx context.Context, req *pb.User, res *pb.User) error {
	message, err := pbToInternal(req)
	if err != nil {
		return err
	}
	if err := storage.Set(message); err != nil {
		return err
	}
	return UserService{}.Read(ctx, req, res)
}

func (UserService) Delete(ctx context.Context, req *pb.User, res *pb.User) error {
	if err := storage.Delete(strconv.FormatUint(req.Id, 10)); err != nil {
		return err
	}
	return nil
}
