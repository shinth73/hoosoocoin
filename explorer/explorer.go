package explorer

import (
	"fmt"
	"github.com/shinth73/hoosoocoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

var templates *template.Template

func home(rw http.ResponseWriter, request *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlcoks()}
	err := templates.ExecuteTemplate(rw, "home", data)
	if err != nil {
		return
	}
}

func add(rw http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		err := templates.ExecuteTemplate(rw, "add", nil)
		if err != nil {
			return
		}
	case "POST":
		err := request.ParseForm()
		if err != nil {
			return
		}
		data := request.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, request, "/", http.StatusPermanentRedirect)
	}
}

func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://127.0.0.1%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
