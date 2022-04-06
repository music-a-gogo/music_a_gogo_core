package elements

import (
	"math"
	"sort"
)

type Intervals []Interval

func (is Intervals) Sort() Intervals {
	c := is
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})
	return c
}

func (is Intervals) ToRelativeIntervals() Intervals {
	temp := make([]Interval, len(is))
	for i, v := range is {
		temp[i] = v.ToRelative()
	}
	return temp
}

func (is Intervals) RemoveDuplicates() Intervals {
	set := make(map[Interval]struct{})
	for _, v := range is {
		_, ok := set[v]
		if !ok {
			set[v] = struct{}{}
		}
	}
	remains := make([]Interval, 0)
	for k := range set {
		remains = append(remains, k)
	}
	return remains
}

func (is Intervals) HasDuplicates() bool {
	set := make(map[Interval]struct{})
	for _, v := range is {
		_, ok := set[v]
		if ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

func (is Intervals) Contains(i ...Interval) bool {
	// Interval slice to map
	targets := make(map[Interval]struct{})
	for _, ii := range i {
		targets[ii] = struct{}{}
	}

	for _, v := range is {
		if _, ok := targets[v]; ok {
			return true
		}
	}
	return false
}

func (is Intervals) Min() Interval {
	min := Interval(math.MaxInt)
	for i := range is {
		if is[i] < min {
			min = is[i]
		}
	}
	return min
}

func (is Intervals) Sum() int {
	sum := 0
	for _, v := range is {
		sum += int(v)
	}
	return sum
}
