package app

import "strings"

func parseInput(rawInput string) shellInput {
	fields := strings.Fields(rawInput)
	if len(fields) == 0 {
		return shellInput{}
	}

	return shellInput{
		Name: fields[0],
		Args: fields[1:],
	}
}
