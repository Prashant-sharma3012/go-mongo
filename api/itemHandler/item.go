package itemHandler

import (
	"encoding/json"
	"io"
	"time"

	"github.com/tryTwo/responses"
)

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

func ItemFromJson(data io.Reader) *Item {
	var item *Item
	json.NewDecoder(data).Decode(&item)
	return item
}

func NewItem() *Item {
	return &Item{}
}

func (i *Item) Validate() error {
	return responses.ClientError("Bad Item object")
}

func (i *Item) Save() ([]byte, error) {
	i.PreSave()
	return SaveItem(i)
}

func (i *Item) Update() ([]byte, error) {
	i.PreUpdate()
	return UpdateItem(i)
}

func (i *Item) Delete() ([]byte, error) {
	// No error handling, becasue too lazy to put one
	return DeleteItem(i)
}

func List(skip int64, limit int64) ([]Item, error) {
	return ItemList(skip, limit)
}
