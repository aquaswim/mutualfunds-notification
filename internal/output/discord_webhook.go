package output

import (
	"aya-money-go/internal/contracts"
	"aya-money-go/internal/utils"
)

type output struct {
	webhookUrl string
	buff       string
}

type dcExecutePayload struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

func (o *output) Print(text string) error {
	o.buff += text
	return nil
}

func (o *output) Flush() error {
	body := &dcExecutePayload{
		Content:  o.buff,
		Username: "aya",
	}
	err := utils.SendJsonRequest("POST", o.webhookUrl, body, nil)
	if err != nil {
		return err
	}

	o.buff = ""
	return nil
}

func NewOutputDiscordWebhook(webhookUrl string) contracts.Output {
	return &output{webhookUrl: webhookUrl, buff: ""}
}
