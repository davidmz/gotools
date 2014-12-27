package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		inFileName  string
		outFileName string
		varName     string
		pkgName     string
		showHelp    bool

		err error
	)

	flag.StringVar(&inFileName, "in", "", "Input file or stdin if empty")
	flag.StringVar(&outFileName, "out", "", "Output file or stdout if empty")
	flag.StringVar(&varName, "var", "foo", "Variable name")
	flag.StringVar(&pkgName, "pkg", "main", "Package name")
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.Parse()

	if showHelp {
		flag.PrintDefaults()
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout

	if inFileName != "" {
		inFile, err = os.Open(inFileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer inFile.Close()
	}

	if outFileName != "" {
		outFile, err = os.Create(outFileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer outFile.Close()
	}

	fmt.Fprintf(outFile, "package %s\n\n", pkgName)
	fmt.Fprintf(outFile, "var %s = []byte(", varName)
	buf := make([]byte, 24)
	firstLine := true
	for {
		n, err := inFile.Read(buf)
		if n > 0 {
			s := strings.Replace(fmt.Sprintf(`\x% x`, buf[:n]), ` `, `\x`, -1)
			if firstLine {
				firstLine = false
				fmt.Fprintf(outFile, "\"%s\"", s)
			} else {
				fmt.Fprintf(outFile, " +\n        \"%s\"", s)
			}
		}
		if err == io.EOF {
			fmt.Fprint(outFile, ")\n")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
	}
}
