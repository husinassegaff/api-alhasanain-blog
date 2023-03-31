package repository

import (
	"api-alhasanain-blog/database"
	"api-alhasanain-blog/structs"
	"github.com/google/uuid"
	"time"
)

func CreateComment(comment structs.Comment) (error, structs.Comment) {

	comment.ID = uuid.New().String()
	comment.CreatedAt = time.Now()

	s := "INSERT INTO public.comment (id, id_user, id_post, content, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, id_user, id_post, content"

	database.Init()

	commentRow := database.DB.QueryRow(s, comment.ID, comment.IDUser, comment.IDPost, comment.Content, comment.CreatedAt)

	err = commentRow.Scan(&comment.ID, &comment.IDUser, &comment.IDPost, &comment.Content)

	if err != nil {
		return err, structs.Comment{}
	}

	return nil, comment
}

func DeleteComment(id string) error {

	s := "DELETE FROM public.comment WHERE id = $1"

	database.Init()

	_, err = database.DB.Exec(s, id)

	if err != nil {
		return err
	}

	return nil
}

func GetAllComment() (error, []structs.Comment) {

	s := "SELECT id, id_user, id_post, content, created_at FROM public.comment"

	database.Init()

	rows, err := database.DB.Query(s)

	if err != nil {
		return err, []structs.Comment{}
	}

	defer rows.Close()

	var comments []structs.Comment

	for rows.Next() {
		var comment structs.Comment

		err = rows.Scan(&comment.ID, &comment.IDUser, &comment.IDPost, &comment.Content, &comment.CreatedAt)

		if err != nil {
			return err, []structs.Comment{}
		}

		comments = append(comments, comment)
	}

	return nil, comments
}
