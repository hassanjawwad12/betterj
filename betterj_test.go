package betterj

import (
	"testing"
)

// Defined the input and output format for the function we created
func TestMinifyJ(t *testing.T) {
	input := `{
		"name": "Hassan Jawwad",
		"age": 22,
		"city": "Lahore"
	}`
	want := `{"name":"Hassan Jawwad","age":22,"city":"Lahore"}`

	t.Run(input, func(t *testing.T) {
		got, err := MinifyJ(input)
		if err != nil {
			panic(err)
		}
		if got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}
func TestBeautifyJ(t *testing.T) {
	input := `{"name":"Hassan Jawwad","age":22,"city":"Lahore"}`
	indent := "  "
	want := `{
  "name": "Hassan Jawwad",
  "age": 22,
  "city": "Lahore"
}`

	t.Run(input, func(t *testing.T) {
		got, err := BeautifyJ(input, indent)
		if err != nil {
			panic(err)
		}
		if got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}
