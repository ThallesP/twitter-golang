package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gosimple/slug"
	"github.com/joho/godotenv"
	"github.com/thallesp/twitter-golang/controller"
	"github.com/thallesp/twitter-golang/entity"
	"github.com/thallesp/twitter-golang/repository"
	"github.com/thallesp/twitter-golang/usecase"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Default().Println("Could not load .env file")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Internal Server Error"
			errname := "internalServerError"

			var e *usecase.Exception

			if errors.As(err, &e) {
				code = e.StatusCode
				message = e.Message
				errname = e.ErrorName
			}

			var fe *fiber.Error

			if errors.As(err, &fe) {
				code = fe.Code
				message = fe.Message
				errname = slug.Make(fe.Message)
			}

			if errname == "internalServerError" {
				fmt.Println(err)
			}

			c.Status(code).JSON(fiber.Map{
				"error":   errname,
				"message": message,
			})

			return nil
		},
	})

	SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed to listen on specified port. err: %s", err.Error())
	}
}

func SetupRoutes(app *fiber.App) {
	pg := SetupDatabase()
	usersRepository := repository.NewPostgresUsersRepository(pg)

	createTweetController := controller.NewCreateTweetController(usecase.NewCreateTweetUseCase(repository.NewPostgresTweetsRepository(pg)))
	getTweetsController := controller.NewGetTweetsController(usecase.NewGetTweetsUseCase(repository.NewPostgresTweetsRepository(pg)))
	registerController := controller.NewRegisterController(usecase.NewRegisterUseCase(usersRepository))
	loginController := controller.NewLoginController(usecase.NewLoginUseCase(usersRepository))
	app.Post("/register", registerController.Handle)

	app.Post("/login", loginController.Handle)
	app.Get("/health", func(c *fiber.Ctx) error {
		err := pg.Ping(context.Background())

		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendStatus(200)
	})
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Post("/tweets", createTweetController.Handle)
	app.Get("/tweets", getTweetsController.Handle)
}

func SetupDatabase() *pg.DB {
	options, err := pg.ParseURL(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("failed to connect to database. Err: %s", err.Error())
	}

	time.Sleep(100 * time.Millisecond)
	db := pg.Connect(options)
	err = createSchema(db)

	if err != nil {
		log.Printf("failed to create schema. Not exiting. Err: %s", err.Error())
	}

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*entity.Tweet)(nil),
		(*entity.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
