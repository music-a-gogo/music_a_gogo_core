# MusicAGoGo
A description of music theory written in Go. Hopefully to be provided as a cornerstone and backend lib for any music related project.

The project name `Music A Gogo` is refferring the famous French song [Jazz À Gogo](https://www.youtube.com/watch?v=qHeA5TWnIks).

# How To Use?

## Key Concepts of `MusicAGoGo`
- Note:
    A note is like a node, but not linked. It contains an integer referring its absolute frequency.

- Interval:
    Distance of 2 notes, but in a more intuitive way (meterred by semitones). In music theory, people tends to use phrase like `Perfect 5th`, `Major 6th`; In this library, it represents an interval of 7 and 9, either positive or negative.

    If two notes are in different octaves. 12 * n will be added to their distance.

- Scale:
    Scale is the sequence of notes. 
    For a major scale, the interval between them will be `2(semitones) - 2 - 1 - 2 - 2 - 2 - 1`, exactly the case of `C D E F G A B C` (C major)

## Useful Methods
tbd

# Make Contributions
The project is on developing, feel free to rasie an issue or a PR.

# License
MIT license. Basically you can do everything you want as long as you give credits to the author, which is me.

