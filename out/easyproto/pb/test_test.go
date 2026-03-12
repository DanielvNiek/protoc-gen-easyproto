package pb

import (
	"reflect"
	"testing"
)

func normalizeSlices(m *Test) {
	if m.Nicknames == nil {
		m.Nicknames = []string{}
	}
	if m.Statuses == nil {
		m.Statuses = []TestEnum{}
	}
}

func TestMarshalUnmarshal(t *testing.T) {
	msg := &Test{
		Name:      "test_name",
		Age:       42,
		Nicknames: []string{"nick1", "nick2"},
		Status:    TestEnum_TEST_ENUM_VALUE_1,
		Statuses:  []TestEnum{TestEnum_TEST_ENUM_UNSPECIFIED, TestEnum_TEST_ENUM_VALUE_2},
	}

	data := msg.MarshalProtobuf(nil)
	if len(data) == 0 {
		t.Fatal("MarshalProtobuf returned empty data")
	}

	var msg2 Test
	if err := msg2.UnmarshalProtobuf(data); err != nil {
		t.Fatalf("UnmarshalProtobuf failed: %v", err)
	}

	normalizeSlices(msg)
	normalizeSlices(&msg2)

	if !reflect.DeepEqual(msg, &msg2) {
		t.Errorf("Unmarshaled message does not match original.\nOriginal: %+v\nUnmarshaled: %+v", msg, &msg2)
	}
}

func TestMarshalUnmarshalEmpty(t *testing.T) {
	msg := &Test{}

	data := msg.MarshalProtobuf(nil)
	// Empty message should marshal to empty bytes in proto3
	if len(data) != 0 {
		t.Errorf("Expected empty data for empty message, got %d bytes", len(data))
	}

	var msg2 Test
	// Put some garbage data in msg2 to ensure it gets cleared
	msg2.Name = "garbage"
	msg2.Age = 100
	msg2.Nicknames = []string{"garbage"}
	msg2.Status = TestEnum_TEST_ENUM_VALUE_2
	msg2.Statuses = []TestEnum{TestEnum_TEST_ENUM_VALUE_1}

	if err := msg2.UnmarshalProtobuf(data); err != nil {
		t.Fatalf("UnmarshalProtobuf failed: %v", err)
	}

	normalizeSlices(msg)
	normalizeSlices(&msg2)

	if !reflect.DeepEqual(msg, &msg2) {
		t.Errorf("Unmarshaled message does not match original.\nOriginal: %+v\nUnmarshaled: %+v", msg, &msg2)
	}
}

func TestMarshalUnmarshalPartial(t *testing.T) {
	msg := &Test{
		Name: "partial",
		// Age: 0,
		// Nicknames: nil,
		Status: TestEnum_TEST_ENUM_VALUE_2,
		// Statuses: nil,
	}

	data := msg.MarshalProtobuf(nil)
	if len(data) == 0 {
		t.Fatal("MarshalProtobuf returned empty data")
	}

	var msg2 Test
	if err := msg2.UnmarshalProtobuf(data); err != nil {
		t.Fatalf("UnmarshalProtobuf failed: %v", err)
	}

	normalizeSlices(msg)
	normalizeSlices(&msg2)

	if !reflect.DeepEqual(msg, &msg2) {
		t.Errorf("Unmarshaled message does not match original.\nOriginal: %+v\nUnmarshaled: %+v", msg, &msg2)
	}
}
