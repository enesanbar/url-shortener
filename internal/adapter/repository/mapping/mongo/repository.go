package mongo

import (
	"context"
	"time"

	"go.uber.org/zap"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/enesanbar/go-service/router"

	"github.com/enesanbar/url-shortener/internal/usecase/mapping/getall"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/enesanbar/go-service/errors"
	"github.com/enesanbar/go-service/log"
	"go.uber.org/fx"

	"github.com/enesanbar/url-shortener/internal/domain"
)

type MappingMongoAdapter struct {
	coll   *mongo.Collection
	logger log.Factory
}

type Params struct {
	fx.In

	Coll   *mongo.Collection `name:"default"`
	Logger log.Factory
}

func NewMappingMongoAdapter(p Params) (*MappingMongoAdapter, error) {
	return &MappingMongoAdapter{
		coll:   p.Coll,
		logger: p.Logger,
	}, nil
}

func (r *MappingMongoAdapter) Store(ctx context.Context, m *domain.Mapping) (*domain.Mapping, error) {
	doc := bson.D{
		{Key: "code", Value: m.Code},
		{Key: "url", Value: m.URL},
		{Key: "expires_at", Value: m.ExpiresAt},
		{Key: "created_at", Value: m.CreatedAt},
		{Key: "updated_at", Value: m.UpdatedAt},
	}
	_, err := r.coll.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	result, err := r.FindByCode(ctx, m.Code)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *MappingMongoAdapter) Update(ctx context.Context, m *domain.Mapping) (*domain.Mapping, error) {
	filter := bson.D{{Key: "code", Value: m.Code}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "url", Value: m.URL},
		{Key: "expires_at", Value: m.ExpiresAt},
		{Key: "updated_at", Value: time.Now()},
	}}}

	res, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if res.ModifiedCount == 0 {
		return nil, errors.Error{Code: errors.ENOTMODIFIED}
	}

	result, err := r.FindByCode(ctx, m.Code)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MappingMongoAdapter) FindAll(ctx context.Context, request *getall.Request) (*router.PagedResponse, error) {
	result := make([]*domain.Mapping, 0)

	filter := bson.M{}
	// Calculate number of documents to skip
	skips := request.PageSize * (request.Page - 1)
	opts := &options.FindOptions{
		Limit: &request.PageSize,
		Skip:  &skips,
	}

	if request.CodeQuery != "" {
		filter["code"] = bson.M{"$regex": primitive.Regex{
			Pattern: ".*" + request.CodeQuery + ".*",
			Options: "i",
		}}
	}

	if request.URLQuery != "" {
		filter["url"] = bson.M{"$regex": primitive.Regex{
			Pattern: ".*" + request.URLQuery + ".*",
			Options: "i",
		}}
	}

	if request.SortBy != "" {
		sortDirection := -1
		if request.SortOrder == "asc" {
			sortDirection = 1
		}
		opts.SetSort(bson.D{{Key: request.SortBy, Value: sortDirection}})
	}

	cursor, err := r.coll.Find(ctx, filter, opts)

	if err != nil {
		return nil, errors.Error{Err: err}
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			r.logger.For(ctx).With(zap.Error(err)).Error("unable to close cursor")
		}
	}()

	for cursor.Next(ctx) {
		var m domain.Mapping
		err := cursor.Decode(&m)
		if err != nil {
			return nil, err
		}

		result = append(result, &m)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Error{Err: err}
	}

	count, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, errors.Error{Err: err}
	}

	return router.NewPagedResponse(result, request.Page, request.PageSize, count), nil
}

func (r *MappingMongoAdapter) FindByCode(ctx context.Context, code string) (*domain.Mapping, error) {
	var result domain.Mapping
	filter := bson.D{{Key: "code", Value: code}}

	err := r.coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, errors.Error{
				Code:    errors.ENOTFOUND,
				Message: "mapping not found",
				Err:     domain.ErrMappingNotFound,
			}
		default:
			return nil, errors.Error{Err: err}
		}
	}

	return &result, nil
}

func (r *MappingMongoAdapter) Delete(ctx context.Context, code string) error {
	filter := bson.D{{Key: "code", Value: code}}

	res, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Error{Err: err}
	}

	if res.DeletedCount == 0 {
		return errors.Error{Code: errors.ENOTMODIFIED}
	}

	return nil
}
