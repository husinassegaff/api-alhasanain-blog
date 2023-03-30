package structs

import "time"

type User struct {
	ID        string     `json:"id_user"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	Token     string     `json:"token"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    string `json:"id_user"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type Category struct {
	ID        string `json:"id_category"`
	Title     string `json:"title"`
	MetaTitle string `json:"meta_title"`
	Slug      string `json:"slug"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Comment struct {
	ID        string `json:"id_comment"`
	IDUser    string `json:"id_user"`
	IDPost    string `json:"id_post"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type Meta struct {
	ID        string `json:"id_meta"`
	Key       string `json:"key"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Post struct {
	ID          string `json:"id_post"`
	IDUser      string `json:"id_user"`
	Title       string `json:"title"`
	MetaTitle   string `json:"meta_title"`
	Slug        string `json:"slug"`
	Content     string `json:"content"`
	Summary     string `json:"summary"`
	Status      string `json:"status"`
	PublishedAt string `json:"published_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Tag struct {
	ID        string `json:"id_tag"`
	IDPost    string `json:"id_post"`
	Title     string `json:"title"`
	MetaTitle string `json:"meta_title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostCategory struct {
	IDPost     string `json:"id_post"`
	IDCategory string `json:"id_category"`
}

type PostMeta struct {
	IDPost string `json:"id_post"`
	IDMeta string `json:"id_meta"`
}
