package priority_queue

import (
	"reflect"
	"testing"
)

type exampleStruct struct {
	priority int
	value    string
}

func (e exampleStruct) lt(other interface{}) bool {
	return e.priority < other.(exampleStruct).priority
}

func TestPriorityQueue_Len(t *testing.T) {
	type testCase struct {
		name string
		pq   PriorityQueue
		want int
	}
	tests := []testCase{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	type testCase struct {
		name string
		pq   PriorityQueue
		args args
		want bool
	}
	tests := []testCase{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	type testCase struct {
		name string
		pq   PriorityQueue
		want interface{}
	}
	tests := []testCase{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	type testCase struct {
		name string
		pq   PriorityQueue
		want interface{}
	}
	tests := []testCase{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	type testCase struct {
		name string
		pq   PriorityQueue
		args args
	}
	tests := []testCase{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Push(tt.args.x)
		})
	}
}

func TestPriorityQueue_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	type testCase struct {
		name string
		pq   PriorityQueue
		args args
	}
	tests := []testCase{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Swap(tt.args.i, tt.args.j)
		})
	}
}
