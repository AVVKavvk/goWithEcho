package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Product struct{
	Id  primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Price int `json:"price" bson:"price"`
	Currency string `json:"currency" bson:"currency"`
	Qunatity string `json:"qunatity" bson:"qunatity"`
	Discount int `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor string `json:"vendor" bson:"vendor"`
	Accessories []string `json:"accessories,omitempty" bson:"accessories,omitempty" ` 
	SkuId string `json:"sku_id" bson:"sku_id"`


}
func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil{
		log.Fatal(err);
	}

	ctx,cancel:=context.WithTimeout(context.Background(),30*time.Second);
	defer cancel();
	err = client.Connect(ctx)
	if err!=nil {
		log.Fatal(err);
	}
	db:=client.Database("tronics");
	collection:=db.Collection("product");

	// var Iphone= Product{
	// 	Id: primitive.NewObjectID(),
	// 	Name: "Iphone",
	// 	Price: 1000,
	// 	Currency: "R",
	// 	Qunatity: "10",
	// 	Vendor: "Apple",
	// 	Discount: 12,
	// 	SkuId: "12345",
	// 	Accessories: []string{"charger","Box","headset"},
	// }
	// var Trimmer= Product{
	// 	Id: primitive.NewObjectID(),
	// 	Name: "Trimmer",
	// 	Price: 600,
	// 	Currency: "R",
	// 	Qunatity: "25",
	// 	Vendor: "Sony",
	// 	Discount: 6,
	// 	SkuId: "4512",
	// 	Accessories: []string{"charger","extras","headset"},
	// }
	// var Laptop= Product{
	// 	Id: primitive.NewObjectID(),
	// 	Name: "Laptop",
	// 	Price: 15000,
	// 	Currency: "R",
	// 	Qunatity: "4",
	// 	Vendor: "Dell",
	// 	Discount: 15,
	// 	SkuId: "686",
	// 	Accessories: []string{"charger","screen","headset"},
	// }
	// res,err:=collection.InsertOne(context.Background(),Iphone)
	// res,err:=collection.InsertOne(context.Background(),Laptop)
	// res,err:=collection.InsertOne(context.Background(),bson.M{
	// 	"name":"vipin",
	// 	" age":"21",
	// 	"hoobies":bson.A{"music","circket"},
	// })
	// if err!=nil{
	// 	log.Fatal(err);
	// }
	// fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())

	// var pro Product;

	// err=collection.FindOne(context.Background(),bson.M{"price":1000}).Decode(&pro)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(pro.Name)

	
	// cursor,err:=collection.Find(context.Background(),bson.M{"price":bson.M{"$gt":890}});
	// if err!=nil{
	// 	log.Fatal(err);
	// }

	// var pro Product;

	// for cursor.Next(context.Background()){
	// 	err=cursor.Decode(&pro);
	// 	if err!=nil {
	// 		log.Fatal(err);
	// 	}
	// 	fmt.Println(pro.Name);
	// }
	// logicalFilter := bson.M{
	// 	"$and": bson.A{
	// 		bson.M{"price": bson.M{"$gt": 890}},
	// 		bson.M{"discount": bson.M{"$gt": 12}},
	// 	},
	// }
	// cursor,err:=collection.Find(context.Background(),logicalFilter)

	// var pro Product;

	// for cursor.Next(context.Background()){
	// 	err=cursor.Decode(&pro);
	// 	if err!=nil{
	// 		log.Fatal(err);
	// 	}
	// 	fmt.Println(pro.Name)
	// }

	// cursor,err:=collection.Find(context.Background(),bson.M{"accessories":bson.M{"$exists":false}})

	// if err!=nil {
	// 	log.Fatal(err);
	// }

	// var pro Product;

	// for cursor.Next(ctx){
	// 	err=cursor.Decode(&pro);
	// 	if err!=nil {
	// 		log.Fatal(err);
	// 	}
	// 	fmt.Println(pro.Name)
	// }

	// cursor,err:=collection.Find(context.Background(),
	// bson.M{"accessories":bson.M{"$all":bson.A{"charger","screen"}}})
	
	// if err!=nil{
	// 	log.Fatal(err)
	// }

	// var pro Product;

	// for cursor.Next(ctx){
	// 	err=cursor.Decode(&pro);
	// 	if err!=nil{
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(pro.Name)
	// }

	// updateFilled:=bson.M{"$set":bson.M{"IsEssential":false}}

	// updatedData,err:=collection.UpdateMany(ctx,bson.M{},updateFilled);

	// if err!=nil{
	// 	log.Fatal(err);
	// }

	// logicalFilter := bson.M{
	// 	"$and": bson.A{
	// 		bson.M{"price": bson.M{"$gt": 890}},
	// 		bson.M{"discount": bson.M{"$gt": 12}},
	// 	},
	// }

	// updateArray,err:= collection.UpdateMany(ctx,logicalFilter,
	// 	bson.M{"$addToSet":bson.M{"accessories":"manual"}})
	
	// if err!=nil {
	// 	log.Fatal(err);
	// }
	// fmt.Println(updateArray.ModifiedCount)


	// updatedNewData,err:=collection.UpdateMany(ctx,bson.M{},
	// bson.M{"$mul":bson.M{"price":1.20}})

	// if err!=nil {
	// 	log.Fatal(err);
	// }

	// fmt.Println(updatedNewData.ModifiedCount)
	

}