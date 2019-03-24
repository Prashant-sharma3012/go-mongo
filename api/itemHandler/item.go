package itemHandler

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/tryTwo/responses"
)

type Item struct {
	ItemId          string    `json:"itemId" bson:"itemId"`
	ItemName        string    `json:"itemName" bson:"itemName"`
	CategoryName    string    `json:"categoryName" bson:"categoryName"`
	SubCategoryName string    `json:"subCategoryName" bson:"subCategoryName"`
	ItemDescription string    `json:"itemDescription" bson:"itemDescription"`
	CreatedBy       string    `json:"createdBy" bson:"createdBy"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"updatedAt"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
}

func (i *Item) PreSave() {
	i.UpdatedAt = time.Now()
	i.CreatedAt = time.Now()
}

func (i *Item) PreUpdate() {
	i.UpdatedAt = time.Now()
}

func ItemFromJson(data io.Reader) *Item {
	var item *Item
	json.NewDecoder(data).Decode(&item)
	return item
}

func ItemToBson(i Item) ([]byte, error) {
	return bson.Marshal(i)
}

func NewItem() *Item {
	return &Item{}
}

func (i *Item) Validate() error {
	return responses.ClientError("Bad Item object")
}

func (i *Item) Save() ([]byte, error) {
	i.PreSave()
	return Save(*i)
}

func (i *Item) Update() ([]byte, error) {
	i.PreUpdate()
	return Update(*i)
}

func (i *Item) Delete() ([]byte, error) {
	// No error handling, becasue too lazy to put one
	return Delete(*i)
}

func ItemList(skip int64, limit int64) ([]Item, error) {
	return List(skip, limit)
}
