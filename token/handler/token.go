package handler

import (
	"context"
	"crypto/rand"
	"time"

	"gitlab.com/orbli/clipboard/token/model"
	pb "gitlab.com/orbli/clipboard/token/proto"
	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	TokenService struct{}
)

var (
	_ pb.TokenServiceHandler = TokenService{}
)

func (TokenService) Create(ctx context.Context, req *pb.Token, res *pb.Token) error {
	message, err := pbToInternal(req)
	if err != nil {
		return err
	}

	message.Token = make([]byte, 32)
	rand.Read(message.Token)

	if message.Secret == nil {
		message.Secret = make([]byte, 32)
		rand.Read(message.Secret)
	}

	if message.ExpireAt == nil {
		expireAt := time.Now().AddDate(0, 0, 90)
		message.ExpireAt = &expireAt
	}

	if err := storage.Set(message); err != nil {
		return err
	}

	pbv, err := internalToPb(message)
	if err != nil {
		return err
	}
	return TokenService{}.Read(ctx, pbv, res)
}

func (TokenService) Read(ctx context.Context, req *pb.Token, res *pb.Token) error {
	v, err := storage.Get(string(req.GetToken()))
	if err != nil {
		return err
	}
	pbv, err := internalToPb(v.(model.Token))
	if err != nil {
		return err
	}
	*res = *pbv
	return nil
}

func (TokenService) Update(ctx context.Context, req *pb.Token, res *pb.Token) error {
	message, err := pbToInternal(req)
	if err != nil {
		return err
	}
	if err := storage.Set(message); err != nil {
		return err
	}
	return TokenService{}.Read(ctx, req, res)
}

func (TokenService) Delete(ctx context.Context, req *pb.Token, res *pb.Token) error {
	if err := storage.Delete(string(req.GetToken())); err != nil {
		return err
	}
	return nil
}