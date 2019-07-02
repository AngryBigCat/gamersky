package models

type Content struct {
	Model
	Id      int
	NewsId  int
	Content string
}

func (content Content) TableName() string {
	return "contents"
}
