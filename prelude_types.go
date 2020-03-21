package redux

import (
	"github.com/wepala/redux/action"
)

type Reducer func(interface{}, action.Action) interface{}
