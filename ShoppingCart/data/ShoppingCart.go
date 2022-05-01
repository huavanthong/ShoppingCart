package data

import "fmt"

/************************ Define error code ************************/
var ErrItemNotFound = fmt.Errorf("Item in ShoppingCart not found")
var UserIdForShoppingCartNotFound = fmt.Errorf("User in ShoppingCart not found")

/************************ Define structure shopping cart ************************/
type ShoppingCart struct {
	UserId int
	Items  []ShoppingCartItem
}

func UserShoppingCart(userID int) *ShoppingCart {
	return &ShoppingCart{
		UserId: userID,
	}
}

func (s *ShoppingCart) AddItems(item ShoppingCartItem) {
	(*s).items = append((*s).items, item)
}

func (s *ShoppingCart) DeleteItems(pId int) {

	for i, p := range (*s).items {
		if p.ProductCatalogueId == pId {
			(*s).items = append((*s).items[:i], (*s).items[i+1])
		}
	}
}

/************************ Define structure shopping cart item ************************/
type ShoppingCartItem struct {
	ProductCatalogueId int
	ProductName        string
	Desscription       string
	Price              Money
}

func NewShoppingCartItem(item ShoppingCartItem) *ShoppingCartItem {
	return &ShoppingCartItem{
		ProductCatalogueId: item.ProductCatalogueId,
		ProductName:        item.ProductName,
		Desscription:       item.Desscription,
		Price:              item.Price,
	}
}

/************************ Define structure money ************************/
type Money struct {
	Currency string
	Amount   float64
}

func NewMoney(currency string, amount float64) *Money {
	return &Money{
		Currency: currency,
		Amount:   amount,
	}
}
