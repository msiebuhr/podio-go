package podio

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Id                 float64  `json:"item_id"`
	App                *App     `json:"app"`
	AppItemId          float64  `json:"app_item_id"`
	FormattedAppItemId string   `json:"app_item_id_formatted"`
	Link               string   `json:"link"`
	Title              string   `json:"title"`
	Files              []*File  `json:"files"`
	Fields             []*Field `json:"fields"`
}

type Field struct {
	FieldId    float64  `json:"field_id"`
	ExternalId string   `json:"external_id"`
	Type       string   `json:"type"`
	Label      string   `json:"label"`
	Values     []*Value `json:"values"`
}

func (f *Field) UnmarshalJSON(data []byte) error {
	var fake struct {
		FieldId    float64  `json:"field_id"`
		ExternalId string   `json:"external_id"`
		Type       string   `json:"type"`
		Label      string   `json:"label"`
		Values     []*Value `json:"values"`
	}

	// Unmarshal regularly
	err := json.Unmarshal(data, &fake)
	if err != nil {
		return err
	}

	f.FieldId = fake.FieldId
	f.ExternalId = fake.ExternalId
	f.Type = fake.Type
	f.Label = fake.Label
	f.Values = fake.Values

	// Does it have a 'type'-key, so we can deduce anything about the values?
	switch f.Type {
	case "app":
		// Hack: Re-encode to JSON and decode as an App
		for i, value := range f.Values {
			reencodedValue, err := json.Marshal(value.Value)
			if err != nil {
				return err
			}
			var newItem Item
			err = json.Unmarshal(reencodedValue, &newItem)
			if err != nil {
				return err
			}
			f.Values[i].Value = newItem
		}
	default:
	}

	return nil
}

type Value struct {
	Value interface{} `json:"value"`
}

type ItemList struct {
	Filtered float64 `json:"filtered"`
	Total    float64 `json:"total"`
	Items    []*Item `json:"items"`
}

func (client *Client) GetItems(app_id float64) (items *ItemList, err error) {
	path := fmt.Sprintf("/item/app/%.0f/filter?fields=items.fields(files)", app_id)
	err = client.Request("POST", path, nil, nil, &items)
	return
}

func (client *Client) GetItemByAppItemId(app_id float64, formatted_app_item_id string) (item *Item, err error) {
	path := fmt.Sprintf("/app/%.0f/item/%s", app_id, formatted_app_item_id)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

func (client *Client) GetItemByExternalID(app_id float64, external_id string) (item *Item, err error) {
	path := fmt.Sprintf("/item/app/%.0f/external_id/%s", app_id, external_id)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

func (client *Client) GetItem(item_id float64) (item *Item, err error) {
	path := fmt.Sprintf("/item/%.0f?fields=files", item_id)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

func (client *Client) CreateItem(app_id float64, external_id string, fieldValues map[string]interface{}) (float64, error) {
	path := fmt.Sprintf("/item/app/%.0f", app_id)
	params := map[string]interface{}{
		"fields": fieldValues,
	}

	if external_id != "" {
		params["external_id"] = external_id
	}

	rsp := &struct {
		ItemId float64 `json:"item_id"`
	}{}
	err := client.RequestWithParams("POST", path, nil, params, rsp)

	return rsp.ItemId, err
}

func (client *Client) UpdateItem(itemId float64, fieldValues map[string]interface{}) error {
	path := fmt.Sprintf("/item/%.0f", itemId)
	params := map[string]interface{}{
		"fields": fieldValues,
	}

	return client.RequestWithParams("PUT", path, nil, params, nil)
}
