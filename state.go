package tastx

import "testing"

type State interface {
	Initialize() error
}

func WithState[S State](s S, fn func(t *testing.T, s S)) func(*testing.T) {
	return func(t *testing.T) {
		if err := s.Initialize(); err != nil {
			t.Fatal("failed to initialize state:", err)
		}
		fn(t, s)
	}
}

func RunWithState[S State](t *testing.T, s S, fn func(t *testing.T, s S)) {
	runFn := WithState(s, fn)
	runFn(t)
}
