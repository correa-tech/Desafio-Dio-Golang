package main

import "fmt"



func main() {

for i := 1; i < 100; i++ {

	multiplosDeTres := i%3 == 0
	multiplosDeCinco := i%5 == 0

	if multiplosDeTres && multiplosDeCinco {
			fmt.Println("numero: PINPAN")
		} else if multiplosDeTres {
			fmt.Println("numero: PIN")
		} else if multiplosDeCinco {
			fmt.Println("numero: PAN")
		} else {
			fmt.Println("numero: ",i)
		}
}

}
