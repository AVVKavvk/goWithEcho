package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
 
type User struct{
	Name string `bson:"name"`
	Age int `bson:"age"`
}
func insertOne(collection *mongo.Collection,user User) (*mongo.InsertOneResult, error){
	res,err:=collection.InsertOne(context.Background(),user);
	return res,err
}

func main() {
	c, err := mongo.Connect(context.Background(),options.Client().ApplyURI("mongodb://localhost:27017"))
	if err !=nil{
		log.Fatalf("Error : %v",err.Error())
	}
	db:=c.Database("UnitTesting");
	col:=db.Collection("Test")
	var user=User{
		Name :"Vipin",
		Age:21,
	}
	res,err:=insertOne(col,user);
	// log.Println(col,err);
	log.Println(res,err)


}