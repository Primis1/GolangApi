
import (
	// "container/heap"
	"flag"
	"fmt"
)

func main() {
	// Define flags

	name := flag.String("name", "T-shirt", "name specifier")
	quantity := flag.Int("quantity", 1, "quantity of products")
	discount:= flag.Bool("discount", false, "discount specifier")

	flag.Parse()


	fmt.Printf("Your clothes is: %v", *name)

	if *quantity > 1 {
		fmt.Printf("You have such: %v", *quantity)
	}
	
	if *discount {
		fmt.Print("You have a discount")
	}

}