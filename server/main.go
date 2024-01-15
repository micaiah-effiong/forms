package main

import (
	"api/form-api/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAnAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAnAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {

			println("here", id)

			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/albums", getAllAlbums)
	router.POST("/albums", addAnAlbum)
	router.GET("/albums/:id", getAnAlbum)

	// forms
	router.Group("/forms")
	router.GET("/forms", route.GetAllForms)
	router.POST("/forms", route.CreateForm)
	router.GET("/forms/:id", route.GetForm)
	router.DELETE("/forms/:id", route.DeleteForm)

	router.Run("localhost:8000")

}
