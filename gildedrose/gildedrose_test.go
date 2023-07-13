package gildedrose_test

import (
	"testing"

	gildedrose "github.com/berdebotond/gildedrose_go/gildedrose"
	pkg "github.com/berdebotond/gildedrose_go/pkg"
)

func Test_Foo(t *testing.T) {
	var items = []*pkg.Item{
		{Name: "foo", SellIn: 0, Quality: 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}
