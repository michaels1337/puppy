package puppy

import "github.com/michaels1337/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func DoJump() string {
	return "JUMP"

}
func Run() string {
	return "Start running"

}
