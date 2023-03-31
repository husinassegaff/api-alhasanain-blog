package repository

import (
	"api-alhasanain-blog/database"
	"api-alhasanain-blog/structs"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"log"
	"strings"
	"time"
)

func CreatePost(post structs.Post) (error, structs.Post) {

	post.ID = uuid.New().String()
	metaTitle := strings.Replace(post.Title, " ", "-", -1)
	post.Slug = slug.Make(post.Title)
	post.CreatedAt = time.Now()

	var publishedAt *time.Time // menginisialisasi variabel publishedAt dengan nilai nil

	if post.Status == "publish" {
		publishedTime := time.Now()  // membuat variabel dengan nilai waktu saat ini
		publishedAt = &publishedTime // mengambil alamat variabel tersebut dan menugaskannya ke variabel publishedAt
	}

	s := "INSERT INTO public.post (id, id_user, title, meta_title, slug, content, summary, status, published_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, id_user, title, meta_title, slug, content, summary, status, published_at, created_at"

	database.Init()

	postRow := database.DB.QueryRow(s, post.ID, post.IDUser, post.Title, metaTitle, post.Slug, post.Content, post.Summary, post.Status, publishedAt, post.CreatedAt)

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

func UpdatePost(post structs.Post) error {
	var setClause []string
	var values []interface{}
	var i = 1

	if post.Title != "" {
		setClause = append(setClause, fmt.Sprintf("title=$%d", i))
		values = append(values, post.Title)
		i++

		setClause = append(setClause, fmt.Sprintf("meta_title=$%d", i))
		values = append(values, strings.Replace(post.Title, " ", "-", -1))
		i++

		setClause = append(setClause, fmt.Sprintf("slug=$%d", i))
		values = append(values, slug.Make(post.Title))
		i++
	}

	if post.Content != "" {
		setClause = append(setClause, fmt.Sprintf("content=$%d", i))
		values = append(values, post.Content)
		i++
	}

	if post.Summary != nil {
		setClause = append(setClause, fmt.Sprintf("summary=$%d", i))
		values = append(values, post.Summary)
		i++
	}

	setClause = append(setClause, fmt.Sprintf("updated_at=$%d", i))
	values = append(values, time.Now())
	i++

	values = append(values, post.ID)

	setStatement := strings.Join(setClause, ", ")
	if setStatement != "" {
		setStatement = " " + setStatement
	}

	query := fmt.Sprintf("UPDATE public.post SET%s WHERE id=$%d", setStatement, i)

	database.Init()

	_, err := database.DB.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func CheckPostById(id string) bool {
	var post structs.Post

	s := "SELECT * FROM public.post WHERE id = $1"

	database.Init()

	row := database.DB.QueryRow(s, id)

	err = row.Scan(&post.ID, &post.IDUser, &post.Title, &post.MetaTitle, &post.Slug, &post.Content, &post.Summary, &post.Status, &post.CreatedAt, &post.UpdatedAt, &post.PublishedAt)

	if err != nil {
		return false
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return true
}
