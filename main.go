package main

import (
	"KevinsProject/OrdersManipulation"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strings"
	"time"
)

type gui struct {
	layoutVerticalBox     *fyne.Container
	topContainer          *fyne.Container
	buttonContainer       *fyne.Container
	filterButton          *widget.Button
	fulfilledChangeButton *widget.Button
	radioGroup1           *widget.RadioGroup
	selectorB             *widget.Select
	progressBar           *fyne.Container
	pBar                  *widget.ProgressBar
	bottomContainer       *fyne.Container
	listItem              *widget.List
}

func newGUI() *gui {
	return &gui{}
}

var orderBy string
var upDown string
var needsChange = false
var myList = OrdersManipulation.GetOrders()

func (g *gui) makeUI() fyne.CanvasObject {
	g.filterButton = widget.NewButtonWithIcon("filter", theme.ConfirmIcon(), func() {
		myList.SortBy(upDown, orderBy)
		needsChange = true
	})
	g.fulfilledChangeButton = widget.NewButtonWithIcon("Status Update", theme.ConfirmIcon(), func() {})
	g.radioGroup1 = widget.NewRadioGroup([]string{"Ascending", "Descending"}, func(s string) {
		upDown = strings.ToLower(s)
	})
	g.radioGroup1.SetSelected("Descending")
	g.selectorB = widget.NewSelect([]string{"DEFAULT (Recent Order)", "Date", "Customer Name", "Address", "Total"}, func(s string) {
		orderBy = s
		switch orderBy {
		case "Customer Name":
			orderBy = "customer name"
		case "Address":
			orderBy = "address"
		case "Date":
			orderBy = "date"
		case "Total":
			orderBy = "total"
		case "DEFAULT (Recent Order)":
			orderBy = "date"
		}
	})
	g.selectorB.SetSelected("DEFAULT (Recent Order)")
	g.buttonContainer = container.NewHBox(
		g.filterButton,
		g.fulfilledChangeButton,
		g.radioGroup1,
		g.selectorB)
	g.pBar = &widget.ProgressBar{Value: 0.000000}
	g.progressBar = container.NewMax(
		g.pBar)
	g.listItem = widget.NewList(func() int { return len(myList) }, func() fyne.CanvasObject {
		return container.New(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
	}, func(id widget.ListItemID, item fyne.CanvasObject) {
		item.(*fyne.Container).Objects[1].(*widget.Label).SetText(strings.TrimSpace(strings.ReplaceAll(fmt.Sprintf("ID: %s\tItem:  %48s", myList[id].Name, myList[id].LineitemName), "\n", "")))
	})
	g.topContainer = container.NewGridWithRows(2,
		g.buttonContainer,
		g.progressBar)
	g.bottomContainer = container.NewMax(
		g.listItem)
	g.layoutVerticalBox = container.NewGridWithRows(2,
		g.topContainer,
		g.bottomContainer)

	return g.layoutVerticalBox
}

func main() {
	a := app.New()
	w := a.NewWindow("Coffee Machine")
	g := newGUI()
	w.Resize(fyne.Size{
		Width:  900,
		Height: 700,
	})
	w.SetContent(g.makeUI())
	go func() {
		for range time.Tick(time.Second) {
			if needsChange {
				g = newGUI()
				w.SetContent(g.makeUI())
				needsChange = false
			}
		}
	}()
	w.ShowAndRun()
}
