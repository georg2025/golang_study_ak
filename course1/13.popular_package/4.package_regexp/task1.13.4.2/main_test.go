package main

import (
	"reflect"
	"testing"
)

func TestCensorAds(t *testing.T) {
	censor := make(map[string]string)
	censor["hello"] = "Hello"
	censor["moo"] = "Forward"
	censor["страйк"] = "Strike"
	censor["macbook pro"] = "Macbook Pro"
	censor["велосипед merida"] = "Телефон Apple"
	ads := []Ad{
		{
			Title:    "HelLo moo",
			Описание: "Its a страйк moo",
		},
		{
			Title:    "ВеЛоСиПед merida",
			Описание: "Продаю MACbook pro",
		},
	}
	result := censorAds(ads, censor)
	expected := []Ad{
		{
			Title:    "Hello Forward",
			Описание: "Its a Strike Forward",
		},
		{
			Title:    "Телефон Apple",
			Описание: "Продаю Macbook Pro",
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
}
