/*
Copyright 2019 The Kubernetes Authors.

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

package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"k8s.io/enhancements/pkg/kepval"
)

const (
	kepsDir = "keps"
)

func TestValidation(t *testing.T) {
	warnings, err := kepval.ValidateRepository(kepsDir)
	require.Nil(t, err)

	t.Logf(
		"KEP validation succeeded, but the following warnings occurred: %v",
		warnings,
	)
}
