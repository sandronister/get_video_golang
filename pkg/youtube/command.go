package youtube

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const cookiesFileEnv = "YTDLP_COOKIES_FILE"

func newCommand(browser, profile string, args ...string) *exec.Cmd {
	commandArgs := make([]string, 0, len(args)+2)

	if cookiesFile := configuredCookiesFile(); cookiesFile != "" {
		commandArgs = append(commandArgs, "--cookies", cookiesFile)
	} else if browser = strings.TrimSpace(browser); browser != "" {
		if profile = strings.TrimSpace(profile); profile != "" {
			browser += ":" + profile
		}
		commandArgs = append(commandArgs, "--cookies-from-browser", browser)
	}

	commandArgs = append(commandArgs, args...)

	return exec.Command("yt-dlp", commandArgs...)
}

func configuredCookiesFile() string {
	return strings.TrimSpace(os.Getenv(cookiesFileEnv))
}

func runCommand(cmd *exec.Cmd) error {
	var stderr bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	err := cmd.Run()
	if err == nil {
		return nil
	}

	if detail := commandErrorDetail(stderr.String()); detail != "" {
		return fmt.Errorf("%s: %w", detail, err)
	}

	return err
}

func commandErrorDetail(output string) string {
	if strings.Contains(output, "could not be decrypted") ||
		strings.Contains(output, "cannot decrypt") ||
		strings.Contains(output, "find-generic-password failed") {
		return fmt.Sprintf(
			"não foi possível desbloquear os cookies no Chaves do macOS; a senha "+
				"solicitada é a do usuário do Mac, não a do YouTube. Como "+
				"alternativa, use %s=/caminho/cookies.txt",
			cookiesFileEnv,
		)
	}

	if strings.Contains(output, "account cookies are no longer valid") {
		return fmt.Sprintf(
			"os cookies do YouTube foram rotacionados e não são mais válidos; "+
				"exporte uma nova sessão anônima e atualize o arquivo indicado em %s",
			cookiesFileEnv,
		)
	}

	if strings.Contains(output, "nsig extraction failed") ||
		strings.Contains(output, "Only images are available for download") {
		return "o yt-dlp não conseguiu resolver o desafio JavaScript do YouTube; " +
			"atualize o yt-dlp e instale o runtime Deno"
	}

	return ""
}
