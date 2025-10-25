package tts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Language represents supported languages
type Language string

const (
	Tagalog  Language = "tl"
	Japanese Language = "ja"
	English  Language = "en"
)

// Voice configuration for each language
var voices = map[Language]string{
	Tagalog:  "fil-PH-AngeloNeural",  // Male
	Japanese: "ja-JP-NanamiNeural",   // Female
	English:  "en-US-AriaNeural",     // Female
}

// Speaker handles text-to-speech operations
type Speaker struct {
	language Language
	voice    string
	tempDir  string
}

// NewSpeaker creates a new Speaker for the specified language
func NewSpeaker(lang Language) (*Speaker, error) {
	voice, ok := voices[lang]
	if !ok {
		return nil, fmt.Errorf("unsupported language: %s", lang)
	}

	tempDir := os.TempDir()

	return &Speaker{
		language: lang,
		voice:    voice,
		tempDir:  tempDir,
	}, nil
}

// Speak converts text to speech and plays it
func (s *Speaker) Speak(text string) error {
	if text == "" {
		return fmt.Errorf("text is empty")
	}

	// Create temporary audio file
	audioFile := filepath.Join(s.tempDir, fmt.Sprintf("tts_%s.mp3", s.language))

	// Generate speech using edge-tts command
	// Use shell to ensure PATH is set correctly
	cmdStr := fmt.Sprintf("edge-tts --voice %s --text %q --write-media %s",
		s.voice, text, audioFile)

	cmd := exec.Command("/bin/bash", "-c", cmdStr)

	// Set environment to include common paths
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = "/Users/" + os.Getenv("USER")
	}

	paths := []string{
		homeDir + "/.pyenv/shims",
		homeDir + "/.local/bin",
		"/usr/local/bin",
		"/usr/bin",
		"/bin",
	}
	pathEnv := "PATH=" + strings.Join(paths, ":")
	cmd.Env = append(os.Environ(), pathEnv)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to generate speech: %w\nOutput: %s", err, string(output))
	}

	// Play audio using afplay (macOS built-in)
	playCmd := exec.Command("afplay", audioFile)
	if err := playCmd.Run(); err != nil {
		// Clean up audio file
		os.Remove(audioFile)
		return fmt.Errorf("failed to play audio: %w", err)
	}

	// Clean up audio file after playing
	os.Remove(audioFile)

	return nil
}

// GetLanguageName returns the human-readable name of the language
func (s *Speaker) GetLanguageName() string {
	names := map[Language]string{
		Tagalog:  "Tagalog",
		Japanese: "日本語",
		English:  "English",
	}
	return names[s.language]
}
