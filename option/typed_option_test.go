// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package option

import (
	"strconv"
	"testing"
	"time"
)

func TestBoolOption(t *testing.T) {
	b := NewBoolOption(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan("true"); err != nil {
		t.Error(err)
	} else if !b.IsBool() || !b.Bool() {
		t.Error(b.Value())
	}
}

func TestInt8Option(t *testing.T) {
	b := NewInt8Option(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan("123"); err != nil {
		t.Error(err)
	} else if !b.IsInt8() || b.Int8() != 123 {
		t.Error(b.Value())
	}
}

func TestFloat64Option(t *testing.T) {
	b := NewFloat64Option(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan("1.2"); err != nil {
		t.Error(err)
	} else if !b.IsFloat64() || b.Float64() != 1.2 {
		t.Error(b.Value())
	}

	if err := b.UnmarshalJSON(strconv.AppendFloat(nil, 2.0, 'f', -1, 64)); err != nil {
		t.Error(err)
	} else if b.Float64() != 2.0 {
		t.Error(b.Float64())
	}
}

func TestStringOption(t *testing.T) {
	b := NewStringOption(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan(123); err != nil {
		t.Error(err)
	} else if !b.IsString() || b.Str() != "123" {
		t.Error(b.Value())
	}
}

func TestInterface(t *testing.T) {
	opts := []Option{
		NewBoolOption(None()),
		NewInt64Option(None()),
		NewFloat64Option(None()),
		NewStringOption(None()),
	}
	values := []interface{}{"true", "123", "1.2", 456}

	for i, opt := range opts {
		if err := opt.Scan(values[i]); err != nil {
			t.Error(err)
		} else if opt.Value() == nil {
			t.Fail()
		}
	}
}

func TestNamedTypedOption(t *testing.T) {
	opt := NewNamedOption("", NewBoolOption(nil))
	if err := opt.Scan("on"); err != nil {
		t.Error(err)
	} else if !opt.IsBool() || !opt.Bool() {
		t.Error(opt)
	}
}

func TestTimeOption(t *testing.T) {
	b := NewTimeOption(None())

	if err := b.Scan("2019-05-13 22:57:42"); err != nil {
		t.Error(err)
	} else if b.Value().(time.Time).Format("2006-01-02 15:04:05") != "2019-05-13 22:57:42" {
		t.Error(b.Value())
	}

	if err := b.UnmarshalJSON([]byte("2019-05-13 22:57:43")); err != nil {
		t.Error(err)
	} else if b.Value().(time.Time).Format("2006-01-02 15:04:05") != "2019-05-13 22:57:43" {
		t.Error(b.Value())
	}

	if err := b.UnmarshalJSON([]byte("0000-00-00 00:00:00")); err != nil {
		t.Error(err)
	} else if b.Value().(time.Time) != (time.Time{}) {
		t.Error(b.Value())
	}
}
