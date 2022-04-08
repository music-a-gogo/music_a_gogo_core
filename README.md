# MusicAGoGo
A description of music theory written in Go. 

Hopefully to be provided as a cornerstone and backend lib for any music-related project (maybe something like a web version of [fretboard-notes](https://github.com/mooxiu/fretboard-notes)).

The project name `Music A Gogo` is referring the famous French song [Jazz À Gogo](https://www.youtube.com/watch?v=qHeA5TWnIks).

# How To Use?

## Key Concepts of `MusicAGoGo`
- Note:
    A note is like a node, but not linked. It contains an integer referring absolute frequency.

- Interval:
    Distance of 2 notes, but in a more intuitive way (metered by semitones). In music theory, people tend to use phrases like `Perfect 5th`, `Major 6th`; In this library, it represents an interval of 7 and 9, either positive or negative.

    If two `Note`s are in different octaves, $12 * n$ will be added to their distance.

- Scale:
    Scale is the sequence of notes. 
    For a major scale, the interval between them will be `2(semitones) - 2 - 1 - 2 - 2 - 2 - 1`, exactly the case of `C D E F G A B C` (C major)

- Chord:
    Chord is defined by a root note and a slice of intervals to root note.


# Make Contributions
The project is on developing. Feel free to raise an issue or a PR.

# License
MIT license. Basically, you can do everything you want as long as you give credits to the author, which is me.

