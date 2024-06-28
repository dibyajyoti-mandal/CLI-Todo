package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Name      string
	Price     int
	Available bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

type Items []item

func (i *Items) Add(name string, price int) {
	book := item{
		Name:      name,
		Price:     price,
		Available: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	*i = append(*i, book)
}

func (i *Items) Sold(index int) error {
	list := *i
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}

	list[index-1].UpdatedAt = time.Now()
	list[index-1].Available = false

	return nil
}

func (i *Items) Delete(index int) error {
	list := *i
	if index <= 0 || index > len(list) {
		return errors.New("invalid index - no item with index exists")
	}
	*i = append(list[:index-1], list[index:]...)
	return nil
}

func (i *Items) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, i)
	if err != nil {
		return err
	}
	return nil
}

func (i *Items) Write(filename string) error {

	data, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (i *Items) Show() {

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: blue("S/No.")},
			{Align: simpletable.AlignCenter, Text: blue("Item")},
			{Align: simpletable.AlignCenter, Text: blue("Price")},
			{Align: simpletable.AlignCenter, Text: blue("Available")},
			{Align: simpletable.AlignCenter, Text: blue("Updated At")},
		},
	}

	var cells [][]*simpletable.Cell
	transac := 0
	for index, item := range *i {
		index++
		text := ""
		if item.Available {
			text = green("Yes")
		} else {
			text = red("Sold out")
			transac += item.Price
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: item.Name},
			{Text: fmt.Sprintf("%d", item.Price)},
			{Text: text},
			{Text: item.UpdatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: blue(fmt.Sprintf("Total transaction: $%s", strconv.Itoa(transac)))},
	}}

	table.SetStyle(simpletable.StyleMarkdown)
	table.Println()
}
