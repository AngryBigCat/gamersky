package models

type News struct {
	Id          int
	Subject     string
	Title       string
	Href        string
	Image       string
	Description string
	PublishAt   int64
}

type Contents struct {
	Id      int
	NewsId  int
	Content string
}
