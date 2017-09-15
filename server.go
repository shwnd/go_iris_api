package main

import (
	"github.com/kataras/iris"
	"os"
	"fmt"
)

//Constants
var api_root string

func generateApiUrl(route string) (string) {
	return api_root + route
}

type Category struct {
	Id int
	Name string
}

type Article struct {
	Id int
	*Category
	Tags []string
	Title string
	Body string
	Published bool
}

func main() {


	if len(os.Args) != 3 {
		fmt.Println("Usage: server <port> <api root>")
		os.Exit(1)
	}

	port := os.Args[1]
	api_root = os.Args[2]
	app := iris.New()
	
    // Load all templates from the "./views" folder
    // where extension is ".html" and parse them
    // using the standard `html/template` package.
    //app.RegisterView(iris.HTML("./views", ".html"))

    // Method:    GET
    // Resource:  http://localhost:8080
    app.Get(generateApiUrl("/"), func(ctx iris.Context) {
        // Bind: {{.message}} with "Hello world!"
        //ctx.ViewData("message", "Hello world!")
        // Render template file: ./views/hello.html
		//ctx.View("hello.html")
		ctas := []Category{}
		ctas = append(ctas, Category{1, "Category 1"})
		ctas = append(ctas, Category{2, "Category 2"})
		ctas = append(ctas, Category{1, "Category 3"})
		ctas = append(ctas, Category{1, "Category 4"})
		ctas = append(ctas, Category{1, "Category 5"})
		
		ctx.JSON(ctas)
    })

    // Method:    GET
    // Resource:  http://localhost:8080/user/42
    app.Get("/user/{id:long}", func(ctx iris.Context) {
        userID, _ := ctx.Params().GetInt64("id")
        ctx.Writef("User ID: %d", userID)
    })

    // Start the server using a network address.
    app.Run(iris.Addr(":" + port))
}