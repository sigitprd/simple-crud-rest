package models

import (
	"github.com/sigitprd/simple-crud-rest/database"
	"github.com/sigitprd/simple-crud-rest/helpers"
	"net/http"
	"strings"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Roles    int    `json:"roles"`
}

func FetchUser(id string) (helpers.Response, error) {
	var obj User
	var arrObj []User
	var res helpers.Response

	if id != "" {
		id = strings.Trim(id, "/")
	}

	connection := database.CreateConnection()

	if id != "" {
		sqlStatement := "SELECT * FROM users WHERE id=?"

		rows, err := connection.Query(sqlStatement, id)
		defer rows.Close()
		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id, &obj.Username, &obj.Email, &obj.Roles)
			if err != nil {
				return res, err
			}
			arrObj = append(arrObj, obj)
		}

		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = arrObj

		return res, nil
	} else {
		sqlStatement := "SELECT * FROM users"

		rows, err := connection.Query(sqlStatement)
		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id, &obj.Username, &obj.Email, &obj.Roles)
			if err != nil {
				return res, err
			}

			arrObj = append(arrObj, obj)
		}

		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = arrObj

		return res, nil
	}
}
