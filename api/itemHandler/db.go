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

type toInsert struct {
	_id primitive.ObjectID
	doc interface{}
}

func setParams() {
	ctx = context.Background()
	itemCollection = db.GetConnection().Collection("item")
}
func Save(i Item) ([]byte, error) {
	setParams()

	res, err := itemCollection.InsertOne(ctx, i)
	if err != nil {
		return nil, err
	}

	return []byte(res.InsertedID.(primitive.ObjectID).Hex()), nil
}

func Update(i Item) ([]byte, error) {
	setParams()
	_, err := itemCollection.UpdateOne(ctx, bson.D{{"itemId", i.ItemId}}, bson.D{{"$set", i}})

	if err != nil {
		return nil, err
	}
	return []byte(i.ItemId + "updated successfully"), nil
}

func Delete(i Item) ([]byte, error) {
	setParams()

	res, err := itemCollection.DeleteOne(ctx, bson.D{{"itemId", i.ItemId}})
	if err != nil {
		return nil, err
	}

	return []byte(strconv.Itoa(int(res.DeletedCount))), nil
}

func List(skip, limit int64) ([]Item, error) {
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
			ItemId:          m["itemId"].(string),
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
