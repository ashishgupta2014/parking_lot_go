package vehicle

type Base struct{
	reg_no string
	color string
	vehicle_type VehicleType
}

func (base *Base) GetReg()string{
	return base.reg_no
}

func (base *Base) GetColor()string{
	return base.color
}

func (base *Base) GetVehicleType()string{
	return base.vehicle_type.FromVehicleType()
}

func Create(vehicle_type string, reg_no string, color string)(*Base, error){
	var vt, err = FromString(vehicle_type)
	if err == nil {
		return &Base{vehicle_type: vt, reg_no:reg_no, color: color}, nil
	}
	return nil, err
}