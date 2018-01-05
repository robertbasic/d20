package handler

import (
	"encoding/base64"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/robertbasic/d20/roll"
)

var tpls = []string{
	"template/layout.html",
	"template/home.html",
}
var tpl = template.Must(template.ParseFiles(tpls...))

type Home struct{}

func (hh *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data []roll.Dice
	defer tpl.Execute(w, data)

	c, err := r.Cookie("dice")
	if err != nil {
		return
	}

	data, err = getDice(c.Value)
	if err != nil {
		log.Println(err)
	}
}

func getDice(c string) ([]roll.Dice, error) {
	data, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}

	var dice []roll.Dice
	err = json.Unmarshal(data, &dice)
	if err != nil {
		return nil, err
	}

	return dice, nil
}
