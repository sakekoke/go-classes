package main

import (
        "classes/controllers"
        "context"
        "fmt"
        _ "github.com/jackc/pgx/v4"
        "github.com/jackc/pgx/v4/pgxpool"
        "github.com/labstack/echo/v4"
        "os"
        "classes/config"
)

func main() {
        e := echo.New()

        // DATABASE QUERIES
        dbPool, err := pgxpool.Connect(context.Background(), config.DatabaseUrl)
        if err != nil {
                fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
                os.Exit(1)
        }
        defer dbPool.Close()

        controller := controllers.Controller{
                DB: dbPool,
        }

        // TABLE CLASS
        students := e.Group("/students")
        students.GET("/getAll", controller.GetAllStudents)
        students.POST("/create", controller.CreateStudent)

        e.Logger.Fatal(e.Start(":1323"))
}

