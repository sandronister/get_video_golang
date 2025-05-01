# Get Video Golang

Este é um projeto simples em Go para baixar vídeos do YouTube usando a ferramenta `yt-dlp`.

## Pré-requisitos

- Go 1.24.1 ou superior.
- Ferramenta [`yt-dlp`](https://github.com/yt-dlp/yt-dlp) instalada e disponível no PATH.

### Instalando o `yt-dlp`

Para instalar o `yt-dlp`, execute o seguinte comando no terminal:

```bash
pip install -U yt-dlp
```

Certifique-se de que o yt-dlp está acessível no PATH do sistema. Você pode verificar a instalação com:

```bash
yt-dlp --version
```

## Como Usar
#### Clone o repositório:

```git
git clone https://github.com/sandronister/get_video_golang.git
cd get_video_golang
````

### Instale as dependências do projeto:
```bash
go mod tidy
````

### Execute o programa:
```basg
go run cmd/main.go
```

Insira a URL do vídeo e o caminho da pasta onde deseja salvar o vídeo quando solicitado.

### Exemplo de Execução
```bash
Enter the video URL: https://www.youtube.com/watch?v=example
Enter the folder path to save the video: ./music
Video downloaded successfully: Video downloaded to [video.mp4](http://_vscodecontentref_/0)
```

### Funcionalidades
Baixa vídeos do YouTube no formato MP4.
Permite especificar o diretório de destino para salvar o vídeo.