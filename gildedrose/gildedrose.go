package gildedrose

import (
	"github.com/berdebotond/gildedrose_go/pkg"
)

func UpdateQuality(items []*pkg.Item) {
	for _, item := range items {
		updateItemQuality(item)
	}
}

func updateItemQuality(item *pkg.Item) {
	switch item.Name {
	case "Aged Brie":
		updateAgedBrie(item)
	case "Backstage passes to a TAFKAL80ETC concert":
		updateBackstagePasses(item)
	case "Sulfuras, Hand of Ragnaros":
		// Legendary item, do nothing
	case "Conjured":
		updateConjured(item)
	default:
		updateNormal(item)
	}

	// Ensure quality is within bounds
	if item.Quality < 0 {
		item.Quality = 0
	} else if item.Quality > 50 && item.Name != "Sulfuras, Hand of Ragnaros" {
		item.Quality = 50
	}

	// Decrease sell-in value for all items except Sulfuras
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn--
	}
}

func updateBackstagePasses(item *pkg.Item) {
	if item.SellIn <= 0 {
		item.Quality = 0
	} else if item.SellIn <= 5 {
		item.Quality += 3
	} else if item.SellIn <= 10 {
		item.Quality += 2
	} else {
		item.Quality++
	}
}

func updateConjured(item *pkg.Item) {
	if item.SellIn <= 0 {
		item.Quality -= 4
	} else {
		item.Quality -= 2
	}
}

func updateAgedBrie(item *pkg.Item) {
	if item.SellIn > 0 {
		item.Quality++
	} else {
		item.Quality += 2
	}

	if item.Quality > 50 {
		item.Quality = 50
	}
}

func updateNormal(item *pkg.Item) {
	if item.SellIn > 0 {
		item.Quality--
	} else {
		item.Quality -= 2
	}

	if item.Quality < 0 {
		item.Quality = 0
	}
}
