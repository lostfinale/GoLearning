package counters

// alertCounter 是一个未公开的类型
type alertCounter int

func New(value int) alertCounter {
	return alertCounter(value)
}
