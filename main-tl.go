package main

import (
	"fmt"
	"log"
	"talkertalker/tts"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	hook "github.com/robotn/gohook"
)

type AppTL struct {
	speaker        *tts.Speaker
	fyneApp        fyne.App
	window         fyne.Window
	textArea       *widget.Entry
	status         *widget.Label
	autoReadCheck  *widget.Check
	speaking       bool
	autoRead       bool
	debounceTimer  *time.Timer
	lastText       string
	activateChan   chan bool
}

func main() {
	// Create speaker for Tagalog
	speaker, err := tts.NewSpeaker(tts.Tagalog)
	if err != nil {
		log.Fatal(err)
	}

	// Create Fyne app with unique ID
	a := app.NewWithID("com.talkertalker.tagalog")
	w := a.NewWindow("TalkerTalker - Tagalog")
	w.Resize(fyne.NewSize(500, 400))

	// Create app instance
	ttsApp := &AppTL{
		speaker:      speaker,
		fyneApp:      a,
		window:       w,
		speaking:     false,
		activateChan: make(chan bool, 1),
	}

	// Setup UI
	ttsApp.setupUI()

	// Setup global hotkey (Ctrl+Shift+H)
	go ttsApp.setupHotkey()

	// Monitor activation channel in main thread
	go func() {
		for range ttsApp.activateChan {
			w.Show()
			w.RequestFocus()
			// Focus on text area
			w.Canvas().Focus(ttsApp.textArea)
		}
	}()

	w.ShowAndRun()
}

func (a *AppTL) setupUI() {
	// Text area
	a.textArea = widget.NewMultiLineEntry()
	a.textArea.SetPlaceHolder("Enter or paste text to read aloud...")
	a.textArea.SetMinRowsVisible(10)

	// Auto-read on text change
	a.textArea.OnChanged = func(text string) {
		if a.autoRead && text != "" && text != a.lastText {
			a.lastText = text
			// Debounce: wait 500ms before reading
			if a.debounceTimer != nil {
				a.debounceTimer.Stop()
			}
			a.debounceTimer = time.AfterFunc(500*time.Millisecond, func() {
				a.speakText(text)
			})
		}
	}

	// Status label
	a.status = widget.NewLabel("Ready")

	// Auto Read checkbox
	a.autoReadCheck = widget.NewCheck("Auto Read", func(checked bool) {
		a.autoRead = checked
		if checked {
			a.status.SetText("Auto Read: ON")
		} else {
			a.status.SetText("Auto Read: OFF")
		}
	})
	a.autoReadCheck.SetChecked(true)

	// Buttons
	speakBtn := widget.NewButton("Speak", func() {
		a.speakText(a.textArea.Text)
	})

	pasteBtn := widget.NewButton("Paste & Speak", func() {
		text, err := clipboard.ReadAll()
		if err != nil {
			a.status.SetText("Failed to read clipboard")
			return
		}
		a.textArea.SetText(text)
		if !a.autoRead {
			a.speakText(text)
		}
	})

	clearBtn := widget.NewButton("Clear", func() {
		a.textArea.SetText("")
		a.lastText = ""
		a.status.SetText("Ready")
	})

	// Info label
	infoText := fmt.Sprintf("Language: %s | Enable 'Auto Read' to automatically speak pasted text",
		a.speaker.GetLanguageName())
	info := widget.NewLabel(infoText)
	info.Wrapping = fyne.TextWrapWord

	// Layout
	buttons := container.NewHBox(speakBtn, pasteBtn, clearBtn, a.autoReadCheck)
	content := container.NewBorder(
		nil,
		container.NewVBox(buttons, a.status, info),
		nil,
		nil,
		a.textArea,
	)

	a.window.SetContent(content)
}

func (a *AppTL) setupHotkey() {
	// Register Ctrl+Shift+H to activate window
	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "h"}, func(e hook.Event) {
		// Send activation signal to main thread
		select {
		case a.activateChan <- true:
		default:
			// Channel full, skip
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func (a *AppTL) speakText(text string) {
	if text == "" {
		a.status.SetText("No text to speak")
		return
	}

	if a.speaking {
		a.status.SetText("Already speaking...")
		return
	}

	a.speaking = true
	a.status.SetText(fmt.Sprintf("Speaking in %s...", a.speaker.GetLanguageName()))

	go func() {
		err := a.speaker.Speak(text)
		// Note: Direct UI updates from goroutine will show warnings but will work
		if err != nil {
			a.status.SetText(fmt.Sprintf("Error: %v", err))
		} else {
			a.status.SetText("Done")
			// Clear text after speaking
			a.textArea.SetText("")
			a.lastText = ""
		}
		a.speaking = false
	}()
}
