package model

import (
	"myteam/types"
	. "fmt"
)

func GetUserAuth(username string, password string) (map[string]interface{}, error) {
	var query string = Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
	row := Connection.QueryRow(query)

	var user types.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return map[string]interface{}{
		"id": user.ID,
		"username": user.Username,
	}, nil
}