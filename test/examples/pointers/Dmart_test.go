package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// rohini &  John are pointers variables
func TestNewUser(t *testing.T) {
	Rohini := NewUser("Rohini")
	John := NewUser("John")
	//checking two diff pointer variables having same address or not
	assert.NotEqual(t, Rohini, John)

	//adding items in the list
	Rohini.AddItem("Cloths")
	Rohini.AddItem("Soaps")
	Rohini.AddItem("Shampoo")
	Rohini.AddItem("Toys")

	list, err := Rohini.GiveMyList()
	assert.Equal(t, 4, len(list))
	assert.Equal(t, nil, err)

	list1, err := John.GiveMyList()
	assert.Equal(t, 0, len(list1))
	assert.Equal(t, ErroEmptyList, err)
}
