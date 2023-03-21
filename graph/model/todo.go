package model

const ObjName = "todo"

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user_id"`
}

func (Todo) IsNode() {}

func (this Todo) GetID() string {
	return this.ID
}
