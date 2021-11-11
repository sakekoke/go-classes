package controllers

import (
	"classes/database/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"log"
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

func (controller *Controller) GetStudents(c echo.Context) error {
	query := `SELECT * FROM students;`
	rows, err := controller.DB.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	var students []models.Student

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, models.Student{
			ID:        values[1].(int32),
			FirstName: values[1].(string),
			LastName:  values[2].(string),
			Gender:    values[3].(int32),
			Status:    values[4].(bool),
		})
	}

	return c.JSON(http.StatusOK, students)
}
