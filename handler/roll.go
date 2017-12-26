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

	var dice []roll.Dice

	for field, value := range r.Form {
		val, err := convv(value)
		if err != nil {
			log.Println(err)
			return
		}

		if val == 0 {
			continue
		}

		var d roll.Dice
		switch field {
		case "d4":
			d = roll.NewD4(val)
		case "d6":
			d = roll.NewD6(val)
		case "d8":
			d = roll.NewD8(val)
		case "d10":
			d = roll.NewD10(val)
		case "d12":
			d = roll.NewD12(val)
		case "d100":
			d = roll.NewD100(val)
		case "d20-normal":
			d = roll.NewD20Normal()
		case "d20-advantage":
			d = roll.NewD20Advantage()
		case "d20-disadvantage":
			d = roll.NewD20Disadvantage()
		}

		dice = append(dice, d)
	}

	log.Println(len(dice))
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
