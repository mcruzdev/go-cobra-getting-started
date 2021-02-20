package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumArr(t *testing.T) {

	var numbers = []string{"1", "1"}
	var expect int64 = 2
	actual := SumArr(numbers)

	assert.Equal(t, expect, actual, "they should be equal")

}

func Test_SumCmd_Creator(t *testing.T) {
	var (
		sumCmd  = CreateSumCmd()
		argsTmp = []string{"sum -c"}
		buffTmp = new(bytes.Buffer)

		expect string
		actual string
	)

	sumCmd.SetOut(buffTmp)
	sumCmd.Flags().Set("creator", "true")
	sumCmd.SetArgs(argsTmp)

	if err := sumCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'sumCmd.Execute()'.", "Error msg: %v", err)
	}

	expect = "Created by: mcruzdev\n"
	actual = buffTmp.String()

	assert.Equal(t, expect, actual, "Command 'sum --creator' should return 'Created by: mcruzdev'")
}

func Test_SumCmd(t *testing.T) {
	var (
		sumCmd  = CreateSumCmd()
		argsTmp = []string{"1", "1", "1"}
		buffTmp = new(bytes.Buffer)

		expect string
		actual string
	)

	sumCmd.SetOut(buffTmp)
	sumCmd.SetArgs(argsTmp)

	if err := sumCmd.Execute(); err != nil {
		assert.FailNowf(t, "Failed to execute 'sumCmd.Execute()'.", "Error msg: %v", err)
	}

	expect = "Sum is: 3\n"
	actual = buffTmp.String()
	assert.Equal(t, expect, actual, "Command 'sum' should return 'Sum is: 3'")
}
