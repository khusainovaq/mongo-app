package entity

type User struct {
	ID       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name     interface{} `json:"name,omitempty" bson:"name,omitempty"`
	Surname  interface{} `json:"surname,omitempty" bson:"surname,omitempty"`
	Username interface{} `json:"username,omitempty" bson:"username,omitempty"`
	Email    interface{} `json:"email,omitempty" bson:"email,omitempty"`
	Number   interface{} `json:"number,omitempty" bson:"number,omitempty"`
}
