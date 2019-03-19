package itemHandler

import (
	"encoding/json"
	"io"

	"github.com/tryTwo/responses"
)

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
