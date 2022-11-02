package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Should be required")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("name")

	if has {
		t.Error("should be invalid because it doesn't have requested field")
	}

	postedData := url.Values{}
	postedData.Add("name", "John")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	has = form.Has("name")
	if !has {
		t.Error("Should be valid ")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	hasLength := form.MinLength("name", 3)
	if hasLength {
		t.Error("Shouldn't be min length valid for non-existing field")
	}

	isError := form.Errors.Get("name")
	if isError == "" {
		t.Error("Should have an error, but didnt get one")
	}

	postedData = url.Values{}
	postedData.Add("name", "John")
	form = New(postedData)

	hasLength = form.MinLength("name", 10)
	if hasLength {
		t.Error("Should not have the requested min length")
	}

	postedData = url.Values{}
	postedData.Add("name", "John")
	form = New(postedData)

	hasLength = form.MinLength("name", 3)
	if !hasLength {
		t.Error("Should have the requested min length")
	}

	isError = form.Errors.Get("name")
	if isError != "" {
		t.Error("Shouldn't have an error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("Shouldn't be valid for non-existing email field")
	}

	postedData = url.Values{}
	postedData.Add("email", "john")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("Shouldn't be valid for non-valid email field")
	}

	postedData = url.Values{}
	postedData.Add("email", "john@john.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Should be valid for valid email field")

	}

}
