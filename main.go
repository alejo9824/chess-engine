package main

import (
	"chess/internal/app"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//game := app.NewGame()
	//fmt.Println(game.ChessBoard)

	r := gin.Default()
	game := app.NewGame()
	// Serve static files (HTML, CSS, JS)
	r.Static("/static", "./static")

	// Route to render the chessboard page
	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	// API endpoint to get chessboard data
	r.GET("/api/chessboard", func(c *gin.Context) {

		c.JSON(http.StatusOK, game.ChessBoard)
	})

	// Start the server
	fmt.Println("Server starting on :8080")
	r.Run(":8080")

}
