package tasks

import "fmt"

var (
	ErrEmptyName   = fmt.Errorf("name cannot be empty")
	ErrEmptyCity   = fmt.Errorf("city cannot be empty")
	ErrNegativeAge = fmt.Errorf("age cannot be negative")
)

type Person struct {
	name string
	age  int
	city string
}

func (p *Person) String() string {
	return fmt.Sprintf("%s\n%d years old\nfrom %s", p.name, p.age, p.city)
}

func New(name, city string, age int) (*Person, error) {
	p := &Person{
		name: name,
		age:  age,
		city: city,
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Person) validate() error {
	if p.name == "" {
		return ErrEmptyName
	}

	if p.city == "" {
		return ErrEmptyCity
	}

	if p.age < 0 {
		return ErrNegativeAge
	}

	return nil
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) GetCity() string {
	return p.city
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetCity(city string) {
	p.city = city
}

func (p *Person) SetAge(age int) {
	p.age = age
}
