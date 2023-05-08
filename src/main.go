package main

import (
	"time"

	"github.com/MamontModesttg-bot-passwords-for_VK/src/bd"
)

func main() {
	time.Sleep(10)
	bd.CreateTable()
	//Вызываем бота
	telegramBot()

}