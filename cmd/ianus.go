
package main

/*
 * TODO: Implement the following command line arguments
 *       -h (help)
 *       -v (verbose)
 *       -t (test config)
 *       -c (alt config file location)
 *       -l (alt log file location)
 */

import (
	"flag"
	"fmt"
)

func main () {
	helpPtr := flag.Bool("h", false, "help")
	verbosePtr := flag.Bool("v", false, "verbose")
	testPtr := flag.Bool("t", false, "test config")
	configPtr := flag.String("c", "~/.config/ianus.conf", "location of config file")
	logPtr := flag.String("l", "/var/log/ianus", "location of log files")

	flag.Parse()


}
