package database

import (
	"context"
	"log"
	"time"

	"graphql_go/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		Username:      "graphql_gp",
		Password:      "password123",
		AuthMechanism: "SCRAM-SHA-256",
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials))

	if err != nil {
		log.Fatal(err)
	}

	_ = client.Connect(ctx)

	return &DB{
		client: client,
	}
}

func (db *DB) Save(input *model.NewDog) *model.Dog {
	collection := db.client.Database("animals").Collection("dogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	return &model.Dog{
		ID:        id,
		Name:      input.Name,
		IsGoodBoi: input.IsGoodBoi,
	}
}

func (db *DB) FindById(ID string) *model.Dog {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}

	collection := db.client.Database("animals").Collection("dogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	dog := model.Dog{}
	res.Decode(&dog)

	return &dog
}

func (db *DB) All() []*model.Dog {
	collection := db.client.Database("animals").Collection("dogs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var dogs []*model.Dog
	for cur.Next(ctx) {
		var dog *model.Dog
		err := cur.Decode(&dog)
		if err != nil {
			log.Fatal(err)
		}
		dogs = append(dogs, dog)
	}

	return dogs
}
