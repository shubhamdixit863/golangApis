package models

type contactinfo struct {
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Website string `json:"website"`
	City    string
	State   string
	Zip     string
}

type reviews struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Date    int32  `json:"date"`
	User    User   `json:"user"`
}

type Listings struct {
	Title       string
	Type        string
	Details     string      `json:"details"`
	ContactInfo contactinfo `json:"contactinfo"`
	Reviews     []reviews   `json:"reviews"`
	Products    []string    `json:"products"`
}
