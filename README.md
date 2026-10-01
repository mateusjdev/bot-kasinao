# Kasinão

Bot que envia o video do kasinão todo sabadaço para um grupo do telegram. (O único kasino liberado no Brasil 🇧🇷).

**Funcionamento:** Toda vez que executado, escolhe um vídeo aleatório de uma pasta e envia para o chat do telegram configurado. É ideal utilizar outras aplicações como cronjob para agendar a tarefa.

O executável foi criado com a intensão de enviar o video do kasinão, mas pode enviar qualquer tipo de vídeo (ex: Sexta-feira, PlayTV).

## Configuração

O programa utiliza uma pasta para ler a configuração:

**Path:** `/usr/local/etc/kasinao.toml`

**Configuração de Exemplo:**

```toml
mediaFolder = "" # Pasta onde estão os arquivos de vídeo
runtimeFile = "tracker.json" # arquivo onde as informações temporárias ficam guardadas
legendas = [ # Legendas/Texto enviado junto com o vídeo (também escolhido aleatoriamente)
    "KASSINOOOO!",
    "AEEEEEE KASINÃO!",
]

# Chat do Telegram
[telegram]
token = "token" # Token do Bot
chatId = 0123456789 # ID do Chat
```

## 🛠️ Compilando

### 📦 Pré-requisitos

Certifique-se de ter o seguinte instalado:

- [Go](https://golang.org/doc/install) (versão ≥ 1.25)
- Para o processo de compilação, [Taskfile](https://taskfile.dev/#/installation) e [Git](https://git-scm.com/downloads) são altamente recomendados.

### 🏗️ Compilando

1. Clone o repositório para sua máquina local:

```shell
git clone https://github.com/mateusjdev/bot-kasinao
cd scruffy
```

1. Compile o projeto:

```shell
go build -o ./build/kasinao
# ou usando Taskfile (recomendado)
task build
```
