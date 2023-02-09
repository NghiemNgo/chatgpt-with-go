package chatGpt

import (
	"go.tienngay/pkg/mongo/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
)

//Repository interface allows us to access the CRUD Operations in mongo here.
type ChatGptRepository interface {
	Insert(chat entities.ChatGpt) (string)
}
type chatGptRepository struct {
	Collection *mongo.Collection
}

//NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) ChatGptRepository {
	return &chatGptRepository{
		Collection: collection,
	}
}

//Insert new data
func (r *chatGptRepository) Insert(chat entities.ChatGpt) (string) {
	result, err := r.Collection.InsertOne(context.TODO(), chat)
	if err != nil {
		fmt.Println(err)
	    return ""
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
	    return oid.String()
	} else {
	    return ""
	}
}