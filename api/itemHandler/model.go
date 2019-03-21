package itemHandler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/tryTwo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var itemCollection *mongo.Collection

func setParams() {
	ctx = context.Background()
	itemCollection = db.GetConnection().Collection("item")
}
func SaveItem(i *Item) ([]byte, error) {
	setParams()

	res, err := itemCollection.InsertOne(ctx, bson.D{
		{"itemName", i.ItemName},
		{"categoryName", i.CategoryName},
		{"subCategoryName", i.SubCategoryName},
		{"itemDescription", i.ItemDescription},
		{"createdBy", i.CreatedBy},
		{"updatedAt", i.UpdatedAt},
		{"createdAt", i.CreatedAt},
	})
	if err != nil {
		return nil, err
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()

	// convert object ID to bytes
	return []byte(insertedID), nil
}

func UpdateItem(i *Item) ([]byte, error) {
	setParams()

	objectId, _ := primitive.ObjectIDFromHex(i.Id)

	_, err := itemCollection.UpdateOne(
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
	if err != nil {
		return nil, err
	}

	return []byte(i.Id + "updated successfully"), nil
}

func DeleteItem(i *Item) ([]byte, error) {
	setParams()
	objectId, _ := primitive.ObjectIDFromHex(i.Id)

	res, err := itemCollection.DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return nil, err
	}

	return []byte(strconv.Itoa(int(res.DeletedCount))), nil

}

func ItemList(skip, limit int64) ([]Item, error) {
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
		}

		items = append(items, item)
	}

	return items, nil
}
