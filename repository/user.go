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

func CreateUser(user structs.User) (error, structs.User) {

	// generate id
	user.ID = uuid.New().String()

	// hash password
	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	// generate token
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		fmt.Println("Error Token: ", err)
	}
	tokenString := base64.StdEncoding.EncodeToString(tokenBytes)
	user.Token = tokenString

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
		fmt.Println("Error scan: ", err)
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

	return nil, results
}
