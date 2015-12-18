package main

import (
	"bufio"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/labstack/gommon/color"
	"math"
	"os"
	"strconv"
	"strings"
)

type color_fun []func(msg interface{}, styles ...string) string

type range_t struct {
	from, to float64
}

var Conf struct {
	Files     []string // files
	Style     string   // --spark
	IsHeader  bool     // -H
	FieldSep  string   // -F
	Width     int      // -w
	Columns   string   // -C
	IsReverse bool     // -R
	Range     range_t  // -r
	CharList  string   // -L
	ColorList color_fun
	Header    string
	Text      string
}

func ParseConf() {
	usage := `
Name:
  stars

Feature:
  * Styles growing continued.
  * The data feed from STDIN processed streamed.

Usage:
    stars [--spark]
          [-C <col>] [-F <fs>] [-L <cl>] [-R] [-r <r>] [-s] [-w <n>]
          [<files>...]
    stars -h | --help
    stars --version

Options:
    --spark    ▁▂▃▅▂▇▁▂▃ style instead of default.
    -C <col>   Given comma separated list of column numbers using as data.
               Example for CSL: 1,2,3 or 1:3 or :-1. [default: :]
    -F <fs>    Use fs for the input field separator [default:  ] that is " ".
    -L <cl>    Character List used by draw recursively. [default: +-*!#]
    -R         Reverse the data format with rows and columns. That makes
               easy to process the standard DSV files.
    -r <r>     Data range of total of each line. Ex. -r "-100:100"
               If not given, automatically scaled if needed.
    -s         Silence mode. Do not output the data itself.
    -w <n>     Holding width for header [default: 40].

    -h --help  Show this screen.
    --version  Show version.

Data Format:
    If no data is given, it'll read from STDIN. The data should like this,

    title a 1 2 3
    title b 2 3 4
    title c 4 5 6

    The <data> line given should be formated like this,

    "title a" 1 2 3  4  5

    title a,1,2,3, 4, 5
`
	arguments, _ := docopt.Parse(usage, nil, true, "stars v0.01", false)
	fmt.Println(arguments)

	// Files of input data
	Conf.Files = arguments["<files>"].([]string)

	// Styles of output
	if arguments["--spark"].(bool) {
		Conf.Style = "spark"
	} else {
		Conf.Style = "star"
	}

	// Field separator
	Conf.FieldSep = arguments["-F"].(string)

	// Holding width for header
	Conf.Width, _ = strconv.Atoi(arguments["-w"].(string))

	// Column numbers using as data
	if arguments["-C"] == nil {
		Conf.Columns = ""
	} else {
		Conf.Columns = arguments["-C"].(string)
	}

	// Reverse the data format with rows and columns
	Conf.IsReverse = arguments["-R"].(bool)

	// Data range of total of each line
	if arguments["-r"] != nil {
		Conf.Range = ParseRange(arguments["-r"].(string))
	} else {
		Conf.Range = range_t{}
	}

	// Character list used by draw recursively
	if arguments["-L"] != nil {
		Conf.CharList = arguments["-L"].(string)
	}

	// Color list used by draw recursively
	Conf.ColorList = color_fun{
		color.Green,
		color.Red,
	}
}

func ParseRange(s string) (ret range_t) {
	r := strings.Split(s, ":")
	fmt.Println(r, len(r))
	ret.from, _ = strconv.ParseFloat(r[0], 64)
	ret.to, _ = strconv.ParseFloat(r[1], 64)
	return
}

func ParseCol(col string) (cols []int) {
	return
}

func ProcessLine(data_list []string) {
	var data []float64
	if Conf.IsHeader {
		Conf.Header = data_list[0]
		// data_list = append(data_list[:0], data_list[1:]...)
		data_list = data_list[1:]
	}
	for _, s := range data_list {
		d, _ := strconv.ParseFloat(s, 64)
		data = append(data, d)
	}
	if Conf.Style == "star" {
		Star(data)
	} else if Conf.Style == "spark" {
		Spark(data)
	}
}

func MinMaxSum(data []float64) (min, max, sum float64) {
	min = math.MaxFloat64
	max = -min
	sum = 0.0
	for _, f := range data {
		sum += f
		if f < min {
			min = f
		}
		if f > max {
			max = f
		}
	}
	return
}

func Star(data []float64) {
	_, _, sum := MinMaxSum(data)
	step := sum / float64(Conf.Width)

	var star string
	for i, f := range data {
		count := int(f / step)
		s := strings.Repeat(string(Conf.CharList[i%len(Conf.CharList)]), count)
		s = Conf.ColorList[i%len(Conf.ColorList)](s)
		star += s
	}

	FormatOut(string(star))
}

func Spark(data []float64) {
	smbs := []rune("▁▂▃▄▅▆▇█")
	min, max, _ := MinMaxSum(data)

	step := (max - min) / float64(len(smbs)-1)
	var spark []rune
	for _, f := range data {
		index := int((f - min) / step)
		spark = append(spark, smbs[index])
	}

	FormatOut(string(spark))
}

func FormatOut(gen string) {
	fmt.Println(gen, "\t|", Conf.Text)
}

func main() {
	ParseConf()

	var in *bufio.Scanner
	if len(Conf.Files) == 0 {
		in = bufio.NewScanner(os.Stdin)
	} else {
		fp, _ := os.Open(Conf.Files[0])
		in = bufio.NewScanner(fp)
	}
	for in.Scan() {
		Conf.Text = in.Text()
		data_list := strings.Split(Conf.Text, Conf.FieldSep)
		ProcessLine(data_list)
	}
}
