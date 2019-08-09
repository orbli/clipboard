package handler

import (
	"context"

	"gitlab.com/orbli/clipboard/clipboard/model"
	pb "gitlab.com/orbli/clipboard/clipboard/proto"
	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	Clipboard struct{}
)

var (
	_ pb.ClipboardHandler = Clipboard{}
)

func (Clipboard) Create(ctx context.Context, req *pb.Message, res *pb.Message) error {
	message, err := pbToInternal(req)
	if err != nil {
		return err
	}
	if err := storage.Set(message); err != nil {
		return err
	}
	return Clipboard{}.Read(ctx, req, res)
}

func (Clipboard) Read(ctx context.Context, req *pb.Message, res *pb.Message) error {
	v, err := storage.Get(req.GetKey())
	if err != nil {
		return err
	}
	pbv, err := internalToPb(v.(model.Message))
	if err != nil {
		return err
	}
	*res = *pbv
	return nil
}

func (Clipboard) Update(ctx context.Context, req *pb.Message, res *pb.Message) error {
	message, err := pbToInternal(req)
	if err != nil {
		return err
	}
	if err := storage.Set(message); err != nil {
		return err
	}
	return Clipboard{}.Read(ctx, req, res)
}

func (Clipboard) Delete(ctx context.Context, req *pb.Message, res *pb.Message) error {
	if err := storage.Delete(req.GetKey()); err != nil {
		return err
	}
	return nil
}
