package main

import (
	_taskHttpDelivery "ApiRest/handlers"
	_taskUsecase "ApiRest/logic"
	_ "ApiRest/repository"
	_taskRepo "ApiRest/repository"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	connStr := "user=postgres password=1234 dbname=postgres host=localhost port=5432 sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	tr := _taskRepo.NewPostgresTaskRepository(dbConn)
	tu := _taskUsecase.NewTaskUsecase(tr, timeoutContext)
	_taskHttpDelivery.NewTaskHandler(e, tu)

	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
