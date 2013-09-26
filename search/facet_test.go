// Copyright 2012 Matthew Baird
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
// limitations under the license.

package search

import (
	//"encoding/json"
	. "github.com/araddon/gou"
	"testing"
)

func TestFacetRegex(t *testing.T) {

	// This is a possible solution for auto-complete
	out, _ := Search("github").Size("0").Facet(
		Facet().Regex("repository.name", "no.*").Size("8"),
	).Result()
	if out == nil || &out.Hits == nil {
		t.Fail()
		return
	}
	//Debug(string(out.Facets))
	fh := NewJsonHelper([]byte(out.Facets))
	facets := fh.Helpers("/repository.name/terms")
	Assert(len(facets) == 8, t, "Should have 8? but was %v", len(facets))
	// for _, f := range facets {
	// 	Debug(f)
	// }
}
