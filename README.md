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

### Exemplo de Execução
```bash
Informe a url do youtube:  https://www.youtube.com/watch?v=example
Informe a pasta que deseja salvar, se não existir será criada:example
Deseja alta qualidade? (s/n): S
Video downloaded successfully: example/Testando baixar video.webm
```

### Funcionalidades
- Baixa vídeos do YouTube no formato MP4.
- Permite especificar o diretório de destino para salvar o vídeo.
- Suporte para download de vídeos em alta qualidade (requer `ffmpeg`).

### Observação
Para vídeos em alta qualidade, o `yt-dlp` utiliza o `ffmpeg` para combinar o vídeo e o áudio, pois eles são baixados separadamente.