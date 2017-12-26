package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/robertbasic/d20/roll"
)

type Roll struct{}

func (rh *Roll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer redirect(w, r)

	if r.Method != "POST" {
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	var dice = make([]roll.Dice, 7)

	for field, value := range r.Form {
		val, err := convv(value)
		if err != nil {
			log.Println(err)
			return
		}

		switch field {
		case "d4":
		case "d6":
			dice = append(dice, roll.NewD6(val))
		case "d8":
		case "d10":
		case "d12":
		case "d20-normal":
		case "d20-advantage":
		case "d20-disadvantage":
		case "d100":
		}
		log.Println(field, val)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func convv(value []string) (int, error) {
	if len(value) != 1 {
		return 0, errors.New("Can not convert submitted value to int")
	}

	var v int
	var val = value[0]

	if val == "on" {
		return 1, nil
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return v, nil
}
