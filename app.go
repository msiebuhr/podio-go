package podio

import "fmt"

type App struct {
	Id       uint   `json:"app_id"`
	Name     string `json:"name"`
	ItemName string `json:"item_name"`
	SpaceId  uint   `json:"space_id"`
}

func (client *Client) GetApps(space_id uint) (apps []App, err error) {
	path := fmt.Sprintf("/app/space/%d?view=micro", space_id)
	err = client.Request("GET", path, nil, nil, &apps)
	return
}

func (client *Client) GetApp(id uint) (app *App, err error) {
	path := fmt.Sprintf("/app/%d?view=micro", id)
	err = client.Request("GET", path, nil, nil, &app)
	return
}

func (client *Client) GetAppBySpaceIdAndSlug(space_id uint, slug string) (app *App, err error) {
	path := fmt.Sprintf("/app/space/%d/%s", space_id, slug)
	err = client.Request("GET", path, nil, nil, &app)
	return
}
