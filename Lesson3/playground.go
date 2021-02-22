package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	botToken := ""
	botAPI := "https://api.telegram.org/"
	botURL := botAPI + botToken
	offset := 0
	for {
		updates, err := getUpdates(botURL, offset)
		if err != nil {
			panic(err)
		}
		for _, update := range updates {
			err = respond(botURL, update)
			if err != nil {
				panic(err)
			}
			offset = update.Update_id + 1
		}
	}
}

func getUpdates(botURL string, offset int) ([]Update, error) {
	allSource, err := http.Get(botURL + "/getUpdates?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer allSource.Body.Close()
	body, err_ := ioutil.ReadAll(allSource.Body)
	if err_ != nil {
		return nil, err_
	}
	result := Respond{}
	_err := json.Unmarshal(body, &result)
	if _err != nil {
		return nil, _err
	}
	return result.Update, _err
}

func respond(botURL string, update Update) error {
	currentMessage := BotMessage{}
	update.Message.Text = strings.ToLower(update.Message.Text)
	if update.Message.Text == "привет" {
		currentMessage.Text = "И тебе привет!"
	} else if update.Message.Text == "как тебя зовут?" {
		currentMessage.Text = "Меня звать Алан! Я Чеченец!"
	} else if update.Message.Text == "как зовут твою жену?" {
		currentMessage.Text = "Ее тоже звать Алан! Она тоже чеченец!"
	} else {
		currentMessage.Text = "Бот не человек, он не понимает такое..."
	}
	currentMessage.ChatID = update.Message.Chat.ChatID
	buffer, err := json.Marshal(currentMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return err
	}
	return nil
}
