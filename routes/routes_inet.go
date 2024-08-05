package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			// "john":    "doe",
			// "admin":   "123456",
			"gofiber": "21022566",
		},
	}))

	// /api/v1
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", c.HelloTest)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.ParamsTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidTest)

	v2 := api.Group("/v2")
	v2.Get("/", c.HelloTestV2)

	v3 := api.Group("/v3")
	v3.Get("/fac/:num", c.Factorial)
	v3.Get("/wut", c.Ascii)
	v3.Post("/register", c.Register)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	dog.Get("/history", c.GetDelectDogs)
	dog.Get("/between", c.GetBetweenDogs)
	//

	//CRUD company
	com := v1.Group("/com")
	com.Get("", c.GetComs)
	com.Get("/filter", c.GetCom)
	com.Post("/", c.AddCom)
	com.Put("/:id", c.UpdateCom)
	com.Delete("/:id", c.RemoveCom)
	com.Get("/history", c.GetDelectComs)
	com.Get("/between", c.GetBetweenComs)
	//

	//CRUD project
	pro := v1.Group("/pro")
	pro.Get("", c.GetUsers)
	pro.Get("/filter", c.GetUser)
	pro.Post("/", c.AddUser)
	pro.Put("/:id", c.UpdateUser)
	pro.Delete("/:id", c.RemoveUser)
	//
}
