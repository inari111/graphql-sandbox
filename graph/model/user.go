package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (User) IsNode() {}

func (this User) GetID() string {
	return this.ID
}
