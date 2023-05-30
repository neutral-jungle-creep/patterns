package structural

import (
	"fmt"
)

//

type Alert struct {
	notifier   Notifier
	message    string
	importance string
}

func newAlert(nType, msg, imp string) *Alert {
	fmt.Println("Starting create alert")
	alert := &Alert{
		notifier:   newNotifier(nType),
		message:    msg,
		importance: imp,
	}
	fmt.Println("Alert created")
	return alert
}

func (a *Alert) sendAlert() {
	a.notifier.send(a.message, a.importance)
}

type Notifier interface {
	send(msg, imp string)
}

func newNotifier(nType string) Notifier {
	if nType == "email" {
		return newEmailNotifier()
	}

	if nType == "telegram" {
		return newTelegramNotifier()
	}

	return nil
}

type EmailNotifier struct {
	address string
}

func newEmailNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (n *EmailNotifier) send(msg, imp string) {
	fmt.Printf("Notifier send alert text `%s` importance `%s` to email\n", msg, imp)
}

type TelegramNotifier struct {
	address string
}

func newTelegramNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (n *TelegramNotifier) send(msg, imp string) {
	fmt.Printf("Notifier send alert text `%s` importance `%s` to telegram\n", msg, imp)
}

func RunFacade() {
	alertEmail := newAlert("email", "Hello", "low")
	alertEmail.sendAlert()

	alertTelegram := newAlert("telegram", "Hello telegram", "medium")
	alertTelegram.sendAlert()
}
