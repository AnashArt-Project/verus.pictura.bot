package main

// https://github.com/go-telegram-bot-api/telegram-bot-api
// https://docs.postmen.com/dhl.html

import (
	"fmt"
	"time"

	"verus.pictura/src/db"
	"verus.pictura/src/logger"
	"verus.pictura/src/value"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(value.BOT_TOKEN)
	OrderInfoMap   map[int64]*db.OrderInfo
	InputState     int = 0
)

func main() {
	logger.ForString(fmt.Sprintf("OK, %v, %v", time.Now().Unix(), time.Now().Weekday()))
	logger.ForError(BotErr)

	// ------ WEBHOOK ------
	// setWebhook(NewBot)
	// updates := NewBot.ListenForWebhook("/" + NewBot.Token)
	// go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	// ------ STANDARD ------
	NewBot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := NewBot.GetUpdatesChan(u)

	OrderInfoMap = make(map[int64]*db.OrderInfo)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				checkCommand(update.Message)
			} else if InputState != 0 {
				checkInputState(update.Message)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я даже не знаю как на это ответить(")
				standardSendMessage(msg)
			}
		} else if update.CallbackQuery != nil {
			checkCallback(update.CallbackQuery)
		}
	}
}

func setWebhook(bot *tgbotapi.BotAPI) {
	webHookInfo, err := tgbotapi.NewWebhookWithCert(fmt.Sprintf("https://%s:%s/%s", value.BOT_ADDRESS, value.BOT_PORT, value.BOT_TOKEN), tgbotapi.FilePath(value.CERT_PATH))
	logger.ForError(err)
	_, err = bot.Request(webHookInfo)
	logger.ForError(err)
	info, err := bot.GetWebhookInfo()
	logger.ForError(err)
	if info.LastErrorDate != 0 {
		// logger.ForError(info.LastErrorMessage)
	}
}

func checkCommand(message *tgbotapi.Message) {
	if message.IsCommand() {
		switch message.Command() {

		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, value.Menu)
			standardSendMessage(msg)

		case "show":
			sendPhotoOctopus(message.Chat.ID)
			msg := tgbotapi.NewMessage(message.Chat.ID, "Samurai Octopus 🐙")
			NewBot.Send(msg)

			sendPhotoShrimp(message.Chat.ID)
			msg = tgbotapi.NewMessage(message.Chat.ID, "Samurai Shrimp 🦐")
			NewBot.Send(msg)

		case "price":
			msg := tgbotapi.NewMessage(message.Chat.ID, value.Price)
			standardSendMessage(msg)

		case "size":
			sendPhotoSize(message.Chat.ID)

		case "order":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Вы можете заказать что либо прямо здесь или обратиться к администратору")
			msg.ReplyMarkup = value.OrderSystem
			standardSendMessage(msg)

		case "help":
			msg := tgbotapi.NewMessage(message.Chat.ID, "В скором времени с вами свяжется наш администратор ...")
			standardSendMessage(msg)

			NewBot.Send(tgbotapi.NewMessage(value.WLANKASPER_ID, "PROBLEM @"+message.From.UserName))
			NewBot.Send(tgbotapi.NewMessage(value.ANASHARMS_ID, "PROBLEM @"+message.From.UserName))

		case "adminMod":
			if message.Chat.ID == value.WLANKASPER_ID || message.Chat.ID == value.ANASHARMS_ID {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Привет Насть)")
				msg.ReplyMarkup = value.AdminSettings
				standardSendMessage(msg)
			}
		}
	}
}

func checkInputState(message *tgbotapi.Message) {
	switch InputState {
	case 1:
		OrderInfoMap[message.Chat.ID].Email = message.Text

		msg := tgbotapi.NewMessage(message.Chat.ID, "Контактный номер телефона:")
		standardSendMessage(msg)
		InputState = 2

	case 2:
		OrderInfoMap[message.Chat.ID].Phone = message.Text

		msg := tgbotapi.NewMessage(message.Chat.ID, "Город: ")
		standardSendMessage(msg)
		InputState = 3

	case 3:
		OrderInfoMap[message.Chat.ID].City = message.Text

		msg := tgbotapi.NewMessage(message.Chat.ID, "Улица: ")
		standardSendMessage(msg)
		InputState = 4

	case 4:
		OrderInfoMap[message.Chat.ID].Street = message.Text

		msg := tgbotapi.NewMessage(message.Chat.ID, "Номер дома: ")
		standardSendMessage(msg)
		InputState = 5

	case 5:
		OrderInfoMap[message.Chat.ID].House = message.Text

		msg := tgbotapi.NewMessage(message.Chat.ID, "Отлично! Проверьте свой заказ и выберите способ оплаты\n\n"+db.ToStringOrderInfo(OrderInfoMap[message.Chat.ID]))
		msg.ReplyMarkup = value.OrderPayment
		standardSendMessage(msg)

		InputState = 0
	}
}

func checkCallback(toCallback *tgbotapi.CallbackQuery) {
	switch toCallback.Data {

	case "telegram":
		standardCallbackCheck(*toCallback)

		OrderInfoMap[toCallback.Message.Chat.ID] = new(db.OrderInfo)
		OrderInfoMap[toCallback.Message.Chat.ID].UserName = toCallback.From.UserName

		msg := tgbotapi.NewMessage(toCallback.Message.Chat.ID, "Выбирай какой принт ты хочешь")
		msg.ReplyMarkup = value.OrderPrint
		sendPhotoPrints(toCallback.Message.Chat.ID)
		standardSendMessage(msg)

	case "admin":
		standardCallbackCheck(*toCallback)
		msg := tgbotapi.NewMessage(toCallback.Message.Chat.ID, "В скором времени с вами свяжется наш администратор ...")
		standardSendMessage(msg)

		NewBot.Send(tgbotapi.NewMessage(value.WLANKASPER_ID, "PROBLEM @"+toCallback.From.UserName))
		NewBot.Send(tgbotapi.NewMessage(value.ANASHARMS_ID, "PROBLEM @"+toCallback.From.UserName))

	case "octopus":
		orderSetPrint("Samurai Octopus 🐙", *toCallback)

	case "shrimp":
		orderSetPrint("Samurai Shrimp 🦐", *toCallback)

	case "S":
		orderSetSize("S", *toCallback)

	case "M":
		orderSetSize("M", *toCallback)

	case "L":
		orderSetSize("L", *toCallback)

	case "card":
		standardCallbackCheck(*toCallback)

		OrderInfoMap[toCallback.Message.Chat.ID].Payment = "Перевод на карту"
		OrderInfoMap[toCallback.Message.Chat.ID].Status = "Ожидание оплаты + Нужно уточнить стоимость доставки!"

		msg := tgbotapi.NewMessage(toCallback.Message.Chat.ID, "К оплате: 3600 + доставка\n\nТинькофф - 5536 9140 3655 4214 (Анастасия Владимировна)")
		standardSendMessage(msg)
		msg = tgbotapi.NewMessage(toCallback.Message.Chat.ID, "В связи с техническими проблемами сайта Почты России мы не можем автоматически расчитать стоимость доставки. В ближайшее время с вами свяжется администратор и уточнить стоимость ;)")
		standardSendMessage(msg)

		msg = tgbotapi.NewMessage(value.WLANKASPER_ID, "NEW ORDER\n\n"+db.ToStringAllOrderInfo(OrderInfoMap[toCallback.Message.Chat.ID]))
		standardSendMessage(msg)

		msg = tgbotapi.NewMessage(value.ANASHARMS_ID, "NEW ORDER\n\n"+db.ToStringAllOrderInfo(OrderInfoMap[toCallback.Message.Chat.ID]))
		standardSendMessage(msg)
	case "busd":
		// TODO

		// ------------------------------------ CALLBACK FOR ADMIN ------------------------------------
	case "add":
		standardCallbackCheck(*toCallback)
		msg := tgbotapi.NewMessage(toCallback.Message.Chat.ID, "Стандартная форма ввода:\nroot/НАЗВАНИЕ_КОЛЛЕКЦИИ/РАЗМЕР/ЦВЕТ/КОЛИЧЕСТВО\n\nПример: root/КРЕВЭД/S/Черный/4")
		standardSendMessage(msg)

		InputState = 1

	case "del":
		standardCallbackCheck(*toCallback)
		msg := tgbotapi.NewMessage(toCallback.Message.Chat.ID, "Стандартная форма ввода:\nroot/НАЗВАНИЕ_КОЛЛЕКЦИИ/РАЗМЕР/ЦВЕТ/КОЛИЧЕСТВО\n\nПример: root/КРЕВЭД/S/Черный/4")
		standardSendMessage(msg)

		InputState = 2

	case "all":
		standardCallbackCheck(*toCallback)
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

func standardSendMessage(msg tgbotapi.MessageConfig) {
	if _, err := NewBot.Send(msg); err != nil {
		logger.ForError(err)
	}
}

func standardCallbackCheck(toCallback tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(toCallback.ID, toCallback.Data)
	if _, err := NewBot.Request(callback); err != nil {
		logger.ForError(err)
	}
}

func orderSetSize(size string, toCallback tgbotapi.CallbackQuery) {
	standardCallbackCheck(toCallback)
	OrderInfoMap[toCallback.Message.Chat.ID].Size = size
	msg := tgbotapi.NewMessage(toCallback.Message.Chat.ID, "Отлично! Сейчас будем заполнять информацию для доставки ;)")
	standardSendMessage(msg)

	msg = tgbotapi.NewMessage(toCallback.Message.Chat.ID, "Контактный email:")
	standardSendMessage(msg)

	InputState = 1
}

func orderSetPrint(print string, toCallback tgbotapi.CallbackQuery) {
	standardCallbackCheck(toCallback)
	OrderInfoMap[toCallback.Message.Chat.ID].Print = print
	msg := tgbotapi.NewMessage(toCallback.Message.Chat.ID, "Теперь нужно выбрать свой размер")
	msg.ReplyMarkup = value.OrderSize
	sendPhotoSize(toCallback.Message.Chat.ID)
	standardSendMessage(msg)
}

func sendPhotoShrimp(sendTo int64) {
	cfg := tgbotapi.NewMediaGroup(sendTo, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_shrimp)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_front)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_back)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_zoom)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_many)),
	})
	messages, err := NewBot.SendMediaGroup(cfg)

	logger.ForError(err)
	if messages == nil {
		// logger.ForString("No received messages")
	}
	if len(messages) != len(cfg.Media) {
		// logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
	}
}

func sendPhotoOctopus(sendTo int64) {
	cfg := tgbotapi.NewMediaGroup(sendTo, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_octopus)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_front)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_back)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_zoom)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_many)),
	})
	messages, err := NewBot.SendMediaGroup(cfg)

	logger.ForError(err)
	if messages == nil {
		// logger.ForString("No received messages")
	}
	if len(messages) != len(cfg.Media) {
		// logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
	}
}

func sendPhotoPrints(sendTo int64) {
	cfg := tgbotapi.NewMediaGroup(sendTo, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_octopus)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_1_many)),
	})
	messages, err := NewBot.SendMediaGroup(cfg)

	logger.ForError(err)
	if messages == nil {
		// logger.ForString("No received messages")
	}
	if len(messages) != len(cfg.Media) {
		// logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
	}

	cfg = tgbotapi.NewMediaGroup(sendTo, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Full_samurai_shrimp)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Product_2_many)),
	})
	messages, err = NewBot.SendMediaGroup(cfg)

	logger.ForError(err)
	if messages == nil {
		// logger.ForString("No received messages")
	}
	if len(messages) != len(cfg.Media) {
		// logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
	}
}

func sendPhotoSize(sendTo int64) {
	cfg := tgbotapi.NewMediaGroup(sendTo, []interface{}{
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Size_s)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Size_m)),
		tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(value.Size_l)),
	})
	messages, err := NewBot.SendMediaGroup(cfg)

	logger.ForError(err)
	if messages == nil {
		// logger.ForString("No received messages")
	}
	if len(messages) != len(cfg.Media) {
		// logger.ForString(fmt.Sprintf("Different number of messages: %v", len(messages)))
	}
}
