package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

// IndexHandler .
func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

// AHandler .
func AHandler(c *gin.Context) {
	c.XML(200, Person{
		FirstName: "abc",
		LastName:  "abc123"})
}

func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.GET("/", AHandler)
	router.Run()
}
