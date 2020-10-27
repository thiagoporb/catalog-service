package main

func init() {
}

// Order Model
type Order struct {
	ID        uint   `json:"id" gorm:"autoIncrement"`
	Username  string `json:"username"`
	ProductId string `json:"product_id"`
	Paid      bool   `json:"paid"`
}

func connect() {

}

func main() {

}
