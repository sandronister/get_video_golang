package youtube

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewCommandUsesBrowserCookies(t *testing.T) {
	t.Setenv(cookiesFileEnv, "")

	cmd := newCommand("chrome", "", "--get-title", "https://example.com/video")
	want := []string{
		"yt-dlp",
		"--cookies-from-browser", "chrome",
		"--get-title",
		"https://example.com/video",
	}

	if !reflect.DeepEqual(cmd.Args, want) {
		t.Fatalf("argumentos inesperados: got %v, want %v", cmd.Args, want)
	}
}

func TestNewCommandPrefersCookiesFile(t *testing.T) {
	t.Setenv(cookiesFileEnv, "/tmp/cookies.txt")

	cmd := newCommand("chrome", "", "--get-title", "https://example.com/video")
	want := []string{
		"yt-dlp",
		"--cookies", "/tmp/cookies.txt",
		"--get-title",
		"https://example.com/video",
	}

	if !reflect.DeepEqual(cmd.Args, want) {
		t.Fatalf("argumentos inesperados: got %v, want %v", cmd.Args, want)
	}
}

func TestNewCommandUsesSelectedBrowserProfile(t *testing.T) {
	t.Setenv(cookiesFileEnv, "")

	cmd := newCommand("chrome", "Profile 1", "--get-title", "https://example.com/video")
	want := []string{
		"yt-dlp",
		"--cookies-from-browser", "chrome:Profile 1",
		"--get-title",
		"https://example.com/video",
	}

	if !reflect.DeepEqual(cmd.Args, want) {
		t.Fatalf("argumentos inesperados: got %v, want %v", cmd.Args, want)
	}
}

func TestCommandErrorDetailDetectsRotatedCookies(t *testing.T) {
	output := "WARNING: The provided YouTube account cookies are no longer valid."

	got := commandErrorDetail(output)
	if !strings.Contains(got, "rotacionados") {
		t.Fatalf("erro inesperado: %q", got)
	}
}

func TestCommandErrorDetailDetectsJavaScriptFailure(t *testing.T) {
	output := "WARNING: nsig extraction failed\nWARNING: Only images are available for download"

	got := commandErrorDetail(output)
	if !strings.Contains(got, "runtime Deno") {
		t.Fatalf("erro inesperado: %q", got)
	}
}
