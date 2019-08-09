package handler

import (
	"gitlab.com/orbli/clipboard/clipboard/model"
	pb "gitlab.com/orbli/clipboard/clipboard/proto"
)

func pbToInternal(pbm *pb.Message) (model.Message, error) {
	return model.Message{
		KeyString: pbm.Key,
		Value:     pbm.Value,
	}, nil
}

func internalToPb(sm model.Message) (*pb.Message, error) {
	return &pb.Message{
		Key:   sm.KeyString,
		Value: sm.Value,
	}, nil
}
