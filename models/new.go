package models

type New struct {
	Id          int
	Subject     string
	Title       string
	Href        string
	Image       string
	Description string
	PublishAt   int64
}

func (new New) TabeName() string {
	return "news"
}
