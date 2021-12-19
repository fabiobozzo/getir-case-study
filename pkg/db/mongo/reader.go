package mongo

import (
	"context"
	"getir-case-study/pkg/db"
	"getir-case-study/pkg/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

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

	// Please read the comment over `aggregateRecords` below
	//cursor, err := r.aggregateRecords(ctx, dateRange, countRange, collection)

	// Suboptimal (but workable) solution follows
	filter := bson.M{
		"createdAt": bson.M{
			"$gte": primitive.NewDateTimeFromTime(dateRange.Start),
			"$lt":  primitive.NewDateTimeFromTime(dateRange.End.Add(time.Hour * 24)),
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []model.Record{}, err
	}

	var records []model.Record

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item bson.M
		if err = cursor.Decode(&item); err != nil {
			return nil, err
		}

		if filterByCountRange(item, &countRange) != nil {
			record, err := r.bsonToModel(item)
			if err != nil {
				return nil, err
			}

			records = append(records, record)
		}
	}

	return records, nil
}

func (r *Reader) bsonToModel(item bson.M) (model.Record, error) {
	record := model.Record{
		Id:         (item["_id"].(primitive.ObjectID)).Hex(),
		Key:        item["key"].(string),
		Value:      item["value"].(string),
		CreatedAt:  (item["createdAt"].(primitive.DateTime)).Time(),
		TotalCount: item["total"].(int),
	}

	return record, nil
}

// This function sucks!
// But it's necessary, in absence of a MongoDB Atlas proper tier.
// Time complexity: O(n^2)
// Space complexity: O(n)
func filterByCountRange(item bson.M, countRange *db.CountRange) bson.M {
	total := 0
	counts := item["counts"].(primitive.A)
	for _, c := range counts {
		total += int(c.(int32))
	}

	if total <= countRange.Max && total >= countRange.Min {
		item["total"] = total

		return item
	}

	return nil
}

// The following code is the ideal implementation of the query
// It would make mongodb perform filtering on the total count too.
// Unfortunately, the error generated is quite clear:
// AtlasError: sum is not allowed in this atlas tier
func (r *Reader) aggregateRecords(
	ctx context.Context,
	dateRange db.DateRange,
	countRange db.CountRange,
	collection *mongo.Collection,
) (*mongo.Cursor, error) {
	matchDateRange := bson.D{{
		"$match", bson.M{
			"createdAt": bson.M{
				"$gte": primitive.NewDateTimeFromTime(dateRange.Start),
				"$lt":  primitive.NewDateTimeFromTime(dateRange.End.Add(time.Hour * 24)),
			},
		},
	}}

	addSumField := bson.D{{
		"sum", bson.M{
			"$sum": "$counts",
		},
	}}

	matchTotalCountRange := bson.D{{
		"$match", bson.M{
			"sum": bson.M{
				"$gte": countRange.Min,
				"$lte": countRange.Max,
			},
		},
	}}

	pipeline := mongo.Pipeline{
		matchDateRange,
		addSumField,
		matchTotalCountRange,
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}
