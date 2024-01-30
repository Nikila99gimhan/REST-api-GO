package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "restapi/docs" // This line is necessary for go-swagger to find your docs!
)

// User represents a user of the API.
type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// users slice to seed user data.
var users = []User{
    {ID: "1", Name: "John Doe"},
    {ID: "2", Name: "Jane Doe"},
}

// @Summary List users
// @Description get users
// @ID get-users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

// @Summary Add a user
// @Description add by json user
// @ID add-user
// @Accept json
// @Produce json
// @Param user body User true "Add user"
// @Success 200 {object} User
// @Router /users [post]
func postUser(c *gin.Context) {
    var newUser User
    if err := c.BindJSON(&newUser); err != nil {
        return
    }
    users = append(users, newUser)
    c.JSON(http.StatusOK, newUser)
}

// @title My API
// @description This is a sample server.
// @version 1.0
// @host localhost:8000
// @BasePath /
func main() {
    router := gin.Default()

    // Use the swaggo middleware to serve the Swagger UI
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Setup the endpoints
    router.GET("/users", getUsers)
    router.POST("/users", postUser)

    // Start the server
    router.Run(":8000")
}
