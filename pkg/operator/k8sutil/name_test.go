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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertMonID(t *testing.T) {
	assert.Equal(t, "a", IndexToName(0))
	assert.Equal(t, "b", IndexToName(1))
	assert.Equal(t, "c", IndexToName(2))
	assert.Equal(t, "z", IndexToName(25))
	assert.Equal(t, "aa", IndexToName(26))
	assert.Equal(t, "ab", IndexToName(27))
	assert.Equal(t, "ac", IndexToName(28))
	assert.Equal(t, "az", IndexToName(51))
	assert.Equal(t, "ba", IndexToName(52))
	assert.Equal(t, "bb", IndexToName(53))
	assert.Equal(t, "bz", IndexToName(77))
	assert.Equal(t, "ca", IndexToName(78))
	assert.Equal(t, "za", IndexToName(676))
	assert.Equal(t, "aaa", IndexToName(702))
	assert.Equal(t, "aaz", IndexToName(727))
	assert.Equal(t, "aba", IndexToName(728))
}
