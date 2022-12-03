package models

import (
	"advent-of-code-2022/m/v2/utils"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type Group []Rucksack

func (g Group) GetBadgeItem() (Item, error) {
	if len(g) < 1 {
		return ' ', errors.New("no rucksacks in group")
	}
	sharedItems := g[0].GetAllItems()
	for _, rucksack := range g[1:] {
		rucksackItems := rucksack.GetAllItems()
		for _, item := range sharedItems {
			if !slices.Contains(rucksackItems, item) {
				sharedItems = utils.Remove(sharedItems, item)
			}
		}
	}

	if len(sharedItems) != 1 {
		return ' ', fmt.Errorf("only expected one shared item, got %#v", sharedItems)
	}

	return sharedItems[0], nil
}

type Rucksack struct {
	CompartmentOne []Item
	CompartmentTwo []Item
}

func (r Rucksack) GetAllItems() []Item {
	uniqueItems := []Item{}
	allItems := append(r.CompartmentOne, r.CompartmentTwo...)
	itemMap := make(map[Item]bool)
	for _, item := range allItems {
		itemMap[item] = true
	}
	for key := range itemMap {
		uniqueItems = append(uniqueItems, key)
	}
	return uniqueItems
}

type Item rune

const indexedAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (i Item) Priority() int {
	return strings.IndexRune(indexedAlphabet, rune(i)) + 1
}
