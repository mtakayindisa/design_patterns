package builder

type BuilderProcess interface {
	SetWheels() BuilderProcess
	SetSeats() BuilderProcess
	SetStructure() BuilderProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector Director
type ManufacturingDirector struct {
	builder BuilderProcess
}

func (f *ManufacturingDirector) Construct(){
	f.builder.SetSeats().SetWheels().SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuilderProcess){
	f.builder = b;
}

//Product
type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}