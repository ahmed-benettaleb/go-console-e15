//go:build gio
// +build gio

package main

import (
	"fmt"
	"image/color"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// runGUI launches a minimal Gio GUI to enter a city and show weather.
func runGUI(apiKey string) {
	go func() {
		w := app.NewWindow(app.Title("Weather - Gio"))
		if err := loop(w, apiKey); err != nil {
			fmt.Println("GUI error:", err)
		}
		// give time for logs to flush
		time.Sleep(100 * time.Millisecond)
	}()
	app.Main()
}

func loop(w *app.Window, apiKey string) error {
	th := material.NewTheme()

	var ops op.Ops
	var cityEditor widget.Editor
	cityEditor.SingleLine = true
	var btn widget.Clickable
	var result string

	for {
		e := <-w.Events()
		switch ev := e.(type) {
		case system.DestroyEvent:
			return ev.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, ev)

			// The layout below will render widgets. We check clicks after rendering the button

			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.H4(th, "Weather - Gio").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &cityEditor, "Entrez une ville").Layout(gtx)
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					btnStyle := material.Button(th, &btn, "Rechercher")
					d := layout.UniformInset(unit.Dp(8)).Layout(gtx, btnStyle.Layout)
					// Check if button was clicked in this frame
					if btn.Clicked(gtx) {
						city := cityEditor.Text()
						if city != "" {
							// call getWeather in a goroutine to avoid blocking the UI
							go func(c string) {
								res, err := getWeather(c, apiKey)
								if err != nil {
									result = err.Error()
								} else {
									result = res
								}
							}(city)
						}
					}
					return d
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// result box
					txt := material.Body1(th, result)
					txt.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 0xff}
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, txt.Layout)
				}),
			)

			ev.Frame(gtx.Ops)
		}
	}
}
