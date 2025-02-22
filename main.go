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

	// API endpoint to move a piece
	r.POST("/api/move", func(c *gin.Context) {
		var move struct {
			X1 int `json:"x1"`
			Y1 int `json:"y1"`
			X2 int `json:"x2"`
			Y2 int `json:"y2"`
		}

		if err := c.ShouldBindJSON(&move); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		game.MovePiece(move.X1, move.Y1, move.X2, move.Y2)
		c.JSON(http.StatusOK, game.ChessBoard)
	})

	// Start the server
	fmt.Println("Server starting on :8080")
	r.Run(":8080")

}
