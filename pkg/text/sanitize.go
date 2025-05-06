package text

import "regexp"

func Sanitize(folder string) string {
	invalidChars := regexp.MustCompile(`[<>:"/\\|?*]`)
	folder = invalidChars.ReplaceAllString(folder, "_")

	folder = regexp.MustCompile(`^[\s.]+|[\s.]+$`).ReplaceAllString(folder, "")

	reservedNames := map[string]bool{
		"CON": true, "PRN": true, "AUX": true, "NUL": true,
		"COM1": true, "COM2": true, "COM3": true, "COM4": true,
		"COM5": true, "COM6": true, "COM7": true, "COM8": true,
		"COM9": true, "LPT1": true, "LPT2": true, "LPT3": true,
		"LPT4": true, "LPT5": true, "LPT6": true, "LPT7": true,
		"LPT8": true, "LPT9": true,
	}
	if reservedNames[folder] {
		folder = folder + "_folder"
	}

	return folder
}
