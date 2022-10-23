/**
 * Copyright 2022 github.com/moizalicious
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gob64

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type data struct {
		Foo string
		Bar string
	}

	expected := data{
		Foo: "ABC",
		Bar: "XYZ",
	}

	s := "Iv+BAwEBBGRhdGEB/4IAAQIBA0ZvbwEMAAEDQmFyAQwAAAAN/4IBA0FCQwEDWFlaAA=="

	var actual data
	if err := Unmarshal(s, &actual); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if actual != expected {
		t.Errorf("Failed to unmarshal data.\nExpected Value:%v\nActual Value:%v", expected, actual)
	}
}
