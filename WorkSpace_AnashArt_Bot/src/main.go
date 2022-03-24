package main

// https://github.com/go-telegram-bot-api/telegram-bot-api

// ------------------- IMPORTS -------------------
import (
	"fmt"
	"net/http"
	"time"

	"AnashArt.bot/db"
	"AnashArt.bot/logger"
	"AnashArt.bot/value"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ------------------- CONSTS -------------------
const CERT_PATH = "/root/WorkSpace_AnashArt_Bot/src/cert.pem"
const KEY_PATH = "/root/WorkSpace_AnashArt_Bot/src/cert.key"
const BOT_TOKEN = "5267887349:AAEr95a2kk8B78h5CO2yv8E-IN9W2FxERi4"
const BOT_ADDRESS = "65.108.154.134"
const BOT_PORT = "8443"
const TELEGRAM_URL = "https://t.me/VerusPicturaBot"

const wlankasperID = 853634511
const anasharmsID = 726736906

// const archiveChatID = 672399763

// ------------------- VARS -------------------
var OrderInfoMap map[int64]*db.OrderInfo
var InputState int = 0

var OrderSystem = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Бот", "telegram"),
		tgbotapi.NewInlineKeyboardButtonData("Администратор", "admin"),
	),
)

// --------- ORDER CHOICE PRINT ---------
var OrderPrint = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Samurai Octopus 🐙", "octopus"),
		tgbotapi.NewInlineKeyboardButtonData("Samurai Shrimp 🦐", "shrimp"),
	),
)

// --------- ORDER CHOICE SIZE ---------
var OrderSize = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("S", "S"),
		tgbotapi.NewInlineKeyboardButtonData("M", "M"),
		tgbotapi.NewInlineKeyboardButtonData("L", "L"),
	),
)

// --------- ORDER CHOICE PAYMENT ---------
var OrderPayment = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		// tgbotapi.NewInlineKeyboardButtonData("Оплата BUSD", "busd"),
		tgbotapi.NewInlineKeyboardButtonData("Перевод на карту", "card"),
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

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	webHookInfo, err := tgbotapi.NewWebhookWithCert(fmt.Sprintf("https://%s:%s/%s", BOT_ADDRESS, BOT_PORT, BOT_TOKEN), tgbotapi.FilePath(CERT_PATH))
	logger.ForError(err)
	_, err = bot.Request(webHookInfo)
	logger.ForError(err)
	info, err := bot.GetWebhookInfo()
	logger.ForError(err)
	if info.LastErrorDate != 0 {
		logger.ForString(fmt.Sprintf("Telegram callback failed: %s", info.LastErrorMessage))
	}
}

func main() {

	// --------- INIT STRUCTS ---------
	OrderInfoMap = make(map[int64]*db.OrderInfo)

	// --------- INIT BOT ---------
	logger.ForString(fmt.Sprintf("OK, %v, %v", time.Now().Unix(), time.Now().Weekday()))
	logger.ForError(BotErr)
	setWebhook(NewBot)

	updates := NewBot.ListenForWebhook("/" + NewBot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	// --------- MESSAGE LOOP ---------
	for update := range updates {

		sendPhotoPrints := func(sendTo int64) {
			cfg := tgbotapi.NewMediaGroup(sendTo, []interface{}{
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_octopus)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_many)),
			})
			messages, err := NewBot.SendMediaGroup(cfg)

			logger.ForError(err)

			if messages == nil {
				logger.ForString("No received messages")
			}

			if len(messages) != len(cfg.Media) {
				logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
			}

			cfg = tgbotapi.NewMediaGroup(sendTo, []interface{}{
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_shrimp)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_many)),
			})
			messages, err = NewBot.SendMediaGroup(cfg)

			logger.ForError(err)

			if messages == nil {
				logger.ForString("No received messages")
			}

			if len(messages) != len(cfg.Media) {
				logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
			}
		}

		sendPhotoOctopus := func() {
			cfg := tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_octopus)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_front)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_back)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_zoom)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_many)),
			})
			messages, err := NewBot.SendMediaGroup(cfg)

			logger.ForError(err)

			if messages == nil {
				logger.ForString("No received messages")
			}

			if len(messages) != len(cfg.Media) {
				logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
			}
		}

		sendPhotoShrimp := func() {
			cfg := tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_shrimp)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_front)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_back)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_zoom)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_many)),
			})
			messages, err := NewBot.SendMediaGroup(cfg)

			logger.ForError(err)

			if messages == nil {
				logger.ForString("No received messages")
			}

			if len(messages) != len(cfg.Media) {
				logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
			}
		}

		sendPhotoSize := func(sendTo int64) {
			cfg := tgbotapi.NewMediaGroup(sendTo, []interface{}{
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Size_s)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Size_m)),
				tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Size_l)),
			})
			messages, err := NewBot.SendMediaGroup(cfg)

			logger.ForError(err)

			if messages == nil {
				logger.ForString("No received messages")
			}

			if len(messages) != len(cfg.Media) {
				logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
			}
		}
		standartSendMessage := func(msg tgbotapi.MessageConfig) {
			NewBot.Send(msg)
		}

		standartCallbackCheck := func() {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := NewBot.Request(callback); err != nil {
				panic(err)
			}
		}

		orderSetSize := func(size string) {
			standartCallbackCheck()
			OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Size = size
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Теперь введите ваш email")
			standartSendMessage(msg)

			InputState = 1
		}

		orderSetPrint := func(print string) {
			standartCallbackCheck()
			OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Print = print
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Теперь нужно выбрать свой размер")
			msg.ReplyMarkup = OrderSize
			sendPhotoSize(update.CallbackQuery.Message.Chat.ID)
			standartSendMessage(msg)
		}

		if update.Message != nil {
			// ---------- Проверка команд ----------
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, value.Menu)
					standartSendMessage(msg)
				case "show":
					sendPhotoOctopus()
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Samurai Octopus 🐙")
					NewBot.Send(msg)

					sendPhotoShrimp()
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Samurai Shrimp 🦐")
					NewBot.Send(msg)

				case "price":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, value.Price)
					standartSendMessage(msg)

				case "size":
					sendPhotoSize(update.Message.Chat.ID)

				case "order":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы можете заказать что либо прямо здесь или обратиться к администратору")
					msg.ReplyMarkup = OrderSystem
					standartSendMessage(msg)
				case "help":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В скором времени с вами свяжется наш администратор ...")
					standartSendMessage(msg)

					NewBot.Send(tgbotapi.NewMessage(wlankasperID, "PROBLEM @"+update.Message.From.UserName))
					NewBot.Send(tgbotapi.NewMessage(anasharmsID, "PROBLEM @"+update.Message.From.UserName))

				case "admin":
					if update.Message.Chat.ID == wlankasperID || update.Message.Chat.ID == anasharmsID {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет Насть)")
						msg.ReplyMarkup = AdminSettings
						standartSendMessage(msg)
					}
				}
			} else if InputState != 0 {
				switch InputState {
				// --------- EMAIL---------
				case 1:
					OrderInfoMap[update.Message.Chat.ID].Email = update.Message.Text

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично, остался последний шаг!\n\nВведите адрес доставки и телефон в формате: \nГород,   Улица,   Дом,   Номер_телефона_для_связи")
					standartSendMessage(msg)
					InputState = 2

				// --------- ADDRESS---------
				case 2:
					OrderInfoMap[update.Message.Chat.ID].ContactInfo = update.Message.Text

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично! Проверьте свой заказ и выберите способ оплаты\n\nВаш заказ:\nПринт - "+OrderInfoMap[update.Message.Chat.ID].Print+"\nРазмер - "+OrderInfoMap[update.Message.Chat.ID].Size+"\nEmail - "+OrderInfoMap[update.Message.Chat.ID].Email+"\nДоставка - "+OrderInfoMap[update.Message.Chat.ID].ContactInfo)
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

			case "telegram":
				standartCallbackCheck()

				OrderInfoMap[update.CallbackQuery.Message.Chat.ID] = new(db.OrderInfo)
				OrderInfoMap[update.CallbackQuery.Message.Chat.ID].UserName = update.CallbackQuery.From.UserName

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выбирай какой принт ты хочешь")
				msg.ReplyMarkup = OrderPrint
				sendPhotoPrints(update.CallbackQuery.Message.Chat.ID)
				standartSendMessage(msg)

			case "admin":
				standartCallbackCheck()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В скором времени с вами свяжется наш администратор ...")
				standartSendMessage(msg)

				NewBot.Send(tgbotapi.NewMessage(wlankasperID, "PROBLEM @"+update.Message.From.UserName))
				NewBot.Send(tgbotapi.NewMessage(anasharmsID, "PROBLEM @"+update.Message.From.UserName))

			case "octopus":
				orderSetPrint("Samurai Octopus 🐙")

			case "shrimp":
				orderSetPrint("Samurai Shrimp 🦐")

			case "S":
				orderSetSize("S")

			case "M":
				orderSetSize("M")

			case "L":
				orderSetSize("L")

			case "card":
				standartCallbackCheck()

				OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Payment = "Перевод на карту"
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Тинькофф - 5536 9140 3655 4214 (Анастасия Владимировна)\n\nПосле перевода вам напишет наш Администратор чтобы подтвердить заказ и сообщит ближайщую дату доставки")
				standartSendMessage(msg)

				msg = tgbotapi.NewMessage(wlankasperID, "NEW ORDER\n\nName: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].UserName+"\nEmail: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Email+"\nAddress: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].ContactInfo+"\nPrint: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Print+"\nSize: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Size+"\nPayment: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Payment+"\nStatus: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Status)
				standartSendMessage(msg)

				msg = tgbotapi.NewMessage(anasharmsID, "NEW ORDER\n\nName: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].UserName+"\nEmail: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Email+"\nAddress: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].ContactInfo+"\nPrint: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Print+"\nSize: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Size+"\nPayment: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Payment+"\nStatus: "+OrderInfoMap[update.CallbackQuery.Message.Chat.ID].Status)
				standartSendMessage(msg)
			case "busd":
				// TODO

				// ------------------------------------ CALLBACK FOR ADMIN ------------------------------------
			case "add":
				standartCallbackCheck()
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Стандартная форма ввода:\nroot/НАЗВАНИЕ_КОЛЛЕКЦИИ/РАЗМЕР/ЦВЕТ/КОЛИЧЕСТВО\n\nПример: root/КРЕВЭД/S/Черный/4")
				standartSendMessage(msg)

				InputState = 1

			case "del":
				standartCallbackCheck()
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Стандартная форма ввода:\nroot/НАЗВАНИЕ_КОЛЛЕКЦИИ/РАЗМЕР/ЦВЕТ/КОЛИЧЕСТВО\n\nПример: root/КРЕВЭД/S/Черный/4")
				standartSendMessage(msg)

				InputState = 2

			case "all":
				standartCallbackCheck()
				// 	str := tree.TreePrint(true, "", "")
				// 	if str != "" {
				// 		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, str)
				// 		standartSendMessage(msg)
				// 	} else {
				// 		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Empty")
				// 		standartSendMessage(msg)
				// 	}
				// }
			}
		}
	}
}
