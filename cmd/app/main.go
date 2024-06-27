package main

import (
	"flag"
	"fmt"
	"os"

	app "github.com/dibyajyoti-mandal/cli-app"
)

const (
	dataFile = ".data.json"
)

func main() {
	add := flag.Bool("add", false, "add new item")
	sell := flag.Int("sell", 0, "mark item sold")
	delete := flag.Int("del", 0, "unlist and item")
	list := flag.Bool("list", false, "list all items")
	flag.Parse()

	items := &app.Items{}
	if err := items.Load(dataFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	var err error
	switch {

	case *list:
		items.Show()

	case *add:

		var price int
		var item string
		fmt.Println("Enter name of item: ")
		fmt.Scan(&item)
		fmt.Println("Enter price of item: ")
		fmt.Scan(&price)
		items.Add(item, price)
		err = items.Write(dataFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Printf("item created: %s - price: %d", item, price)
		items.Show()

	case *sell > 0:
		err = items.Sold(*sell)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if err == nil {
			err = items.Write(dataFile)
		}

	case *delete > 0:
		err = items.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if err == nil {
			err = items.Write(dataFile)
		}

	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
