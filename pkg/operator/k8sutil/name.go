/*
Copyright 2018 The Rook Authors. All rights reserved.
 Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 	http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package k8sutil

import "fmt"

const (
	maxPerChar = 26
)

// IndexToName converts an index to a daemon name based on as few letters of the alphabet as possible.
// For example:
//  0 -> a
//  1 -> b
//  25 -> z
//  26 -> aa
func IndexToName(index int) string {
	var result string
	for {
		i := index % maxPerChar
		c := 'z' - maxPerChar + i + 1
		result = fmt.Sprintf("%c%s", c, result)
		if index < maxPerChar {
			break
		}
		// subtract 1 since the character conversion is zero-based
		index = (index / maxPerChar) - 1
	}
	return result
}
