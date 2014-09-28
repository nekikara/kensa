package kensa

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

const (
	DATA_FILES = "data/*.data"
	answer     = "answer/%s.answer"
)

var extension, _ = regexp.Compile("rb|py|pl")

type Spec string

func NewSpec() Spec {
	f, _ := filepath.Glob("*")
	var name Spec
	for _, s := range f {
		if extension.MatchString(s) {
			name = Spec(s)
			break
		}
	}
	return name
}

var quote, _ = regexp.Compile("\"")

func (s Spec) ExecuteScript(file string) string {
	script := "./" + string(s)
	cmd := exec.Command(script)
	f, err := os.Open(file)
	if err != nil {
		log.Println("cannot open [[data file]].\n", err)
	}
	defer f.Close()
	cmd.Stdin = f
	out, err := cmd.Output()
	if err != nil {
		log.Printf("\n%sのデータ量が少ないかもしれません", file)
	}
	return "\n" + quote.ReplaceAllString(string(out), "")
}

func ListUpData() []string {
	files, err := filepath.Glob(DATA_FILES)
	if err != nil {
		log.Fatal("Globでエラー: ", err)
	}
	return files
}

func ShowAnswer(file string) string {
	filename := getFileName(file)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("CANNOT OPEN FILE or DO NOT EXIST ", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var str string
	for scanner.Scan() {
		str += scanner.Text() + "\n"
	}
	return "\n" + str
}

var number, _ = regexp.Compile("\\d")

func getFileName(file string) string {
	num := number.FindString(file)
	return fmt.Sprintf(answer, num)
}
