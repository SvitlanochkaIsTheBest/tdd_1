package main

import (
	"testing"
)

func TestToRoman_One(t *testing.T) {
	result, err := toRoman(1)
	expected := "I"
	if err != nil {
		t.Errorf("toRoman(1) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(1) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Two(t *testing.T) {
	result, err := toRoman(2)
	expected := "II"
	if err != nil {
		t.Errorf("toRoman(2) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(2) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Three(t *testing.T) {
	result, err := toRoman(3)
	expected := "III"
	if err != nil {
		t.Errorf("toRoman(3) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(3) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Four(t *testing.T) {
	result, err := toRoman(4)
	expected := "IV"
	if err != nil {
		t.Errorf("toRoman(4) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(4) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Five(t *testing.T) {
	result, err := toRoman(5)
	expected := "V"
	if err != nil {
		t.Errorf("toRoman(5) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(5) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Six(t *testing.T) {
	result, err := toRoman(6)
	expected := "VI"
	if err != nil {
		t.Errorf("toRoman(6) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(6) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Seven(t *testing.T) {
	result, err := toRoman(7)
	expected := "VII"
	if err != nil {
		t.Errorf("toRoman(7) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(7) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Eight(t *testing.T) {
	result, err := toRoman(8)
	expected := "VIII"
	if err != nil {
		t.Errorf("toRoman(8) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(8) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Nine(t *testing.T) {
	result, err := toRoman(9)
	expected := "IX"
	if err != nil {
		t.Errorf("toRoman(9) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(9) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Ten(t *testing.T) {
	result, err := toRoman(10)
	expected := "X"
	if err != nil {
		t.Errorf("toRoman(10) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(10) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Eleven(t *testing.T) {
	result, err := toRoman(11)
	expected := "XI"
	if err != nil {
		t.Errorf("toRoman(11) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(11) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Twelve(t *testing.T) {
	result, err := toRoman(12)
	expected := "XII"
	if err != nil {
		t.Errorf("toRoman(12) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(12) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Thirteen(t *testing.T) {
	result, err := toRoman(13)
	expected := "XIII"
	if err != nil {
		t.Errorf("toRoman(13) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(13) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Fourteen(t *testing.T) {
	result, err := toRoman(14)
	expected := "XIV"
	if err != nil {
		t.Errorf("toRoman(14) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(14) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Fifteen(t *testing.T) {
	result, err := toRoman(15)
	expected := "XV"
	if err != nil {
		t.Errorf("toRoman(15) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(15) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Sixteen(t *testing.T) {
	result, err := toRoman(16)
	expected := "XVI"
	if err != nil {
		t.Errorf("toRoman(16) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(16) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Seventeen(t *testing.T) {
	result, err := toRoman(17)
	expected := "XVII"
	if err != nil {
		t.Errorf("toRoman(17) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(17) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Eighteen(t *testing.T) {
	result, err := toRoman(18)
	expected := "XVIII"
	if err != nil {
		t.Errorf("toRoman(18) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(18) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Nineteen(t *testing.T) {
	result, err := toRoman(19)
	expected := "XIX"
	if err != nil {
		t.Errorf("toRoman(19) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(19) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Twenty(t *testing.T) {
	result, err := toRoman(20)
	expected := "XX"
	if err != nil {
		t.Errorf("toRoman(20) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(20) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentyOne(t *testing.T) {
	result, err := toRoman(21)
	expected := "XXII" // XXI
	if err != nil {
		t.Errorf("toRoman(21) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(21) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentyTwo(t *testing.T) {
	result, err := toRoman(22)
	expected := "XXII"
	if err != nil {
		t.Errorf("toRoman(22) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(22) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentyThree(t *testing.T) {
	result, err := toRoman(23)
	expected := "XXIII"
	if err != nil {
		t.Errorf("toRoman(23) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(23) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentyFour(t *testing.T) {
	result, err := toRoman(24)
	expected := "XXIV"
	if err != nil {
		t.Errorf("toRoman(24) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(24) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentyFive(t *testing.T) {
	result, err := toRoman(25)
	expected := "XXV"
	if err != nil {
		t.Errorf("toRoman(25) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(25) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentySix(t *testing.T) {
	result, err := toRoman(26)
	expected := "XXVI"
	if err != nil {
		t.Errorf("toRoman(26) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(26) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentySeven(t *testing.T) {
	result, err := toRoman(27)
	expected := "XXVII"
	if err != nil {
		t.Errorf("toRoman(27) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(27) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentyEight(t *testing.T) {
	result, err := toRoman(28)
	expected := "XXVIII"
	if err != nil {
		t.Errorf("toRoman(28) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(28) == %s, expected %s", result, expected)
	}
}

func TestToRoman_TwentyNine(t *testing.T) {
	result, err := toRoman(29)
	expected := "XXIX"
	if err != nil {
		t.Errorf("toRoman(29) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(29) == %s, expected %s", result, expected)
	}
}

func TestToRoman_Thirty(t *testing.T) {
	result, err := toRoman(30)
	expected := "XXX"
	if err != nil {
		t.Errorf("toRoman(30) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(30) == %s, expected %s", result, expected)
	}
}

func TestToRoman_ThirtyOne(t *testing.T) {
	result, err := toRoman(31)
	expected := "XXXI"
	if err != nil {
		t.Errorf("toRoman(31) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(31) == %s, expected %s", result, expected)
	}
}

func TestToRoman_ThirtyTwo(t *testing.T) {
	result, err := toRoman(32)
	expected := "XXXII"
	if err != nil {
		t.Errorf("toRoman(32) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(32) == %s, expected %s", result, expected)
	}
}

func TestToRoman_ThirtyThree(t *testing.T) {
	result, err := toRoman(33)
	expected := "XXXIII"
	if err != nil {
		t.Errorf("toRoman(33) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(33) == %s, expected %s", result, expected)
	}
}

func TestToRoman_ThirtyFour(t *testing.T) {
	result, err := toRoman(34)
	expected := "XXXIV"
	if err != nil {
		t.Errorf("toRoman(34) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(34) == %s, expected %s", result, expected)
	}
}

func TestToRoman_ThirtyFive(t *testing.T) {
	result, err := toRoman(35)
	expected := "XXXV"
	if err != nil {
		t.Errorf("toRoman(35) returned error: %s", err)
	} else if result != expected {
		t.Errorf("toRoman(35) == %s, expected %s", result, expected)
	}
}

func TestToRoman_ThirtySix(t *testing.T) {
	num := 2058
	result, err := toRoman(num)
	expected := "MMLVIII"
	if err != nil {
		t.Errorf("toRoman(%d) returned error: %s", num, err)
	} else if result != expected {
		t.Errorf("toRoman(%d) == %s, expected %s", num, result, expected)
	}
}

func TestToRoman_ThirtySeven(t *testing.T) {
	num := 1493
	result, err := toRoman(num)
	expected := "MCDXCIII"
	if err != nil {
		t.Errorf("toRoman(%d) returned error: %s", num, err)
	} else if result != expected {
		t.Errorf("toRoman(%d) == %s, expected %s", num, result, expected)
	}
}

func TestToRoman_ThirtyEight(t *testing.T) {
	num := 3844
	result, err := toRoman(num)
	expected := "MMMDCCCXLIV"
	if err != nil {
		t.Errorf("toRoman(%d) returned error: %s", num, err)
	} else if result != expected {
		t.Errorf("toRoman(%d) == %s, expected %s", num, result, expected)
	}
}

func TestToRoman_ThirtyNine(t *testing.T) {
	num := 1562
	result, err := toRoman(num)
	expected := "MDLXII"
	if err != nil {
		t.Errorf("toRoman(%d) returned error: %s", num, err)
	} else if result != expected {
		t.Errorf("toRoman(%d) == %s, expected %s", num, result, expected)
	}
}

func TestToRoman_Forty(t *testing.T) {
	num := 4000 // More than 3999
	result, err := toRoman(num)
	expected := "MMMM"
	if err != nil {
		t.Errorf("toRoman(%d) returned error: %s", num, err)
	} else if result != expected {
		t.Errorf("toRoman(%d) == %s, expected %s", num, result, expected)
	}
}
