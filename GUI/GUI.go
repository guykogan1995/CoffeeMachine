package GUI

import (
	"KevinsProject/OrdersManipulation"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"strconv"
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
	radioGroup2           *widget.RadioGroup
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
var listType string
var needsChange = false
var myList = OrdersManipulation.GetOrders()
var myList2 = myList.GetUnFulfilledOrders()
var hasHappened = false
var oldSelect string
var oldSelect2 string
var oldSelect3 string
var oldSelect4 string
var oldSelect5 string

func RunGUI() {
	a := app.New()
	w := a.NewWindow("ShopifyDropShipPOS-Vunio V1.0")
	g := newGUI()
	w.Resize(fyne.Size{
		Width:  1000,
		Height: 700,
	})
	w.CenterOnScreen()
	w.SetFixedSize(true)
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

func (g *gui) makeUI() fyne.CanvasObject {
	g.filterButton = widget.NewButtonWithIcon("filter", theme.ConfirmIcon(), func() {
		if listType == "fulfilled" {
			myList2 = myList.GetFulfilledOrders()
			listType = ""
		} else if listType == "unfulfilled" {
			myList2 = myList.GetUnFulfilledOrders()
			listType = ""
		} else if listType == "rejected" {
			myList2 = OrdersManipulation.GetOrders()
		}
		oldSelect4 = ""
		oldSelect = g.radioGroup1.Selected
		oldSelect2 = g.selectorB.Selected
		oldSelect3 = g.selectorB2.Selected
		oldSelect5 = g.radioGroup2.Selected
		myList2 = myList2.SortBy(upDown, orderBy)
		needsChange = true
	})
	g.fulfilledChangeButton = widget.NewButtonWithIcon("Status Update", theme.ConfirmIcon(), func() {})
	entry := widget.NewEntry()
	g.entryBox = &widget.Form{Items: []*widget.FormItem{widget.NewFormItem("Search", entry)},
		OnSubmit: func() {
			switch searchBy {
			case "customer name":
				myList2 = myList2.GetOrdersByName(entry.Text)
			case "item name":
				myList2 = myList2.GetOrdersByItemName(entry.Text)
			}
			oldSelect = g.radioGroup1.Selected
			oldSelect2 = g.selectorB.Selected
			oldSelect3 = g.selectorB2.Selected
			oldSelect4 = entry.Text
			oldSelect5 = g.radioGroup2.Selected
			needsChange = true
		},
		OnCancel: func() {
			oldSelect4 = ""
			myList2 = myList.GetUnFulfilledOrders()
			needsChange = true
			hasHappened = false
			listType = ""
		}}
	g.entryBox.CancelText = "Clear All"
	g.loginBox = &widget.Form{Items: []*widget.FormItem{widget.NewFormItem("Username", widget.NewEntry()), widget.NewFormItem("Password", widget.NewPasswordEntry())}, OnSubmit: func() {}, OnCancel: func() {}}
	g.radioGroup1 = widget.NewRadioGroup([]string{"Ascending", "Descending"}, func(s string) {
		upDown = strings.ToLower(s)
	})
	g.radioGroup2 = widget.NewRadioGroup([]string{"Unfulfilled", "Fulfilled", "Rejected"}, func(s string) {
		listType = strings.ToLower(s)
	})
	g.selectorB = widget.NewSelect([]string{"Date", "Customer Name", "Address", "Total"}, func(s string) {
		orderBy = strings.ToLower(s)
	})
	g.selectorB2 = widget.NewSelect([]string{"Date", "Customer Name", "Address", "Item Name"}, func(s string) {
		searchBy = strings.ToLower(s)
	})
	if needsChange == true {
		g.radioGroup1.SetSelected(oldSelect)
		g.selectorB.SetSelected(oldSelect2)
		g.selectorB2.SetSelected(oldSelect3)
		g.radioGroup2.SetSelected(oldSelect5)
		entry.SetText(oldSelect4)
	}
	if hasHappened == false {
		g.radioGroup1.SetSelected("Descending")
		g.selectorB.SetSelected("Date")
		g.selectorB2.SetSelected("Customer Name")
		g.radioGroup2.SetSelected("Unfulfilled")
		hasHappened = true
	}
	g.buttonContainer = container.NewHBox(
		g.fulfilledChangeButton,
		g.filterButton,
		g.radioGroup1,
		g.selectorB,
		g.entryBox,
		g.selectorB2,
		g.radioGroup2,
		g.loginBox)

	g.pBar = widget.NewTextGrid()
	searchStr = fmt.Sprintf("%s %s %s %s %s %s \t%s", truncateText("ID:", 12), truncateText("Item:", 52), truncateText("SKU:", 20), truncateText("Quantity:", 12), truncateText("Status:", 16), "Value:", truncateText("Customer:", 16))
	tGrid := widget.NewTextGrid()
	tGrid.SetText(searchStr)
	g.pBar.SetText(searchStr)
	g.progressBar = container.NewBorder(nil, canvas.NewRectangle(colornames.Gray), nil, nil, container.NewScroll(g.pBar))
	g.listItem = widget.NewList(func() int { return len(myList2.Orders) }, func() fyne.CanvasObject {
		return container.NewMax(
			widget.NewTextGridFromString(""),
		)
	}, func(id widget.ListItemID, item fyne.CanvasObject) {
		itemStr := fmt.Sprintf("%s  %s  \t%s  \t%d  \t\t%s  \t$%s \t%s", truncateText(strconv.Itoa(int(myList2.Orders[id].ID)), 11), truncateText(strings.TrimSpace(myList2.Orders[id].LineItems[0].Name), 48), truncateText(strings.TrimSpace(myList2.Orders[id].LineItems[0].SKU), 17), myList2.Orders[id].LineItems[0].Quantity, truncateText(fmt.Sprintf("%v", myList2.Orders[id].FulfillmentStatus), 11), myList2.Orders[id].TotalPrice, truncateText(myList2.Orders[id].Customer.FirstName+" "+myList2.Orders[id].Customer.LastName, 15))
		itemStr = strings.ReplaceAll(itemStr, "\n", "")
		item.(*fyne.Container).Objects[0].(*widget.TextGrid).SetText(itemStr)
	})
	g.listItem.OnSelected = func(id widget.ListItemID) {
		itemStr := fmt.Sprintf("%s  %s  \t%s  \t%d  \t\t%s  \t$%s \t%s", truncateText(strconv.Itoa(int(myList2.Orders[id].ID)), 11), truncateText(strings.TrimSpace(myList2.Orders[id].LineItems[0].Name), 48), truncateText(strings.TrimSpace(myList2.Orders[id].LineItems[0].SKU), 17), myList2.Orders[id].LineItems[0].Quantity, truncateText(fmt.Sprintf("%v", myList2.Orders[id].FulfillmentStatus), 11), myList2.Orders[id].TotalPrice, truncateText(myList2.Orders[id].Customer.FirstName+" "+myList2.Orders[id].Customer.LastName, 15))
		itemStr = strings.ReplaceAll(itemStr, "\n", "")
		g.pBar.SetText(g.pBar.Text() + "\n" + itemStr)
	}
	g.topContainer = container.NewGridWithRows(2,
		g.buttonContainer,
		g.progressBar)
	g.bottomContainer = container.NewBorder(tGrid, nil, nil,
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
