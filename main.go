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
	selectorB2            *widget.Select
	progressBar           *fyne.Container
	pBar                  *widget.ProgressBar
	bottomContainer       *fyne.Container
	listItem              *widget.List
	entryBox              *widget.Form
	loginBox              *widget.Form
}

func newGUI() *gui {
	return &gui{}
}

var orderBy string
var searchBy string
var searchStr string
var upDown string
var needsChange = false
var label = widget.NewLabel("")
var myList = OrdersManipulation.GetOrders()

func (g *gui) makeUI() fyne.CanvasObject {
	g.filterButton = widget.NewButtonWithIcon("filter", theme.ConfirmIcon(), func() {
		myList = myList.SortBy(upDown, orderBy)
		needsChange = true
	})
	g.fulfilledChangeButton = widget.NewButtonWithIcon("Status Update", theme.ConfirmIcon(), func() {})
	entry := widget.NewEntry()
	g.entryBox = &widget.Form{Items: []*widget.FormItem{widget.NewFormItem("Search", entry), widget.NewFormItem("Searched: ", label)},
		OnSubmit: func() {
			switch searchBy {
			case "customer name":
				myList = myList.GetOrdersByName(entry.Text)
			case "item name":
				myList = myList.GetOrdersByItemName(entry.Text)
			}
			str := entry.Text
			label = widget.NewLabel(str)
			needsChange = true
		},
		OnCancel: func() {
			entry.Text = ""
			label = widget.NewLabel(entry.Text)
			myList = OrdersManipulation.GetOrders()
			needsChange = true
		}}
	g.entryBox.CancelText = "Clear All"
	g.loginBox = &widget.Form{Items: []*widget.FormItem{widget.NewFormItem("Username", widget.NewEntry()), widget.NewFormItem("Password", widget.NewPasswordEntry())}, OnSubmit: func() {}, OnCancel: func() {}}
	g.radioGroup1 = widget.NewRadioGroup([]string{"Ascending", "Descending"}, func(s string) {
		upDown = strings.ToLower(s)
	})
	g.radioGroup1.SetSelected("Descending")
	g.selectorB = widget.NewSelect([]string{"Date", "Customer Name", "Address", "Total"}, func(s string) {
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
		default:
			orderBy = "date"
		}
	})
	g.selectorB2 = widget.NewSelect([]string{"Date", "Customer Name", "Address", "Item Name"}, func(s string) {
		searchBy = s
		switch searchBy {
		case "Customer Name":
			searchBy = "customer name"
		case "Address":
			searchBy = "address"
		case "Date":
			searchBy = "date"
		case "Item Name":
			searchBy = "item name"
		default:
			searchBy = "customer name"
		}
	})
	g.selectorB.SetSelected("Date")
	g.selectorB2.SetSelected("Customer Name")
	g.buttonContainer = container.NewHBox(
		g.fulfilledChangeButton,
		g.filterButton,
		g.radioGroup1,
		g.selectorB,
		g.entryBox,
		g.selectorB2,
		g.loginBox)
	g.pBar = &widget.ProgressBar{Value: 0.000000}
	g.progressBar = container.NewMax(
		g.pBar)
	g.listItem = widget.NewList(func() int { return len(myList) }, func() fyne.CanvasObject {
		return container.New(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
	}, func(id widget.ListItemID, item fyne.CanvasObject) {
		item.(*fyne.Container).Objects[1].(*widget.Label).SetText(strings.TrimSpace(strings.ReplaceAll(fmt.Sprintf("ID: %s\tItem:  %.10s\tItemSKU: %s\tQuantity: %.1f\tStatus: %s\tTotal: $%.2f", myList[id].Name, myList[id].LineitemName, myList[id].LineitemSku, myList[id].LineitemQuantity, myList[id].FulfillmentStatus, myList[id].Total), "\n", "")))
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

func truncateText(s string, max int) string {
	return s[:max]
}

func main() {
	a := app.New()
	w := a.NewWindow("Coffee Machine")
	g := newGUI()
	w.Resize(fyne.Size{
		Width:  1000,
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
