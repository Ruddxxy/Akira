package gui

import (
	"akira/internal/alias"
	"akira/internal/filecomplete"
	"fmt"
	"os/user"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func RunGUI() {
	a := app.New()
	w := a.NewWindow("Akira CLI Autocomplete Tool")

	aliasesData := alias.GetAliases()
	list := widget.NewList(
		func() int {
			count := 0
			for _, cmds := range aliasesData {
				count += len(cmds)
			}
			return count
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			//Flatten the map to a list
			flatten := []string{}
			for group, cmds := range aliasesData {
				for key, cmd := range cmds {
					flatten = append(flatten, fmt.Sprintf("%s: %s -> %s", group, key, cmd))
				}
			}
			if i < len(flatten) {
				o.(*widget.Label).SetText(flatten[i])
			}
		},
	)

	addButton := widget.NewButton("Add Alias", func() {
		groupEntry := widget.NewEntry()
		keyEntry := widget.NewEntry()
		cmdEntry := widget.NewEntry()
		dialog.ShowForm("Add Alias", "Add", "Cancel", []*widget.FormItem{
			{Text: "Group", Widget: groupEntry},
			{Text: "Key", Widget: keyEntry},
			{Text: "Command", Widget: cmdEntry},
		}, func(b bool) {
			if b && groupEntry.Text != "" && keyEntry.Text != "" && cmdEntry.Text != "" {
				group := groupEntry.Text
				key := keyEntry.Text
				cmd := cmdEntry.Text
				//Update the aliases map
				data := alias.GetAliases()
				if data[group] == nil {
					data[group] = make(map[string]string)
				}
				data[group][key] = cmd

				//Save
				usr, _ := user.Current()
				alias.SaveAliases(filepath.Join(usr.HomeDir))
				list.Refresh()
			}
		}, w)
	})

	suggestEntry := widget.NewEntry()
	suggestionsList := widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			// No suggestions yet
		},
	)
	suggestEntry.OnChanged = func(text string) {
		results := filecomplete.GetSuggestions(text)
		suggestionsList.Length = func() int {
			return len(results)
		}
		suggestionsList.UpdateItem = func(id widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(results[id])
		}
		suggestionsList.Refresh()
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Aliases", container.NewVBox(list, addButton)),
		container.NewTabItem("Suggestions", container.NewBorder(suggestEntry, nil, nil, nil, suggestionsList)),
	)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
