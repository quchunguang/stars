package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `
Name:
  stars

Usage:
  stars [--spark] [--chars <cs>] [--header] [--width=<n>] <data>...
  stars [--spark] [--chars <cs>] [--header] [--width=<n>] [--cols <csl>] [--reverse]
  stars -h | --help
  stars --version

Options:
  --header      Including the description at beginning of each line of data.
  --spark       Format like ▁▂▃▅▂▇ .
  --width=<n>   Hold width for draw [default: 20].
  --chars <cs>  Characters used by draw recursively. [default: '+-*!#']
  --cols <csl>  Given comma separated list of column numbers. Example for CSL,
                1,2,3 or -3,-2,-1.
  --reverse     Reverse the data format with rows and columns. That makes easy
                to process the standard DSV files.
  -h --help     Show this screen.
  --version     Show version.

Data Format:
  If no data is given, read line by line from STDIN. The data should like this,

  title a,1,2,3
  title b,2,3,4
  title c,4,5,6

  title a 1 2 3
  title b 2 3 4
  title c 4 5 6

  The <data> line given should be formated like this,

  1 2 3
  1 2 3; 2 3 4; 4 5 6
  title a,1,2,3;title b,2,3,4;title c,4,5,6
`
	arguments, _ := docopt.Parse(usage, nil, true, "stars v0.01", false)
	fmt.Println(arguments)
}
