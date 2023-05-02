package utils

import (
	"math/rand"
	"time"

	aw "github.com/deanishe/awgo"
)

func NewCopyableItem(wf *aw.Workflow, title, val string) {
	NewCopyableItemWithDescription(wf, title, val, val)
}

func NewCopyableItemWithDescription(wf *aw.Workflow, title, description, val string) {
	it := wf.NewItem(title)
	it.Subtitle(description)
	it.Arg(val)
	it.Valid(true)
	it.Copytext(val)
}

func GetTimeFromID(id int64) time.Time {
	var res = id >> 32
	return time.Unix(res, 0)
}

func GetIDFromTime(t time.Time) int64 {
	id := t.Unix()
	return id << 32
}

func GenID(t time.Time) int64 {
	id := t.Unix()
	randNum := rand.Int31()
	return id<<32 | int64(randNum)
}
