package mapconv

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type dummyTagged struct {
	A          string `mapconv:"ValueA.A"`
	B          string `mapconv:"ValueA.ValueB.B"`
	C          string `mapconv:"ValueA.ValueB.ValueC.C"`
	Pointer    *time.Time
	Slice      []string
	NoTag      string
	unexported string
}

type dummyNaked struct {
	ValueA *struct {
		A      string
		ValueB *struct {
			B      string
			ValueC *struct {
				C string
			}
		}
	}
	Pointer    *time.Time
	Slice      []string
	NoTag      string
	unexported string
}

func TestToNaked(t *testing.T) {
	zeroTime := time.Unix(0, 0)
	tests := []struct {
		input  *dummyTagged
		output *dummyNaked
		err    error
	}{
		{
			input: &dummyTagged{
				A:          "A",
				B:          "B",
				C:          "C",
				Pointer:    &zeroTime,
				Slice:      []string{"a", "b", "c"},
				NoTag:      "NoTag",
				unexported: "unexported",
			},
			output: &dummyNaked{
				ValueA: &struct {
					A      string
					ValueB *struct {
						B      string
						ValueC *struct {
							C string
						}
					}
				}{
					A: "A",
					ValueB: &struct {
						B      string
						ValueC *struct {
							C string
						}
					}{
						B: "B",
						ValueC: &struct {
							C string
						}{
							C: "C",
						},
					},
				},
				Pointer: &zeroTime,
				Slice:   []string{"a", "b", "c"},
				NoTag:   "NoTag",
			},
		},
	}

	for _, tt := range tests {
		output := &dummyNaked{}
		err := ConvertTo(tt.input, output)
		require.Equal(t, tt.err, err)
		if err == nil {
			require.Equal(t, tt.output, output)
		}
	}

}

func TestFromNaked(t *testing.T) {

	tests := []struct {
		output *dummyTagged
		input  *dummyNaked
		err    error
	}{
		{
			output: &dummyTagged{
				A:     "A",
				B:     "B",
				C:     "C",
				NoTag: "NoTag",
			},
			input: &dummyNaked{
				ValueA: &struct {
					A      string
					ValueB *struct {
						B      string
						ValueC *struct {
							C string
						}
					}
				}{
					A: "A",
					ValueB: &struct {
						B      string
						ValueC *struct {
							C string
						}
					}{
						B: "B",
						ValueC: &struct {
							C string
						}{
							C: "C",
						},
					},
				},
				NoTag: "NoTag",
			},
		},
	}

	for _, tt := range tests {
		output := &dummyTagged{}
		err := ConvertFrom(tt.input, output)
		require.Equal(t, tt.err, err)
		if err == nil {
			require.Equal(t, tt.output, output)
		}
	}

}

type dummySlice struct {
	Slice []*dummySliceInner `json:",omitempty"`
}

type dummySliceInner struct {
	Value string             `json:",omitempty"`
	Slice []*dummySliceInner `json:",omitempty"`
}

type dummyExtractInnerSlice struct {
	Values       []string `json:",omitempty" mapconv:"[]Slice.Value"`
	NestedValues []string `json:",omitempty" mapconv:"[]Slice.[]Slice.Value"`
}

func TestExtractInnerSlice(t *testing.T) {
	tests := []struct {
		input  *dummySlice
		expect *dummyExtractInnerSlice
	}{
		{
			input: &dummySlice{
				Slice: []*dummySliceInner{
					{Value: "value1"},
					{Value: "value2"},
					{
						Value: "value3",
						Slice: []*dummySliceInner{
							{Value: "value4"},
							{Value: "value5"},
						},
					},
				},
			},
			expect: &dummyExtractInnerSlice{
				Values:       []string{"value1", "value2", "value3"},
				NestedValues: []string{"value4", "value5"},
			},
		},
	}

	for _, tt := range tests {
		output := &dummyExtractInnerSlice{}
		err := ConvertFrom(tt.input, output)

		require.NoError(t, err)
		require.Equal(t, tt.expect, output)
	}
}

func TestInsertInnerSlice(t *testing.T) {
	tests := []struct {
		input  *dummyExtractInnerSlice
		output *dummySlice
	}{
		{
			input: &dummyExtractInnerSlice{
				Values:       []string{"value1", "value2", "value3"},
				NestedValues: []string{"value4", "value5"},
			},
			output: &dummySlice{
				Slice: []*dummySliceInner{
					{Value: "value1"},
					{Value: "value2"},
					{Value: "value3"},
					{
						Slice: []*dummySliceInner{
							{Value: "value4"},
						},
					},
					{
						Slice: []*dummySliceInner{
							{Value: "value5"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		output := &dummySlice{}
		err := ConvertTo(tt.input, output)

		require.NoError(t, err)
		require.Equal(t, tt.output, output)
	}
}

type hasDefaultSource struct {
	Field string `mapconv:"Field,default=default-value"`
}

type hasDefaultDest struct {
	Field string
}

func TestDefaultValue(t *testing.T) {
	tests := []struct {
		input  *hasDefaultSource
		output *hasDefaultDest
	}{
		{
			input: &hasDefaultSource{},
			output: &hasDefaultDest{
				Field: "default-value",
			},
		},
	}

	for _, tt := range tests {
		output := &hasDefaultDest{}
		err := ConvertTo(tt.input, output)
		require.NoError(t, err)
		require.Equal(t, tt.output, output)
	}
}

type multipleSource struct {
	Field string `mapconv:"Field1/Field2"`
}

type multipleDest struct {
	Field1 string
	Field2 string
}

func TestMultipleDestination(t *testing.T) {
	tests := []struct {
		input  *multipleSource
		output *multipleDest
	}{
		{
			input: &multipleSource{
				Field: "value",
			},
			output: &multipleDest{
				Field1: "value",
				Field2: "value",
			},
		},
	}

	for _, tt := range tests {
		output := &multipleDest{}
		err := ConvertTo(tt.input, output)
		require.NoError(t, err)
		require.Equal(t, tt.output, output)
	}
}

type recursiveSource struct {
	Field *recursiveSourceChild `mapconv:",recursive"`
}

type recursiveSourceChild struct {
	Field1 string `mapconv:"Dest1"`
	Field2 string `mapconv:"Dest2"`
}

type recursiveDest struct {
	Field *recursiveDestChild
}

type recursiveDestChild struct {
	Dest1 string
	Dest2 string
}

type recursiveSourceSlice struct {
	Fields []*recursiveSourceChild `mapconv:"[]Slice,recursive"`
}

type recursiveDestSlice struct {
	Slice []*recursiveDestChild
}

func TestRecursive(t *testing.T) {
	tests := []struct {
		input  *recursiveSource
		expect *recursiveDest
	}{
		{
			input: &recursiveSource{
				Field: &recursiveSourceChild{
					Field1: "value1",
					Field2: "value2",
				},
			},
			expect: &recursiveDest{
				Field: &recursiveDestChild{
					Dest1: "value1",
					Dest2: "value2",
				},
			},
		},
	}

	for _, tt := range tests {
		dest := &recursiveDest{}
		err := ConvertTo(tt.input, dest)
		require.NoError(t, err)
		require.Equal(t, tt.expect, dest)

		// reverse
		source := &recursiveSource{}
		err = ConvertFrom(tt.expect, source)
		require.NoError(t, err)
		require.Equal(t, tt.input, source)
	}
}

func TestRecursiveSlice(t *testing.T) {
	tests := []struct {
		input  *recursiveSourceSlice
		output *recursiveDestSlice
	}{
		{
			input: &recursiveSourceSlice{
				Fields: []*recursiveSourceChild{
					{
						Field1: "value1",
						Field2: "value2",
					},
					{
						Field1: "value3",
						Field2: "value4",
					},
				},
			},
			output: &recursiveDestSlice{
				Slice: []*recursiveDestChild{
					{
						Dest1: "value1",
						Dest2: "value2",
					},
					{
						Dest1: "value3",
						Dest2: "value4",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		output := &recursiveDestSlice{}
		err := ConvertTo(tt.input, output)
		require.NoError(t, err)
		require.Equal(t, tt.output, output)

		// reverse
		source := &recursiveSourceSlice{}
		err = ConvertFrom(tt.output, source)
		require.NoError(t, err)
		require.Equal(t, tt.input, source)
	}
}
