package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/Syfaro/telegram-bot-api"
)

func telegramBot() {

    //Создаем бота
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
    if err != nil {
		log.Fatal(err,bot)
        panic(err)
    }

    //Устанавливаем время обновления
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

	log.Println("bot started ^__^")
    //Получаем обновления от бота 
    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        fmt.Println(update.Message.Chat)
        if update.Message == nil {
            continue
        }

        //Проверяем что от пользователья пришло именно текстовое сообщение
        if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

            switch update.Message.Text {
            case "/start":

                //Отправлем сообщение
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a wikipedia bot, i can search information in a wikipedia, send me something what you want find in Wikipedia.")
                bot.Send(msg)

			default:
				fmt.Print(3)
		}
    }
}
}