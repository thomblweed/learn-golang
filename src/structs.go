package car

import "fmt"

func MyCar() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "Ford",
		Model: "Escort",
	}

	fmt.Println(myCar.Make)
	fmt.Println(myCar.Model)
}
