package ab5logging

import (
	"fmt"
	"strings"
)

func Horizontal(r, g, b, re, ge, be int, msg string) string {
	rb := r
	gb := g
	bb := b
	lines := strings.Split(msg, "\n")

	max := 0
	for _, line := range lines {
		l := len(line)
		if l > max {
			max = l
		}
	}

	rd := (re - r) / (max)
	gd := (ge - g) / (max)
	bd := (be - b) / (max)

	var a string

	for _, line := range lines {
		chars := strings.Split(line, "")
		for _, char := range chars {
			a += fmt.Sprintf("\033[38;2;%d;%d;%dm%v\033[0m", r, g, b, char)
			r = int(r + rd)
			g = int(g + gd)
			b = int(b + bd)
		}
		a += "\n"
		r = rb
		g = gb
		b = bb
	}

	return a[0 : len(a)-1]
}

func Vertical(r, g, b, re, ge, be int, msg string) string {
	var a []string
	splitted := strings.Split(msg, "\n")
	lines := len(splitted)
	rd := (re - r) / lines
	gd := (ge - g) / lines
	bd := (be - b) / lines
	for _, line := range strings.Split(msg, "\n") {
		a = append(a, fmt.Sprintf("\033[38;2;%v;%v;%vm%v", r, g, b, line))
		r += rd
		g += gd
		b += bd
	}
	return strings.Join(a, "\n")
}
