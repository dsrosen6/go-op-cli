package opcli

import "testing"

func TestRead(t *testing.T) {
	o := ReadOptions{
		Reference: "op://Private/Zoom/username",
	}
	out, err := Read(o)
	if err != nil {
		t.Errorf("error with Read: %v", err)
	}

	t.Log(out)
}
