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
	flag.Parse()
	items := &app.Items{}
	if err := items.Load(dataFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch {
	case *add:
		items.Add("Book 2", 150)
		err := items.Write(dataFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *sell > 0:
		err := items.Sold(*sell)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = items.Write(dataFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)

	}

}

//add - add item by passing name and price
//sell=idx sell item at index = idx
