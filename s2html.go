package main

import (
	"bufio"
	"fmt"
	"html"
	"os"
	"regexp"
	"strings"
)

func printSection(s string, i string) {
	fmt.Printf("<li id='%s' class='section'>%s</li>\n", i, html.EscapeString(s))
}
func printData(s string, i string) {
	fmt.Printf("<li id='%s' class='data'>%s</li>\n", i, html.EscapeString(s))
}
func printInfo(s string) {
	fmt.Printf("<li class='info'>%s</li>\n", html.EscapeString(s))
}

func main() {

	// stdin to scanner
	scanner := bufio.NewScanner(os.Stdin)

	// pre compile regexp
	rs := regexp.MustCompile(`^[[:xdigit:]]{8} <(.+)>:$`)
	ri := regexp.MustCompile(`^[ ]+([[:xdigit:]]+):\t[[:xdigit:] ]+\t[a-z].*$`)
	rc := regexp.MustCompile(`^(.+); 0x([[:xdigit:]]+) <([^\+]+)(.*)>$`)
	rd := regexp.MustCompile(`^[ ]+([[:xdigit:]]+):\t[[:xdigit:] ]{47}.*$`)

	printInstruction := func(s string, i string) {
		// match call/jmp check
		m := rc.FindStringSubmatch(s)
		if m != nil {
			fmt.Printf("<li id='%s' class='instruction'>%s ; 0x<a href='#%s'>%s</a> &lt;<a href='#%s'>%s</a>%s&gt;</li>\n", i, m[1], m[2], m[2], m[3], m[3], m[4])
		} else {
			fmt.Printf("<li id='%s' class='instruction'>%s</li>\n", i, html.EscapeString(s))
		}

	}
	readline := func(s string) {
		s = strings.Trim(s, "\r\n")

		// match section check
		m := rs.FindStringSubmatch(s)
		if m != nil {
			printSection(s, m[1])
			return
		}

		// match instruction check
		m = ri.FindStringSubmatch(s)
		if m != nil {
			printInstruction(s, m[1])
			return
		}

		// match data check
		m = rd.FindStringSubmatch(s)
		if m != nil {
			printData(s, m[1])
			return
		}

		// other -> info
		printInfo(s)

	}

	for scanner.Scan() {
		readline(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
