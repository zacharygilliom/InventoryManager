package cli

import (
	"bufio"
	"fmt"
	"os"
)

// Initialize will build and populate our database tables.
func Initialize() {
	populateTables()
}

// Populate our Inventory table...
func populateTables() {
	fmt.Println("Welcome to the inventory manager system.\nLets get you set up with our systems.")
	fmt.Println("----------------------------------------")
	fmt.Println("First thing we will need to do is\npopulate your inventory.")
	fmt.Println("Please enter how many items\nyou will be adding and hit 'Enter'\nwhen you are done.")
	fmt.Println("----------------------------------------")
	numOfItems := getIntInput()

	// Make a slice to store the user's inventory items in.
	inventoryItems := make([]string, numOfItems)
	fmt.Println("Please list the exact names of\nthe items that you have in your inventory.\nPlease hit 'Enter' after each item.")
	for i := 0; i < numOfItems; i++ {
		fmt.Printf("Please enter item %v: ", i+1)
		invItem := getUserInventory()
		inventoryItems = append(inventoryItems, invItem)
	}
	fmt.Println(inventoryItems)

}

func getIntInput() int {
	var text int
	fmt.Scanln(&text)
	return text
}

func getUserInventory() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return line
}
