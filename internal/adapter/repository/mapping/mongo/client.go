package mongo

import (
	"context"

	"go.uber.org/zap"

	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/persistance/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

type ClientParams struct {
	fx.In

	Logger    log.Factory
	Config    *mongodb.Config `name:"default"`
	Connector *mongodb.Connector
}

type Client struct {
	cfg       *mongodb.Config
	client    *mongo.Client
	connector *mongodb.Connector
	logger    log.Factory
}

func NewMongoClient(p ClientParams) (*Client, error) {
	client, err := p.Connector.Connect(p.Config)
	if err != nil {
		return nil, err
	}
	return &Client{
		cfg:    p.Config,
		client: client,
		logger: p.Logger,
	}, nil
}

func (c *Client) Connect(ctx context.Context) error {
	c.logger.For(ctx).
		With(zap.String("db", c.cfg.Host)).
		Info("connecting to the database")
	client, err := c.connector.Connect(c.cfg)
	if err != nil {
		return err
	}
	c.client = client
	c.logger.For(ctx).
		With(zap.String("db", c.cfg.Host)).
		Info("connected to the database")

	return nil
}

func (c *Client) Close(ctx context.Context) error {
	c.logger.For(ctx).
		With(zap.String("db", c.cfg.Host)).
		Info("closing connection to the database")
	return c.client.Disconnect(ctx)
}

func (c *Client) Start(ctx context.Context) error {
	return nil
}

func (c *Client) Name() string {
	return c.cfg.Name
}
