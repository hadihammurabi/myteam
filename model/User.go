package model

import (
	. "fmt"
	"myteam/types"
)

func GetUsers() ([]types.User, error) {
	var users []types.User

	rows, err := Connection.Query("SELECT * FROM users")
	if err != nil {
		Println(err.Error())
	}
	
	defer rows.Close()
	
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, err
	}
	
	return users, nil
}

func GetUserByID(id string) (types.User, error) {
	var query string = Sprintf("SELECT * FROM users WHERE id=%s", id)
	row := Connection.QueryRow(query)

	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func StoreUser(user types.User) (types.User, error) {
	_, err := Connection.Query(Sprintf("INSERT INTO users VALUES(null, '%s', '%s')", user.Username, user.Password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id string, data types.User) (types.User, error) {
	olduser, err := GetUserByID(id)

	if len(data.Username) <= 0 {
		data.Username = olduser.Username
	}
	if len(data.Password) <= 0 {
		data.Password = olduser.Password
	}

	var query string = Sprintf("UPDATE users SET username='%s', password='%s' WHERE id=%s", data.Username, data.Password, id)
	_, err = Connection.Query(query)

	if err != nil {
		return data, err
	}

	newuser, err := GetUserByID(id)

	if err != nil {
		return data, err
	}

	return newuser, nil
}

func DeleteUser(id string) (types.User, error) {
	user, err := GetUserByID(id)
	if err != nil {
		return user, err
	}

	var query string = Sprintf("DELETE FROM users WHERE id=%s", id)
	_, err = Connection.Query(query)

	if err != nil {
		return user, err
	}

	return user, nil
}