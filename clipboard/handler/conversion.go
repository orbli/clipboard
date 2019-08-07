package handler

import (
	pb "gitlab.com/orbli/clipboard/clipboard/proto"
	"gitlab.com/orbli/clipboard/clipboard/storage"
)

func pbToInternal(pbm *pb.Message) (storage.Message, error) {
	return storage.Message{
		Key:   pbm.Key,
		Value: pbm.Value,
	}, nil
}

func internalToPb(sm storage.Message) (*pb.Message, error) {
	return &pb.Message{
		Key:   sm.Key,
		Value: sm.Value,
	}, nil
}
