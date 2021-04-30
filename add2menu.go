// Creates an .desktop entry in the KDE applications folder
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Info struct {
	Name string
	Path string
	Icon string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execTemplate(info Info, temp *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	err := temp.Execute(buf, info)
	return buf.String(), err
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Creating a menu entry")
	fmt.Println("---------------------")

	fmt.Println("Whats the name of the application?")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	fmt.Println("Where is the executable located?")
	path, _ := reader.ReadString('\n')
	path = strings.Replace(path, "\n", "", -1)

	fmt.Println("Any icon? (provide path)")
	icon, _ := reader.ReadString('\n')
	icon = strings.Replace(icon, "\n", "", -1)

	info := Info{name, path, icon}
	fileTemplate, err := template.New("path").Parse("/usr/share/applications/{{.Name}}.desktop")
	check(err)
	tmpl, err := template.New("desktop").Parse("[Desktop Entry]\nName={{.Name}}\nComment=\nExec={{.Path}}\nIcon={{.Icon}}\nTerminal=false\nType=Application")
	check(err)

	fileName, err := execTemplate(info, fileTemplate)
	check(err)
	fileContent, err := execTemplate(info, tmpl)
	check(err)

	file, err := os.Create(fileName)

	n, err := file.WriteString(fileContent)

	fmt.Printf("wrote %d bytes\n", n)

}
