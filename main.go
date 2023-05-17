package main

import (
	"KevinsProject/GUI"
	"KevinsProject/ShopifyAPI"
)

func main() {
	ShopifyAPI.GetDataForJSON("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")
	GUI.RunGUI()
}
