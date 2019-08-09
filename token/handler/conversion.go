package handler

import (
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/orbli/clipboard/token/model"
	pb "gitlab.com/orbli/clipboard/token/proto"
)

func pbToInternal(pbm *pb.Token) (model.Token, error) {
	rt := model.Token{
		Token:  pbm.Token,
		Secret: pbm.Secret,
		Parent: pbm.Parent,
		Data:   pbm.Data,
	}

	if expireAt, err := ptypes.Timestamp(pbm.GetExpireAt()); err == nil {
		rt.ExpireAt = &expireAt
	}

	return rt, nil
}

func internalToPb(sm model.Token) (*pb.Token, error) {
	rt := &pb.Token{
		Token:  sm.Token,
		Secret: sm.Secret,
		Parent: sm.Parent,
		Data:   sm.Data,
	}

	if sm.ExpireAt != nil {
		if expireAt, err := ptypes.TimestampProto(*sm.ExpireAt); err == nil {
			rt.ExpireAt = expireAt
		} else {
			return nil, err
		}
	}

	return rt, nil
}
