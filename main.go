package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/beevik/ntp"
)

func main() {
	os.Setenv("LC_ALL", "C")
	os.Setenv("LANG", "C")

	myApp := app.New()
	window := myApp.NewWindow("NTP Time Sync")

	statusLabel := widget.NewLabel("Status: Ready")
	ntpTimeLabel := widget.NewLabel("NTP time: -")
	localTimeLabel := widget.NewLabel("Local time: -")
	timeDiffLabel := widget.NewLabel("Difference: -")
	serverEntry := widget.NewEntry()
	serverEntry.SetText("time.apple.com")

	showWarning := func(message string) {
		dialog.NewInformation("Warning", message, window).Show()
	}

	syncTime := func() {
		statusLabel.SetText("Status: Synchronizing...")

		go func() {
			servers := []string{
				serverEntry.Text,
				"time.google.com",
				"ntp1.stratum2.ru",
				"pool.ntp.org",
			}

			var bestTime time.Time
			var err error

			for _, server := range servers {
				response, e := ntp.QueryWithOptions(server, ntp.QueryOptions{Timeout: 3 * time.Second})
				if e == nil {
					bestTime = response.Time
					err = nil
					break
				}
				err = e
			}

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "NTP Time Sync",
				Content: "Synchronization complete.",
			})

			if err != nil {
				statusLabel.SetText("Error: " + err.Error())
				showWarning("Failed to get NTP time")
				return
			}

			now := time.Now()
			diff := bestTime.Sub(now)
			statusLabel.SetText("Status: Synchronized")
			ntpTimeLabel.SetText(fmt.Sprintf("NTP time: %s", bestTime.Format("2006-01-02 15:04:05")))
			localTimeLabel.SetText(fmt.Sprintf("Local time: %s", now.Format("2006-01-02 15:04:05")))
			timeDiffLabel.SetText(fmt.Sprintf("Difference: %.3f seconds", diff.Seconds()))

			if runtime.GOOS == "darwin" {
				if diff.Abs() > 5*time.Second {
					showWarning(fmt.Sprintf(
						"Large time difference detected (%.1f sec).\n"+
							"To sync time on macOS:\n\n"+
							"1. Open Terminal\n"+
							"2. Run: sudo sntp -sS %s\n"+
							"3. Enter your password",
						diff.Seconds(),
						serverEntry.Text,
					))
				}
			}
		}()
	}

	syncBtn := widget.NewButton("Sync Time", syncTime)

	content := container.NewVBox(
		widget.NewLabel("NTP Time Synchronization"),
		widget.NewLabel("NTP Server:"),
		serverEntry,
		statusLabel,
		ntpTimeLabel,
		localTimeLabel,
		timeDiffLabel,
		container.NewHBox(
			syncBtn,
		),
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(500, 300))
	window.ShowAndRun()
}
