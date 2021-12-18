package mongo

import (
	"context"
	"getir-case-study/pkg/db"
	"getir-case-study/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reader struct {
	db             *mongo.Database
	collectionName string
}

func NewReader(db *mongo.Database, collectionName string) db.Reader {
	return &Reader{
		db:             db,
		collectionName: collectionName,
	}
}

func (r *Reader) RecordsByDateAndCountRange(
	ctx context.Context,
	dateRange db.DateRange,
	countRange db.CountRange,
) ([]model.Record, error) {
	collection := r.db.Collection(r.collectionName)

	filter := bson.D{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var records []bson.M
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	return r.bsonToModel(records)
}

func (r *Reader) bsonToModel(data []bson.M) ([]model.Record, error) {
	var records []model.Record

	return records, nil
}
