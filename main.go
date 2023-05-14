package main

import (
	"KevinsProject/OrdersManipulation"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	pBar                  *widget.TextGrid
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
		orderBy = strings.ToLower(s)
	})
	g.selectorB2 = widget.NewSelect([]string{"Date", "Customer Name", "Address", "Item Name"}, func(s string) {
		searchBy = strings.ToLower(s)
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
	g.pBar = widget.NewTextGrid()
	g.pBar.SetText(fmt.Sprintf("%s %13s %47s %18s %8s %14s", "ID:", "Item:", "SKU:", "Quantity:", "Status:", "Value:"))
	g.progressBar = container.NewMax(
		g.pBar)
	g.listItem = widget.NewList(func() int { return len(myList) }, func() fyne.CanvasObject {
		return container.NewMax(
			widget.NewButton("", nil),
		)
	}, func(id widget.ListItemID, item fyne.CanvasObject) {
		itemStr := fmt.Sprintf("%s  %s  %s  %.1f  %s  $%.2f", truncateText(myList[id].Name, 11), truncateText(myList[id].LineitemName, 48), strings.TrimSpace(truncateText(myList[id].LineitemSku, 5)), myList[id].LineitemQuantity, myList[id].FulfillmentStatus, myList[id].Total)
		itemStr = strings.ReplaceAll(itemStr, "\n", "")
		item.(*fyne.Container).Objects[0].(*widget.Button).Alignment = widget.ButtonAlignLeading
		item.(*fyne.Container).Objects[0].(*widget.Button).SetText(itemStr)
		item.(*fyne.Container).Objects[0].(*widget.Button).OnTapped = func() {
			//item.(*fyne.Container).Objects[0].(*widget.Button).Text
			g.pBar.SetText(g.pBar.Text() + "\n" + itemStr)
		}
		println(itemStr)
	})
	searchStr = fmt.Sprintf("%s %10s %49s %25s %14s %14s", "ID:", "Item:", "SKU:", "Quantity:", "Status:", "Value:")
	println(searchStr)
	g.topContainer = container.NewGridWithRows(2,
		g.buttonContainer,
		g.progressBar)
	g.bottomContainer = container.NewBorder(widget.NewLabel(searchStr), nil, nil,
		nil,
		g.listItem)
	g.layoutVerticalBox = container.NewGridWithRows(2,
		g.topContainer,
		g.bottomContainer)

	return g.layoutVerticalBox
}

func truncateText(s string, max int) string {
	if max > len(s) {
		s = fmt.Sprintf("%s%s", s, strings.Repeat(" ", max-len(s)-1))
	} else {
		s = s[:max]
	}
	return s
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
