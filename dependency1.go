package dependency1

import (
	"fmt"

	dependancy2 "github.com/PanosVasilopoulos92/learn-to-code-code-dependency-2"
)

func SayWelcome() string {
	return "Welcome!"
}

func SayGoodBye() string {
	return "Goodbye!"
}

func Leaving() string {
	return dependancy2.WhenYouLeave(SayGoodBye())
}

func FromVersion1_1_0() {
	fmt.Println("I' m from version 1.1.0")
}

func FromVersion1_2_0() {
	fmt.Println("I' m from version 1.2.0")
}
