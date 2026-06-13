# Get Video Golang

Este é um projeto simples em Go para baixar vídeos do YouTube usando a ferramenta `yt-dlp`.

## Pré-requisitos

- Go 1.24.1 ou superior.
- Ferramenta [`yt-dlp`](https://github.com/yt-dlp/yt-dlp) instalada e disponível no PATH.
- [`ffmpeg`](https://ffmpeg.org/download.html) instalado e disponível no PATH (necessário para baixar vídeos em alta qualidade).

### Instalando o `yt-dlp`

Para instalar o `yt-dlp`, execute o seguinte comando no terminal:

```bash
pip install -U yt-dlp
```

Certifique-se de que o `yt-dlp` está acessível no PATH do sistema. Você pode verificar a instalação com:

```bash
yt-dlp --version
```

O YouTube muda o player com frequência. Mantenha o `yt-dlp` atualizado; o canal
nightly é recomendado pelo próprio projeto para usuários regulares:

```bash
yt-dlp --update-to nightly
```

### Instalando o `ffmpeg`

Para baixar e instalar o `ffmpeg`, siga as instruções no site oficial: [FFmpeg Downloads](https://ffmpeg.org/download.html).

Após a instalação, certifique-se de que o `ffmpeg` está acessível no PATH do sistema. Você pode verificar a instalação com:

```bash
ffmpeg -version
```

## Como Usar

#### Clone o repositório:

```bash
git clone https://github.com/sandronister/get_video_golang.git
cd get_video_golang
```

### Instale as dependências do projeto:
```bash
go mod tidy
```

### Execute o programa:
```bash
go run cmd/main.go
```

Quando solicitado, insira a URL do vídeo, o caminho da pasta onde deseja salvar o vídeo e se deseja baixar em alta resolução (sim/não).

O programa também solicita o navegador em que sua conta do YouTube está
autenticada, o perfil que deve ser usado e pede confirmação antes de passar os
cookies ao `yt-dlp`. O padrão é `chrome` com o perfil padrão. O `yt-dlp` não
expõe o e-mail da conta Google; a confirmação identifica a origem da sessão,
como `chrome:Profile 1`. Para usar um arquivo de cookies em vez do navegador:

```bash
YTDLP_COOKIES_FILE=/caminho/cookies.txt go run cmd/main.go
```

No macOS, o sistema pode solicitar a senha do usuário do Mac para liberar a
chave `Chrome Safe Storage`. Essa não é a senha do Google ou do YouTube. Caso os
cookies não possam ser descriptografados, use `YTDLP_COOKIES_FILE` com um
arquivo exportado no formato Netscape.

O YouTube rotaciona cookies de abas abertas. Para criar um arquivo que continue
válido:

1. Abra uma única janela anônima e entre no YouTube.
2. Na mesma aba, acesse `https://www.youtube.com/robots.txt`.
3. Exporte apenas os cookies de `youtube.com` no formato Netscape.
4. Feche a janela anônima e não volte a abrir essa sessão.

Versões atuais do `yt-dlp` também exigem um runtime JavaScript para resolver os
desafios do YouTube. No macOS com Homebrew:

```bash
brew update
brew upgrade yt-dlp
brew install deno
```

### Exemplo de Execução
```bash
Informe a url do youtube:  https://www.youtube.com/watch?v=example
Informe a pasta que deseja salvar, se não existir será criada:example
Deseja alta qualidade? (s/n): S
Video downloaded successfully: example/Testando baixar video.webm
```

### Funcionalidades
- Baixa vídeos do YouTube no formato MP4.
- Prioriza vídeo H.264 e áudio AAC para maior compatibilidade com players.
- Permite especificar o diretório de destino para salvar o vídeo.
- Suporte para download de vídeos em alta qualidade (requer `ffmpeg`).

### Observação
Para vídeos em alta qualidade, o `yt-dlp` utiliza o `ffmpeg` para combinar o vídeo e o áudio, pois eles são baixados separadamente.
