//go:build ignore

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"strconv"
	"syscall"
)

var unameFunc = syscall.Uname

// KernelVersion returns major and minor kernel version numbers, parsed from
// the syscall.Uname's Release field, or 0, 0 if the version can't be obtained
// or parsed.
//
// Currently only implemented for Linux.
func main() {
	var uname syscall.Utsname
	if err := unameFunc(&uname); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", uname.Release)

	var (
		values    [3]int
		value, vi int
	)
	for _, c := range uname.Release {
		if '0' <= c && c <= '9' {
			value = (value * 10) + int(c-'0')
		} else {
			// Note that we're assuming N.N.N here.
			// If we see anything else, we are likely to mis-parse it.
			values[vi] = value
			vi++
			if vi >= len(values) {
				break
			}
			value = 0
		}
	}
	ver := fmt.Sprintf("%s.%s.%s", strconv.Itoa(values[0]), strconv.Itoa(values[1]), strconv.Itoa(values[2]))

	fmt.Printf("major: %s, minor: %s, patch: %s, %s\n", strconv.Itoa(values[0]), strconv.Itoa(values[1]), strconv.Itoa(values[2]), ver)
}
