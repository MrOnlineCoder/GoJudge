package db

import (
	"errors"
	"log"
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
		INSERT INTO "users" (username, fullname, password, access) VALUES
		(
			$1,
			$2,
			$3,
			$4
		);
	`;

	_, err := Maindb.Exec(createSql, user.Username, user.Fullname, user.Password, user.Access);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to create user: %s\n", err.Error());
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
		log.Printf("[DB] ERROR: Failed to find user for auth: %s\n", err.Error());
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
		log.Printf("[DB] ERROR: Failed to get user #%d: %s\n", id, err.Error());
		return user, errors.New("Database query failed.");
 }

 return user, nil
}

func GetAllUsers() ([]User, error) {
	const getAllSql = `
		SELECT "id", "username", "fullname", "access" FROM "users"
	`;

	list := []User{}

	rows, err := Maindb.Query(getAllSql);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to fetch users: %s\n", err.Error());
		return list, errors.New("Database query failed.");
	}

	for rows.Next() {
    u := User{}
    err := rows.Scan(&u.Id, &u.Username, &u.Fullname, &u.Access)
    if err != nil {
      continue
    }
    list = append(list, u)
  }

	rows.Close()

	return list, nil
}

func UpdateUserPassword(user_id int, pass string) bool {
	const updateSql = `
		UPDATE "users" SET "password" = $1 WHERE "id" = $2
	`;

	_, err := Maindb.Exec(updateSql, pass, user_id);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to update user #%d password: %s\n", user_id, err.Error());
		return false
	}

	return true
}

func DeleteUser(user_id int) bool {
	const deleteSql = `
		DELETE FROM "users" WHERE "id" = $1
	`;

	_, err := Maindb.Exec(deleteSql, user_id);

	if err != nil {
		log.Printf("[DB] ERROR: Failed to delete user #%d: %s\n", user_id, err.Error());
		return false
	}

	return true
}