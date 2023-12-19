package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//	type ClickableImage struct {
//		Path string
//	}
//
//	func (c *ClickableImage) Tapped(_ *fyne.PointEvent) {
//		showImageInSameWindow(c.Path)
//	}
const imageHeight = 300

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Gallery")

	// Create a list of image file paths (replace with your own images)
	imagePaths := []string{"img/1.jpg", "img/2.png", "img/3.jpg",
		"img/4.png", "img/5.jpg", "img/6.jpg", "img/1.jpg", "img/2.png", "img/3.jpg",
		"img/4.png", "img/5.jpg", "img/6.jpg", "img/1.jpg", "img/2.png", "img/3.jpg",
		"img/4.png", "img/5.jpg", "img/6.jpg", "img/1.jpg", "img/2.png", "img/3.jpg",
		"img/4.png", "img/5.jpg", "img/6.jpg",
		"img/7.jpg"}

	// Create a grid layout for the images
	gridLayout := container.New(layout.NewGridLayout(3))

	// Add images to the grid
	for _, path := range imagePaths {
		image := canvas.NewImageFromFile(path)
		image.FillMode = canvas.ImageFillContain
		image.SetMinSize(fyne.NewSize(0, imageHeight))

		openButton := widget.NewButton("", func() {
			showImageInSameWindow(myWindow, path)
		})
		box := container.NewPadded(openButton, image)
		gridLayout.Add(box)
	}

	// Wrap the grid in a scroller for scrolling functionality
	scrollContainer := container.NewScroll(gridLayout)
	myWindow.SetContent(scrollContainer)
	myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.ShowAndRun()
}

// Function to display the clicked image in the same window
func showImageInSameWindow(window fyne.Window, path string) {
	image := canvas.NewImageFromFile(path)
	content := container.New(layout.NewCenterLayout(), image)
	image.FillMode = canvas.ImageFillContain
	window.SetContent(content)
}
