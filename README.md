# KeyOfDay

Terminal executable that outputs a musical key, incrementing chromatically each day.

I've heard about two great instrumentalists who practiced chromatically: Charlie Parker and John Coltrane. I play violin most every day to kill a little time, and the only strict method I follow is playing a variety of patterns (scales, arpeggios, double stops, triple stops, melodies, chord progressions, etc), while focusing heavily on counting, in the day's particular key. Each day the key increases by a halfstep (same thing as chromatically). This Go program allows me to check the key of the day. More importantly it allows me to forget the key of the day if I take a break from the practice routine.

The program is hardcoded for September 9th, 2021 being the key of C.

## Installation
* `git clone` the repository
* ensure `go` is available in your `$PATH`
* execute `go run .` to run the program
* execute `go install` tto add `keyOfDay` executable to `$GOBIN`
