package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"encoding/json"

	"encoding/base64"

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

		var sides int
		switch field {
		case "d4":
			sides = 4
		case "d6":
			sides = 6
		case "d8":
			sides = 8
		case "d10":
			sides = 10
		case "d12":
			sides = 12
		case "d100":
			sides = 100
		case "d20-normal":
			sides = 20
		case "d20-advantage":
			sides = 20
		case "d20-disadvantage":
			sides = 20
		}

		d := roll.NewDice(sides, val)

		dice = append(dice, *d)
	}
	setCookie(dice, w)
}

func setCookie(dice []roll.Dice, w http.ResponseWriter) {
	j, err := json.Marshal(dice)
	if err != nil {
		log.Println(err)
	}
	c := http.Cookie{
		Name:  "dice",
		Value: base64.StdEncoding.EncodeToString(j),
	}
	http.SetCookie(w, &c)
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
