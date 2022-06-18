package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
var (
	P1 = Person {
		firstName: "a",
		lastName: "a",
		birthDay: time.Now(),
	}
	P2 = Person {
		firstName: "b",
		lastName: "b",
		birthDay: time.Now(),
	}
	P3 = Person {
		firstName: "c",
		lastName: "c",
		birthDay: time.Now(),
	}
)

func TestSwapPeople(t *testing.T) {
	people := People {P1, P2, P3}
	people.Swap(0, 2)

	if (P1 != people[2]) && (P3 != people[0]) && (P2 == people[1]) {
		t.Errorf("persons didn't swap")
	}
}

func TestSwapPeopleWithSameIndex(t *testing.T) {
	people := People {P1, P2, P3}
	people.Swap(0, 0)

	if !((P1 == people[0]) && (P2 == people[1]) && (P3 == people[2])) {
		t.Errorf("persons should not to swap")
	}
}

func TestPeopleLen(t *testing.T) {
	t.Parallel()
	table := map[string]struct {
		Persons People
		Expected int
	} {
		"empty": {People{}, 0},
		"single person": {People{P1}, 1},
		"many persons": {People{P1, P2, P3}, 3},
	}
	for name, tcase := range table {
		t.Run(name, func(t *testing.T) {
			if tcase.Persons.Len() != tcase.Expected {
				t.Errorf("Wrong len result for People")
			}
		})
	}
}

var ErrMsgForTestLessFunc = "Wrong result of Less function"

func TestPeopleLess(t *testing.T) {
	t.Parallel()

	date := time.Now()
	firstName := "a"
	lastName := "a"

	table := map[string]struct {
		Persons People
		Expected bool
	} {
		"same date, same firstName, diff lastName less": {
			People{
				Person {
					firstName: firstName,
					lastName: lastName + "b",
					birthDay: date,
				},
				Person {
					firstName: firstName,
					lastName: lastName + "c",
					birthDay: date,
				},
			},
			true,
		},
		"same date, diff firstName, diff lastName less": {
			People{
				Person {
					firstName: firstName + "b",
					lastName: lastName + "c",
					birthDay: date,
				},
				Person {
					firstName: firstName + "c",
					lastName: lastName + "b",
					birthDay: date,
				},
			},
			true,
		},
		"diff date, diff firstName, diff lastName less": {
			People{
				Person {
					firstName: firstName + "c",
					lastName: lastName + "c",
					birthDay: date,
				},
				Person {
					firstName: firstName + "b",
					lastName: lastName + "b",
					birthDay: date.Add(time.Duration(1)),
				},
			},
			false,
		},
	}
	for name, tcase := range table {
		t.Run(name, func (t *testing.T) {
			if tcase.Persons.Less(0, 1) != tcase.Expected {
				t.Error(ErrMsgForTestLessFunc)
			}
			if tcase.Persons.Less(1, 0) == tcase.Expected {
				t.Error(ErrMsgForTestLessFunc)
			}
		})
	}
}
