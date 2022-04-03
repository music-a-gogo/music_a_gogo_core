package elements

import "sort"

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

func (is Intervals) Contains(i Interval) bool {
	for _, v := range is {
		if v == i {
			return true
		}
	}
	return false
}
