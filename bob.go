package bob

import "strings"

// package main

//"fmt"

//import "regex"

func checkString(str string) bool {
	for _, letter := range str {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			return true
		}
	}
	return false
}

func normalLetter(str string) bool {
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			return true
		}
	}
	return false
}

func checkQuestion(str string) bool {
	if len(str) == 0 {
		return false
	} else if strings.Contains(str, "?") {
		qindex := strings.Index(str, "?")
		theRest := str[qindex+1:]
		if isBlank(theRest) == false {
			return true
		}
	}
	return false
}

func isBlank(str string) bool {
	for _, letter := range str {
		if letter != ' ' && letter != '\t' && letter != '\n' && letter != '\r' {
			return true
		}
	}
	return false
}

var answer string

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if checkString(remark) {
		if normalLetter(remark) {
			if checkQuestion(remark) {
				answer = "Sure."
			} else {
				answer = "Whatever."
			}
		} else if checkQuestion(remark) {
			answer = "Calm down, I know what I'm doing!"
		} else {
			answer = "Whoa, chill out!"
		}
	} else if checkQuestion(remark) {
		answer = "Sure."
	} else if string(remark) == "" || isBlank(remark) == false {
		answer = "Fine. Be that way!"
	} else {
		answer = "Whatever."
	}
	return answer
}
