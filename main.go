package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name  string `json:"name" form:"name" `
	Age   int    `json:"age" form:"age"`
	Hobby string `json:"hobby" form:"hobby"`
}

var (
	students map[string]Student
)

func main() {
	students = make(map[string]Student)

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("test/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	//get
	r.GET("/ping", func(c *gin.Context) {

		// s := Student{
		// 	"zzy",
		// 	23,
		// 	"nnjing",
		// }

		var s []Student
		for _, value := range students {
			s = append(s, value)
		}

		c.JSON(200, s)
	})

	//create
	r.POST("/ping", func(c *gin.Context) {

		var s Student

		if err := c.ShouldBind(&s); err != nil { // 不管是form,queryString,还是json，都自动判断接收
			fmt.Println(err.Error())
		} else {
			// fmt.Println(s)
			students[s.Name] = s
			// students = append(students, s)
			fmt.Println(students)

		}
	})

	//update
	r.PUT("/ping", func(c *gin.Context) {})

	//delete
	r.DELETE("/ping", func(c *gin.Context) {

		var s Student

		if err := c.ShouldBind(&s); err != nil { // 不管是form,queryString,还是json，都自动判断接收
			fmt.Println(err.Error())
		} else {
			delete(students, s.Name)
		}

	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("hello")
}
