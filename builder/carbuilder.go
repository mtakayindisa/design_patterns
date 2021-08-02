package builder

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuilderProcess{
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuilderProcess {
	c.v.Seats = 4
	return c
}

func (c *CarBuilder) SetStructure() BuilderProcess {
	c.v.Structure = "Toyota Hilux"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
