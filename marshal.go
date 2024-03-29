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
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

// Marshal encodes a value to gob which is then encoded to base64 string.
func Marshal(v interface{}) (string, error) {
	buff := bytes.Buffer{}
	err := gob.NewEncoder(&buff).Encode(v)
	if err != nil {
		return "", err
	}

	s := base64.StdEncoding.EncodeToString(buff.Bytes())

	return s, nil
}
