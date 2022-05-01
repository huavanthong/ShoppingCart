package data

// Create storage for event
type ShoppingCarts []*ShoppingCart

/*********************************************
Implement interface from IShoppingCartStore.go
*********************************************/
// Get(userId int): Returns a ShoppingCart with userId.
// func (s *ShoppingCarts) Get(userId int) ShoppingCart {

// }

// // Raise: Add a event with eventName
// func (s *ShoppingCarts) Save(shoppingCart ShoppingCart) {

// }

// GetProducts returns a list of products
func GetShoppingCarts() ShoppingCarts {
	return shoppingCartList
}

// GetUserById returns a single shopping cart which matches the userid from the
// database.
// If a shopping cart is not found this function returns a UserIdForShoppingCartNotFound error
func GetUserById(id int) (*ShoppingCart, error) {
	i := findIndexByUserID(id)
	if id == -1 {
		return nil, UserIdForShoppingCartNotFound
	}

	return shoppingCartList[i], nil
}

/************************ Internal function for Product ************************/
// findIndex finds the index of a shopping cart with userid in the database
// returns -1 when no userid for shopping cart can be found
func findIndexByUserID(userId int) int {
	for i, s := range shoppingCartList {
		if s.UserId == userId {
			return i
		}
	}

	return -1
}

/************************ Storage Shopping Cart ************************/
// shoppingCart is a hard coded list of shopping cart user for this
// example data source
var shoppingCartList = []*ShoppingCart{
	&ShoppingCart{
		UserId: 42,
		Items: []ShoppingCartItem{
			{
				ProductCatalogueId: 1,
				ProductName:        "Basic t-shirt",
				Description:        "a quiet t-shirt",
				Price: Money{
					Currency: "eur",
					Amount:   40,
				},
			},
			{
				ProductCatalogueId: 2,
				ProductName:        "Fancy shirt",
				Description:        "a loud t-shirt",
				Price: Money{
					Currency: "eur",
					Amount:   50,
				},
			},
		},
	},
}
