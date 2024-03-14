package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyConfig struct {
	Foo    string
	Bar    string
	Nested *MyNestedConfig
}

type MyNestedConfig struct {
	SomeInt int
}

func TestConfigBuilderOptions_Merge(t *testing.T) {

	tests := []struct {
		name           string
		configA        MyConfig
		configB        MyConfig
		expectedResult MyConfig
	}{
		{
			name: "ShouldMergeTopLevelField",
			configA: MyConfig{
				Foo: "hello",
			},
			configB: MyConfig{
				Bar: "world",
			},
			expectedResult: MyConfig{
				Foo: "hello",
				Bar: "world",
			},
		},
		{
			name: "ShouldMergeNestedStruct",
			configA: MyConfig{
				Foo: "hello",
				Bar: "world",
			},
			configB: MyConfig{
				Nested: &MyNestedConfig{SomeInt: 1},
			},
			expectedResult: MyConfig{
				Foo:    "hello",
				Bar:    "world",
				Nested: &MyNestedConfig{SomeInt: 1},
			},
		},
	}

	for _, tt := range tests {
		opts := &ConfigBuilderOptions[MyConfig]{}
		t.Run(tt.name, func(st *testing.T) {

			opts.Merge(&tt.configA)
			opts.Merge(&tt.configB)

			assert.Equal(st, tt.expectedResult, opts.Config)

		})
	}
}
