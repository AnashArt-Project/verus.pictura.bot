package value

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// -------------- STANDARD MSG --------------
const (
	Menu  = "Привет! Я бот Verus Pictura и вот что я умею:\n\n  /start - Начальное меню 🧾\n  /show - Наши коллекции ✨\n  /price - Прайс-Лист 💸 \n  /size - Размеры 👀\n  /order - Оформить заказ 📦\n  /help - Вызвать Администратора ⁉️\n\nНаш официальный сайт: https://AnashArt.ru\nНаш официальный Instagram: https://www.instagram.com/p/CbBBrGRMt8w/?utm_medium=copy_link"
	Price = "Price List:\n\n  Футболка 'Samurai Octopus 🐙' - 3600₽\n  Футболка 'Samurai Shrimp 🦐' - 3600₽"
)

// -------------- PHOTO PATH --------------
const (
	Full_samurai_octopus = "img/full_samurai_octopus.jpg"
	Product_1_front      = "img/product_1_front.jpg"
	Product_1_back       = "img/product_1_back.jpg"
	Product_1_zoom       = "img/product_1_zoom.jpg"
	Product_1_many       = "img/product_1_many.jpg"

	Full_samurai_shrimp = "img/full_samurai_shrimp.png"
	Product_2_front     = "img/product_2_front.png"
	Product_2_back      = "img/product_2_back.png"
	Product_2_zoom      = "img/product_2_zoom.png"
	Product_2_many      = "img/product_2_many.png"

	Size_s = "img/size_s.png"
	Size_m = "img/size_m.png"
	Size_l = "img/size_l.png"
)

// -------------- BOT INFO --------------
const (
	CERT_PATH    = "/root/WorkSpace_AnashArt_Bot/certs/cert.pem"
	KEY_PATH     = "/root/WorkSpace_AnashArt_Bot/certs/cert.key"
	BOT_TOKEN    = "5267887349:AAEr95a2kk8B78h5CO2yv8E-IN9W2FxERi4"
	BOT_ADDRESS  = "65.108.154.134"
	BOT_PORT     = "8443"
	TELEGRAM_URL = "https://t.me/VerusPicturaBot"

	WLANKASPER_ID = 853634511
	ANASHARMS_ID  = 726736906
)

var (
	OrderSystem = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Бот", "telegram"),
			tgbotapi.NewInlineKeyboardButtonData("Администратор", "admin"),
		),
	)

	OrderPrint = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Samurai Octopus 🐙", "octopus"),
			tgbotapi.NewInlineKeyboardButtonData("Samurai Shrimp 🦐", "shrimp"),
		),
	)

	OrderSize = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("S", "S"),
			tgbotapi.NewInlineKeyboardButtonData("M", "M"),
			tgbotapi.NewInlineKeyboardButtonData("L", "L"),
		),
	)

	OrderPayment = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			// tgbotapi.NewInlineKeyboardButtonData("Оплата BUSD", "busd"),
			tgbotapi.NewInlineKeyboardButtonData("Перевод на карту", "card"),
		),
	)
	AdminSettings = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Что-то добавить", "add"),
			tgbotapi.NewInlineKeyboardButtonData("Что-то удалить", "del"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Показать все", "all"),
		),
	)
)
