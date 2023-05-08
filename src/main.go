package main

import (
	"github.com/MamontModest/bot_for_vk/bd"
	"github.com/MamontModest/bot_for_vk/telegram_bot"
	"log"
	"os"
)

func main() {
	//даём время контейнеру создать базу данных
	os.Setenv("TOKEN", "6101250366:AAFd8UZBbg2pTv8ic1KTRTMo77vRdluQFyg")
	os.Setenv("DBNAME", "bot_vk")
	os.Setenv("PORT", "5432")
	os.Setenv("USER", "postgres")
	os.Setenv("PASSWORD", "QWertas1122")
	os.Setenv("HOST", "localhost")
	os.Setenv("SSLMODE", "disable")
	err := bd.CreateTables()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	//Вызываем бота
	telegram_bot.TelegramBot()

}
