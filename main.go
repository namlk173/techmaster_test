package main

import (
	"fmt"
	"irishttp/model"
	"log"

	"github.com/kataras/iris/v12"
)

var quizzes = []model.Quiz{}
var err error

func main() {
	app := iris.New()
	quizzes, err = model.GetQuiz("data/Cauhoimau2.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quizzes)

	app.RegisterView(iris.HTML("./view", ".html").Reload(true))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index")
	})

	app.Post("/submit", func(ctx iris.Context) {
		var quiz model.Quiz
		err := ctx.ReadJSON(&quiz)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			_ = ctx.JSON(err)
			return
		}
		quizzes = append(quizzes, quiz)
		_ = ctx.JSON("Thành công")
	})

	app.Get("/get", func(ctx iris.Context) {
		_ = ctx.JSON(quizzes)
	})

	app.Listen(":3000")
}
