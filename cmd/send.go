/*
Copyright © 2026 Mateus Santana <mateusjuniordev@gmail.com>
*/
package cmd

import (
	"fmt"
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
		cfg, err := internal.ReadConfig()
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
		// Verificar arquivos
		foundFiles, err := internal.ScanFiles(cfg.MediaFolder)
		if err != nil {
			return err
		}
		// Escolher arquivo
		rntme, err := internal.ReadRuntimeData(cfg.RuntimeFile)
		if err != nil {
			return err
		}
		err = rntme.AddFiles(foundFiles)
		if err != nil {
			return err
		}
		arquivo, err := rntme.EscolherArquivoAleatorio()
		if err != nil {
			return err
		}
		caption := rntme.EscolherFraseAleatoria()
		// Enviar arquivo
		testes, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			return err
		}

		if testes {
			fmt.Printf("arquivo: %s\n", arquivo)
			fmt.Printf("mensagem: %s\n", caption)
		} else {
			err = tgBot.SendFile(arquivo, caption)
			if err != nil {
				return err
			}
		}

		// Salvar arquivo
		err = rntme.Save(cfg.RuntimeFile)
		return err
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().BoolP("dry-run", "d", false, "Para testes, não envia o vídeo.")
	// MAYBE: ignore-scan?
}
