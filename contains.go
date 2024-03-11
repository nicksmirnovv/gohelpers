package gohelpers

import ()

func ContainsInArray[Arr comparable](s []Arr, str Arr) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
