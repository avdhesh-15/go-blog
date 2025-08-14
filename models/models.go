package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // don't expose password
	Role     string `json:"role"`
}

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int    `json:"authorId"`
}
