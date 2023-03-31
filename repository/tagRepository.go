package repository

import (
	"api-alhasanain-blog/database"
	"api-alhasanain-blog/structs"
	"database/sql"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

func CreateTag(tag structs.Tag) (error, structs.Tag) {

	tag.ID = uuid.New().String()
	tag.MetaTitle = strings.Replace(tag.Title, " ", "-", -1)
	tag.CreatedAt = time.Now()

	s := "INSERT INTO public.tag(id, id_post, title, meta_title, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id, id_post, title, meta_title"

	database.Init()

	row := database.DB.QueryRow(s, tag.ID, tag.IDPost, tag.Title, tag.MetaTitle, tag.CreatedAt)

	err = row.Scan(&tag.ID, &tag.IDPost, &tag.Title, &tag.MetaTitle)
	if err != nil {
		return err, structs.Tag{}
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, tag
}

func DeleteTag(id string) error {

	s := "DELETE FROM public.tag WHERE id = $1"

	database.Init()

	_, err = database.DB.Exec(s, id)
	if err != nil {
		return err
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil
}

func GetAllTag() ([]structs.Tag, error) {

	s := "SELECT id, id_post, title, meta_title, created_at FROM public.tag"

	database.Init()

	rows, err := database.DB.Query(s)
	if err != nil {
		return nil, err
	}

	var tags []structs.Tag

	for rows.Next() {
		var tag structs.Tag

		err = rows.Scan(&tag.ID, &tag.IDPost, &tag.Title, &tag.MetaTitle, &tag.CreatedAt)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return tags, nil
}
