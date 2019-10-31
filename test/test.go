package main

import (
	"github.com/fatezhou/tools/sys"
	"html/template"
	"os"
	"strings"
)

type headerData struct{
	Str string
	i int
}

func Add(a int32)int32{
	return a + a
}

func TestTemplate(){
	dir := sys.GetProcessDir()
	files := dir + "/main.html"
	funcMap := template.FuncMap{"Add":Add}
	tHeader, err := template.New("main.html").Funcs(funcMap).ParseFiles(files, dir + "/header.html")
	if err != nil{
		panic(err)
	}else{
		headerData := headerData{Str:"hello", i:1}
		tHeader.Execute(os.Stdout, headerData)
	}
}

func TestIndex(){
	filePath:= "xxxx.xxxx/x/1234.js"
	npos := strings.LastIndex(filePath, "a")
	print(npos)
	ext := filePath[npos+1:]
	print(ext)
}

func main(){
	TestIndex()
}

