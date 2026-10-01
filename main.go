package main

import (
	"log"
	"mateusjdev/bot-kasinao/internal"
)

func main() {
	// carregar configuração
	cfg, err := internal.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	tgBot, err := internal.NewTelegramConnection(
		cfg.TelegramConfig.Token,
		cfg.TelegramConfig.ChatID,
	)
	if err != nil {
		log.Fatal(err)
	}
	// Escolher arquivo
	cntnt, err := internal.ReadRuntimeData(cfg.RuntimeFile)
	if err != nil {
		log.Fatal(err)
	}
	arquivo, err := cntnt.EscolherArquivoAleatorio()
	if err != nil {
		log.Fatal(err)
	}
	caption := cntnt.EscolherFraseAleatoria()
	// Enviar arquivo
	err = tgBot.SendFile(arquivo, caption)
	if err != nil {
		log.Fatal(err)
	}
	// Salvar arquivo
	err = cntnt.Save(cfg.RuntimeFile)
	if err != nil {
		log.Fatal(err)
	}
}
