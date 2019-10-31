package app

import(
	f_http "github.com/fatezhou/tools/http"
	sys "github.com/fatezhou/tools/sys"
	"html/template"
	"net/http"
)

type Server struct{
	t *template.Template
}

func (s *Server)Index(data map[string]string, writer http.ResponseWriter){
	s.t.Execute(writer, nil)
}

func BindFunction(name string)*template.Template{
	funcMap := template.FuncMap{}
	t := template.New(name).Funcs(funcMap)
	return t
}

func LoadTemplate(tWithFunc *template.Template, name ...string)*template.Template{
	html, _ := tWithFunc.ParseFiles(name...)
	return html
}

var httpServer f_http.HttpServer

func (s *Server)Init(){
	httpServer.AddHandle("/index", s.Index)
	t := BindFunction("index.html")
	s.t = LoadTemplate(t, sys.GetProcessDir() + "/static/html/index.html")
	httpServer.SetStaticPath("/js/", sys.GetProcessDir() + "/static/js/")
	httpServer.Run("0.0.0.0", 7788)
}




