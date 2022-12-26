package user

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func UpdateUsersHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	id := c.Param("id")

	stmt, err := db.Prepare("UPDATE users SET age=$2 WHERE id=$1 RETURNING *;")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare update user statement:" + err.Error()})
	}

	row := stmt.QueryRow(id, u.Age)
	err = row.Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "error execute update " + err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}
