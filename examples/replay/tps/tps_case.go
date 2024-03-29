package tps

import (
	"log"

	"github.com/metaworking/channeld/pkg/channeldpb"
	"github.com/metaworking/channeld/pkg/client"
	"github.com/metaworking/channeld/pkg/replay"
)

func Run() {
	rc, err := replay.CreateReplayClientByConfigFile("./tps/case-config.json")
	if err != nil {
		log.Panicf("failed to create replay client: %v\n", err)
		return
	}

	rc.SetNeedWaitMessageCallback(func(msgType channeldpb.MessageType, msgPack *channeldpb.MessagePack, c *client.ChanneldClient) bool {
		return msgType == channeldpb.MessageType_SUB_TO_CHANNEL
	})

	rc.Run()

}
