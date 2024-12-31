package betterj

import (
	"log"
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
			log.Printf("Error has occured: %v\n", err)
			panic(err)
		}
		if got != want {
			log.Printf("Function failed: expected %v, got %v\n", want, got)
		} else {
			log.Println("Function has passed successfully")
		}
	})
	log.Println("Finished Testing MinifyJ Function")
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
			log.Printf("Error has occured: %v\n", err)
			panic(err)
		}
		if got != want {
			log.Printf("Function failed: expected %v, got %v\n", want, got)
		} else {
			log.Println("Function has passed successfully")
		}
	})

	log.Println("Finished Testing BeautifyJ Function")
}


