// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package blockRange

import (
	"fmt"

	"encoding/json"
	"testing"
)

func TestPointToPointTypeBlock(t *testing.T) {
	point := &Point{Block: 100}
	result := getPointType(point)

	if result != BlockRangeBlockNumber {
		t.Error("Bad point type returned")
	}
}

func TestPointToPointTypeSpecial(t *testing.T) {
	point := &Point{Special: "london"}
	result := getPointType(point)

	if result != BlockRangeSpecial {
		t.Error("Bad point type returned")
	}
}

func TestPointToPointTypeDate(t *testing.T) {
	point := &Point{Date: "2021-10-03"}
	result := getPointType(point)

	if result != BlockRangeDate {
		t.Error("Bad point type returned")
	}
}

func TestModifierToModifierTypeStep(t *testing.T) {
	modifier := &Modifier{Step: 15}
	result := getModifierType(modifier)

	if result != BlockRangeStep {
		t.Error("Bad modifier type returned")
	}
}

func TestModifierToModifierTypePeriod(t *testing.T) {
	modifier := &Modifier{Period: "daily"}
	result := getModifierType(modifier)

	if result != BlockRangePeriod {
		t.Error("Bad modifier type returned")
	}
}

func TestNewBlocks(t *testing.T) {
	blockRange, err := New("10-1000:10")
	if err != nil {
		t.Error(err)
	}

	if blockRange.StartType != BlockRangeBlockNumber {
		t.Error("StartType is not block number")
	}

	if blockRange.Start.Block != 10 {
		t.Errorf("Wrong start")
	}

	if blockRange.EndType != BlockRangeBlockNumber {
		t.Error("EndType is not block number")
	}

	if blockRange.End.Block != 1000 {
		t.Error("Wrong end")
	}

	if blockRange.ModifierType != BlockRangeStep {
		t.Error("ModifierType is not step")
	}

	if blockRange.Modifier.Step != 10 {
		t.Error("Wrong modifier")
	}
}

func TestNewSpecial(t *testing.T) {
	blockRange, err := New("london:weekly")

	if err != nil {
		t.Error(err)
	}

	if blockRange.StartType != BlockRangeSpecial {
		t.Error("StartType is not special")
	}

	if blockRange.Start.Special != "london" {
		t.Errorf("Wrong start")
	}

	if blockRange.EndType != BlockRangeNotDefined {
		t.Error("EndType is not notdefined")
	}

	if blockRange.ModifierType != BlockRangePeriod {
		t.Error("ModifierType is not period")
	}

	if blockRange.Modifier.Period != "weekly" {
		t.Error("Wrong modifier")
	}
}

func TestHandleParserErrors(t *testing.T) {
	_, modifierErr := New("10-100:biweekly")

	if me, ok := modifierErr.(*WrongModifierError); ok {
		if me.Token != "biweekly" {
			t.Errorf("Wrong token: %s", me.Token)
		}
	} else {
		t.Error("Returned error is not WrongModifier")
		t.Error(modifierErr)
	}
}

func TestBlockRange_UnmarshalJSON(t *testing.T) {
	type SomeRecord struct {
		Blocks BlockRange `json:"blocks"`
	}

	var record SomeRecord
	source := []byte(`{"blocks":"000000000-10567003"}`)

	err := json.Unmarshal(source, &record)
	if err != nil {
		t.Error(err)
	}

	if record.Blocks.StartType != BlockRangeBlockNumber {
		t.Errorf("Wrong StartType %d", record.Blocks.StartType)
	}

	if record.Blocks.EndType != BlockRangeBlockNumber {
		t.Errorf("Wrong EndType %d", record.Blocks.EndType)
	}

	if record.Blocks.Start.Block != uint(0) {
		t.Error("Wrong start value")
	}

	if record.Blocks.End.Block != uint(10567003) {
		t.Errorf("Wrong end value %d", record.Blocks.End.Block)
	}
}

func TestToString(t *testing.T) {
	br, err := New("1234")
	if err != nil {
		t.Errorf("Could not parse block")
	}
	expected := "{\"StartType\":0,\"Start\":{\"Block\":1234,\"Date\":\"\",\"Special\":\"\"},\"EndType\":5,\"End\":{\"Block\":0,\"Date\":\"\",\"Special\":\"\"},\"ModifierType\":5,\"Modifier\":{\"Step\":0,\"Period\":\"\"}}\n"
	got := fmt.Sprintf("%s\n", br.MarshalJSON())
	if got != expected {
		t.Errorf("String printer for blockRange not equal to expected")
	}
}
