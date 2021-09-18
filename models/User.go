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

func StoreUser(username, email, roles string) (helpers.Response, error) {
	var res helpers.Response

	//newRoles, err := strconv.Atoi(roles)
	//if err != nil {
	//	return res, err
	//}

	connection := database.CreateConnection()

	sqlStatement := "INSERT users (username, email, roles) VALUES (?,?,?)"

	stmt, err := connection.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, roles)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_insert_id": lastInsertId,
	}

	return res, nil
}
