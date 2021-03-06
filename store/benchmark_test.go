package store

import (
	"testing"

	"github.com/wepala/redux/action"
	"github.com/wepala/redux/store"
)

func Benchmark_v1_storeDispatch(b *testing.B) {
	counter := func(state interface{}, action action.Action) interface{} {
		if state == nil {
			return 0
		}
		switch action.Type {
		case "INC":
			return state.(int) + 1
		case "DEC":
			return state.(int) - 1
		default:
			return state
		}
	}
	store := store.New(counter)
	for i := 0; i < b.N; i++ {
		store.Dispatch(action.New("INC"))
	}
}

func Benchmark_storeDispatch(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.Dispatch(1)
	}
}

func Benchmark_v1_storeGetState(b *testing.B) {
	counter := func(state interface{}, action action.Action) interface{} {
		if state == nil {
			return 0
		}
		switch action.Type {
		case "INC":
			return state.(int) + 1
		case "DEC":
			return state.(int) - 1
		default:
			return state
		}
	}
	store := store.New(counter)
	for i := 0; i < b.N; i++ {
		store.GetState("func1")
	}
}

func Benchmark_storeStateOf(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.StateOf(counter)
	}
}

func Benchmark_v1_storeSubscribe(b *testing.B) {
	counter := func(state interface{}, action action.Action) interface{} {
		if state == nil {
			return 0
		}
		switch action.Type {
		case "INC":
			return state.(int) + 1
		case "DEC":
			return state.(int) - 1
		default:
			return state
		}
	}
	store := store.New(counter)
	for i := 0; i < b.N; i++ {
		store.Subscribe(func() {})
	}
}

func Benchmark_storeSubscribe(b *testing.B) {
	counter := func(state int, action int) int {
		return state + action
	}
	store := New(counter)
	for i := 0; i < b.N; i++ {
		store.Subscribe(func(action interface{}) {})
	}
}
