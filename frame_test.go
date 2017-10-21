package main

import "testing"

func TestFrameCompleted(t *testing.T) {

	// non strike first roll
	frame := NewFrame()
	frame.AddRoll(3)
	if frame.FrameCompleted() == true {
		t.Errorf("Frame should not be Over")
	}

	//non strike non spare 2nd roll
	frame = NewFrame()
	frame.AddRoll(3)
	frame.AddRoll(2)

	if frame.FrameCompleted() == false {
		t.Errorf("Frame should be Over")
	}

	// strike no bonus rolls
	frame = NewFrame()
	frame.AddRoll(10)

	if frame.FrameCompleted() == true {
		t.Errorf("strike frame should not be Over")
	}

	// strike 1 bonus roll
	frame = NewFrame()
	frame.AddRoll(10)
	frame.AddRoll(3)

	if frame.FrameCompleted() == true {
		t.Errorf("strike frame should not be over")
	}

	// strike 2 bonus rolls
	frame = NewFrame()
	frame.AddRoll(10)

	frame.AddRoll(3)
	frame.AddRoll(3)

	if frame.FrameCompleted() == false {
		t.Errorf("strike frame should be over")
	}

	//spare no bonus roll
	frame = NewFrame()
	frame.AddRoll(4)
	frame.AddRoll(6)

	if frame.FrameCompleted() == true {
		t.Errorf("spare frame should not be over")
	}

	//spare no with 1 bonus roll
	frame = NewFrame()
	frame.AddRoll(4)
	frame.AddRoll(6)
	frame.AddRoll(6)

	if frame.FrameCompleted() == false {
		t.Errorf("spare frame should be over")
	}
}

func TestAddRoll(t *testing.T) {
	//ignores negative numbers
	frame := NewFrame()

	frame.AddRoll(-3)
	if frame.Score() != 0 {
		t.Errorf("Added a negative score")
	}

	//ignores anything over 10 for the first roll
	frame = NewFrame()

	frame.AddRoll(5)

	if frame.Score() != 5 {
		t.Errorf("Fail to add score")
	}
}
