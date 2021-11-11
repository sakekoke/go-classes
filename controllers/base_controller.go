package controllers

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type (
	Controller struct {
		DB *pgxpool.Pool
	}
)

func (controller *Controller) CreateStudent(c echo.Context) error {
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	gender := c.FormValue("gender")

	query := `INSERT INTO students
	("first_name", "last_name", "gender") 
	VALUES ($1, $2, $3);`
	_, err := controller.DB.Exec(context.Background(), query, firstName, lastName, gender)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return c.String(http.StatusOK, "student created!")
}




