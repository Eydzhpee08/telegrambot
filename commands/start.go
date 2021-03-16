package commands

import (
	"fmt"
	"github.com/tomassirio/bitcoinTelegram/utils"
)

func GetStart() (string, error) {
 p, err:= utils.GetApiCall()
	return fmt.Sprintln("/summary", p.Text), err
}