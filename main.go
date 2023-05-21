package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/svgo"
)

func main() {
	width, height := 1600, 800
	squareSize := 100

	f, err := os.Create("animated_square.svg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	canvas := svg.New(f)
	canvas.Start(width, height)

	// Define the linear gradient
	canvas.Def()
	canvas.LinearGradient("gradient-blue", 0, 0, 0, 100, []svg.Offcolor{
		{Offset: 0, Color: "#6083CB", Opacity: 1.0},
		{Offset: 50, Color: "#3E70CA", Opacity: 1.0},
		{Offset: 100, Color: "#2E61BA", Opacity: 1.0},
	})
	canvas.DefEnd()

	// Draw a square with the gradient
	canvas.Def()
	canvas.Gid("square")
	canvas.Rect(0, 0, squareSize, squareSize, "fill:url(#gradient-blue)")
	canvas.Text(squareSize/2, squareSize/2+5,
		"Hello World",
		"text-anchor:middle; font-size:14px; fill:white; ")
	canvas.Gend()
	canvas.DefEnd()

	animationStyle := `
	<style>
		@keyframes moveSquare {
			from {transform: translateX(100%);}
			to {transform: translateX(0%);}
		}
	</style>
	`

	canvas.Writer.Write([]byte("<!-- Add animation style -->\n"))
	canvas.Writer.Write([]byte(animationStyle))

	canvas.Gstyle("animation: moveSquare 3s linear forwards;")
	canvas.Use(0, height/2-squareSize/2, "#square")
	canvas.Gend()

	canvas.End()
}
