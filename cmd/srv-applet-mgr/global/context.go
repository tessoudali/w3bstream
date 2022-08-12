package global

import (
	"context"

	conflog "github.com/iotexproject/Bumblebee/conf/log"
	confmqtt "github.com/iotexproject/Bumblebee/conf/mqtt"
	"github.com/iotexproject/Bumblebee/kit/sqlx"
	"github.com/iotexproject/Bumblebee/x/contextx"
)

type keyDatabase struct{}

func WithDatabase(db sqlx.DBExecutor) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, keyDatabase{}, db)
	}
}

func WithDatabaseContext(ctx context.Context) context.Context {
	return contextx.WithValue(ctx, keyDatabase{}, postgres)
}

func DBExecutorFromContext(ctx context.Context) sqlx.DBExecutor {
	return ctx.Value(keyDatabase{}).(sqlx.DBExecutor).WithContext(ctx)
}

type keyLogger struct{}

func WithLogger(l conflog.Logger) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, keyLogger{}, l)
	}
}

func WithLoggerContext(ctx context.Context) context.Context {
	return contextx.WithValue(ctx, keyLogger{}, logger)
}

func LoggerFromContext(ctx context.Context) conflog.Logger {
	return ctx.Value(keyLogger{}).(conflog.Logger)
}

type keyMqtt struct{}

func WithMqtt(mqtt *confmqtt.Broker) contextx.WithContext {
	return func(ctx context.Context) context.Context {
		return contextx.WithValue(ctx, keyMqtt{}, mqtt)
	}
}

func WithMqttContext(ctx context.Context) context.Context {
	return contextx.WithValue(ctx, keyMqtt{}, mqtt)
}

func MqttFromContext(ctx context.Context) *confmqtt.Broker {
	return ctx.Value(keyMqtt{}).(*confmqtt.Broker)
}
