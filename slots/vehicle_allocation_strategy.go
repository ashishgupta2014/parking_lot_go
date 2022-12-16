package slots

import (
	"github.com/ashishgupta2014/parking_lot_go/parking_lot/vehicle"
	heap "github.com/ashishgupta2014/parking_lot_go/parking_lot/minheap"
)

func reserved_vehicle_type_slot(slot int)  vehicle.VehicleType{
	var vehicle_type vehicle.VehicleType
	switch {
	case slot == 0:
		vehicle_type, _ = vehicle.FromString("TRUCK")
	case slot == 1 || slot == 2:
		vehicle_type, _ = vehicle.FromString("BIKE")
	default:
		vehicle_type, _ = vehicle.FromString("CAR")
	}
	
	return vehicle_type
}

func max_reserved_bike_space(floors int)  int{
	return floors*2
}

func max_reserved_truck_space(floors int)  int{
	return floors*1
}

func max_reserved_car_space(floors, slots int)  int{
	return floors*slots - max_reserved_bike_space(floors) - max_reserved_truck_space(floors)
}

func VehicleSpaceFilter(spaces *map[string]*FloorSlots, floors, slots int)  map[string]*heap.StringMinHeap{
	result := make(map[string]*heap.StringMinHeap, 3)
	result["BIKE"] = vehicle_filter("BIKE", max_reserved_bike_space(floors), spaces)
	result["CAR"] = vehicle_filter("CAR", max_reserved_car_space(floors, slots), spaces)
	result["TRUCK"] = vehicle_filter("TRUCK", max_reserved_truck_space(floors), spaces)
	return result
}

func vehicle_filter(vehicle_name string, size int, spaces *map[string]*FloorSlots)*heap.StringMinHeap{
	var vehicle_type vehicle.VehicleType
	minHeap := heap.NewMinHeap(size)
	vehicle_type, _ = vehicle.FromString(vehicle_name)
	for key, value := range *spaces{
		if value.reserved_vehicle_type == vehicle_type{
			minHeap.Insert(key)
		}
	} 
	return minHeap
}