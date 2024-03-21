package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyConfig struct {
	Foo         string
	Bar         string
	Nested      *MyNestedConfig
	OtherNested MyNestedConfig
}

type MyNestedConfig struct {
	SomeInt        int
	SomeOtherValue string
}

func TestConfigBuilderOptions_Merge(t *testing.T) {

	tests := []struct {
		name           string
		configA        MyConfig
		configB        MyConfig
		expectedResult *MyConfig
	}{
		{
			name: "ShouldMergeTopLevelField",
			configA: MyConfig{
				Foo: "hello",
			},
			configB: MyConfig{
				Bar: "world",
			},
			expectedResult: &MyConfig{
				Foo: "hello",
				Bar: "world",
			},
		},
		{
			name: "ShouldMergeNestedStruct",
			configA: MyConfig{
				Foo:         "hello",
				Bar:         "world",
				Nested:      &MyNestedConfig{SomeOtherValue: "foobar"},
				OtherNested: MyNestedConfig{SomeInt: 6},
			},
			configB: MyConfig{
				Nested: &MyNestedConfig{SomeInt: 1},
			},
			expectedResult: &MyConfig{
				Foo:         "hello",
				Bar:         "world",
				Nested:      &MyNestedConfig{SomeInt: 1, SomeOtherValue: "foobar"},
				OtherNested: MyNestedConfig{SomeInt: 6},
			},
		},
	}

	for _, tt := range tests {
		var cfg MyConfig
		opts := &ConfigBuilderOptions[MyConfig]{
			Config: &cfg,
		}
		t.Run(tt.name, func(st *testing.T) {

			opts.Merge(&tt.configA)
			opts.Merge(&tt.configB)

			assert.Equal(st, tt.expectedResult, opts.Config)

		})
	}
}
