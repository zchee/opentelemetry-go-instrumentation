// Copyright The OpenTelemetry Authors
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

package utils

import (
	"syscall"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func TestGetLinuxKernelVersion(t *testing.T) {
	oldUnameFn := unameFn
	unameFn = func(buf *syscall.Utsname) error {
		buf.Release = [65]int8{54, 46, 52, 46, 49, 54, 45, 108, 105, 110, 117, 120, 107, 105, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		return nil
	}
	t.Cleanup(func() {
		unameFn = oldUnameFn
	})

	got, err := GetLinuxKernelVersion()
	if err != nil {
		t.Errorf("GetLinuxKernelVersion() error = %v", err)
		return
	}

	want := version.Must(version.NewVersion("6.4.16"))
	assert.Equal(t, want, got)
}
