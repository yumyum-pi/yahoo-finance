package yahooFinace

import (
	"fmt"
)

const version = "0.0.1"
const authorName = "Vivek Rawat"
const program_name = "yahooFinance"

func printProgramDetails() {
	fmt.Printf(
		"Program Name: %s\nAuthor Name: %s\nVersion: %s\n",
		program_name,
		authorName,
		version,
	)
}
