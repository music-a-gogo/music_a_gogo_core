package utility

/*
	If we have a slice []string{1, 2, 3}, left rolling 1 step will be []string{2, 3, 1}.
*/
func SliceLeftRolling(s []int, step int) []int {
	if step == 0 {
		return s
	}
	if step > 0 && step < len(s) {
		return append(s[step:], s[:step]...)
	} else {
		step = step % len(s)
		if step < 0 {
			step = len(s) + step
		}
		return SliceLeftRolling(s, step)
	}
}
