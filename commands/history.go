package commands

import (
	"fmt"
	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetHistory() (string, error) {
 p, err:= utils.GetApiCall()
	return fmt.Sprintln("https://www.kaspersky.ru/resource-center/definitions/what-is-bitcoin", p.Text), err
}