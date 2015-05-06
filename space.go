package podio

import "fmt"

type Space struct {
	Id   float64 `json:"space_id"`
	Slug string  `json:"url_label"`
	Name string  `json:"name"`
}

func (client *Client) GetSpaces(org_id float64) (spaces []Space, err error) {
	path := fmt.Sprintf("/org/%.0f/space", org_id)
	err = client.Request("GET", path, nil, nil, &spaces)
	return
}

func (client *Client) GetSpace(id float64) (space *Space, err error) {
	path := fmt.Sprintf("/space/%.0f", id)
	err = client.Request("GET", path, nil, nil, &space)
	return
}

func (client *Client) GetSpaceByOrgIdAndSlug(org_id float64, slug string) (space *Space, err error) {
	path := fmt.Sprintf("/space/org/%.0f/%s", org_id, slug)
	err = client.Request("GET", path, nil, nil, &space)
	return
}
