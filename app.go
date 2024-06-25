package app

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Name      string
	Price     int
	Available bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Items []item

func (i *Items) Add(name string, price int) {
	book := item{
		Name:      name,
		Price:     price,
		Available: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
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
		return errors.New("invalid index")
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
