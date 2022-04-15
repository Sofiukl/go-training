package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type employee struct {
	EmpNo     string   `json:"emp_no,omitempty" bson:"emp_no,omitempty"`
	FirstName string   `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Skills    []string `json:"skills,omitempty" bson:"skills,omitempty"`
}

const host = "localhost"
const port = "27017"
const dbname = "sofikuldb"

func main() {
	emp01 := employee{EmpNo: "E01", FirstName: "Sofikul", LastName: "Mallick", Skills: []string{"java", "golang"}}
	connectionString := fmt.Sprintf("mongodb://%s:%s/%s", host, port, dbname)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	dbConnectMsg := fmt.Sprintf("Connected to DB %s", connectionString)
	fmt.Println(dbConnectMsg)
	DB := client.Database(dbname)
	res, err := DB.Collection("employee").InsertOne(context.TODO(), emp01)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("Document inserted with id %s", res.InsertedID))
	}
}
