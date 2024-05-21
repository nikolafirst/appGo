package links

import (
	"context"
	"fmt"
	"time"
	"work/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "links"

func New(db *mongo.Database, timeout time.Duration) *Repository {
	return &Repository{db: db, timeout: timeout}
}

type Repository struct {
	db      *mongo.Database
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateReq) (database.Link, error) {
	var l database.Link
	// implement me
	// change
	if _, err := r.db.Collection(collection).InsertOne(ctx, l); err != nil {
		return l, fmt.Errorf("mongo InsertOne: %w", err)
	}

	return l, nil
}

func (r *Repository) FindByUserAndURL(ctx context.Context, link, userID string) ([]database.Link, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var links []database.Link
	// implement me
	// change
	cursor, err := r.db.Collection(collection).Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("mongo Find: %w", err)
	}

	for cursor.Next(ctx) {
		var l database.Link
		if err := cursor.Decode(&l); err != nil {
			return nil, fmt.Errorf("mongo Decode: %w", err)
		}
		links = append(links, l)
	}

	return links, nil
}

func (r *Repository) FindByCriteria(ctx context.Context, criteria Criteria) ([]database.Link, error) {
	return nil, nil
}
