package utils

import (
	"fmt"
	"testing"
)

func Test_Trim(t *testing.T) {
	expected := "aty"
	value_for_test := fmt.Sprintf("     %s     ", expected)
	trim_string := Trim(value_for_test)

	if expected != trim_string {
		t.Errorf("A função Trim está retornando: %s != %s", expected, trim_string)
	}
}

func Test_ToLower(t *testing.T) {
	expected := "aty"
	value_for_test := "ATY"
	tolower_string := ToLower(value_for_test)

	if expected != tolower_string {
		t.Errorf("A função ToLower está retornando: %s != %s", expected, tolower_string)
	}
}

func Test_CapitalizeSentence(t *testing.T) {
	expected := "Tainan Felipe Dos Santos"
	value_for_test := "tainan felipe dos santos"
	capitalizeSentence_string := CapitalizeSentence(value_for_test)

	if expected != capitalizeSentence_string {
		t.Errorf("A função CapitalizeSentence está retornando: %s != %s", expected, capitalizeSentence_string)
	}
}
