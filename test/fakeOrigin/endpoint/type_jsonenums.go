// generated by jsonenums -type Type; DO NOT EDIT

package endpoint

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"encoding/json"
	"fmt"
)

var (
	_TypeNameToValue = map[string]Type{
		"InvalidType": InvalidType,
		"Vod":         Vod,
		"Live":        Live,
		"Event":       Event,
		"Static":      Static,
		"Dir":         Dir,
		"Testing":     Testing,
	}

	_TypeValueToName = map[Type]string{
		InvalidType: "InvalidType",
		Vod:         "Vod",
		Live:        "Live",
		Event:       "Event",
		Static:      "Static",
		Dir:         "Dir",
		Testing:     "Testing",
	}
)

func init() {
	var v Type
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TypeNameToValue = map[string]Type{
			interface{}(InvalidType).(fmt.Stringer).String(): InvalidType,
			interface{}(Vod).(fmt.Stringer).String():         Vod,
			interface{}(Live).(fmt.Stringer).String():        Live,
			interface{}(Event).(fmt.Stringer).String():       Event,
			interface{}(Static).(fmt.Stringer).String():      Static,
			interface{}(Dir).(fmt.Stringer).String():         Dir,
			interface{}(Testing).(fmt.Stringer).String():     Testing,
		}
	}
}

// MarshalJSON is generated so Type satisfies json.Marshaler.
func (r Type) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _TypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Type: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Type satisfies json.Unmarshaler.
func (r *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Type should be a string, got %s", data)
	}
	v, ok := _TypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Type %q", s)
	}
	*r = v
	return nil
}