package main

type stateTableType []map[string]int

func startStateTable() stateTableType {

	stateTable := make(stateTableType, 25)
	stateTable[0] = map[string]int{"EOF": 6, "numero": 16, "letra": 3, ",": 15, ";": 14, "*": 11, "+": 11, "-": 11, "/": 11, "(": 12, ")": 13, "{": 4, "<": 9, ">": 7, "=": 8, "\"": 1, " ": 5, "\n": 5, "\t": 5}
	stateTable[1] = map[string]int{"numero": 1, "letra": 1, ",": 1, ";": 1, ":": 1, ".": 1, "!": 1, "?": 1, "\\": 1, "*": 1, "+": 1, "-": 1, "/": 1, "(": 1, ")": 1, "{": 1, "}": 1, "[": 1, "]": 1, "<": 1, ">": 1, "=": 1, "'": 1, "\"": 2, "_": 1, "\t": 1, " ": 1}
	stateTable[2] = map[string]int{}
	stateTable[3] = map[string]int{"numero": 3, "letra": 3, "_": 3}
	stateTable[4] = map[string]int{"numero": 4, "letra": 4, ",": 4, ";": 4, ":": 4, ".": 4, "!": 4, "?": 4, "\\": 4, "*": 4, "+": 4, "-": 4, "/": 4, "(": 4, ")": 4, "{": 4, "}": 5, "[": 4, "]": 4, "<": 4, ">": 4, "=": 4, "'": 4, "\"": 4, "_": 4, "\t": 4, " ": 4}
	stateTable[5] = map[string]int{}
	stateTable[6] = map[string]int{}
	stateTable[7] = map[string]int{"=": 8}
	stateTable[8] = map[string]int{}
	stateTable[9] = map[string]int{"-": 10, ">": 8, "=": 8}
	stateTable[10] = map[string]int{}
	stateTable[11] = map[string]int{}
	stateTable[12] = map[string]int{}
	stateTable[13] = map[string]int{}
	stateTable[14] = map[string]int{}
	stateTable[15] = map[string]int{}
	stateTable[16] = map[string]int{"numero": 16, "e": 19, "E": 19, ".": 17}
	stateTable[17] = map[string]int{"numero": 18}
	stateTable[18] = map[string]int{"numero": 18, "e": 19, "E": 19}
	stateTable[19] = map[string]int{"numero": 22, "+": 20, "-": 20}
	stateTable[20] = map[string]int{"numero": 21}
	stateTable[21] = map[string]int{"numero": 21}
	stateTable[22] = map[string]int{"numero": 22}

	return stateTable
}
