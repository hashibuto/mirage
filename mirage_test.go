package mirage

import (
	"testing"
)

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	NumPages int    `json:"num_pages"`
}

func TestReflect(t *testing.T) {
	book := &Book{}
	meta := Reflect(book, "")
	bookKeys := NewStringSet(meta.Keys())
	if !bookKeys.Has("Title") || !bookKeys.Has("Author") || !bookKeys.Has("NumPages") {
		t.Errorf("Item is missing one or more expected keys")
	}
}

func TestGetValueByName(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "")
	metaIo := meta.Io()
	title, err := metaIo.ValueFromName("Title")
	if err != nil {
		t.Error(err)
		return
	}

	if title != "Mr. Black" {
		t.Error("Problem reading value from name")
	}
}

func TestGetValueByTagName(t *testing.T) {
	book := &Book{
		Title:    "Mrs. White",
		Author:   "Mr. Black",
		NumPages: 444,
	}
	meta := Reflect(book, "json")
	metaIo := meta.Io()
	title, err := metaIo.ValueFromTagKey("title")
	if err != nil {
		t.Error(err)
		return
	}

	if title != "Mrs. White" {
		t.Error("Problem reading value from name")
	}
}

func TestGetValueByNameNewObj(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "")
	metaIo := meta.NewIo(&Book{
		Title:    "Mr. Red",
		Author:   "Mrs. Black",
		NumPages: 333,
	})
	title, err := metaIo.ValueFromName("Title")
	if err != nil {
		t.Error(err)
		return
	}

	if title != "Mr. Red" {
		t.Error("Problem reading value from name")
	}
}

func TestGetValueByTagNameNewObj(t *testing.T) {
	book := &Book{
		Title:    "Mrs. White",
		Author:   "Mr. Black",
		NumPages: 444,
	}
	meta := Reflect(book, "json")
	metaIo := meta.NewIo(&Book{
		Title:    "Mrs. Black",
		Author:   "Mr. Red",
		NumPages: 333,
	})
	title, err := metaIo.ValueFromTagKey("title")
	if err != nil {
		t.Error(err)
		return
	}

	if title != "Mrs. Black" {
		t.Error("Problem reading value from name")
	}
}

func TestSetValueByName(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "")
	metaIo := meta.Io()
	err := metaIo.SetValueByName("Title", "Bonnie Wagner")
	if err != nil {
		t.Error(err)
		return
	}

	if book.Title != "Bonnie Wagner" {
		t.Error("Problem setting value from name")
	}
}

func TestSetValueByNameNewObj(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "")

	book2 := &Book{
		Title:    "Sir. Robin",
		Author:   "Mr. Python",
		NumPages: 222,
	}

	metaIo := meta.NewIo(book2)
	err := metaIo.SetValueByName("NumPages", 334)
	if err != nil {
		t.Error(err)
		return
	}

	if book.NumPages != 333 {
		t.Error("Original object was damaged despite creating new IO object")
	}

	if book2.NumPages != 334 {
		t.Error("Problem setting value to new object")
	}
}
