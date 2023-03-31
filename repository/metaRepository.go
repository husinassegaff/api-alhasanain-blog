package repository

import (
	"api-alhasanain-blog/database"
	"api-alhasanain-blog/structs"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

func CreateMeta(meta structs.Meta) (error, structs.Meta) {

	meta.ID = uuid.New().String()
	meta.CreatedAt = time.Now()

	s := "INSERT INTO public.meta (id, key, content, created_at) VALUES ($1, $2, $3, $4) RETURNING id,key,content"

	database.Init()

	metaRow := database.DB.QueryRow(s, meta.ID, meta.Key, meta.Content, meta.CreatedAt)

	err = metaRow.Scan(&meta.ID, &meta.Key, &meta.Content)

	if err != nil {
		return err, structs.Meta{}
	}

	return nil, meta
}

func UpdateMeta(meta structs.Meta) error {
	var setClause []string
	var values []interface{}
	var i = 1

	if meta.Key != "" {
		setClause = append(setClause, fmt.Sprintf("key=$%d", i))
		values = append(values, meta.Key)
		i++
	}

	if meta.Content != "" {
		setClause = append(setClause, fmt.Sprintf("content=$%d", i))
		values = append(values, meta.Content)
		i++
	}

	setClause = append(setClause, fmt.Sprintf("updated_at=$%d", i))
	values = append(values, time.Now())
	i++

	values = append(values, meta.ID)

	setStatement := strings.Join(setClause, ", ")
	if setStatement != "" {
		setStatement = " " + setStatement
	}

	s := fmt.Sprintf("UPDATE public.meta SET%s WHERE id=$%d", setStatement, i)

	database.Init()

	_, err := database.DB.Exec(s, values...)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMeta(id string) error {
	s := "DELETE FROM public.meta WHERE id=$1"

	database.Init()

	_, err := database.DB.Exec(s, id)

	if err != nil {
		return err
	}

	return nil
}

func GetAllMeta() (error, []structs.Meta) {
	s := "SELECT id, key, content, created_at, updated_at FROM public.meta"

	database.Init()

	rows, err := database.DB.Query(s)

	if err != nil {
		return err, []structs.Meta{}
	}

	defer rows.Close()

	var metas []structs.Meta

	for rows.Next() {
		var meta structs.Meta

		err = rows.Scan(&meta.ID, &meta.Key, &meta.Content, &meta.CreatedAt, &meta.UpdatedAt)

		if err != nil {
			return err, []structs.Meta{}
		}

		metas = append(metas, meta)
	}

	return nil, metas
}

func GetMetaById(id string) (error, structs.Meta) {
	s := "SELECT id, key, content, created_at, updated_at FROM public.meta WHERE id=$1"

	database.Init()

	row := database.DB.QueryRow(s, id)

	var meta structs.Meta

	err = row.Scan(&meta.ID, &meta.Key, &meta.Content, &meta.CreatedAt, &meta.UpdatedAt)

	if err != nil {
		return err, structs.Meta{}
	}

	return nil, meta
}
