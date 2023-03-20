package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
	User   *User  `json:"user"`
}

func (Todo) IsNode() {}

func (this Todo) GetID() string {
	return this.ID
}
