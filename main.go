package main

import (
	"fmt"
	"golang-bwa/handler"
	"golang-bwa/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

type User struct {
	Id       int
	Name     string
	Email    string
	Age      int
	IsActive bool
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
	fmt.Println("DB connected")

	var userRepository user.Repository = user.NewRepository(db)
	// fmt.Println(userRepository)
	// user, err := userRepository.Save(user.User{Name: "Nikko"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var userService user.Service = user.NewService(userRepository)

	// user, err := userService.RegisterUser(user.RegisterUserInput{Name: "Gooo", Email: "goleng2@gmail.com", Password: "12345678"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(user)

	var userHandler = handler.NewUserHandler(userService)

	router := gin.Default()

	apiV1 := router.Group("api/v1")
	apiV1.GET("users", userHandler.Index)
	apiV1.POST("users", userHandler.RegisterUser)

	router.Run()
}

func hello(c *gin.Context) {
	users := []User{
		User{
			27,
			"nikko",
			"nikko@gmail.com",
			10,
			true,
		},
		User{
			1,
			"oleg",
			"oleg@gmail.com",
			20,
			false,
		},
	}

	type album struct {
		ID     string  `json:"id"`
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	}
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	fmt.Println(users)
	fmt.Println(albums)
	c.IndentedJSON(http.StatusOK, users)
}

func users(c *gin.Context) {

	dsn := "root:@tcp(127.0.0.1:3306)/golang_bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var users []user.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
	// fmt.Println(data)
	// fmt.Println(len(users))
	// fmt.Fprint(c.Writer, users)
}
