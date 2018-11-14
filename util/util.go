// Copyright 2018 Aleksandr Demakin. All rights reserved.

package util

import "sort"

// AddSortedSliceString inserts a string into a sorted slice.
// Returns new slice.
func AddSortedSliceString(sl []string, s string) []string {
	if idx := sort.SearchStrings(sl, s); idx >= len(sl) || sl[idx] != s {
		sl = append(sl[:idx], append([]string{s}, sl[idx:]...)...)
	}
	return sl
}
