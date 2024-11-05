package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPersonCreation(t *testing.T) {

	tt := []struct {
		testName string
		name     string
		age      int
		city     string
		expected *Person
		err      error
	}{
		{
			testName: "Test correct creation",
			name:     "Nikita",
			age:      25,
			city:     "Moscow",
			expected: &Person{
				name: "Nikita",
				age:  25,
				city: "Moscow",
			},
		},
		{
			testName: "Test with empty name",
			name:     "",
			age:      25,
			city:     "Moscow",
			err:      ErrEmptyName,
		},
		{
			testName: "Test with empty city",
			name:     "Nikita",
			age:      25,
			city:     "",
			err:      ErrEmptyCity,
		},
		{
			testName: "Test with negative age",
			name:     "Nikita",
			age:      -25,
			city:     "Moscow",
			err:      ErrNegativeAge,
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.testName,
			func(t *testing.T) {
				p, err := New(tc.name, tc.city, tc.age)
				assert.Equal(t, tc.expected, p)
				assert.Equal(t, tc.err, err)
			},
		)
	}

}

func TestPersonGetters(t *testing.T) {

	p := &Person{
		name: "Nikita",
		age:  25,
		city: "Moscow",
	}

	assert.Equal(t, "Nikita", p.GetName())
	assert.Equal(t, 25, p.GetAge())
	assert.Equal(t, "Moscow", p.GetCity())
}

func TestPersonSetters(t *testing.T) {

	p := &Person{
		name: "Nikita",
		age:  25,
		city: "Moscow",
	}

	name, city := "Boris", "Saint-Petersburg"
	age := 20

	p.SetName(name)
	p.SetCity(city)
	p.SetAge(age)

	assert.Equal(t, name, p.GetName())
	assert.Equal(t, age, p.GetAge())
	assert.Equal(t, city, p.GetCity())
}

func TestPersonFormatting(t *testing.T) {
	p, _ := New("Nikita", "Moscow", 25)

	require.Equal(t, "Nikita\n25 years old\nfrom Moscow", p.String())
}
