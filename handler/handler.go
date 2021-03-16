package handler

import (
	"github.com/tomassirio/bitcoinTelegram/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := commands.GetPrice()
		b.Send(m.Chat, "Текущая цена BTC: USA$ "+res)
	}

	commandMap["/historic"] = func(m *tb.Message) {
		res, g, _ := commands.GetHistoric()
		b.Send(m.Chat, "Цена BTC по сравнению со вчерашним днемсоставляет:"+res)
		b.Send(m.Chat, g)
	}

	commandMap["/summary"] = func(m *tb.Message) {
		p, h, _ := commands.GetSummary()
		b.Send(m.Chat, "Текущая цена BTC: USA$ "+p+"\nЦена BTC по сравнению со вчерашним днем​​ составляет: "+h)
	}
	commandMap["/start"] = func(m *tb.Message) {
		p, _:= commands.GetStart()
		b.Send(m.Chat,"Что такое Bitcoin? /history")
		b.Send(m.Chat,"Сегоднящная цена Bitcoin-a /price")
		b.Send(m.Chat,"Курс Bitcoin по сравнению со вчеращнего дня /historic " + p)
	}
	commandMap["/history"] = func(m *tb.Message) {
		p, _:= commands.GetHistory()
		b.Send(m.Chat,"Что такое Bitcoin? "+p)
	}

	return commandMap
}
