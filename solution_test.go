package square

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalcSquare(t *testing.T) {
	sidesNum := MyType{
		sides: 0,
	}
	fmt.Println(CalcSquare(10.0, sidesNum))

	type test struct {
		name string
		input float64
		sep   MyType
		want  float64
	}

	tests := []test{
		{name:"Test1",input: 10.0, sep: MyType{
			sides: 0,
		}, want: 314.1592653589793},
		{name:"Test2",input: 10.0, sep: MyType{
			sides: 3,
		}, want: 43.30127018922193},
		{name:"Test3",input: 10.0, sep: MyType{
			sides: 4,
		}, want: 100},
		{name:"Test4",input: 10.0, sep: MyType{
			sides: 11,
		}, want: 0.0},
	}

	for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := CalcSquare(tc.input, tc.sep)
            if !reflect.DeepEqual(tc.want, got) {
                t.Fatalf("expected: %v, got: %v", tc.want, got)
            }
        })
    }
}