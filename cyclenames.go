package main

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}
