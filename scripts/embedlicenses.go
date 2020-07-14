package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	out, _ := os.Create("licenses.go")
	out.Write([]byte("package main \n\nvar licenseString=`"))
	out.WriteString("LICENSES\n")
	out.WriteString("========\n")
	defer out.Close()
	filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if strings.HasPrefix("LICENSE", info.Name()) {
			rel, err := filepath.Rel(cwd, path)
			if err != nil {
				panic(err)
			}
			dir := filepath.Dir(rel)
			out.WriteString("FILES: " + dir + "\n")
			c, err := ioutil.ReadFile(path)
			str := string(c)
			str = strings.ReplaceAll(str, "`", "`"+" + \"`\" + `") // escape backticks in license text for go src
			if err != nil {
				panic(err)
			}
			out.WriteString(str)
			out.WriteString("\n\n")

		}
		return nil
	})
	out.Write([]byte("`\n"))
}
