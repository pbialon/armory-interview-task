package priority_queue

type ordered interface {
	eq(other ordered) bool
	ne(other ordered) bool
	gt(other ordered) bool
	lt(other ordered) bool
	ge(other ordered) bool
	le(other ordered) bool
}
