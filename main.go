package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"four-layer-todo-app/controller"
	"four-layer-todo-app/repository"
	"four-layer-todo-app/service"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://LocalDeveloper:f8104adbecc0ec@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.Background())

	db := client.Database("db_todos")

	// Connect to MySQL
	connectionString := "root:emx2bPxI52jUe133GWlkHoyESLzQ@tcp(localhost:3307)/parcialGo?parseTime=true"
	mySqlDB, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer func(oracleDB *sqlx.DB) {
		err := oracleDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(mySqlDB)

	err = mySqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize layers(todo_layer)
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	// (doctor_layer)
	doctorRepo := repository.NewMySqlDentistRepository(mySqlDB)
	doctorService := service.NewDentistService(doctorRepo)
	doctorController := controller.NewDentistController(doctorService)

	// Setup Fiber
	app := fiber.New()

	// Routes
	app.Post("/todos", todoController.CreateTodo)
	app.Get("/todos", todoController.GetTodos)
	app.Post("/dentists", doctorController.CreateDentist)
	app.Get("/dentists", doctorController.GetAllDentists)

	log.Println(app.Listen(":3000"))
}
