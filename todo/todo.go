package todo

import (
	"encoding/json"
	"os"
)

type Item struct {
	Text     string
	Priority int
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)

	if err != nil {
		return []Item{}, err
	}

	var items []Item

	err = json.Unmarshal(b, &items)

	if err != nil {
		return []Item{}, err
	}

	return items, nil
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (i *Item) SetPriority(priority int) {
	switch priority {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}
