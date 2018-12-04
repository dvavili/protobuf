// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package legacy

import (
	"reflect"
	"testing"

	pragma "github.com/golang/protobuf/v2/internal/pragma"
	pref "github.com/golang/protobuf/v2/reflect/protoreflect"
	ptype "github.com/golang/protobuf/v2/reflect/prototype"
	cmp "github.com/google/go-cmp/cmp"

	proto2_20160225 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v0.0.0-20160225-2fc053c5"
	proto2_20160519 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v0.0.0-20160519-a4ab9ec5"
	proto2_20180125 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.0.0-20180125-92554152"
	proto2_20180430 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.1.0-20180430-b4deda09"
	proto2_20180814 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.2.0-20180814-aa810b61"
	proto2_20181126 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.2.1-20181126-8d0c54c1"
	proto3_20160225 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v0.0.0-20160225-2fc053c5"
	proto3_20160519 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v0.0.0-20160519-a4ab9ec5"
	proto3_20180125 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.0.0-20180125-92554152"
	proto3_20180430 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.1.0-20180430-b4deda09"
	proto3_20180814 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.2.0-20180814-aa810b61"
	proto3_20181126 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.2.1-20181126-8d0c54c1"
)

func mustLoadFileDesc(b []byte, _ []int) pref.FileDescriptor {
	fd, err := ptype.NewFileFromDescriptorProto(loadFileDesc(b), nil)
	if err != nil {
		panic(err)
	}
	return fd
}

func TestDescriptor(t *testing.T) {
	var tests []struct{ got, want pref.Descriptor }

	fileDescP2_20160225 := mustLoadFileDesc(new(proto2_20160225.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto2_20160225.SiblingEnum(0))),
		want: fileDescP2_20160225.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto2_20160225.Message_ChildEnum(0))),
		want: fileDescP2_20160225.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.SiblingMessage))),
		want: fileDescP2_20160225.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_ChildMessage))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message))),
		want: fileDescP2_20160225.Messages().ByName("Message"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_NamedGroup))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("NamedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_OptionalGroup))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("OptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_RequiredGroup))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("RequiredGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_RepeatedGroup))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("RepeatedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_OneofGroup))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("OneofGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_ExtensionOptionalGroup))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("ExtensionOptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160225.Message_ExtensionRepeatedGroup))),
		want: fileDescP2_20160225.Messages().ByName("Message").Messages().ByName("ExtensionRepeatedGroup"),
	}}...)

	fileDescP3_20160225 := mustLoadFileDesc(new(proto3_20160225.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto3_20160225.SiblingEnum(0))),
		want: fileDescP3_20160225.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto3_20160225.Message_ChildEnum(0))),
		want: fileDescP3_20160225.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20160225.SiblingMessage))),
		want: fileDescP3_20160225.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20160225.Message_ChildMessage))),
		want: fileDescP3_20160225.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20160225.Message))),
		want: fileDescP3_20160225.Messages().ByName("Message"),
	}}...)

	fileDescP2_20160519 := mustLoadFileDesc(new(proto2_20160519.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto2_20160519.SiblingEnum(0))),
		want: fileDescP2_20160519.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto2_20160519.Message_ChildEnum(0))),
		want: fileDescP2_20160519.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.SiblingMessage))),
		want: fileDescP2_20160519.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_ChildMessage))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message))),
		want: fileDescP2_20160519.Messages().ByName("Message"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_NamedGroup))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("NamedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_OptionalGroup))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("OptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_RequiredGroup))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("RequiredGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_RepeatedGroup))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("RepeatedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_OneofGroup))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("OneofGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_ExtensionOptionalGroup))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("ExtensionOptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20160519.Message_ExtensionRepeatedGroup))),
		want: fileDescP2_20160519.Messages().ByName("Message").Messages().ByName("ExtensionRepeatedGroup"),
	}}...)

	fileDescP3_20160519 := mustLoadFileDesc(new(proto3_20160519.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto3_20160519.SiblingEnum(0))),
		want: fileDescP3_20160519.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto3_20160519.Message_ChildEnum(0))),
		want: fileDescP3_20160519.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20160519.SiblingMessage))),
		want: fileDescP3_20160519.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20160519.Message_ChildMessage))),
		want: fileDescP3_20160519.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20160519.Message))),
		want: fileDescP3_20160519.Messages().ByName("Message"),
	}}...)

	fileDescP2_20180125 := mustLoadFileDesc(new(proto2_20180125.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto2_20180125.SiblingEnum(0))),
		want: fileDescP2_20180125.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto2_20180125.Message_ChildEnum(0))),
		want: fileDescP2_20180125.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.SiblingMessage))),
		want: fileDescP2_20180125.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_ChildMessage))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message))),
		want: fileDescP2_20180125.Messages().ByName("Message"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_NamedGroup))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("NamedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_OptionalGroup))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("OptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_RequiredGroup))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("RequiredGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_RepeatedGroup))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("RepeatedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_OneofGroup))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("OneofGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_ExtensionOptionalGroup))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("ExtensionOptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180125.Message_ExtensionRepeatedGroup))),
		want: fileDescP2_20180125.Messages().ByName("Message").Messages().ByName("ExtensionRepeatedGroup"),
	}}...)

	fileDescP3_20180125 := mustLoadFileDesc(new(proto3_20180125.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto3_20180125.SiblingEnum(0))),
		want: fileDescP3_20180125.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto3_20180125.Message_ChildEnum(0))),
		want: fileDescP3_20180125.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180125.SiblingMessage))),
		want: fileDescP3_20180125.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180125.Message_ChildMessage))),
		want: fileDescP3_20180125.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180125.Message))),
		want: fileDescP3_20180125.Messages().ByName("Message"),
	}}...)

	fileDescP2_20180430 := mustLoadFileDesc(new(proto2_20180430.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto2_20180430.SiblingEnum(0))),
		want: fileDescP2_20180430.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto2_20180430.Message_ChildEnum(0))),
		want: fileDescP2_20180430.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.SiblingMessage))),
		want: fileDescP2_20180430.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_ChildMessage))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message))),
		want: fileDescP2_20180430.Messages().ByName("Message"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_NamedGroup))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("NamedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_OptionalGroup))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("OptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_RequiredGroup))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("RequiredGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_RepeatedGroup))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("RepeatedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_OneofGroup))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("OneofGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_ExtensionOptionalGroup))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("ExtensionOptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180430.Message_ExtensionRepeatedGroup))),
		want: fileDescP2_20180430.Messages().ByName("Message").Messages().ByName("ExtensionRepeatedGroup"),
	}}...)

	fileDescP3_20180430 := mustLoadFileDesc(new(proto3_20180430.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto3_20180430.SiblingEnum(0))),
		want: fileDescP3_20180430.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto3_20180430.Message_ChildEnum(0))),
		want: fileDescP3_20180430.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180430.SiblingMessage))),
		want: fileDescP3_20180430.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180430.Message_ChildMessage))),
		want: fileDescP3_20180430.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180430.Message))),
		want: fileDescP3_20180430.Messages().ByName("Message"),
	}}...)

	fileDescP2_20180814 := mustLoadFileDesc(new(proto2_20180814.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto2_20180814.SiblingEnum(0))),
		want: fileDescP2_20180814.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto2_20180814.Message_ChildEnum(0))),
		want: fileDescP2_20180814.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.SiblingMessage))),
		want: fileDescP2_20180814.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_ChildMessage))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message))),
		want: fileDescP2_20180814.Messages().ByName("Message"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_NamedGroup))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("NamedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_OptionalGroup))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("OptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_RequiredGroup))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("RequiredGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_RepeatedGroup))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("RepeatedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_OneofGroup))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("OneofGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_ExtensionOptionalGroup))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("ExtensionOptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20180814.Message_ExtensionRepeatedGroup))),
		want: fileDescP2_20180814.Messages().ByName("Message").Messages().ByName("ExtensionRepeatedGroup"),
	}}...)

	fileDescP3_20180814 := mustLoadFileDesc(new(proto3_20180814.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto3_20180814.SiblingEnum(0))),
		want: fileDescP3_20180814.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto3_20180814.Message_ChildEnum(0))),
		want: fileDescP3_20180814.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180814.SiblingMessage))),
		want: fileDescP3_20180814.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180814.Message_ChildMessage))),
		want: fileDescP3_20180814.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20180814.Message))),
		want: fileDescP3_20180814.Messages().ByName("Message"),
	}}...)

	fileDescP2_20181126 := mustLoadFileDesc(new(proto2_20181126.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto2_20181126.SiblingEnum(0))),
		want: fileDescP2_20181126.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto2_20181126.Message_ChildEnum(0))),
		want: fileDescP2_20181126.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.SiblingMessage))),
		want: fileDescP2_20181126.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_ChildMessage))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message))),
		want: fileDescP2_20181126.Messages().ByName("Message"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_NamedGroup))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("NamedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_OptionalGroup))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("OptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_RequiredGroup))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("RequiredGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_RepeatedGroup))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("RepeatedGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_OneofGroup))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("OneofGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_ExtensionOptionalGroup))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("ExtensionOptionalGroup"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto2_20181126.Message_ExtensionRepeatedGroup))),
		want: fileDescP2_20181126.Messages().ByName("Message").Messages().ByName("ExtensionRepeatedGroup"),
	}}...)

	fileDescP3_20181126 := mustLoadFileDesc(new(proto3_20181126.Message).Descriptor())
	tests = append(tests, []struct{ got, want pref.Descriptor }{{
		got:  loadEnumDesc(reflect.TypeOf(proto3_20181126.SiblingEnum(0))),
		want: fileDescP3_20181126.Enums().ByName("SiblingEnum"),
	}, {
		got:  loadEnumDesc(reflect.TypeOf(proto3_20181126.Message_ChildEnum(0))),
		want: fileDescP3_20181126.Messages().ByName("Message").Enums().ByName("ChildEnum"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20181126.SiblingMessage))),
		want: fileDescP3_20181126.Messages().ByName("SiblingMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20181126.Message_ChildMessage))),
		want: fileDescP3_20181126.Messages().ByName("Message").Messages().ByName("ChildMessage"),
	}, {
		got:  loadMessageDesc(reflect.TypeOf(new(proto3_20181126.Message))),
		want: fileDescP3_20181126.Messages().ByName("Message"),
	}}...)

	// TODO: We need a test package to compare descriptors.
	type list interface {
		Len() int
		pragma.DoNotImplement
	}
	opts := cmp.Options{
		cmp.Transformer("", func(x list) []interface{} {
			out := make([]interface{}, x.Len())
			v := reflect.ValueOf(x)
			for i := 0; i < x.Len(); i++ {
				m := v.MethodByName("Get")
				out[i] = m.Call([]reflect.Value{reflect.ValueOf(i)})[0].Interface()
			}
			return out
		}),
		cmp.Transformer("", func(x pref.Descriptor) map[string]interface{} {
			out := make(map[string]interface{})
			v := reflect.ValueOf(x)
			for i := 0; i < v.NumMethod(); i++ {
				name := v.Type().Method(i).Name
				if m := v.Method(i); m.Type().NumIn() == 0 && m.Type().NumOut() == 1 {
					switch name {
					case "Index":
						// Ignore index since legacy descriptors have no parent.
					case "Options":
						// Ignore descriptor options since protos are not cmperable.
					case "Enums", "Messages", "Extensions":
						// Ignore nested message and enum declarations since
						// legacy descriptors are all created standalone.
					case "OneofType", "ExtendedType", "EnumType", "MessageType":
						// Avoid descending into a dependency to avoid a cycle.
						// Just record the full name if available.
						//
						// TODO: Cycle support in cmp would be useful here.
						v := m.Call(nil)[0]
						if !v.IsNil() {
							out[name] = v.Interface().(pref.Descriptor).FullName()
						}
					default:
						out[name] = m.Call(nil)[0].Interface()
					}
				}
			}
			return out
		}),
		cmp.Transformer("", func(v pref.Value) interface{} {
			return v.Interface()
		}),
	}

	for _, tt := range tests {
		t.Run(string(tt.want.FullName()), func(t *testing.T) {
			if diff := cmp.Diff(&tt.want, &tt.got, opts); diff != "" {
				t.Errorf("descriptor mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
