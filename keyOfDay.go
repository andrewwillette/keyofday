package main

import (
	"fmt"
	"time"
)

type musicalNote int

const (
	A musicalNote = iota
	Bb
	B
	C
	Db
	D
	Eb
	E
	F
	Gb
	G
	Ab
)

const (
	musicalKeys = 12
)

func (note musicalNote) getNoteName() string {
	switch note {
	case A:
		return "A"
	case Bb:
		return "Bb"
	case B:
		return "B"
	case C:
		return "C"
	case Db:
		return "Db"
	case D:
		return "D"
	case Eb:
		return "Eb"
	case E:
		return "E"
	case F:
		return "F"
	case Gb:
		return "Gb"
	case G:
		return "G"
	case Ab:
		return "Ab"
	}
	panic(fmt.Sprintf("unsupported note in getNoteName(): %d", note))
}

type practiceDay struct {
	musicalNote musicalNote
	date        time.Time
}

func (pday *practiceDay) getCurrentDayKey() musicalNote {
	currentDay := time.Now()
	today := time.Date(currentDay.Year(), currentDay.Month(), currentDay.Day(), 0, 0, 0, 0, time.UTC)
	diff := today.Sub(pday.date)
	daysSince := int(diff.Hours() / 24)
	return musicalNote((int(pday.musicalNote) + daysSince) % musicalKeys)
}

func newPracticeDay(date time.Time, key musicalNote) *practiceDay {
	return &practiceDay{date: date, musicalNote: key}
}

func main() {
	// key of day calculated from given base, sept 11th 2021 was key of D
	sept_9_2021 := time.Date(2021, 9, 9, 0, 0, 0, 0, time.UTC)
	practiceDay := newPracticeDay(sept_9_2021, C)
	todayKey := practiceDay.getCurrentDayKey()
	fmt.Println(todayKey.getNoteName())
}
