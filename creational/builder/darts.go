package builder

type Flight int

const (
	Standard = iota
    Shape
	Slim
	Teardrop
)

type Shaft int

const (
	// 8Flightの規格
	MM19_0 = iota
	MM22_5
	MM26_0
	MM29_5
	MM33_0
	MM44_0
)

type Tip int

const (
	Normal = iota
	Short
)

type Darts struct {
	Barrel string
	Flight Flight
    Shaft Shaft
	Tip Tip
}

