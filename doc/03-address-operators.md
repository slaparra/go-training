# Address operators


- [Golang: Address operators](https://golang.org/ref/spec#Address_operators)
- [Golang book: pointers](https://www.golang-book.com/books/intro/8)
    - [& and * operators](https://www.golang-book.com/books/intro/8#section1)
- [Examples](../todd-mcleod/04-addresses-and-pointers)
    - [Scan and memory addresses](../todd-mcleod/04-addresses-and-pointers/scan-and-memory-addresses.go)
    - [Tour golang pointers](https://tour.golang.org/moretypes/1)
   
Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1): [result](../todd-mcleod/04-addresses-and-pointers/swap-example.go)

```
func main() {
	var x int = 1
	var y int = 2

	//y is the value &y is the address
	fmt.Printf("Y address is %v and value %d\n", &y, y)

	swap(&x, &y)    //we pass the var addresses as arguments instead of the values
	                //because if we pass the variables without ampersand new 
	                //variables would be created with the same content
	                //and we couldn't swap the previous values (without return any value)
}

func swap(x *int, y *int) {
	//here, the value of x and y are the addresses of the variables, not the content
	//and *x and *y are the values of these addresses
	
	tmp := *x  //now tmp is 1
	
	//so here &x is the address that contain the address of the value 1
```
