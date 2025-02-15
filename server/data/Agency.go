package data

type Agency struct {
	InternalId   int    `json:"-"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	ShortName    string `json:"shortName"`
	SortableName string `json:"sortableName"`
	Slug         string `json:"slug"`
}
