package orc

import (
	"bytes"
	"encoding/json"
	"math/big"
	"testing"

	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func TestDecimal(t *testing.T) {

	d := Decimal{big.NewInt(-8361232), 4}

	if v := d.Float64(); v != -836.1232 {
		t.Errorf("Test failed, expected -836.1232 got %v", v)
	}

	if v := d.Float32(); v != -836.1232 {
		t.Errorf("Test failed, expected -836.1232 got %v", v)
	}

	byt, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	expected := []byte(`-836.1232`)
	if !bytes.Equal(byt, expected) {
		t.Errorf("Test failed, expected %s got %s", expected, byt)
	}

}

func TestDecimalMarshalBSONValue(t *testing.T) {
	d := Decimal{big.NewInt(-8361232), 4}
	btype, byt, err := d.MarshalBSONValue()
	if err != nil {
		t.Errorf("Test failed, expected no error but got %v", err)
	}

	expected := "{\"$numberDecimal\":\"-836.1232\"}"
	actual := bsoncore.Value{Type: btype, Data: byt}.String()
	if expected != actual {
		t.Errorf("Test failed, expected %s but got %s", expected, actual)
	}
}
