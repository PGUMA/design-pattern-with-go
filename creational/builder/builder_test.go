package builder

import (
	"fmt"
)

func ExampleDartsBuilder() {
	bpDarts := NewDartsWithBP("H2 REVIVE").
	WithFlight(Teardrop).
	WithShaft(MM22_5).
	WithTip(Normal).
	Build()

	fmt.Println("barrel=", bpDarts.Barrel)
	fmt.Println("flight=", bpDarts.Flight)
	fmt.Println("shaft=", bpDarts.Shaft)
	fmt.Println("tip=", bpDarts.Tip)

	// Output:
	// barrel= H2 REVIVE
	// flight= 3
	// shaft= 1
	// tip= 0
}