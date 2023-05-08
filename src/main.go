package main

import (
	"time"

	"github.com/MamontModest/bot_for_vk/src/bd"
)

func main() {
	time.Sleep(10)
	bd.CreateTable()
	//Вызываем бота
	telegramBot()

}