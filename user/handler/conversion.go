package handler

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"gitlab.com/orbli/clipboard/user/model"
	pb "gitlab.com/orbli/clipboard/user/proto"
)

func pbToInternal(pbm *pb.User) (model.User, error) {
	rt := model.User{
		Id:   pbm.Id,
		Name: pbm.Name,
	}

	if metadata, err := proto.Marshal(pbm.GetMetadata()); err == nil {
		rt.Metadata = metadata
	}

	return rt, nil
}

func internalToPb(sm model.User) (*pb.User, error) {
	rt := &pb.User{
		Id:   sm.Id,
		Name: sm.Name,
	}

	metadata := &any.Any{}
	if err := proto.Unmarshal(sm.Metadata, metadata); err == nil {
		rt.Metadata = metadata
	}

	return rt, nil
}
