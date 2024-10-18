package structformat_test

import (
	"fmt"
	"strconv"

	"github.com/gentlemanautomaton/structformat"
	"github.com/gentlemanautomaton/structformat/fieldformat"
)

// PersonFormat describes a format for people.
type PersonFormat struct {
	ID     fieldformat.Options
	Name   fieldformat.Options
	Age    fieldformat.Options
	Email  fieldformat.Options
	Status fieldformat.Options
}

// Age is a person's age in years.
type Age int

// String returns a string representation of the age.
func (age Age) String() string {
	if age == 0 {
		return ""
	}
	return strconv.Itoa(int(age))
}

// Person holds information about a person.
type Person struct {
	ID     string
	Name   string
	Age    Age
	Email  string
	Status string
}

// String returns a string representation of a person.
func (p Person) String() string {
	return p.Format(PersonFormat{})
}

// Format returns a string representation of a person with the given
// formatting options.
func (p Person) Format(format PersonFormat) string {
	var builder structformat.Builder
	builder.ApplyRules(structformat.InferRules(
		format.ID,
		format.Name,
		format.Age,
		format.Email,
		format.Status,
	))

	builder.WriteField(p.ID, fieldformat.Primary, format.ID)
	builder.WriteField(p.Name, format.Name)
	builder.WriteField(p.Age.String(), fieldformat.Label("Age"), fieldformat.Note, format.Age)
	builder.WriteField(p.Email, fieldformat.Note, format.Email)
	builder.Divide()
	builder.WriteField(p.Status, format.Status)

	return builder.String()
}

func Example() {
	// Prepare a list of people.
	people := []Person{
		{ID: "1591", Name: "Alice", Age: 57, Email: "alice@example.com", Status: "Online"},
		{ID: "122520", Name: "Bob", Age: 53, Email: "bob@example.com", Status: "On Vacation"},
		{Name: "Eve", Age: 34, Email: "eve@example.com", Status: "Listening"},
		{ID: "128", Name: "Mallory", Age: 29, Email: "mallory@example.com", Status: "Tinkering"},
		{ID: "172340", Name: "Doug", Age: 1},
		{Name: "Felix", Status: "Busy"},
	}

	// Print the ID and email address for each person.
	idAndEmail := PersonFormat{
		ID: fieldformat.Options{
			Include:   true,
			Alignment: fieldformat.Right,
		},
		Email: fieldformat.Combine(fieldformat.Include, fieldformat.Standard),
	}

	for _, person := range people {
		idAndEmail.ID.AdjustWidth(len(person.ID))
	}

	fmt.Println("People (ID, Email):")
	for _, person := range people {
		fmt.Println(person.Format(idAndEmail))
	}

	// Print all fields except for email.
	// Use a column width for ID and name.
	allExceptID := PersonFormat{
		ID:  fieldformat.Exclude.Options(),
		Age: fieldformat.Right.Options(),
	}
	allExceptID.Age.Padding = "0"

	for _, person := range people {
		allExceptID.Name.AdjustWidth(len(person.Name))
		allExceptID.Age.AdjustWidth(len(person.Age.String()))
		allExceptID.Email.AdjustWidth(len(person.Email))
	}

	fmt.Println("People (Name, Age, Email, Status):")
	for i, person := range people {
		fmt.Printf("%d: %s\n", i, person.Format(allExceptID))
	}

	// Output:
	// People (ID, Email):
	//   1591: alice@example.com
	// 122520: bob@example.com
	//         eve@example.com
	//    128: mallory@example.com
	// 172340
	//
	// People (Name, Age, Email, Status):
	// 0: Alice   (Age: 57, alice@example.com):   Online
	// 1: Bob     (Age: 53, bob@example.com):     On Vacation
	// 2: Eve     (Age: 34, eve@example.com):     Listening
	// 3: Mallory (Age: 29, mallory@example.com): Tinkering
	// 4: Doug    (Age: 01)
	// 5: Felix:                                  Busy
}
