package main

func isBadVersionShim(version int, bad int) bool {
	return version <= bad 
}

func FirstBadMain(n int, bad int) int {
	if n == 1 {
		return 1
	}

	for i := n; i > 0; i-- {
        if !isBadVersionShim(i, bad) {
			return i+1
		}
	}
	return 1
}