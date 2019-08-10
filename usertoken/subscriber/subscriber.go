package subscriber

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	pbToken "gitlab.com/orbli/clipboard/token/proto"
	pbEvent "gitlab.com/orbli/clipboard/util/proto/event"
	// "github.com/micro/go-micro/util/log"
	// "github.com/micro/go-micro/metadata"
)

type (
	UserSubscriber struct {
		Client client.Client
	}
)

func (s *UserSubscriber) Process(ctx context.Context, event *pbEvent.Event) error {
	// md, _ := metadata.FromContext(ctx)
	// log.Logf("Received event %+v with metadata %+v\n", event, md)
	if event.Action == "DELETE" {
		tokenService := pbToken.NewTokenService("orbli.micro.token", s.Client)
		target := fmt.Sprintf("%s@User", event.Message)
		_, err := tokenService.DeleteParentedTokens(ctx, &pbToken.Token{Parent: target})
		if err != nil {
			return err
		}
	}
	return nil
}
