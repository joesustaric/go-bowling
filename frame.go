package main

type Frame struct {
	rolls []int
}

func NewFrame() Frame {
	return Frame{rolls: []int{}}
}

func (f *Frame) AddRoll(roll int) {
	f.rolls = append(f.rolls, roll)
}

func (f *Frame) FrameCompleted() bool {
	// 2 bowls that are < 10 is Over
	if f.openFrame() {
		return true
	}

	//strike with both bonus rolls
	if f.strike() && len(f.rolls) == 3 {
		return true
	}

	//spare with one bonus roll
	if f.spare() && len(f.rolls) == 3 {
		return true
	}

	return false
}

func (f *Frame) spare() bool {
	if len(f.rolls) >= 2 && (f.rolls[0]+f.rolls[1] == 10) {
		return true
	}
	return false
}

func (f *Frame) openFrame() bool {
	if len(f.rolls) == 2 && (f.rolls[0]+f.rolls[1] < 10) {
		return true
	}
	return false
}

func (f *Frame) strike() bool {
	if len(f.rolls) > 1 && f.rolls[0] == 10 {
		return true
	}
	return false
}

func (f *Frame) Score() int {
	frameScore := 0
	for _, roll := range f.rolls {
		frameScore = frameScore + roll
	}
	return frameScore
}
