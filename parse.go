// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"strings"

	"github.com/Unknwon/com"
	dry "github.com/ungerik/go-dry"
)

func main() {
	lines, err := dry.FileGetLines("spending.md")
	if err != nil {
		log.Fatalln(err)
	}

	var (
		total = 0
		eat   = 0
	)

	for _, line := range lines {
		if strings.Contains(line, "@еда") {
			eat += com.StrTo(strings.Fields(line)[0]).MustInt()
		}
		if len(line) > 0 && line[0] != '#' {
			total += com.StrTo(strings.Fields(line)[0]).MustInt()
		}
	}

	log.Printf("total: %d\neat:%d", total, eat)
}
