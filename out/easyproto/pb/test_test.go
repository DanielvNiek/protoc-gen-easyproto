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
		Tags: map[string]string{
			"key1": "val1",
			"key2": "",     // test empty string
			"":     "val3", // test empty key
		},
		EnumMap: map[int32]TestEnum{
			1: TestEnum_TEST_ENUM_VALUE_1,
			2: TestEnum_TEST_ENUM_UNSPECIFIED, // test default enum
			0: TestEnum_TEST_ENUM_VALUE_2,     // test default key
		},
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
