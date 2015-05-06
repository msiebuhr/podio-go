package podio

import "fmt"

type App struct {
	Id       float64 `json:"app_id"`
	Name     string  `json:"name"`
	ItemName string  `json:"item_name"`
	SpaceId  float64 `json:"space_id"`
}

func (client *Client) GetApps(space_id float64) (apps []App, err error) {
	path := fmt.Sprintf("/app/space/%.0f?view=micro", space_id)
	err = client.Request("GET", path, nil, nil, &apps)
	return
}

func (client *Client) GetApp(id float64) (app *App, err error) {
	path := fmt.Sprintf("/app/%.0f?view=micro", id)
	err = client.Request("GET", path, nil, nil, &app)
	return
}

func (client *Client) GetAppBySpaceIdAndSlug(space_id float64, slug string) (app *App, err error) {
	path := fmt.Sprintf("/app/space/%.0f/%s", space_id, slug)
	err = client.Request("GET", path, nil, nil, &app)
	return
}
