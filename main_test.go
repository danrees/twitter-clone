package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_setup(t *testing.T) {
	tests := []struct {
		name string
		want string
		//want1 *gin.Engine
	}{
		// TODO: Add test cases.
		{
			"basic test",
			":8080",
			//router.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := setup()
			assert.Equal(t, tt.want, got)
			//assert.Equal(t, tt.want1, got1)
		})
	}
}
