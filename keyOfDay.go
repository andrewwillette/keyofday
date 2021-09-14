package main

import (
	"fmt"
	"time"
)

type musicalNote int

const (
	A  musicalNote = iota
	Bb             = iota
	B              = iota
	C              = iota
	Db             = iota
	D              = iota
	Eb             = iota
	E              = iota
	F              = iota
	Gb             = iota
	G              = iota
	Ab             = iota
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
	musicalNote int
	date        time.Time
}

func (pday *practiceDay) getCurrentDayKey() musicalNote {
	currentDay := time.Now()
	zeroedCurrentDay := time.Date(currentDay.Year(), currentDay.Month(), currentDay.Day(), 0, 0, 0, 0, time.UTC)
	diff := zeroedCurrentDay.Sub(pday.date)
	daysSince := int(diff.Hours() / 24)
	return musicalNote(pday.musicalNote + daysSince)
}

func (pday *practiceDay) getDateKey(date time.Time) musicalNote {
	panic("not implemented")
}

func main() {
	sep11 := time.Date(2021, 9, 11, 0, 0, 0, 0, time.UTC)
	anchorDay := new(practiceDay)
	anchorDay.musicalNote = D
	anchorDay.date = sep11
	fmt.Println(anchorDay.getCurrentDayKey().getNoteName())
}
