package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aasmall/asciigraph"
)

var (
	height   uint
	width    uint
	offset   uint = 3
	caption  string
	strdata  string
	strxdata string
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "%s expects data points from stdin. Invalid values are logged to stderr.\n", os.Args[0])
	}
	flag.UintVar(&height, "h", height, "`height` in text rows, 0 for auto-scaling")
	flag.UintVar(&width, "w", width, "`width` in columns, 0 for auto-scaling")
	flag.UintVar(&offset, "o", offset, "`offset` in columns, for the label")
	flag.StringVar(&caption, "c", caption, "`caption` for the graph")
	flag.StringVar(&strdata, "d", strdata, "'data' for the graph")
	flag.StringVar(&strxdata, "x", strxdata, "'x axis data' for the graph")
	flag.Parse()

	data := make([]float64, 0, 64)
	xdata := make([]float64, 0, 64)
	var s *bufio.Scanner
	if strdata == "" {
		s = bufio.NewScanner(os.Stdin)
	} else {
		s = bufio.NewScanner(strings.NewReader(strdata))
	}
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Text()
		p, err := strconv.ParseFloat(word, 64)
		if err != nil {
			log.Printf("ignore %q: cannot parse value", word)
			continue
		}
		data = append(data, p)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	if strxdata != "" {
		s = bufio.NewScanner(strings.NewReader(strxdata))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			word := s.Text()
			p, err := strconv.ParseFloat(word, 64)
			if err != nil {
				log.Printf("ignore %q: cannot parse value", word)
				continue
			}
			xdata = append(xdata, p)
		}
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if len(data) == 0 {
		log.Fatal("no data")
	}

	plot := asciigraph.Plot(data,
		asciigraph.Height(int(height)),
		asciigraph.Width(int(width)),
		asciigraph.Offset(int(offset)),
		asciigraph.Caption(caption),
		asciigraph.FixedXAxis(xdata))

	fmt.Println(plot)
}
