package models

type CreateBook struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	CreatedAt string `json:"created_at"`
}

type GetBook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	CreatedAt string `json:"created_at"`
}

type UpdateBook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
