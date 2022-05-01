package data

import "fmt"

/************************ Define error code ************************/
var ErrItemNotFound = fmt.Errorf("Item in ShoppingCart not found")
var UserIdForShoppingCartNotFound = fmt.Errorf("User in ShoppingCart not found")

/************************ Define structure shopping cart ************************/
type ShoppingCart struct {
	UserId int                `json:"userid"`
	Items  []ShoppingCartItem `json:"shoppingcartitem"`
}

func UserShoppingCart(userID int) *ShoppingCart {
	return &ShoppingCart{
		UserId: userID,
	}
}

// func (s *ShoppingCart) AddItems(item ShoppingCartItem) {
// 	(*s).items = append((*s).items, item)
// }

// func (s *ShoppingCart) DeleteItems(pId int) {

// 	for i, p := range (*s).items {
// 		if p.ProductCatalogueId == pId {
// 			(*s).items = append((*s).items[:i], (*s).items[i+1])
// 		}
// 	}
// }

/************************ Define structure shopping cart item ************************/
type ShoppingCartItem struct {
	ProductCatalogueId int    `json:"productcatalogueid"`
	ProductName        string `json:"productname"`
	Description        string `json:"description"`
	Price              Money  `json:"money"`
}

func NewShoppingCartItem(item ShoppingCartItem) *ShoppingCartItem {
	return &ShoppingCartItem{
		ProductCatalogueId: item.ProductCatalogueId,
		ProductName:        item.ProductName,
		Description:        item.Description,
		Price:              item.Price,
	}
}

/************************ Define structure money ************************/
type Money struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

func NewMoney(currency string, amount float64) *Money {
	return &Money{
		Currency: currency,
		Amount:   amount,
	}
}
