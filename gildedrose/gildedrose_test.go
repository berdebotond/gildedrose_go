package gildedrose

import (
	"testing"

	"github.com/berdebotond/gildedrose_go/pkg"
)

func TestUpdateQuality(t *testing.T) {
	items := []*pkg.Item{
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		{Name: "Conjured", SellIn: 3, Quality: 6},
		{Name: "Normal Item", SellIn: 5, Quality: 7},
	}

	UpdateQuality(items)

	if items[0].Quality != 1 {
		t.Errorf("Expected quality of Aged Brie to be 1, got %d", items[0].Quality)
	}

	if items[1].Quality != 21 {
		t.Errorf("Expected quality of Backstage passes to be 21, got %d", items[1].Quality)
	}

	if items[2].Quality != 80 {
		t.Errorf("Expected quality of Sulfuras to be 80, got %d", items[2].Quality)
	}

	if items[3].Quality != 4 {
		t.Errorf("Expected quality of Conjured item to be 4, got %d", items[3].Quality)
	}

	if items[4].Quality != 6 {
		t.Errorf("Expected quality of Normal Item to be 6, got %d", items[4].Quality)
	}
}

func TestUpdateAgedBrie(t *testing.T) {
	item := &pkg.Item{Name: "Aged Brie", SellIn: 2, Quality: 0}
	updateAgedBrie(item)
	if item.Quality != 1 {
		t.Errorf("Expected quality to be 1, got %d", item.Quality)
	}
}

func TestUpdateAgedBrieWithMaxQuality(t *testing.T) {
	item := &pkg.Item{Name: "Aged Brie", SellIn: 2, Quality: 50}
	updateAgedBrie(item)
	if item.Quality != 50 {
		t.Errorf("Expected quality to be 1, got %d", item.Quality)
	}
}

func TestUpdateBackstagePasses(t *testing.T) {
	item := &pkg.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20}
	updateBackstagePasses(item)
	if item.Quality != 21 {
		t.Errorf("Expected quality to be 21, got %d", item.Quality)
	}
}

func TestUpdateConjured(t *testing.T) {
	item := &pkg.Item{Name: "Conjured", SellIn: 3, Quality: 6}
	updateConjured(item)
	if item.Quality != 4 {
		t.Errorf("Expected quality to be 4, got %d", item.Quality)
	}
}

func TestUpdateNormal(t *testing.T) {
	item := &pkg.Item{Name: "Normal Item", SellIn: 5, Quality: 7}
	updateNormal(item)
	if item.Quality != 6 {
		t.Errorf("Expected quality to be 6, got %d", item.Quality)
	}
}

func TestUpdateQualityMultipleDays(t *testing.T) {
	items := []*pkg.Item{
		&pkg.Item{Name: "Aged Brie", SellIn: 2, Quality: 0},
		&pkg.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		&pkg.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		&pkg.Item{Name: "Conjured", SellIn: 3, Quality: 6},
		&pkg.Item{Name: "Normal Item", SellIn: 5, Quality: 7},
	}

	for i := 0; i < 5; i++ {
		UpdateQuality(items)
	}

	if items[0].Quality != 8 {
		t.Errorf("Expected quality of Aged Brie to be 8, got %d", items[0].Quality)
	}

	if items[1].Quality != 25 {
		t.Errorf("Expected quality of Backstage passes to be 25, got %d", items[1].Quality)
	}

	if items[2].Quality != 80 {
		t.Errorf("Expected quality of Sulfuras to be 80, got %d", items[2].Quality)
	}

	if items[3].Quality != 0 {
		t.Errorf("Expected quality of Conjured item to be 0, got %d", items[3].Quality)
	}

	if items[4].Quality != 2 {
		t.Errorf("Expected quality of Normal Item to be 2, got %d", items[4].Quality)
	}
}

func TestUpdateQualityAdditionalCases(t *testing.T) {
	items := []*pkg.Item{
		&pkg.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 20},
		&pkg.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20},
		&pkg.Item{Name: "Aged Brie", SellIn: 2, Quality: 49},
		&pkg.Item{Name: "Normal Item", SellIn: 0, Quality: 1},
	}

	UpdateQuality(items)

	if items[0].Quality != 23 {
		t.Errorf("Expected quality of Backstage passes to be 23, got %d", items[0].Quality)
	}

	if items[1].Quality != 22 {
		t.Errorf("Expected quality of Backstage passes to be 22, got %d", items[1].Quality)
	}

	if items[2].Quality != 50 {
		t.Errorf("Expected quality of Aged Brie to be 50, got %d", items[2].Quality)
	}

	if items[3].Quality != 0 {
		t.Errorf("Expected quality of Normal Item to be 0, got %d", items[3].Quality)
	}
}

func TestUpdateBackstagePassesEdgeCases(t *testing.T) {
	items := []*pkg.Item{
		&pkg.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 20},
		&pkg.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49},
	}

	UpdateQuality(items)

	if items[0].Quality != 0 {
		t.Errorf("Expected quality of Backstage passes to be 0, got %d", items[0].Quality)
	}

	if items[1].Quality != 50 {
		t.Errorf("Expected quality of Backstage passes to be 50, got %d", items[1].Quality)
	}
}
