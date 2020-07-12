package model

type Person struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
}

type Address struct {
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
}
