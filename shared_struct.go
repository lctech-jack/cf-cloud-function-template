package cloudfunc

import "time"

type StringValue struct {
	Value string
}

type IntValue struct {
	Value int
}

type TimeValue struct {
	Value time.Time
}

type BoolValue struct {
	Value bool
}
