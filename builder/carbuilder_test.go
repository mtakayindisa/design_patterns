package builder

import "testing"

func TestCarBuilder_GetVehicle(t *testing.T) {
	manufacturingDirector := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingDirector.SetBuilder(carBuilder)
	manufacturingDirector.Construct()

	car := carBuilder.GetVehicle()
	wheels := car.Wheels
	if wheels != 4 {
		t.Errorf("Car wheels should be 4, but found %d", wheels)
	}
}
