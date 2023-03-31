package repository

import (
	"api-alhasanain-blog/database"
	"database/sql"
	"log"
)

func SetToken(email string, token string) (err error, t string) {

	s := "UPDATE public.user SET token = $1 WHERE email = $2"

	database.Init()

	_, err = database.DB.Exec(s, token, email)

	if err != nil {
		return err, ""
	}

	return nil, token
}

func GetToken(email string) (err error, token string) {

	s := "SELECT token FROM public.user WHERE email = $1"

	database.Init()

	row := database.DB.QueryRow(s, email)

	err = row.Scan(&token)
	if err != nil {
		return err, ""
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, token
}

func GetRoleWithToken(token string) (err error, role string) {

	s := "SELECT role FROM public.user WHERE token = $1"

	database.Init()

	row := database.DB.QueryRow(s, token)

	err = row.Scan(&role)
	if err != nil {
		return err, ""
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, role
}

func GetEmailWithToken(token string) (err error, email string) {

	s := "SELECT email FROM public.user WHERE token = $1"

	database.Init()

	row := database.DB.QueryRow(s, token)

	err = row.Scan(&email)
	if err != nil {
		return err, ""
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, email
}

func GetIdWithToken(token string) (err error, id string) {

	s := "SELECT id FROM public.user WHERE token = $1"

	database.Init()

	row := database.DB.QueryRow(s, token)

	err = row.Scan(&id)
	if err != nil {
		return err, ""
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, id
}
