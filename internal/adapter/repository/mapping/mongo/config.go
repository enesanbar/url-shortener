package mongo

import (
	"github.com/enesanbar/go-service/config"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/persistance/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

func NewConfig(cfg config.Config, logger log.Factory) (*mongodb.Config, error) {
	return mongodb.NewConfig(cfg, "datasources.mongo.default")
}

type DBParams struct {
	fx.In

	Config *mongodb.Config `name:"default"`
	Conn   *Client         `name:"default"`
}

func NewDatabase(p DBParams) *mongo.Database {
	return p.Conn.client.Database(p.Config.Name)
}

type CollParams struct {
	fx.In

	DB *mongo.Database `name:"default"`
}

func NewMongoCollection(p CollParams) *mongo.Collection {
	return p.DB.Collection("mappings")
}
