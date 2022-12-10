package key

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

// set the root practice day for calculating today's key
var sept_9_2021 = time.Date(2021, 9, 9, 0, 0, 0, 0, getChicagoTimeZone())
var rootPracticeDay = newPracticeDay(sept_9_2021, C)

func (pday *practiceDay) getCurrentDayKey() musicalNote {
	currentDay := time.Now()
	currentDay = currentDay.In(getChicagoTimeZone())
	today := time.Date(currentDay.Year(), currentDay.Month(), currentDay.Day(), 0, 0, 0, 0, getChicagoTimeZone())
	diff := today.Sub(pday.date)
	daysSinceRoot := int(diff.Hours() / 24)
	return musicalNote((int(pday.musicalNote) + daysSinceRoot) % musicalKeys)
}

func newPracticeDay(date time.Time, key musicalNote) *practiceDay {
	return &practiceDay{date: date, musicalNote: key}
}

func GetKeyOfDay() string {
	todayKey := rootPracticeDay.getCurrentDayKey()
	return todayKey.getNoteName()
}

func getChicagoTimeZone() *time.Location {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		panic(err)
	}
	return loc
}
