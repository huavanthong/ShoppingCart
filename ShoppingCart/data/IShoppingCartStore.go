package data

type IShoppingCartStore interface {
	Get(userId int) ShoppingCart
	Save(shoppingCart ShoppingCart)
}
