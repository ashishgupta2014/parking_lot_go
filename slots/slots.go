package slots
import (
	"github.com/ashishgupta2014/parking_lot_go/parking_lot/vehicle"
	"fmt"
)

type FloorSlots struct{
	floor int
	slot int
	slot_id string
	reserved_vehicle_type vehicle.VehicleType
	vehicle *vehicle.Base
}

func Create(floor int, slot int, prefix string) *FloorSlots{
	return &FloorSlots{
		floor: floor,
		slot: slot,
		slot_id: fmt.Sprintf("%s_%d_%d", prefix, floor, slot),
		reserved_vehicle_type: reserved_vehicle_type_slot(slot),
	}
}

func (floor_slot *FloorSlots)  IsOccpied()bool{
	return floor_slot.vehicle == nil
	
}

func (floor_slot *FloorSlots) GetSlotId() string{
	return floor_slot.slot_id
}

func (floor_slot *FloorSlots) GetParkedVehicleType()string{
	if !floor_slot.IsOccpied(){
		return floor_slot.vehicle.GetVehicleType()
	}
	return ""
}

func (floor_slot *FloorSlots) Park(vehicle *vehicle.Base){
	if floor_slot.IsOccpied(){
		floor_slot.vehicle = vehicle
	}
}

func (floor_slot *FloorSlots) UnPark(slot_id string){
	if !floor_slot.IsOccpied(){
		floor_slot.vehicle = nil
	}
}

func (floor_slot *FloorSlots) GetReservedVehicleType() string{
	return floor_slot.reserved_vehicle_type.FromVehicleType()
}