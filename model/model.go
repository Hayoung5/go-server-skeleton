package model

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client       *mongo.Client
	colEmployees *mongo.Collection
}

type Employee struct {
	Name string `bson:"name"`
	ID   int    `bson:"id"`
}

func NewModel() (*Model, error) {
	r := &Model{}
	return r, nil

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-ready")
		r.colEmployees = db.Collection("tEmployee")
	}

	return r, nil
}

func (p *Model) GetEmployee() []Employee {
	filter := bson.D{}
	cursor, err := p.colEmployees.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var employees []Employee
	for _, result := range employees {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	return employees
}
