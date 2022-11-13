package mirage

import (
	"reflect"
	"testing"
	"time"
)

type Book struct {
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	NumPages int     `json:"num_pages"`
	ISBN     *string `json:"isbn"`
}

type Journal struct {
	Date  *time.Time `json:"date"`
	Title string     `json:"title"`
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
		return
	}

	if book2.NumPages != 334 {
		t.Error("Problem setting value to new object")
		return
	}
}

func TestKindByName(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "")

	info, err := meta.InfoByName("Author")
	if err != nil {
		t.Error(err)
		return
	}
	if info.Kind != reflect.String {
		t.Error("Incorrectly reported kind")
		return
	}
}

func TestKindByTagKey(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "json")

	info, err := meta.InfoByTagKey("title")
	if err != nil {
		t.Error(err)
		return
	}
	if info.Kind != reflect.String {
		t.Error("Incorrectly reported kind")
		return
	}
}

func TestKindByNameNullable(t *testing.T) {
	isbn := "1234-1234"
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
		ISBN:     &isbn,
	}
	meta := Reflect(book, "")

	info, err := meta.InfoByName("ISBN")
	if err != nil {
		t.Error(err)
		return
	}
	if !info.IsPointer {
		t.Error("Incorrectly reported pointer type")
		return
	}
	if info.Kind != reflect.String {
		t.Error("Incorrectly reported kind")
		return
	}
}

func TestNilPointerCheck(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "")
	isNilPtr, _ := meta.Io().IsNilPointerByName("ISBN")
	if !isNilPtr {
		t.Errorf("Unable to detect nil pointer")
		return
	}
}

func TestInstantiate(t *testing.T) {
	book := &Book{
		Title:    "Mr. Black",
		Author:   "Mrs. White",
		NumPages: 333,
	}
	meta := Reflect(book, "")
	meta.Io().InstantiateByName("ISBN")

	isNilPtr, _ := meta.Io().IsNilPointerByName("ISBN")
	if isNilPtr {
		t.Errorf("Should not be a nil pointer")
		return
	}
}

func TestInstantiate2(t *testing.T) {
	type Inner struct {
		X string
	}

	type Outer struct {
		Inner *Inner
	}

	out := &Outer{}
	meta := Reflect(out, "")
	meta.Io().InstantiateByName("Inner")

	if out.Inner.X != "" {
		t.Errorf("Problem instantiating field")
	}
}
