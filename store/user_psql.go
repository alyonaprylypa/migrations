package store

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"petbook/models"
)

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetUsers(db *sqlx.DB) ([]models.User, error) {
	rows, err := db.Query("select * from users")
	logErr(err)
	defer rows.Close()
	users := []models.User{}
	err = sqlx.StructScan(rows, &users)
	if err != nil {
		logErr(err)
		return users, fmt.Errorf("cannot scan users from db: %v", err)
	}
	return users, nil
}

func GetUser(db *sqlx.DB, user *models.User) error {
	err := db.QueryRowx("select * from users where email=$1", user.Email).StructScan(user)
	if err != nil {
		logErr(err)
		return fmt.Errorf("cannot scan user from db: %v", err)
	}
	return nil
}

func Register(db *sqlx.DB, user *models.User) error {
	_, err := db.Exec("insert into users (email,firstname, lastname, login ,password) values ($1,$2,$3, $4, $5)",
		user.Email, user.Firstname, user.Lastname, user.Login, user.Password)
	if err != nil {
		return fmt.Errorf("cannot affect rows in users in db: %v", err)
	}
	return nil
}

func ChangePassword(db *sqlx.DB, user *models.User, newPassword string) error {
	res, err := db.Exec("UPDATE users SET password=$1 WHERE email = $2",
		newPassword, user.Email)
	if err != nil {
		return fmt.Errorf("cannot update users in db: %v", err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot affect rows in users in db: %v", err)
	}
	if num != 1 {
		return fmt.Errorf("cannot find this user")
	}
	return nil
}

// func Register(db *sqlx.DB, user models.User) int {
// 	err := db.QueryRow("insert into teammate (firstname,lastname,password) values ($1,$2,$3) RETURNING id_user;",
// 		user.Firstname, user.Lastname, user.Password).Scan(&user.ID)
// 	logErr(err)
// 	models.AddUser(&user)
// 	return user.ID
// }

func Login(db *sqlx.DB, userСhecking *models.User, userFromBase *models.User) error {
	err := db.QueryRow("select password from users where email=$1", userСhecking.Email).Scan(userFromBase.Password)
	if err == sql.ErrNoRows {
		return fmt.Errorf("cannot find such user: %v", err)
	}
	return nil
}
