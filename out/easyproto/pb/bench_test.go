package pb

import (
	"testing"

	gopb "github.com/DanielvNiek/protoc-gen-easyproto/goout/easyproto/pb"
	"google.golang.org/protobuf/proto"
)

func BenchmarkMarshalEasyProto(b *testing.B) {
	msg := &Test{
		Name:      "test_name",
		Age:       42,
		Nicknames: []string{"nick1", "nick2"},
		Status:    TestEnum_TEST_ENUM_VALUE_1,
		Statuses:  []TestEnum{TestEnum_TEST_ENUM_UNSPECIFIED, TestEnum_TEST_ENUM_VALUE_2},
		Tags: map[string]string{
			"key1": "val1",
			"key2": "",
			"":     "val3",
		},
		EnumMap: map[int32]TestEnum{
			1: TestEnum_TEST_ENUM_VALUE_1,
			2: TestEnum_TEST_ENUM_UNSPECIFIED,
			0: TestEnum_TEST_ENUM_VALUE_2,
		},
	}
	b.ReportAllocs()
	b.ResetTimer()
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf = msg.MarshalProtobuf(buf[:0])
	}
}

func BenchmarkUnmarshalEasyProto(b *testing.B) {
	msg := &Test{
		Name:      "test_name",
		Age:       42,
		Nicknames: []string{"nick1", "nick2"},
		Status:    TestEnum_TEST_ENUM_VALUE_1,
		Statuses:  []TestEnum{TestEnum_TEST_ENUM_UNSPECIFIED, TestEnum_TEST_ENUM_VALUE_2},
		Tags: map[string]string{
			"key1": "val1",
			"key2": "",
			"":     "val3",
		},
		EnumMap: map[int32]TestEnum{
			1: TestEnum_TEST_ENUM_VALUE_1,
			2: TestEnum_TEST_ENUM_UNSPECIFIED,
			0: TestEnum_TEST_ENUM_VALUE_2,
		},
	}
	data := msg.MarshalProtobuf(nil)
	b.ReportAllocs()
	b.ResetTimer()
	var msg2 Test
	for i := 0; i < b.N; i++ {
		_ = msg2.UnmarshalProtobuf(data)
	}
}

func BenchmarkMarshalGoProto(b *testing.B) {
	msg := &gopb.Test{
		Name:      "test_name",
		Age:       42,
		Nicknames: []string{"nick1", "nick2"},
		Status:    gopb.TestEnum_TEST_ENUM_VALUE_1,
		Statuses:  []gopb.TestEnum{gopb.TestEnum_TEST_ENUM_UNSPECIFIED, gopb.TestEnum_TEST_ENUM_VALUE_2},
		Tags: map[string]string{
			"key1": "val1",
			"key2": "",
			"":     "val3",
		},
		EnumMap: map[int32]gopb.TestEnum{
			1: gopb.TestEnum_TEST_ENUM_VALUE_1,
			2: gopb.TestEnum_TEST_ENUM_UNSPECIFIED,
			0: gopb.TestEnum_TEST_ENUM_VALUE_2,
		},
	}
	b.ReportAllocs()
	b.ResetTimer()
	var buf []byte
	var err error
	for i := 0; i < b.N; i++ {
		buf, err = proto.MarshalOptions{}.MarshalAppend(buf[:0], msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalGoProto(b *testing.B) {
	msg := &gopb.Test{
		Name:      "test_name",
		Age:       42,
		Nicknames: []string{"nick1", "nick2"},
		Status:    gopb.TestEnum_TEST_ENUM_VALUE_1,
		Statuses:  []gopb.TestEnum{gopb.TestEnum_TEST_ENUM_UNSPECIFIED, gopb.TestEnum_TEST_ENUM_VALUE_2},
		Tags: map[string]string{
			"key1": "val1",
			"key2": "",
			"":     "val3",
		},
		EnumMap: map[int32]gopb.TestEnum{
			1: gopb.TestEnum_TEST_ENUM_VALUE_1,
			2: gopb.TestEnum_TEST_ENUM_UNSPECIFIED,
			0: gopb.TestEnum_TEST_ENUM_VALUE_2,
		},
	}
	data, _ := proto.Marshal(msg)
	b.ReportAllocs()
	b.ResetTimer()
	var msg2 gopb.Test
	for i := 0; i < b.N; i++ {
		msg2.Reset()
		_ = proto.Unmarshal(data, &msg2)
	}
}
