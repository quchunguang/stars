package main

import (
	"bufio"
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"strings"
)

var arguments map[string]interface{}

func main() {
	usage := `
Name:
  stars

Usage:
    stars [--spark]
          [-F <fs>] [-H] [-S <cs>] [-w <n>] [-R] [-C <col>]
          [<data>...]
    stars -h | --help
    stars --version

Options:
    --spark    ▁▂▃▅▂▇▁▂▃ style.
    -C <col>   Given comma separated list of column numbers.
               Example for CSL: 1,2,3 or -3,-2,-1.
    -F <fs>    Use fs for the input field separator [default: ' '].
    -H         Including the description at beginning of each line of data.
    -R         Reverse the data format with rows and columns. That makes
               easy to process the standard DSV files.
    -S <cs>    Characters used by draw recursively. [default: '+-*!#']
    -w <n>     Hold width for draw [default: 20].
    -h --help  Show this screen.
    --version  Show version.

Data Format:
    If no data is given, it'll read from STDIN. The data should like this,

    title a 1 2 3
    title b 2 3 4
    title c 4 5 6

    The <data> line given should be formated like this,

    "title a" 1 2 3

    title a,1,2,3
`
	arguments, _ = docopt.Parse(usage, nil, true, "stars v0.01", false)
	fmt.Println(arguments)

	var in *bufio.Scanner
	data := strings.Join(arguments["<data>"].([]string), ",")
	if data == "" {
		in = bufio.NewScanner(os.Stdin)
	} else {
		in = bufio.NewScanner(strings.NewReader(data))
	}
	for in.Scan() {
		fmt.Println(in.Text())
	}
}
