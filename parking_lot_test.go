package main

import (
    "testing"
    "errors"
    "github.com/stretchr/testify/assert"
)
 
func TestParkingLotParkBus(t *testing.T) {
    name := "ParkingLotTest"
    floors := 2
    slots := 6
    parking_lot := create(name, floors, slots)

    vehicle_type := "BUS"
    reg_no := "RG1234"
    color := "White"
    _, actual_err := parking_lot.park(vehicle_type, reg_no, color)
    expected_err := errors.New("unknown vehicle type "+ vehicle_type)
    assert.Equal(t, actual_err, expected_err)

    _, actual_err = parking_lot.unpark("ParkingLotTest_0_0")
    expected_err = errors.New("Parking Space is Free")
    assert.Equal(t, actual_err, expected_err)
    
}


func TestParkingLotParkCar(t *testing.T) {
    name := "ParkingLotTest"
    floors := 2
    slots := 6
    parking_lot := create(name, floors, slots)

    vehicle_type := "CAR"
    reg_no := "RG1234"
    color := "White"
    slot_id, _ := parking_lot.park(vehicle_type, reg_no, color)
    expected_slot_id := "ParkingLotTest_0_3"
    assert.Equal(t, slot_id, expected_slot_id)

    unpark_slot_id, _ := parking_lot.unpark(expected_slot_id)
    assert.Equal(t, unpark_slot_id, expected_slot_id)

    _, actual_err := parking_lot.unpark("ParkingLotTest_3_0")
    expected_err := errors.New("Invliad Parking ID")
    assert.Equal(t, actual_err, expected_err)
    
}


func TestParkingLotParkBike(t *testing.T) {
    name := "ParkingLotTest"
    floors := 2
    slots := 6
    parking_lot := create(name, floors, slots)

    vehicle_type := "BIKE"
    reg_no := "RG1234"
    color := "White"
    slot_id, _ := parking_lot.park(vehicle_type, reg_no, color)
    expected_slot_id := "ParkingLotTest_0_1"
    assert.Equal(t, slot_id, expected_slot_id)

    unpark_slot_id, _ := parking_lot.unpark(expected_slot_id)
    assert.Equal(t, unpark_slot_id, expected_slot_id)
    
}


func TestParkingLotParkTruck(t *testing.T) {
    name := "ParkingLotTest"
    floors := 2
    slots := 6
    parking_lot := create(name, floors, slots)

    vehicle_type := "TRUCK"
    reg_no := "RG1234"
    color := "White"
    slot_id, _ := parking_lot.park(vehicle_type, reg_no, color)
    expected_slot_id := "ParkingLotTest_0_0"
    assert.Equal(t, slot_id, expected_slot_id)

    unpark_slot_id, _ := parking_lot.unpark(expected_slot_id)
    assert.Equal(t, unpark_slot_id, expected_slot_id)
    
}


func TestParkingLotParkingFree(t *testing.T) {
    var vehicle_type, reg_no, color string
    name := "ParkingLotTest"
    floors := 2
    slots := 6
    parking_lot := create(name, floors, slots)

    vehicle_type = "TRUCK"
    reg_no = "RG1234"
    color = "White"
    _, _ = parking_lot.park(vehicle_type, reg_no, color)

    vehicle_type = "TRUCK"
    reg_no = "RG5678"
    color = "White"
    _, _ = parking_lot.park(vehicle_type, reg_no, color)

    vehicle_type = "TRUCK"
    reg_no = "RG1376"
    color = "White"
    _, actual_err := parking_lot.park(vehicle_type, reg_no, color)
    expected_err := errors.New("Parking Space not avilable")
    assert.Equal(t, actual_err, expected_err)
    
    
}