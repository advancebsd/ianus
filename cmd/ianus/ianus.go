package main

/*
 * TODO: Implement the following command line arguments
 *       -h (help)
 *       -v (verbose)
 *       -t (test config)
 *       -c (alt config file location)
 *       -l (alt log file location)
 *
 * TODO: Implement parsing configurations
 *       Parse nodes and configurations
 *
 * TODO: Static generator implementation
 *       Generate gemtext from markdown
 *       Generate html from gemtext
 *
 * TODO: Implement LRU caching
 *
 * TODO: Server instances
 *       Gemini Server over TLS
 *       HTTP server
 */

import (
	"flag"
	"fmt"
	"os"

	gemtextRender "github.com/advancebsd/ianus/gemtextRender"
	lex "github.com/advancebsd/ianus/markdownLexer"
)

func main() {
	helpPtr := flag.Bool("h", false, "help")
	verbosePtr := flag.Bool("v", false, "verbose")
	testPtr := flag.Bool("t", false, "test config")
	configPtr := flag.String("c", "~/.config/ianus.conf", "location of config file")
	logPtr := flag.String("l", "/var/log/ianus", "location of log files")

	flag.Parse()

	/* temporary for testing */
	fmt.Println("h: ", *helpPtr)
	fmt.Println("v: ", *verbosePtr)
	fmt.Println("t: ", *testPtr)
	fmt.Println("c: ", *configPtr)
	fmt.Println("l: ", *logPtr)

	if *helpPtr {
		fmt.Println("Ianus server help")
		fmt.Println("\th - help")
		fmt.Println("\tv - verbose")
		fmt.Println("\tt - test configure file")
		fmt.Println("\tc <file_location> - specify alternative configuration file")
		fmt.Println("\tl <directory_location> - specify alternative log directory")
		os.Exit(0)
	}

	var lex = new(lex.Lexer)
	var gr = new(gemtextRender.GemtextRender)

	// code just to use these modules for compiling
	lex.InitializeLexer("hello")
	gr.InitializeGemtextRender(nil)

}
