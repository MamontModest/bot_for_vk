package main

import (
	"github.com/MamontModest/bot_for_vk/bd"
	"github.com/MamontModest/bot_for_vk/telegram_bot"
	"log"
	"time"
)

func main() {
	//даем развернутся бд
	time.Sleep(10 * time.Second)
	err := bd.CreateTables()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	//Вызываем бота
	telegram_bot.TelegramBot()

}
