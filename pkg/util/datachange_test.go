package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type args[T any] struct {
		obj T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    map[string]any
		wantErr assert.ErrorAssertionFunc
	}
	type testStruct struct {
		Name string
		Age  float64
	}
	tests := []testCase[testStruct]{
		{
			name: "test1",
			args: args[testStruct]{
				obj: testStruct{
					Name: "test",
					Age:  18,
				},
			},
			want: map[string]any{
				"Name": "test",
				"Age":  18.0,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToMapF1(tt.args.obj)
			if !tt.wantErr(t, err, fmt.Sprintf("StructToMapF1(%v)", tt.args.obj)) {
				return
			}
			fmt.Printf("got: %v\n", got)
			fmt.Printf("want: %v\n", tt.want)
			assert.Equalf(t, tt.want, got, "StructToMapF1(%v)", tt.args.obj)
		})
	}
}
