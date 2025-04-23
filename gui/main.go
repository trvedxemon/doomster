package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	//setup basic view
	a := app.New()
	w := a.NewWindow("Doomster v0.0.1b")

	text := canvas.NewText("http://localhost:", color.White)
	input1 := widget.NewEntry()
	input1.SetPlaceHolder("/Endpoint address")
	input2 := widget.NewEntry()
	input2.SetPlaceHolder("Server port")
	startButton := widget.NewButton("Start server", func() {StartServer(w, input2.Text, input1.Text)})
	initcont := container.New(layout.NewGridLayout(4), text, input2, input1, startButton)
	content := container.New(layout.NewPaddedLayout(), initcont)
	w.SetContent(content)
	w.ShowAndRun()
}

func StartServer(w fyne.Window, p string, ep string){
	port, err := strconv.Atoi(p)
		if (err != nil){
			panic(err)
		}
	srv := Server{port: port, endpoint: ep}
	srv.Start()
	addr := "Server runs on http://localhost:" + p + srv.endpoint
	text := canvas.NewText(addr, color.White)
	stopButton := widget.NewButton("Stop", func (){
		text := canvas.NewText("http://localhost:", color.White)
		input1 := widget.NewEntry()
		input1.SetPlaceHolder("/Endpoint address")
		input2 := widget.NewEntry()
		input2.SetPlaceHolder("Server port")
		startButton := widget.NewButton("Start server", func() {StartServer(w, p, ep)})
		initcont := container.New(layout.NewGridLayout(4), text, input2, input1, startButton)
		content := container.New(layout.NewPaddedLayout(), initcont)
		srv.Stop()
		w.SetContent(content)
	})
	cont := container.New(layout.NewGridLayout(2), text, stopButton)
	container := container.New(layout.NewPaddedLayout(), cont)
	w.SetContent(container)
}