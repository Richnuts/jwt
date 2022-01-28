package entities

type Book struct {
	Id     int    `json:"id" form:"id"`
	Tittle string `json:"tittle" form:"tittle"`
	Author string `json:"author" form:"author"`
	Status bool   `json:"status" form:"status"`
}
