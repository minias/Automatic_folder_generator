package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

type folderMap map[string]interface{}

type Dirs struct {
	Dirs []Dir `xml:"dir"`
}

// the dir struct, this contains our permission attribute
type Dir struct {
	Name       string `xml:"name,attr"`
	Permission string `xml:"permission,attr"`
}

func main() {
	app := &cli.App{
		Name:  "auto_dir",
		Usage: "Read the yml file to create a folder automatically.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   "sample.xml",
				Usage:   "make dir import file",
			},
		},
		Action: func(cCtx *cli.Context) error {
			filename := cCtx.String("input")
			//folderMap, err := readYML(filename)
			dir, err := readXML(filename)
			if err != nil {
				fmt.Printf("read file %s #%v \n", filename, err)
			}
			fmt.Printf("-> %s\n", dir)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// read yml file
//
//lint:ignore U1000 Ignore unused function
func readYML(f string) (folderMap, error) {
	yamlFile, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	obj := make(map[string]interface{})
	if err = yaml.Unmarshal(yamlFile, obj); err != nil {
		return nil, err
	} else {
		return obj, nil
	}
}

// read xml file
func readXML(f string) (Dirs, error) {
	fmt.Printf("file: %s\n", f)
	var dirs Dirs
	// Open our xmlFile
	xmlFile, err := os.Open(f)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return dirs, err
	}
	byteValue, _ := io.ReadAll(xmlFile)
	fmt.Printf("byteValue: %s\n", byteValue)
	defer xmlFile.Close()

	err = xml.Unmarshal(byteValue, &dirs)
	if err != nil {
		fmt.Println(err)
		return dirs, err
	}
	fmt.Printf("dir cnt: %v\n", len(dirs.Dirs))
	for i := 0; i < len(dirs.Dirs); i++ {
		fmt.Println("Permission: " + dirs.Dirs[i].Permission)
		fmt.Println("Name: " + dirs.Dirs[i].Name)
	}

	return dirs, nil
}
