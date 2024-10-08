package main

import (
	"log"
	"sync"
	"time"

	"github.com/channeldorg/channeld/examples/unity-mirror-tanks/tankspb"
	"github.com/channeldorg/channeld/pkg/channeldpb"
	"github.com/channeldorg/channeld/pkg/client"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// var ctxKeyTanksChannelData struct{}
// var ctxKeyClientNetId struct{}

const (
	ctxKeyTanksChannelData = 1
	ctxKeyClientNetId      = 2
)

func TanksInitFunc(c *client.ChanneldClient, data *clientData) {
	data.ctx[ctxKeyTanksChannelData] = &tankspb.TankGameChannelData{}
	c.AddMessageHandler(uint32(channeldpb.MessageType_CHANNEL_DATA_UPDATE), wrapTanksChannelDataUpateHandle(data))
}

var TanksClientActions = []*clientAction{
	{
		name:        "updateTankTransform",
		probability: 1,
		minInterval: time.Millisecond * 200,
		perform: func(client *client.ChanneldClient, data *clientData) bool {
			tanksChannelData, ok := data.ctx[ctxKeyTanksChannelData].(*tankspb.TankGameChannelData)
			if !ok {
				return false
			}

			netId, ok := data.ctx[ctxKeyClientNetId].(uint32)
			if !ok {
				log.Printf("netId is not set in the ctx for client %d!\n", client.Id)
				return true
			}

			transform, exists := tanksChannelData.TransformStates[netId]
			if !exists {
				return false
			}
			if transform.Removed {
				delete(tanksChannelData.TransformStates, netId)
				return true
			}

			transform.Rotation.Y += data.rnd.Float32() * 0.1
			/*
				if transform.Position == nil {
					log.Println("transform.position is not initialized yet!")
					return false
				}
				pos := transform.Position
				pos.Z += 0.1
				log.Printf("updating transform (netId=%d) to %s\n", netId, pos.String())
			*/

			any, err := anypb.New(&tankspb.TankGameChannelData{
				TransformStates: map[uint32]*channeldpb.TransformState{
					netId: {
						//Position: transform.Position,
						Rotation: transform.Rotation,
					},
				},
			})
			if err != nil {
				log.Println(err)
				return false
			}
			client.Send(0, channeldpb.BroadcastType_NO_BROADCAST, uint32(channeldpb.MessageType_CHANNEL_DATA_UPDATE), &channeldpb.ChannelDataUpdateMessage{
				Data: any,
			}, nil)

			return true
		},
	},
	{
		name:        "fire",
		probability: 0,
		minInterval: time.Millisecond * 200,
		perform: func(client *client.ChanneldClient, data *clientData) bool {

			return true
		},
	},
	{
		name:        "ping",
		probability: 0,
		minInterval: time.Millisecond * 200,
		perform: func(client *client.ChanneldClient, data *clientData) bool {

			return true
		},
	},
}

// netId -> connId
var tanksNetIdMapping sync.Map

func wrapTanksChannelDataUpateHandle(data *clientData) client.MessageHandlerFunc {
	return func(client *client.ChanneldClient, channelId uint32, m client.Message) {
		tanksChannelData, ok := data.ctx[ctxKeyTanksChannelData].(*tankspb.TankGameChannelData)
		if !ok {
			log.Println("tanksChannelData is not initialized in the ctx!")
			return
		}
		updateMsg, _ := m.(*channeldpb.ChannelDataUpdateMessage)
		var channelData tankspb.TankGameChannelData
		updateMsg.Data.UnmarshalTo(&channelData)
		proto.Merge(tanksChannelData, &channelData)

		if _, exists := data.ctx[ctxKeyClientNetId].(uint32); exists {
			//log.Printf("received transform (netId=%d): %s\n", netId, tanksChannelData.TransformStates[netId])
			return
		}

		// Randomly pick a tank that is not taken by other client (netId is not registered in tanksNetIdMapping)
		for netId := range tanksChannelData.TankStates {
			if _, ok := tanksNetIdMapping.Load(netId); !ok {
				tanksNetIdMapping.Store(netId, client.Id)
				data.ctx[ctxKeyClientNetId] = netId
				log.Printf("set netId=%d for client %d\n", netId, client.Id)
				break
			}
		}
	}
}
