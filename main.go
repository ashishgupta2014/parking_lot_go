package main

import (
	//"fmt"
	"github.com/ashishgupta2014/parking_lot_go/parking_lot/helper"
)

func main()  {
	content := helper.InputReader("public/input_1.txt")
	CommandExecutor(content)
	// parking_lot := create("pl_1234", 2, 6)
	// fmt.Println("Free Spaces: ",parking_lot.free_space_count())
	// slot_id := parking_lot.park("TRUCK", "abc", "black")
	// fmt.Println("Parked Spaces: ",parking_lot.parked_space_count())
	// parking_lot.unpark(slot_id)
	// fmt.Println("Free Spaces: ",parking_lot.free_space_count())
}

