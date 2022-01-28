package entities

type Product struct {
	Id     int  `json:"id" form:"id"`
	UserId int  `json:"user_id" form:"user_id"`
	Price  int  `json:"price" form:"price"`
	User   User `json:"user" form:"user"`
}
