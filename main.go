package main

import (
	"api-requester/context"
	"api-requester/tui/components/main_page"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ctx, err := context.NewAppContext()
	if err != nil {
		ctx.Logger.Fatalln(err)
	}

	p := tea.NewProgram(main_page.NewModel(ctx), tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		ctx.Logger.Fatalln(err)
	}
}
