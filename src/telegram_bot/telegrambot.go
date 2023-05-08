package telegram_bot

import (
	"fmt"
	"github.com/MamontModest/bot_for_vk/bd"
	"github.com/MamontModest/bot_for_vk/gen"
	"github.com/MamontModest/bot_for_vk/model"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"reflect"
)

func TelegramBot() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err, bot)
		panic(err)
	}
	log.Println("bot started ^__^")
	//устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	UsersStatus := make(model.UsersStatus)
	LastUserService := make(model.LastUserService)
	LastUserCommand := make(model.LastUserCommand)

	for update := range updates {
		fmt.Println(update.Message.Chat)
		if update.Message == nil {
			continue
		}
		//Проверяем что от пользователья пришло именно текстовое сообщение
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			uid := int(update.Message.Chat.ID)
			switch update.Message.Text {
			case "/start":
				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Привет , я бот для генерации паролей и хранения паролей к сервисам\n"+
						"/set (установить пароль и логин к сервису)\n"+
						"/del (удалить пароль и логин к сервису)\n"+
						"/get (получить пароль и логин к сервису)\n"+
						"/gen (сгенерировать пароль к сервису)\n")
				bot.Send(msg)
				LastUserService[uid] = model.NewUser(uid)

				msg = tgbotapi.NewMessage(int64(uid), "Выберите команду")
				msg.ReplyMarkup = KeyBoardLobby()
				bot.Send(msg)

			case "/set":
				//управляем статусом + сбрасываем кэш
				LastUserCommand[uid] = "set"
				UsersStatus[uid] = 1
				LastUserService[uid] = model.NewUser(uid)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вводи название сервиса, для сохранения.")
				msg.ReplyMarkup = nil
				bot.Send(msg)

			case "/get":
				//управляем статусом + сбрасываем кэш
				LastUserCommand[uid] = "get"
				UsersStatus[uid] = 1
				LastUserService[uid] = model.NewUser(uid)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вводи название сервиса, для поиска.")
				msg.ReplyMarkup = nil
				bot.Send(msg)

			case "/del":
				//управляем статусом + сбрасываем кэш
				LastUserCommand[uid] = "del"
				UsersStatus[uid] = 1
				LastUserService[uid] = model.NewUser(uid)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вводи название сервиса, для удаления.")
				msg.ReplyMarkup = nil
				bot.Send(msg)

			case "/gen":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "your password\n"+gen.Password())
				bot.Send(msg)

			default:
				switch LastUserCommand[uid] {

				case "set":
					switch UsersStatus[uid] {
					case 1:
						//добавляем в кэш + обновляем статус
						LastUserService[uid].Service = update.Message.Text
						UsersStatus[uid] = 2
						msg := tgbotapi.NewMessage(int64(uid), "Напишите логин")
						bot.Send(msg)

					case 2:
						LastUserService[uid].Login = update.Message.Text
						UsersStatus[uid] = 3
						msg := tgbotapi.NewMessage(int64(uid), "Напишите пароль")
						bot.Send(msg)

					case 3:
						LastUserService[uid].Password = update.Message.Text
						UsersStatus[uid] = 4

						button1 := tgbotapi.NewKeyboardButton("Да")
						button2 := tgbotapi.NewKeyboardButton("Нет")
						data := tgbotapi.NewKeyboardButtonRow(button1, button2)
						response := tgbotapi.NewReplyKeyboard(data)

						str := fmt.Sprintf(
							"Данные верны ?\nСервис : %s\nЛогин : %s\nПароль : %s",
							LastUserService[uid].Service, LastUserService[uid].Login, LastUserService[uid].Password)
						msg := tgbotapi.NewMessage(int64(uid), str)
						msg.ReplyMarkup = response
						bot.Send(msg)
					case 4:
						if update.Message.Text == "Да" {
							err := bd.CreatePassword(LastUserService[uid])
							LastUserCommand[uid] = ""
							UsersStatus[uid] = 0

							if err != nil {
								log.Println(err)
								msg := tgbotapi.NewMessage(int64(uid), "Сохранить логин и пароль не удалось , попробуйте поменять логин или пароль")
								msg.ReplyMarkup = KeyBoardLobby()
								bot.Send(msg)
								break
							}
							msg := tgbotapi.NewMessage(int64(uid), "Пароль успешно сохранён")
							msg.ReplyMarkup = KeyBoardLobby()
							bot.Send(msg)
							break

						}
						LastUserCommand[uid] = ""
						UsersStatus[uid] = 0

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Попробуем сначала \nВводи название сервиса")
						bot.Send(msg)
					}
				case "get":
					switch UsersStatus[uid] {
					case 1:
						LastUserService[uid].Service = update.Message.Text
						err := bd.SearchPassword(LastUserService[uid])
						if err != nil {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не правильный ввод сервиса попробуйте снова")
							bot.Send(msg)

							LastUserCommand[uid] = ""
							UsersStatus[uid] = 0
							break
						}
						str := fmt.Sprintf(
							"Вот твои данные ! Время до их удаленния 10 минут\nСервис : %s\nЛогин : %s\nПароль : %s",
							LastUserService[uid].Service, LastUserService[uid].Login, LastUserService[uid].Password)

						LastUserCommand[uid] = ""
						UsersStatus[uid] = 0

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
						msg.ReplyMarkup = KeyBoardLobby()
						bot.Send(msg)
					}
				case "del":
					switch UsersStatus[uid] {
					case 1:
						LastUserService[uid].Service = update.Message.Text
						err := bd.DeletePassword(LastUserService[uid])
						if err != nil {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не правильный ввод сервиса попробуйте снова")
							bot.Send(msg)

							LastUserCommand[uid] = ""
							UsersStatus[uid] = 0
							break
						}
						str := fmt.Sprintf(
							"Сервис: %s успешно удалён .",
							LastUserService[uid].Service)

						LastUserCommand[uid] = ""
						UsersStatus[uid] = 0

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
						msg.ReplyMarkup = KeyBoardLobby()
						bot.Send(msg)
					}

				}
			}
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я умею принимать только сообщения"))
		}
	}
}

func KeyBoardLobby() tgbotapi.ReplyKeyboardMarkup {
	button1 := tgbotapi.NewKeyboardButton("/set")
	button2 := tgbotapi.NewKeyboardButton("/del")
	button3 := tgbotapi.NewKeyboardButton("/get")
	button4 := tgbotapi.NewKeyboardButton("/gen")
	data1 := tgbotapi.NewKeyboardButtonRow(button1, button2)
	data2 := tgbotapi.NewKeyboardButtonRow(button3, button4)
	response := tgbotapi.NewReplyKeyboard(data1, data2)
	return response
}
