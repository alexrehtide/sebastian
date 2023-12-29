package mongoconnection

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, ops MongoOptions) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", ops.Host, ops.Port)))
	if err != nil {
		return nil, err
	}
	return client.Database(ops.Name), nil
}

type MongoOptions struct {
	Host string
	Port string
	Name string
}
