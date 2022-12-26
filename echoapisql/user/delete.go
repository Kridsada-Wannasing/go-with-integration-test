package user

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1;")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare delete statement " + err.Error()})
	}

	if _, err := stmt.Exec(id); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't execute delete statement " + err.Error()})
	}

	return c.JSON(http.StatusNoContent, &User{})
}
