package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Gallery")

	// Create a list of image file paths (replace with your own images)
	imagePaths := []string{"img/1.jpg", "img/2.png", "img/3.jpg"}

	// Create a grid layout for the images
	gridLayout := container.New(layout.NewGridLayout(3))

	// Add images to the grid
	for _, path := range imagePaths {
		image := canvas.NewImageFromFile(path)
		//image.Resize(fyne.NewSize(10, 10))
		image.FillMode = canvas.ImageFillContain
		gridLayout.Add(image)
	}

	myWindow.SetContent(gridLayout)
	myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.ShowAndRun()
}
