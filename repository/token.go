package repository

import "api-alhasanain-blog/database"

func SetToken(email string, token string) (err error, t string) {

	s := "UPDATE public.user SET token = NULL WHERE email = $1"

	database.Init()

	_, err = database.DB.Exec(s, email)

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

	return nil, token
}
