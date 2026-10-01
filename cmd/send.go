/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"mateusjdev/bot-kasinao/internal"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:     "enviar",
	Aliases: []string{"send"},
	Short:   "Envia o vídeo.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// carregar configuração
		cfg, err := internal.InitConfig()
		if err != nil {
			return err
		}
		tgBot, err := internal.NewTelegramConnection(
			cfg.TelegramConfig.Token,
			cfg.TelegramConfig.ChatID,
		)
		if err != nil {
			return err
		}
		// Escolher arquivo
		cntnt, err := internal.ReadRuntimeData(cfg.RuntimeFile)
		if err != nil {
			return err
		}
		arquivo, err := cntnt.EscolherArquivoAleatorio()
		if err != nil {
			return err
		}
		caption := cntnt.EscolherFraseAleatoria()
		// Enviar arquivo
		enviar, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			return err
		}

		if !enviar {
			err = tgBot.SendFile(arquivo, caption)
			if err != nil {
				return err
			}
		}

		// Salvar arquivo
		err = cntnt.Save(cfg.RuntimeFile)
		return err
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolP("dry-run", "d", false, "Para testes, não envia o vídeo.")
}
