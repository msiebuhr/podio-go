package podio

// ByLine describes the creator of a Podio object
type ByLine struct {
	Id         uint   `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	AvatarType string `json:"avatar_type"`
	AvatarId   uint   `json:"avatar_id"`
	Image      File   `json:"image"`
	LastSeenOn Time   `json:"last_seen_on"`

	Avatar int `json:"avatar"` //deprecated
}

// Via describes the source of a Podio object
type Via struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Display bool   `json:"display"`
}
