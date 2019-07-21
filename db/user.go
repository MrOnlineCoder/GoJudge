package db

import (
	"errors"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"-"`
	Access int `json:"access"`
}

func CreateUser(user User) bool {
	const createSql = `
		INSERT INTO 'users' (username, fullname, password, access) VALUES
		(
			$1,
			$2,
			$3,
			$4
		);
	`;

	_, err := Maindb.Exec(createSql, user.Username, user.Fullname, user.Password, user.Access);

	if err != nil {
		return false
	}

	return true
}

func FindAuthUser(username string, password string) (User, error) {
	const findUserSql = `
		SELECT * FROM "users" WHERE
		"username" = $1 AND
		"password" = $2
	`;

	row := Maindb.QueryRow(findUserSql, username, password);

	user := User{};

	err := row.Scan(&user.Id, &user.Username, &user.Fullname, &user.Password, &user.Access);

	if err != nil {
		return user, errors.New("Database query failed.");
 }

 return user, nil
}

func GetUser(id int) (User, error) {
	const getUserSql = `
		SELECT * FROM "users" WHERE "id" = $1
	`;

	row := Maindb.QueryRow(getUserSql, id);

	user := User{};

	err := row.Scan(&user.Id, &user.Username, &user.Fullname, &user.Password, &user.Access);

	if err != nil {
		return user, errors.New("Database query failed.");
 }

 return user, nil
}