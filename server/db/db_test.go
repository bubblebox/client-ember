package db_test

import (
	"os"
	"testing"
	"time"

	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/model"
)

var database = &db.DB{}

func TestGetItem(t *testing.T) {
	item, err := database.GetItem("url")

	if err != nil {
		t.Error(err)
	}

	if item.Code != "url" {
		t.Error("GetItem() did not return an Item with the proper ID.")
	}
}

func TestGetItems(t *testing.T) {
	items, err := database.GetItems()

	if err != nil {
		t.Error(err)
	}

	if len(items) != 2 {
		t.Errorf("Expected to find 2 items, got %d instead.", len(items))
	}
}

func TestGetItemNotFound(t *testing.T) {
	item, err := database.GetItem("nope")

	if item != nil {
		t.Error("Expected item to be `nil` when it does not exist.")
	}

	if err == nil {
		t.Error("Expected an error when unable to find an Item by ID.")
	}
}

func TestSaveItem(t *testing.T) {
	var err error
	var actual *model.Item

	item := &model.Item{
		Code:      "bam",
		Content:   "Lorem Ipsum",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	}

	err = database.SaveItem(item)

	if err != nil {
		t.Error("Did not expect an error saving Item", err)
	}

	actual, err = database.GetItem("bam")
	if actual == nil || err != nil {
		t.Error("Expected to be able to retrieve just saved Item.")
	}

	database.DeleteItem(item.Code)
}

func TestDeleteItem(t *testing.T) {
	database.SaveItem(&model.Item{
		Code:      "bam",
		Content:   "Lorem Ipsum",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	})

	item, _ := database.GetItem("bam")
	if item == nil {
		t.Error("Expected item to be persisted.")
	}

	err := database.DeleteItem("bam")

	if err != nil {
		t.Error("failed deleting item")
		t.Error(err)
	}

	actual, err := database.GetItem("bam")

	if actual != nil {
		t.Error("Expected item to be `nil` when it does not exist.")
	}

	if err == nil {
		t.Error("Expected an error when unable to find an Item by ID.")
	}
}

func setup() {
	database.Open("testing.db")

	items := []*model.Item{
		&model.Item{Code: "url", Type: model.URLItemType, Content: "https://ariejan.net", CreatedAt: time.Now()},
		&model.Item{Code: "txt", Type: model.TextItemType, Content: "Lorem ipsum", CreatedAt: time.Now()},
	}

	for _, item := range items {
		database.SaveItem(item)
	}
}

func teardown() {
	os.Remove(database.Path())
	database.Close()
}

func TestMain(m *testing.M) {
	setup()

	retCode := m.Run()

	teardown()

	os.Exit(retCode)
}
