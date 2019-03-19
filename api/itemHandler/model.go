package itemHandler

import (
	"context"
	"log"
	"time"

	"github.com/tryTwo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var itemCollection *mongo.Collection

type Item struct {
	Id              string    `json:"itemId"`
	ItemName        string    `json:"itemName"`
	CategoryName    string    `json:"categoryName"`
	SubCategoryName string    `json:"subCategoryName"`
	ItemDescription string    `json:"itemDescription"`
	CreatedBy       string    `json:"createdBy`
	UpdatedAt       time.Time `json:"updatedAt"`
	CreatedAt       time.Time `json:"createdAt"`
}

func (i *Item) PreSave() {
	i.UpdatedAt = time.Now()
	i.CreatedAt = time.Now()
}

func (i *Item) PreUpdate() {
	i.UpdatedAt = time.Now()
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
		{"UpdatedAt", i.UpdatedAt},
		{"CreatedAt", i.CreatedAt},
	})
}

func (i *Item) Update() (*mongo.UpdateResult, error) {
	setParams()
	i.PreUpdate()

	return itemCollection.UpdateOne(
		ctx,
		i.Id,
		bson.D{
			{"$set", bson.D{
				{"itemName", i.ItemName},
				{"categoryName", i.CategoryName},
				{"subCategoryName", i.SubCategoryName},
				{"itemDescription", i.ItemDescription},
				{"createdBy", i.CreatedBy},
				{"UpdatedAt", i.UpdatedAt},
			},
			},
		})
}

func (i *Item) Delete() (*mongo.DeleteResult, error) {
	setParams()
	return itemCollection.DeleteOne(ctx, bson.D{{"_id", i.Id}})
}

func setParams() {
	ctx = context.Background()
	itemCollection = db.GetConnection().Collection("item")
}

func List(skip int64, limit int64) []Item {
	setParams()
	itemCur, _ := itemCollection.Find(ctx, bson.M{}, options.Find().SetSkip(skip).SetLimit(limit))
	defer itemCur.Close(ctx)

	var items []Item
	for itemCur.Next(nil) {
		item := Item{}
		err := itemCur.Decode(&item)
		if err != nil {
			log.Fatal("Decode error ", err)
		}
		items = append(items, item)
	}

	return items
}
