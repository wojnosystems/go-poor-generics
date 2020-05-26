package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/wojnosystems/poor-generic/pkg"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	topError := (&cli.App{
		Name:  "poor-generics",
		Usage: "generate generics the poor-man's way in GoLang",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "namesToPrimitiveTypes",
						Aliases:  []string{"m"},
						Usage:    "Int=int,Int64=int64,Bool=bool",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "templateFile",
						Aliases:  []string{"t"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "outFile",
						Aliases:  []string{"o"},
						Value:    "-",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "package",
						Aliases:  []string{"p"},
						Required: true,
					},
				},
				Action: func(context *cli.Context) (err error) {
					nameMap, err := keyValuePairToMap(context.String("namesToPrimitiveTypes"))
					if err != nil {
						return
					}

					var source *template.Template
					source, err = template.ParseFiles(context.String("templateFile"))
					if err != nil {
						return
					}

					outputFile, outputDeferClose, err := getGeneratorOutputStream(context.String("outFile"))
					if err != nil {
						return
					}
					defer outputDeferClose()
					return pkg.Generate(context.String("package"), nameMap, source, outputFile)
				},
			},
		},
	}).Run(os.Args)
	if topError != nil {
		log.Panic(topError.Error())
	}
}

func keyValuePairToMap(value string) (nameMap map[string]string, err error) {
	nameMap = make(map[string]string)
	if len(value) == 0 {
		return
	}
	elements := strings.Split(value, ",")
	for _, elementPair := range elements {
		keyAndValue := strings.Split(elementPair, "=")
		if len(keyAndValue) != 2 {
			if len(keyAndValue) < 2 {
				err = fmt.Errorf("both a key and value are required, only a key was found, separate with equals (=)")
			} else {
				err = fmt.Errorf("only 1 key and 1 value are permitted, you cannot have more than 1 equal sign for each key assignment")
			}
			return
		} else {
			nameMap[keyAndValue[0]] = keyAndValue[1]
		}
	}
	return
}

func getGeneratorOutputStream(outFilePath string) (writer io.WriteCloser, closer func(), err error) {
	closer = func() {}
	if outFilePath == "-" {
		writer = os.Stdout
	} else {
		writer, err = os.OpenFile(outFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return
		}
		closer = func() { _ = writer.Close() }
	}
	return
}
