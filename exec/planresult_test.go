// Copyright © 2016 Asteris, LLC
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

package exec_test

import (
	"testing"

	"github.com/asteris-llc/converge/exec"
	"github.com/stretchr/testify/assert"
)

var planResult1 = &exec.PlanResult{
	Path:          "moduleA/submodule1",
	CurrentStatus: "status",
	WillChange:    true,
}

var planResult2 = &exec.PlanResult{
	Path:          "moduleB/submodule1",
	CurrentStatus: "status",
	WillChange:    false,
}

func TestPlanResultString(t *testing.T) {
	t.Parallel()

	expected1 := "+ moduleA/submodule1:\n\tCurrently: status\n\tWill Change: true"
	expected2 := "- moduleB/submodule1:\n\tCurrently: status\n\tWill Change: false"
	assert.Equal(t, expected1, planResult1.String())
	assert.Equal(t, expected2, planResult2.String())
}

func TestPlanResultPretty(t *testing.T) {
	t.Parallel()

	expected1 := "\x1b[32m+\x1b[0m \x1b[33mmoduleA/submodule1\x1b[0m:\n\tCurrently: status\n\tWill Change: \x1b[33mtrue\x1b[0m"
	expected2 := "\x1b[31m-\x1b[0m \x1b[34mmoduleB/submodule1\x1b[0m:\n\tCurrently: status\n\tWill Change: \x1b[34mfalse\x1b[0m"
	assert.Equal(t, expected1, planResult1.Pretty())
	assert.Equal(t, expected2, planResult2.Pretty())
}