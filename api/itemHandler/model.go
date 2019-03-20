package itemHandler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tryTwo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var itemCollection *mongo.Collection

type Item struct {
	Id              string `json:"_id"`
	ItemName        string `json:"itemName"`
	CategoryName    string `json:"categoryName"`
	SubCategoryName string `json:"subCategoryName"`
	ItemDescription string `json:"itemDescription"`
	CreatedBy       string `json:"createdBy`
	UpdatedAt       string `json:"updatedAt"`
	CreatedAt       string `json:"createdAt"`
}

func (i *Item) PreSave() {
	i.UpdatedAt = time.Now().String()
	i.CreatedAt = time.Now().String()
}

func (i *Item) PreUpdate() {
	i.UpdatedAt = time.Now().String()
}

func (i *Item) Save() (*mongo.InsertOneResult, error) {
	setParams()
	i.PreSave()

	return itemCollection.InsertOne(ctx, bson.D{
		{"itemName", i.ItemName},
		{"categoryName", i.CategoryName},
		{"subCategoryName", i.SubCategoryName},
		{"itemDescription", i.ItemDescription},
		{"createdBy", i.CreatedBy},
		{"updatedAt", i.UpdatedAt},
		{"createdAt", i.CreatedAt},
	})
}

func (i *Item) Update() (*mongo.UpdateResult, error) {
	setParams()
	i.PreUpdate()

	fmt.Println("object ID from bahar se")
	fmt.Println(i.Id)
	objectId, _ := primitive.ObjectIDFromHex(i.Id)

	return itemCollection.UpdateOne(
		ctx,
		bson.D{{"_id", objectId}},
		bson.D{
			{"$set", bson.D{
				{"itemName", i.ItemName},
				{"categoryName", i.CategoryName},
				{"subCategoryName", i.SubCategoryName},
				{"itemDescription", i.ItemDescription},
				{"createdBy", i.CreatedBy},
				{"updatedAt", i.UpdatedAt},
			},
			},
		})
}

func (i *Item) Delete() (*mongo.DeleteResult, error) {
	setParams()
	// No error handling, becasue too lazy to put one
	objectId, _ := primitive.ObjectIDFromHex(i.Id)

	return itemCollection.DeleteOne(ctx, bson.D{{"_id", objectId}})
}

func setParams() {
	ctx = context.Background()
	itemCollection = db.GetConnection().Collection("item")
}

func List(skip int64, limit int64) ([]Item, error) {
	setParams()

	itemCur, err := itemCollection.Find(ctx, bson.M{}, options.Find().SetSkip(skip).SetLimit(limit))
	defer itemCur.Close(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var items []Item
	for itemCur.Next(nil) {
		elem := &bson.D{}
		err := itemCur.Decode(elem)

		if err != nil {
			log.Fatal("Decode error ", err)
		}

		m := elem.Map()

		item := Item{
			Id:              m["_id"].(primitive.ObjectID).Hex(),
			ItemName:        m["itemName"].(string),
			CategoryName:    m["categoryName"].(string),
			SubCategoryName: m["subCategoryName"].(string),
			ItemDescription: m["itemDescription"].(string),
			CreatedBy:       m["createdBy"].(string),
			UpdatedAt:       m["updatedAt"].(string),
			CreatedAt:       m["createdAt"].(string),
		}

		items = append(items, item)
	}
	return items, nil
}
