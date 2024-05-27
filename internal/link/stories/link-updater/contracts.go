package link_updater

import (
	"appGo/internal/database"
	"context"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repository interface {
	FindByID(ctx context.Context, id primitive.ObjectID) (database.Link, error)
	Update(ctx context.Context, req database.UpdateLinkReq) (database.Link, error)
}

type amqpConsumer interface {
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (
		<-chan amqp.Delivery,
		error,
	)
}

type logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
}
