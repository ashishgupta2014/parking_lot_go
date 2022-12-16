package main

import (
	"strings"
	"fmt"
	"strconv"
)

func CommandExecutor(commands string){
	arr := strings.Split(string(commands), "\n")
	var parking_lot *ParkingLot
	for index := range arr{
		fmt.Println(arr[index])
		cmd := strings.Split(arr[index], " ")
		if cmd[0] == "create_parking_lot"{
			floors, _ := strconv.Atoi(cmd[2])
			slots, _ :=  strconv.Atoi(cmd[3])
			parking_lot = create(cmd[1], floors, slots)
		}else if cmd[0] == "display" && cmd[1] == "free_count"{
			fmt.Println(parking_lot.free_count(cmd[2]))
		}else if cmd[0] == "display" && cmd[1] == "free_slots"{
			fmt.Println(parking_lot.free_slots(cmd[2]))
		}else if cmd[0] == "display" && cmd[1] == "occupied_slots"{
			fmt.Println(parking_lot.occupied_slots(cmd[2]))
		}else if cmd[0] == "park_vehicle"{
			fmt.Println(parking_lot.park(cmd[1], cmd[2], cmd[3]))
		}else if cmd[0] == "unpark_vehicle"{
			fmt.Println(parking_lot.unpark(cmd[1]))
		} 
	}
}