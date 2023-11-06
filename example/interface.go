//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock_example
package example

type Item struct {
	ID   string
	Name string
}

type Shop interface {
	GetItem(id string) (Item, error)
	GetItems() ([]Item, error)
	CreateItem(id, name string, price int) error
	UpdateItem(id, name string, price int) error
	DeleteItem(id string) error
	DeleteItems(ids ...string) error
}
