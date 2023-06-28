package priority_queue

import (
	"reflect"
	"testing"
)

type testStruct struct {
	priority int
	value    string
}

func (e testStruct) lt(other interface{}) bool {
	return e.priority < other.(*testStruct).priority
}

func TestPriorityQueue_Len(t *testing.T) {
	type testCase struct {
		name string
		pq   PriorityQueue
		want int
	}
	tests := []testCase{
		{
			name: "empty priority queue should return length 0",
			pq:   PriorityQueue{},
			want: 0,
		}, {
			name: "priority queue with one element should return length 1",
			pq: PriorityQueue{
				items: []Item{
					&testStruct{
						priority: 1,
						value:    "test",
					},
				},
			},
			want: 1,
		},
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
		{
			name: "first element has lower priority",
			pq: PriorityQueue{
				items: []Item{
					&testStruct{
						priority: 1,
						value:    "test1",
					}, &testStruct{
						priority: 2,
						value:    "test2",
					},
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
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
		{
			name: "empty priority queue should return nil",
			pq:   PriorityQueue{},
			want: nil,
		}, {
			name: "priority queue with one element should return that element",
			pq: PriorityQueue{
				items: []Item{
					&testStruct{
						priority: 1,
						value:    "test",
					},
				},
			},
			want: &testStruct{
				priority: 1,
				value:    "test",
			},
		}, {
			name: "priority queue with two elements should return the first element",
			pq: PriorityQueue{
				items: []Item{
					&testStruct{
						priority: 1,
						value:    "test1",
					}, &testStruct{
						priority: 2,
						value:    "test2",
					},
				},
			},
			want: &testStruct{
				priority: 1,
				value:    "test1",
			},
		},
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
		name   string
		pq     PriorityQueue
		args   args
		length int
	}
	tests := []testCase{
		{
			name: "pushing one element to empty priority queue",
			pq:   PriorityQueue{},
			args: args{
				x: &testStruct{
					priority: 1,
					value:    "test",
				},
			},
			length: 1,
		}, {
			name: "pushing element to non-empty priority queue",
			pq: PriorityQueue{
				items: []Item{
					&testStruct{
						priority: 1,
						value:    "test1",
					},
				},
			},
			args: args{
				x: &testStruct{
					priority: 2,
					value:    "test2",
				},
			},
			length: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Push(tt.args.x)
			if tt.pq.Len() != tt.length {
				t.Errorf("After Push(), got length %v, want %v", tt.pq.Len(), tt.length)
			}
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
		want PriorityQueue
	}
	tests := []testCase{
		{
			name: "swapping two elements",
			pq: PriorityQueue{
				items: []Item{
					&testStruct{
						priority: 1,
						value:    "test1",
					}, &testStruct{
						priority: 2,
						value:    "test2",
					},
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			want: PriorityQueue{
				items: []Item{
					&testStruct{
						priority: 2,
						value:    "test2",
					}, &testStruct{
						priority: 1,
						value:    "test1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Swap(tt.args.i, tt.args.j)
		})
	}
}
