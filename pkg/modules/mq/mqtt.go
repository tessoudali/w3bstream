package mq

import (
	"context"
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	conflog "github.com/iotexproject/Bumblebee/conf/log"
	confmqtt "github.com/iotexproject/Bumblebee/conf/mqtt"
	"github.com/iotexproject/Bumblebee/x/mapx"

	"github.com/iotexproject/w3bstream/pkg/depends/protocol/eventpb"
	"github.com/iotexproject/w3bstream/pkg/types"
)

type OnMessage func(ctx context.Context, channel string, data *eventpb.Event) (interface{}, error)

var channels = mapx.New[string, *ChannelContext]()

type ChannelContext struct {
	ctx    context.Context
	cancel context.CancelFunc
	logger conflog.Logger
	Name   string
	cli    *confmqtt.Client
	hdl    OnMessage
}

func (ctx *ChannelContext) Run() {
	_, _l := ctx.logger.Start(ctx.ctx, "Channel Run")
	defer _l.End()
	for {
		select {
		case <-ctx.ctx.Done():
			_l.Info("channel closed")
			return
		default:
			_ = ctx.cli.WithTopic(ctx.Name).Subscribe(
				func(cli mqtt.Client, msg mqtt.Message) {
					_, l := ctx.logger.Start(ctx.ctx, "OnMessage:"+ctx.Name)
					defer l.End()

					pl := msg.Payload()
					ev := &eventpb.Event{}
					err := json.Unmarshal(pl, ev)
					if err != nil {
						ctx.logger.Error(err)
						return
					}
					_, err = ctx.hdl(ctx.ctx, ctx.Name, ev)
					if err != nil {
						ctx.logger.Error(err)
					}
					l.WithValues("payload", ev).Info("sub handled")
				},
			)
		}
	}
}

func (ctx *ChannelContext) Stop() { ctx.cancel() }

func CreateChannel(ctx context.Context, prjName string, hdl OnMessage) error {
	broker := types.MustMqttBrokerFromContext(ctx)

	cli, err := broker.Client(prjName)
	if err != nil {
		return err
	}

	cctx := &ChannelContext{
		Name:   prjName,
		logger: types.MustLoggerFromContext(ctx),
		cli:    cli,
		hdl:    hdl,
	}
	cctx.ctx, cctx.cancel = context.WithCancel(ctx)
	channels.Store(prjName, cctx)

	go cctx.Run()

	return nil
}

func StopChannel(prjName string) {
	c, ok := channels.LoadAndRemove(prjName)
	if !ok {
		return
	}
	c.Stop()
}
