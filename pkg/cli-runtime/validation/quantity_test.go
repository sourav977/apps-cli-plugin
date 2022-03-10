/*
Copyright 2019 VMware, Inc.

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

package validation_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	clitesting "github.com/vmware-tanzu/apps-cli-plugin/pkg/cli-runtime/testing"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/cli-runtime/validation"
)

func TestQuantity(t *testing.T) {
	tests := []struct {
		name     string
		expected validation.FieldErrors
		value    string
	}{{
		name:     "valid",
		expected: validation.FieldErrors{},
		value:    "2",
	}, {
		name:     "valid units",
		expected: validation.FieldErrors{},
		value:    "2M",
	}, {
		name:     "empty",
		expected: validation.ErrInvalidValue("", clitesting.TestField),
		value:    "",
	}, {
		name:     "invalid",
		expected: validation.ErrInvalidValue("/", clitesting.TestField),
		value:    "/",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := test.expected
			actual := validation.Quantity(test.value, clitesting.TestField)
			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Errorf("%s() = (-expected, +actual): %s", test.name, diff)
			}
		})
	}
}
