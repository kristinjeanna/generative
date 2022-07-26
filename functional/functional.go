package functional

type Function[T, R any] func(T) R
type BiFunction[T, U, R any] func(T, U) R
type TriFunction[T, U, V, R any] func(T, U, V) R
type Supplier[R any] func() R
type Consumer[T any] func(T)
type Predicate[T any] Function[T, bool]

type Float64Supplier Supplier[float64]
type Float64BiFunction[R any] BiFunction[float64, float64, R]
