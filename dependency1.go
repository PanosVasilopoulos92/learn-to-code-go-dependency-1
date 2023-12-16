package dependency1

import (
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
