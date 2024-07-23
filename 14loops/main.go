package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	//     fmt.Println(days[d])
	// }

	// for i := range days {
	//     fmt.Println(days[i])
	// }

	// for _, day := range days {
	//     fmt.Printf("index is v and value is %v\n", day)
	// }

	routeValue := 1

	for routeValue < 10 {

		if routeValue == 2 {
			goto lco
		}

		if routeValue == 5 {
			routeValue++
			continue
		}

		fmt.Println("Value is: ", routeValue)
		routeValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")
}
