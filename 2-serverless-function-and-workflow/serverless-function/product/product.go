package product

import "fmt"

func GetById(productId string) string {
	return fmt.Sprintf("The product name for your product id %s is Starfruit Explosion", productId)
}
