package main

import (
	"bufio"
	"fmt"
	"github.com/docopt/docopt-go"
	"math"
	"os"
	"strconv"
	"strings"
)

var Conf struct {
	Data      []string // data
	Style     string   // --spark
	IsHeader  bool     // -H
	FieldSep  string   // -F
	Width     int      // -w
	Columns   string   // -C
	IsReverse bool     // -R
	CharList  string   // -L
}

func ParseConf() {
	usage := `
Name:
  stars

Usage:
    stars [--spark]
          [-F <fs>] [-H] [-L <cl>] [-w <n>] [-R] [-C <col>]
          [<data>...]
    stars -h | --help
    stars --version

Options:
    --spark    ▁▂▃▅▂▇▁▂▃ style instead of default.
    -C <col>   Given comma separated list of column numbers.
               Example for CSL: 1,2,3 or -3,-2,-1.
    -F <fs>    Use fs for the input field separator [default:  ] that is " ".
    -H         Including the description at beginning of each line of data.
    -R         Reverse the data format with rows and columns. That makes
               easy to process the standard DSV files.
    -L <cl>    Character List used by draw recursively. [default: +-*!#]
    -w <n>     Hold width for draw [default: 40].
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

	Conf.Data = arguments["<data>"].([]string)
	if arguments["--spark"].(bool) {
		Conf.Style = "spark"
	} else {
		Conf.Style = "star"
	}
	Conf.IsHeader = arguments["-H"].(bool)
	Conf.FieldSep = arguments["-F"].(string)
	Conf.Width, _ = strconv.Atoi(arguments["-w"].(string))
	if arguments["-C"] == nil {
		Conf.Columns = ""
	} else {
		Conf.Columns = arguments["-C"].(string)
	}
	Conf.IsReverse = arguments["-R"].(bool)
	Conf.CharList = arguments["-L"].(string)

	fmt.Println(Conf)
}

func ProcessLine(data_list []string) {
	var data []float64
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
		} else if f > max {
			max = f
		}
	}
	return
}

func Star(data []float64) {
	_, _, sum := MinMaxSum(data)
	step := sum / float64(Conf.Width)

	var star []byte
	for i, f := range data {
		count := int(f / step)
		for j := 0; j < count; j++ {
			star = append(star, Conf.CharList[i%len(Conf.CharList)])
		}
	}
	fmt.Println(string(star))
}

func Spark(data []float64) {
	smbs := []rune("▁▂▃▄▅▆▇█")
	min, max, _ := MinMaxSum(data)
	step := (max - min) / float64(len(smbs)-1)
	var spark []rune
	for _, f := range data {
		index := int((f - min) / step)
		fmt.Print(index)
		spark = append(spark, smbs[index])
	}
	fmt.Println(string(spark))
}

func main() {
	ParseConf()

	if len(Conf.Data) == 0 {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			data_list := strings.Split(in.Text(), Conf.FieldSep)
			ProcessLine(data_list)
		}
	} else {
		if Conf.FieldSep == " " {
			// By default, the space separator will cause data already a list.
			ProcessLine(Conf.Data)
		} else {
			sdata := strings.Join(Conf.Data, "")
			data_list := strings.Split(sdata, Conf.FieldSep)
			ProcessLine(data_list)
		}
	}
}
