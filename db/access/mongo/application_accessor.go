package mongo

import (
	"context"
	"versioner/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ApplicationsAccessor mongo implementation of ApplicationAccessor
type MongoApplicationsAccessor struct {
	collection *mongo.Collection
}

// NewMongoApplicationsAccessor constructs an MongoApplicationsAccessor
func NewMongoApplicationsAccessor(collection *mongo.Collection) *MongoApplicationsAccessor {
	return &MongoApplicationsAccessor{
		collection: collection,
	}
}

// ApplicationVersionExists returns whether a given application is registered as having a version
func (a *MongoApplicationsAccessor) ApplicationVersionExists(id string) (bool, error) {
	filter := bson.D{{Key: "id", Value: id}}
	count, err := a.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetApplicationVersion returns a version document for an application
func (a *MongoApplicationsAccessor) GetApplicationVersion(id string) (models.ApplicationVersion, error) {
	var doc models.ApplicationVersion
	filter := bson.D{{Key: "id", Value: id}}

	result := a.collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return models.ApplicationVersion{}, result.Err()
	}

	err := result.Decode(&doc)
	if err != nil {
		return models.ApplicationVersion{}, err
	}

	return doc, nil
}

// UpdateApplicationVersion updates an application version
func (a *MongoApplicationsAccessor) UpdateApplicationVersion(version models.ApplicationVersion) error {
	filter := bson.D{{Key: "id", Value: version.ID}}
	_, err := a.collection.ReplaceOne(context.Background(), filter, version)
	if err != nil {
		return err
	}

	return nil
}
