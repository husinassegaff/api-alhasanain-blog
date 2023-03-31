package repository

import (
	"api-alhasanain-blog/database"
	"api-alhasanain-blog/structs"
	"database/sql"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"log"
	"time"
)

func CreatePost(post structs.Post) (error, structs.Post) {

	post.ID = uuid.New().String()
	post.Slug = slug.Make(post.Title)
	post.CreatedAt = time.Now()

	var publishedAt *time.Time // menginisialisasi variabel publishedAt dengan nilai nil

	if post.Status == "publish" {
		publishedTime := time.Now()  // membuat variabel dengan nilai waktu saat ini
		publishedAt = &publishedTime // mengambil alamat variabel tersebut dan menugaskannya ke variabel publishedAt
	}

	s := "INSERT INTO public.post (id, id_user, title, meta_title, slug, content, summary, status, published_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, id_user, title, meta_title, slug, content, summary, status, published_at, created_at"

	database.Init()

	postRow := database.DB.QueryRow(s, post.ID, post.IDUser, post.Title, post.MetaTitle, post.Slug, post.Content, post.Summary, post.Status, publishedAt, post.CreatedAt)

	err = postRow.Scan(&post.ID, &post.IDUser, &post.Title, &post.MetaTitle, &post.Slug, &post.Content, &post.Summary, &post.Status, &post.PublishedAt, &post.CreatedAt)

	if err != nil {
		return err, structs.Post{}
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, post
}
