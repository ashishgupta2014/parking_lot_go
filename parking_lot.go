package main

import (
	"errors"
	area "github.com/ashishgupta2014/parking_lot_go/parking_lot/slots"
	heap "github.com/ashishgupta2014/parking_lot_go/parking_lot/minheap"
	vehicle "github.com/ashishgupta2014/parking_lot_go/parking_lot/vehicle"
)

type ParkingLot struct{
	prefix string
	spaces map[string]*area.FloorSlots
	space_min_heap map[string]*heap.StringMinHeap
}

func compute_parking_lot_area(floors int, slots int, prefix string) map[string]*area.FloorSlots{
	spaces := make(map[string]*area.FloorSlots, floors)
	for floor:=0; floor<floors; floor++{
		for slot:=0; slot<slots; slot++{
			space := area.Create(floor, slot, prefix)
			spaces[space.GetSlotId()] = space
		}
	}
	return spaces
}

func create(prefix string, floors int, slots int)  *ParkingLot{
	parking_spaces := compute_parking_lot_area(floors, slots, prefix)	
	vehcile_spaces := area.VehicleSpaceFilter(&parking_spaces, floors, slots)
	return &ParkingLot{prefix: prefix, 
		spaces: parking_spaces,
		space_min_heap: vehcile_spaces,
	}
}

func (parking_lot *ParkingLot) park(vehicle_type string, reg_no string, color string) (string, error){
	vh, err := vehicle.Create(vehicle_type, reg_no, color)
	if err == nil {
		minheap := parking_lot.space_min_heap[vehicle_type]
		if minheap.GetHeapSize() > 0{
			slot_id := minheap.Remove()
			slot := parking_lot.spaces[slot_id]
			slot.Park(vh)
			return slot_id, nil
		}
		return "", errors.New("Parking Space not avilable")
	}
	return "", err
}

func (parking_lot *ParkingLot) unpark(slot_id string)(string, error){
	slot := parking_lot.spaces[slot_id]
	if slot == nil{
		return "", errors.New("Invliad Parking ID")
	}
	if !slot.IsOccpied(){
		vehicle_type := slot.GetParkedVehicleType()
		slot.UnPark(slot_id)
		minheap := parking_lot.space_min_heap[vehicle_type]
		minheap.Insert(slot_id)
		return slot_id, nil
	}
	
	return "", errors.New("Parking Space is Free")
}

func (parking_lot *ParkingLot) free_count(vehicle_type string)int{
	return parking_lot.space_min_heap[vehicle_type].GetHeapSize()
	
}

func (parking_lot *ParkingLot) free_slots(vehicle_type string)[]string{
	spaces := parking_lot.spaces
	arr := make([]string, len(spaces))
	i := 0
	for slot_id, slot := range spaces{
		if slot.GetReservedVehicleType() == vehicle_type  && slot.IsOccpied(){
			arr[i] = slot_id
			i++
		}
	}
	return arr
}

func (parking_lot *ParkingLot) occupied_slots(vehicle_type string)[]string{
	spaces := parking_lot.spaces
	arr := make([]string, len(spaces))
	i := 0
	for slot_id, slot := range spaces{
		if slot.GetReservedVehicleType() == vehicle_type  && !slot.IsOccpied(){
			arr[i] = slot_id
			i++
		}
	}
	return arr
}