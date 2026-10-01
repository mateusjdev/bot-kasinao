package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type MediaFile struct {
	Path        string `mapstructure:"path" json:"path"`
	SentCounter int    `mapstructure:"counter" json:"counter"`
}

type Caption struct {
	Text        string `mapstructure:"text" json:"text"`
	SentCounter int    `mapstructure:"counter" json:"counter"`
}

type RuntimeData struct {
	Files   []MediaFile `mapstructure:"files" json:"files"`
	Message []Caption   `mapstructure:"messages" json:"messages"`
}

func ReadRuntimeData(filepath string) (*RuntimeData, error) {
	viperData := viper.New()
	if path.Ext(filepath) != ".json" {
		return nil, errors.New("Invalid ContentTracker file type")
	}
	viperData.SetConfigFile(filepath)
	viperData.SetConfigType("json")
	if runtime.GOOS == "windows" {
		// TODO: Windows
		viperData.AddConfigPath("./")
	} else {
		viperData.AddConfigPath("/usr/local/etc/")
	}

	if err := viperData.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error content file: %w", err)
	}

	var data RuntimeData
	if err := viperData.Unmarshal(&data); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct: %v", err)
	}

	return &data, nil
}
func (data *RuntimeData) EscolherArquivoAleatorio() (string, error) {

	tamanho := len(data.Files)
	if tamanho <= 0 {
		return "", errors.New("Não é possível escolher com 0 opções.")
	}
	v := rand.IntN(tamanho)
	file := data.Files[v]
	file.SentCounter++
	return file.Path, nil
}

func (data *RuntimeData) EscolherFraseAleatoria() string {
	tamanho := len(data.Message)
	if tamanho <= 0 {
		return ""
	}
	v := rand.IntN(tamanho)
	text := data.Message[v]
	text.SentCounter++
	return ""
}

func (data *RuntimeData) Save(filepath string) error {
	// Marshal into JSON (or use yaml.Marshal)
	jdata, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Save to file
	err = os.WriteFile(filepath, jdata, 0644)
	return err
}
