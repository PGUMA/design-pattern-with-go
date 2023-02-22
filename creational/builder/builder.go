package builder

type DartsBuilder interface {
    WithFlight(f Flight) DartsBuilder
	WithShaft(s Shaft) DartsBuilder
	WithTip(t Tip) DartsBuilder
    Build() *Darts
}

type dartsBuilder struct {
    barrel string
    subFlight Flight
    subShaft Shaft
    subTips Tip
}

func (b *dartsBuilder) WithFlight(f Flight) DartsBuilder {
     b.subFlight = f
	 return b
}
func (b *dartsBuilder) WithShaft(s Shaft) DartsBuilder {
    b.subShaft = s
	return b
}
func (b *dartsBuilder) WithTip(t Tip) DartsBuilder {
    b.subTips = t
	return b
}

func (b *dartsBuilder) Build() *Darts {
	return &Darts{
        Barrel: b.barrel,
        Flight: b.subFlight,
        Shaft: b.subShaft,
        Tip: b.subTips,
    }
}

func NewDartsWithBP(barrel string) DartsBuilder {
    return &dartsBuilder{
        barrel: barrel,
    }
}