package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/projectdiscovery/gologger"
)

var (
	inputFile *string
	toLower   *bool
)

func gen() {

	body, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		gologger.Error().Msgf(err.Error())
		return
	}

	var alts []string
	for _, line := range strings.Split(string(body), "\n") {

		if line != "" {

			if *toLower {
				line = strings.ToLower(line)
			}

			parts := strings.Split(line, " ")

			letters1 := strings.Split(parts[0], "") // get slice of letters
			letters2 := strings.Split(parts[1], "")

			alt1 := parts[0] + "." + parts[1]
			alt2 := parts[1] + "." + parts[0]
			alt3 := parts[0] + "_" + parts[1]
			alt4 := parts[1] + "_" + parts[0]
			alt5 := letters1[0] + "." + parts[1]
			alt6 := letters2[0] + "." + parts[0]
			alt7 := letters1[0] + "_" + parts[1]
			alt8 := letters2[0] + "_" + parts[0]
			alt9 := letters1[0] + parts[1]
			alt10 := letters2[0] + parts[0]
			alt11 := parts[0] + parts[1]
			alt12 := parts[1] + parts[0]

			alts = append(alts, alt1, alt2, alt3, alt4, alt5, alt6, alt7, alt8, alt9, alt10, alt11, alt12)

		} else {
			continue
		}
	}

	for _, e := range alts {
		fmt.Printf("%s\n", e)
	}

}

func main() {

	inputFile = flag.String("i", "", "")
	toLower = flag.Bool("to-lower", false, "")
	flag.Usage = func() {
		fmt.Printf("Usage:\n\t" +
			"-i, <INPUT_FILE>       File with first and last names.\n\t" +
			"-to-lower              Conversion to lower case\n",
		)
	}
	flag.Parse()

	if *inputFile == "" {
		fmt.Printf("no input file!\n")
		flag.Usage()
		return
	}

	gen()

}
