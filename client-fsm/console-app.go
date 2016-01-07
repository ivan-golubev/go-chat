package main

import (
	"github.com/gizak/termui"
	"fmt"
)

func main(){
	strs := []string{
		"[1] [send message](fg-blue)",
		"[2] [invite friend](fg-blue)",
		"[0] [logout](fg-red)"
	}

	ls := termui.NewList()
	ls.Items = strs
	ls.ItemFgColor = termui.ColorYellow
	ls.BorderLabel = "Chat"
	ls.Height = 5
	ls.Width = 25

	// // build
    termui.Body.AddRows(
        termui.NewRow(
        	termui.NewCol(6, 0, ls)
    	)
            // ui.NewCol(6, 0, widget1))
        // ui.NewRow(
        //     ui.NewCol(3, 0, widget2),
        //     ui.NewCol(3, 0, widget30, widget31, widget32),
        //     ui.NewCol(6, 0, widget4))
    )

    // calculate layout
    termui.Body.Align()

    termui.Render(termui.Body)

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	termui.Handle("/sys/kbd/1", func(termui.Event) {
		fmt.Println("Sending a message...")		
	})
	termui.Handle("/sys/kbd/2", func(termui.Event) {
		fmt.Println("Inviting a friend...")
	})
	termui.Handle("/sys/kbd/0", func(termui.Event) {		
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/q", func(termui.Event) {		
		termui.StopLoop()
	})
	termui.Loop()
}