package json

import "testing"

const jsonExampleAbsMachinePos string = `{"r":{"mpo":{"x":3.000,"y":0.000,"z":0.000,"a":0.000,"b":0.000,"c":0.000}},"f":[3,0,10]}`

func TestParseResponse(t *testing.T) {
	dut, err := ParseResponse([]byte(jsonExampleAbsMachinePos))
	if err != nil {
		t.Error(err)
		t.Error("Parse failed!")
		t.FailNow()
	}
	if dut.ResponseData.AbsoluteMachinePosition == nil {
		t.Error("TestParseResponse>jsonExampleAbsMachinePos>==nil")
		t.FailNow()
	} else {
		if dut.ResponseData.AbsoluteMachinePosition.X != 3.0 {
			t.Error("TestParseResponse>jsonExampleAbsMachinePos>X!=3")
			t.FailNow()
		}
	}
}

func TestUpdateFrom(t *testing.T) {
	dst := TResponse{}
	src, _ := ParseResponse([]byte(jsonExampleAbsMachinePos))
	dst.UpdateFrom(src)
	if dst.ResponseData.AbsoluteMachinePosition == nil {
		t.Error("AbsMachPos is nil")
		t.FailNow()
	} else {
		if dst.ResponseData.AbsoluteMachinePosition.X != 3.0 {
			t.Error("AbsMachPos != 3")
			t.FailNow()
		}
	}
}
