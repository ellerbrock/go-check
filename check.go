package check

import (
	"log"

	"github.com/fatih/color"
)

func Error(err error) {
	if err != nil {
		// ⚠️ ❌  ✘
		color.Red("\n")
		color.Red("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
		color.Red("▓▓           ✘ ЄɌɌϴɌ ✘           ▓▓")
		color.Red("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
		color.Red("\n")
		log.Fatal(err)
	}
}
