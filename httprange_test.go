/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
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

package main

import "testing"

// Test parseRequestRange()
func TestParseRequestRange(t *testing.T) {
	// Test success cases.
	successCases := []struct {
		rangeString  string
		firstBytePos int64
		lastBytePos  int64
		length       int64
	}{
		{"bytes=2-5", 2, 5, 4},
		{"bytes=2-20", 2, 9, 8},
		{"bytes=2-", 2, 9, 8},
		{"bytes=-4", 6, 9, 4},
	}

	for _, successCase := range successCases {
		hrange, err := parseRequestRange(successCase.rangeString, 10)
		if err != nil {
			t.Fatalf("expected: <nil>, got: %s", err)
		}

		if hrange.firstBytePos != successCase.firstBytePos {
			t.Fatalf("expected: %d, got: %d", successCase.firstBytePos, hrange.firstBytePos)
		}

		if hrange.lastBytePos != successCase.lastBytePos {
			t.Fatalf("expected: %d, got: %d", successCase.lastBytePos, hrange.lastBytePos)
		}
		if hrange.getLength() != successCase.length {
			t.Fatalf("expected: %d, got: %d", successCase.length, hrange.getLength())
		}
	}

	// Test invalid range strings.
	invalidRangeStrings := []string{"", "2-5", "bytes= 2-5", "bytes=2 - 5", "bytes=0-0,-1", "bytes=5-2", "bytes=2-5 ", "bytes=2--5"}
	for _, rangeString := range invalidRangeStrings {
		if _, err := parseRequestRange(rangeString, 10); err == nil {
			t.Fatalf("expected: an error, got: <nil>")
		}
	}

	// Test error range strings.
	errorRangeString := "bytes=-0"
	if _, err := parseRequestRange(errorRangeString, 10); err != errInvalidRange {
		t.Fatalf("expected: %s, got: %s", errInvalidRange, err)
	}
}
