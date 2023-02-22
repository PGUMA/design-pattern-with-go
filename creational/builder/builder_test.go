package builder

import "testing"

func TestBuilderNew(t *testing.T) {
	
	bpDarts := NewDartsWithBP("H2 REVIVE").
	WithFlight(Teardrop).
	WithShaft(MM22_5).
	WithTip(Normal).
	Build()

	if bpDarts.Barrel != "H2 REVIVE" {
		t.Fatal("ダーツの生成に失敗しました。バレルの値が期待値と異なります")
	}
	if bpDarts.Flight != Teardrop {
		t.Fatal("ダーツの生成に失敗しました。フライトの値が期待値と異なります")
	}
	if bpDarts.Shaft != MM22_5 {
		t.Fatal("ダーツの生成に失敗しました。シャフトの値が期待値と異なります")
	}
	if bpDarts.Tip != Normal {
		t.Fatal("ダーツの生成に失敗しました。チップの値が期待値と異なります")
	}
	
}