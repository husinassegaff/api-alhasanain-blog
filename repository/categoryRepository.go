package repository

import (
	"api-alhasanain-blog/database"
	"api-alhasanain-blog/structs"
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"strings"
	"time"
)

func CreateCategory(category structs.Category) (error, structs.Category) {

	category.ID = uuid.New().String()
	category.MetaTitle = strings.Replace(category.Title, " ", "-", -1)
	category.Slug = slug.Make(category.Title)
	category.CreatedAt = time.Now()

	s := "INSERT INTO public.category (id, title, meta_title, slug, content, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, title, meta_title, slug, content"

	database.Init()

	categoryRow := database.DB.QueryRow(s, category.ID, category.Title, category.MetaTitle, category.Slug, category.Content, category.CreatedAt)

	err = categoryRow.Scan(&category.ID, &category.Title, &category.MetaTitle, &category.Slug, &category.Content)

	if err != nil {
		return err, structs.Category{}
	}

	return nil, category
}

func UpdateCategory(category structs.Category) error {
	var setClause []string
	var values []interface{}
	var i = 1

	if category.Title != "" {
		setClause = append(setClause, fmt.Sprintf("title=$%d", i))
		values = append(values, category.Title)
		i++

		setClause = append(setClause, fmt.Sprintf("meta_title=$%d", i))
		values = append(values, strings.Replace(category.Title, " ", "-", -1))
		i++

		setClause = append(setClause, fmt.Sprintf("slug=$%d", i))
		values = append(values, slug.Make(category.Title))
		i++
	}

	if category.Content != "" {
		setClause = append(setClause, fmt.Sprintf("content=$%d", i))
		values = append(values, category.Content)
		i++
	}

	setClause = append(setClause, fmt.Sprintf("updated_at=$%d", i))
	values = append(values, time.Now())
	i++

	values = append(values, category.ID)

	setStatement := strings.Join(setClause, ", ")
	if setStatement != "" {
		setStatement = " " + setStatement
	}

	s := fmt.Sprintf("UPDATE public.category SET%s WHERE id=$%d", setStatement, i)

	database.Init()

	_, err = database.DB.Exec(s, values...)

	if err != nil {
		return err
	}

	return nil
}

func CheckCategoryById(id string) error {
	s := "SELECT id FROM public.category WHERE id=$1"

	database.Init()

	var category structs.Category

	err = database.DB.QueryRow(s, id).Scan(&category.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id string) error {
	s := "DELETE FROM public.category WHERE id=$1"

	database.Init()

	_, err = database.DB.Exec(s, id)

	if err != nil {
		return err
	}

	return nil
}

func GetAllCategory() (error, []structs.Category) {
	s := "SELECT id, title, meta_title, slug, content FROM public.category"

	database.Init()

	rows, err := database.DB.Query(s)

	if err != nil {
		return err, []structs.Category{}
	}

	defer rows.Close()

	var categories []structs.Category

	for rows.Next() {
		var category structs.Category

		err = rows.Scan(&category.ID, &category.Title, &category.MetaTitle, &category.Slug, &category.Content)

		if err != nil {
			return err, []structs.Category{}
		}

		categories = append(categories, category)
	}

	return nil, categories
}

func GetCategoryById(id string) (error, structs.Category) {
	s := "SELECT id, title, meta_title, slug, content FROM public.category WHERE id=$1"

	database.Init()

	var category structs.Category

	err = database.DB.QueryRow(s, id).Scan(&category.ID, &category.Title, &category.MetaTitle, &category.Slug, &category.Content)

	if err != nil {
		return err, structs.Category{}
	}

	return nil, category
}
