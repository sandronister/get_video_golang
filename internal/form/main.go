package form

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sandronister/get_video_golang/pkg/text"
	"github.com/sandronister/get_video_golang/pkg/youtube/types"
)

func Start() *types.Input {
	reader := bufio.NewReader(os.Stdin)
	var kind string = "v"

	fmt.Println("Informe a url do youtube: ")
	urlPath := readLine(reader)

	fmt.Println("Informe a pasta que deseja salvar, se não existir será criada: ")
	folderPath := readLine(reader)

	fmt.Println("Deseja somente o áudio? (s/n): ")
	audio := readLine(reader)

	fmt.Println("Deseja alta qualidade? (s/n): ")
	quality := readLine(reader)

	fmt.Println("Navegador autenticado no YouTube (chrome/firefox/edge/safari, padrão: chrome): ")
	browser := readLine(reader)
	if browser == "" {
		browser = "chrome"
	}

	fmt.Println("Perfil do navegador (ex.: Default ou Profile 1; vazio usa o perfil padrão): ")
	profile := readLine(reader)

	cookieSource := browser
	if profile != "" {
		cookieSource += ":" + profile
	}

	fmt.Printf("A sessão do YouTube será lida de %s. Confirma? (s/n):\n", cookieSource)
	if text.Sanitize(readLine(reader)) != "s" {
		fmt.Println("Operação cancelada.")
		return nil
	}

	if text.Sanitize(audio) == "s" {
		kind = "a"
	}

	return &types.Input{
		Url:     urlPath,
		Path:    text.Sanitize(folderPath),
		Kind:    kind,
		Quality: text.Sanitize(quality),
		Browser: browser,
		Profile: profile,
	}
}

func readLine(reader *bufio.Reader) string {
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}
