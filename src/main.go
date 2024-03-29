package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

const imageHeight = 200

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
	for i, path := range imagePaths {
		p := path
		image := canvas.NewImageFromFile(path)
		image.FillMode = canvas.ImageFillContain
		image.ScaleMode = canvas.ImageScaleFastest
		image.SetMinSize(fyne.NewSize(0, imageHeight))
		openImage := func() {
			showImageInSameWindow(myWindow, p, gridLayout)
		}
		box := container.NewPadded(widget.NewButton(strconv.Itoa(i), openImage), image)
		gridLayout.Add(box)
	}

	// Wrap the grid in a scroller for scrolling functionality
	scrollContainer := container.NewScroll(gridLayout)
	myWindow.SetContent(scrollContainer)
	myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.ShowAndRun()
}

// Function to display the clicked image in the same window
func showImageInSameWindow(window fyne.Window, path string, gridLayout *fyne.Container) {
	image := canvas.NewImageFromFile(path)
	image.FillMode = canvas.ImageFillContain
	closeButton := widget.NewButton("close", func() {
		scrollContainer := container.NewScroll(gridLayout)
		window.SetContent(scrollContainer)
		window.Show()
	})
	content := container.New(layout.NewPaddedLayout(), closeButton, image)
	window.SetContent(content)
}
