package controllers

import (
	"classes/database/models"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
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

func (controller *Controller) UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	gender := c.FormValue("gender")

	query := `UPDATE students
	SET "first_name" = $1, "last_name" = $2, "gender" = $3 
	WHERE "id" = $4;`
	_, err := controller.DB.Exec(context.Background(), query, firstName, lastName, gender, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return c.String(http.StatusOK, "student updated!")
}

func (controller *Controller) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	query := `DELETE FROM students
	WHERE "id" = $1;`
	_, err := controller.DB.Exec(context.Background(), query, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return c.String(http.StatusOK, "student deleted!")
}

func (controller *Controller) GetAllStudents(c echo.Context) error {
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
			ID:        values[0].(int32),
			FirstName: values[1].(string),
			LastName:  values[2].(string),
			Gender:    values[3].(int32),
			Status:    values[4].(bool),
		})
	}

	return c.JSON(http.StatusOK, students)
}

