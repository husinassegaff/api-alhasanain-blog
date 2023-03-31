package repository

import (
	"api-alhasanain-blog/database"
	"api-alhasanain-blog/structs"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

var (
	err error
)

func RegisterUser(user structs.User) (error, structs.User) {

	// generate id
	user.ID = uuid.New().String()

	// hash password
	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	// generate token
	user.Token = ""

	// if role not assigned, set role to user
	if user.Role == "" {
		user.Role = "user"
	}

	// set created at
	user.CreatedAt = time.Now()

	database.Init()

	// Check if email already exists
	var count int
	row := database.DB.QueryRow("SELECT COUNT(*) FROM public.user WHERE email = $1", user.Email)
	err = row.Scan(&count)
	if err != nil {
		return err, structs.User{}
	}

	if count > 0 {
		return errors.New("email already exists"), structs.User{}
	}

	s := "INSERT INTO public.user (id, name, email, password, role, created_at, token) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, email, role, token"

	userRow := database.DB.QueryRow(s, user.ID, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.Token)

	err = userRow.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.Token)

	if err != nil {
		return err, structs.User{}
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return err, user
}

func GetAllUser() (err error, results []structs.User) {

	s := "SELECT * FROM public.user"

	database.Init()

	rows, _ := database.DB.Query(s)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var user structs.User
		var updatedAt sql.NullTime
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &updatedAt, &user.Role, &user.Password, &user.Token)
		if err != nil {
			return err, nil
		}

		results = append(results, user)
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, results
}

func GetUserById(id string) (err error, user structs.User) {

	s := "SELECT * FROM public.user WHERE id = $1"

	database.Init()

	row := database.DB.QueryRow(s, id)

	var updatedAt sql.NullTime
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &updatedAt, &user.Role, &user.Password, &user.Token)
	if err != nil {
		return errors.New("user not found"), structs.User{}
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, user
}

func LoginUser(email string, password string) (err error, user structs.User) {

	// check if user already login
	s := "SELECT token FROM public.user WHERE email = $1"

	database.Init()

	row := database.DB.QueryRow(s, email)

	var token string
	err = row.Scan(&token)

	if err != nil {
		return errors.New("email or password is wrong"), structs.User{}
	}

	if token != "" {
		return errors.New("user already login"), structs.User{}
	}

	// hash password
	hash := md5.Sum([]byte(password))
	password = hex.EncodeToString(hash[:])

	s = "SELECT * FROM public.user WHERE email = $1 AND password = $2"

	row = database.DB.QueryRow(s, email, password)

	var updatedAt sql.NullTime
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &updatedAt, &user.Role, &user.Password, &user.Token)
	if err != nil {
		return errors.New("email or password is wrong"), structs.User{}
	}

	//generate token
	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		fmt.Println("Error Token: ", err)
	}
	tokenString := base64.StdEncoding.EncodeToString(tokenBytes)
	err, user.Token = SetToken(email, tokenString)

	if err != nil {
		return errors.New("error set token"), structs.User{}
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(database.DB)

	return nil, user
}
