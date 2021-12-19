/*
 ------------------- ЗАДАЧИ -------------------
 Основные:
	Передалать под webhook

Дополнительно:
	Доделать Систему Админов
 ----------------------------------------------
*/

package main

// ------------------- IMPORTS -------------------
import (
	"io/ioutil"
	"log"

	"AnashArt.bot/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ------------------- CONSTS -------------------
const botAPI = "5015857552:AAGHOqwNAeeJ4Su0Rnlu9UOAz6MaO3IDpng"

const octopusPATH = "img/octopus.jpg"
const shrimpPATH = "img/shrimp.jpg"

const wlankasperID = 853634511
const anasharmsID = 726736906

// ------------------- KEYBOARDS-------------------
// --------- ORDER CHOICE SYSTEM ---------
var OrderSystem = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("В Telegram", "Telegram"),
		tgbotapi.NewInlineKeyboardButtonURL("В Insagram", "https://www.instagram.com/anash.art/"),
	),
)

// --------- ORDER CHOICE PRINT ---------
var OrderPrint = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("МОКРИЙ", "МОКРИЙ"),
		tgbotapi.NewInlineKeyboardButtonData("КРЕВЭД", "КРЕВЭД"),
	),
)

// --------- ORDER CHOICE SIZE ---------
var OrderSize = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("S", "S"),
		tgbotapi.NewInlineKeyboardButtonData("M", "M"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("L", "L"),
		tgbotapi.NewInlineKeyboardButtonData("XL", "XL"),
	),
)

// --------- ORDER CHOICE PAYMENT ---------
var OrderPayment = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		// tgbotapi.NewInlineKeyboardButtonData("Telegram Pay", "TelegramPay"),
		tgbotapi.NewInlineKeyboardButtonData("Перевод на карту", "Перевод"),
	),
)

// ------------------------------------ ADMIN KEYBOARDS ------------------------------------
// --------- ADMIN SETTINGS ---------
var AdminSettings = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Что-то добавить", "add"),
		tgbotapi.NewInlineKeyboardButtonData("Что-то удалить", "del"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать все", "all"),
	),
)

// ------------------- MAPS -------------------
var OrderInfoMap map[int64]*db.OrderInfo
var InputState int = 0

// ------------------- FUNCS -------------------
func init() {
	OrderInfoMap = make(map[int64]*db.OrderInfo)
}

func main() {

	// --------- INIT BOT ---------
	bot, err := tgbotapi.NewBotAPI(botAPI)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// --------- CHECK NEW MESSAGE LOOP ---------
	for update := range updates {

		standartSendMessage := func(msg tgbotapi.MessageConfig) {
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		}

		standartCallbackCheck := func() {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
		}

		orderSetSize := func(size string) {
			standartCallbackCheck()
			OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Size = size
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Теперь введите ваш email")
			standartSendMessage(msg)

			InputState = 5
		}

		orderSetPrint := func(print string) {
			standartCallbackCheck()
			OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Print = print
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Теперь нужно выбрать свой размер")
			msg.ReplyMarkup = OrderSize
			standartSendMessage(msg)
		}

		sendPhoto := func(path string) {
			photoBytes, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			photoFileBytes := tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: photoBytes,
			}
			bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, photoFileBytes))
		}

		if update.Message != nil {
			if update.Message.IsCommand() {
				if update.Message.Command() == "start" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот AnashArt и вот что я умею:\n\n  /start - Начальное меню 🧾\n  /show - Покажу все наши коллекции ✨\n  /price - Покажу Прайс-Лист 💸\n  /order - Оформлю заказ 📦\n  /help - Позову Администратора ⁉️\n\nНаш официальный сайт: https://AnashArt.ru\nНаш официальный Instagram: https://www.instagram.com/anash.art/")
					standartSendMessage(msg)

				} else if update.Message.Command() == "show" {
					sendPhoto(octopusPATH)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Принт МОКРИЙ")
					standartSendMessage(msg)

					sendPhoto(shrimpPATH)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Принт МОКРИЙ")
					standartSendMessage(msg)

				} else if update.Message.Command() == "help" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В скором времени с вами свяжется наш администратор ...")
					standartSendMessage(msg)

					bot.Send(tgbotapi.NewMessage(wlankasperID, "PROBLEM @"+update.Message.From.UserName))
					bot.Send(tgbotapi.NewMessage(anasharmsID, "PROBLEM @"+update.Message.From.UserName))

				} else if update.Message.Command() == "price" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Price List:\n\n  Футболка 'МОКРИЙ' - 3590\n  Футболка 'КРЕВЭД' - 3590")
					standartSendMessage(msg)

				} else if update.Message.Command() == "order" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы можете заказать что либо прямо здесь или обратиться к администратору в нашем Instagram")
					msg.ReplyMarkup = OrderSystem
					standartSendMessage(msg)

				} else if update.Message.Command() == "admin" && update.Message.Chat.ID == wlankasperID || update.Message.Chat.ID == wlankasperID {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет Насть)")
					msg.ReplyMarkup = AdminSettings
					standartSendMessage(msg)

				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я даже не знаю как на это ответить(")
					standartSendMessage(msg)
				}

			} else if InputState != 0 {
				switch InputState {
				case 1:

				case 2:

				case 3:

				case 4:

					// --------- EMAIL---------
				case 5:
					OrderInfoMap[update.Message.Chat.ID].Email = update.Message.Text

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично, остался последний шаг!\n\nВведите адрес доставки и телефон в формате: \nГород, Улица, Дом, Номер_телефона_для_связи")
					standartSendMessage(msg)
					InputState = 6

					// --------- ADDRES---------
				case 6:
					OrderInfoMap[update.Message.Chat.ID].Addres = update.Message.Text

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично! Проверьте свой заказ и выберите способ оплаты\n\nВаш заказ:\nПринт - "+OrderInfoMap[update.Message.Chat.ID].Print+"\nРазмер - "+OrderInfoMap[update.Message.Chat.ID].Size+"\nEmail - "+OrderInfoMap[update.Message.Chat.ID].Email+"\nДоставка - "+OrderInfoMap[update.Message.Chat.ID].Addres)
					msg.ReplyMarkup = OrderPayment
					standartSendMessage(msg)

					InputState = 0
				}

			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я даже не знаю как на это ответить(")
				standartSendMessage(msg)
			}
		}

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {

			case "Telegram":
				standartCallbackCheck()

				OrderInfoMap[update.CallbackQuery.Message.Chat.ID] = new(db.OrderInfo)
				OrderInfoMap[update.CallbackQuery.Message.Chat.ID].UserName = update.CallbackQuery.From.UserName

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выбирай какой принт ты хочешь")
				msg.ReplyMarkup = OrderPrint
				standartSendMessage(msg)

			case "МОКРИЙ":
				orderSetPrint("МОКРИЙ")

			case "КРЕВЭД":
				orderSetPrint("КРЕВЭД")

			case "S":
				orderSetSize("S")

			case "M":
				orderSetSize("M")

			case "L":
				orderSetSize("L")

			case "XL":
				orderSetSize("XL")

			case "Перевод":
				standartCallbackCheck()

				OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Payment = "Перевод на карту"
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "СберБанк - 1000 1000 1000 1000\nТинькофф - 1000 1000 1000 1000\n\nПосле перевода вам напишет наш Администратор чтобы подтвердить заказ и сообщит ближайщую дату доставки")
				standartSendMessage(msg)

				msg = tgbotapi.NewMessage(wlankasperID, "NEW ORDER\n\nName: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].UserName+"\nEmail: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Email+"\nAddres: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Addres+"\nPrint: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Print+"\nSize: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Size+"\nPayment: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Payment+"\nStatus: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Status)
				standartSendMessage(msg)

				msg = tgbotapi.NewMessage(anasharmsID, "NEW ORDER\n\nName: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].UserName+"\nEmail: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Email+"\nAddres: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Addres+"\nPrint: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Print+"\nSize: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Size+"\nPayment: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Payment+"\nStatus: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Status)
				standartSendMessage(msg)

				// ------------------------------------ CALLBACK FOR ADMIN ------------------------------------
			case "add":
				standartCallbackCheck()
				// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Меню:\n  1 - Добавить Принт\n  2 - Добавить размер\n  3 - Добавить Цвет")
				// standartSendMessage(msg)
				tree := db.Tree{}
				tree.Add("root/Collection 1/S/3")
				tree.Add("root/Collection 1/XL/7")
				tree.Add("root/Collection 2/xx/45")
				tree.Add("root/Collection 3/Ss/35")

				str := tree.TreePrint(true, "", "")
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, str)
				standartSendMessage(msg)

				// InputState = 1

				// case "del":
				// 	standartCallbackCheck()
				// 	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введи размер (Заглавными английскими: S, XLL, ...): ")
				// 	standartSendMessage(msg)

				// 	InputState = 2

				// case "all":
				// 	standartCallbackCheck()
				// 	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введи размер (Заглавными английскими: S, XLL, ...): ")

				// 	var temp string
				// 	for key, value := range Products_db {
				// 		temp = "Принт: " + key + "Размеры: "
				// 	}
				// 	standartSendMessage(msg)
			}
		}
	}
}
