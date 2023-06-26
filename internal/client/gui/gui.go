package gui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/api/events"
	"github.com/vasiliyantufev/gophkeeper/internal/client/component"
	"github.com/vasiliyantufev/gophkeeper/internal/client/component/form"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/table"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/errors"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/labels"
)

func InitGUI(log *logrus.Logger, application fyne.App, client *events.Event) {
	window := application.NewWindow("GophKeeper")
	window.Resize(fyne.NewSize(250, 80))
	var dataTblLoginPassword = [][]string{{"NAME", "DESCRIPTION", "LOGIN", "PASSWORD", "CREATED AT", "UPDATED AT"}}
	var dataTblText = [][]string{{"NAME", "DESCRIPTION", "DATA", "CREATED AT", "UPDATED AT"}}
	var dataTblCard = [][]string{{"NAME", "DESCRIPTION", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED AT", "UPDATED AT"}}
	var radioOptions = []string{"Login", "Registration"}
	var accessToken = model.Token{}
	var password string
	var exist bool
	var valid bool
	var layout string
	var err error
	layout = "01/02/2006 15:04:05"
	//---------------------------------------------------------------------- containers
	var containerRadio *fyne.Container
	var containerFormLogin *fyne.Container
	var containerFormRegistration *fyne.Container
	var containerFormLoginPassword *fyne.Container
	var containerFormText *fyne.Container
	var containerFormCard *fyne.Container
	//---------------------------------------------------------------------- buttons
	var buttonAuth *widget.Button
	var buttonTop *widget.Button
	var buttonLoginPassword *widget.Button
	var buttonText *widget.Button
	var buttonCard *widget.Button
	var buttonLoginPasswordAdd *widget.Button
	var buttonTextAdd *widget.Button
	var buttonCardAdd *widget.Button
	//---------------------------------------------------------------------- tabs
	var containerTabs *container.AppTabs
	var tblLoginPassword *widget.Table
	var tblText *widget.Table
	var tblCard *widget.Table
	var tabLoginPassword *container.TabItem
	var tabText *container.TabItem
	var tabCard *container.TabItem
	//---------------------------------------------------------------------- entries init
	separator := widget.NewSeparator()
	usernameLoginEntry := widget.NewEntry()
	passwordLoginEntry := widget.NewPasswordEntry()
	usernameRegistrationEntry := widget.NewEntry()
	passwordRegistrationEntry := widget.NewPasswordEntry()
	passwordConfirmationRegistrationEntry := widget.NewPasswordEntry()
	textNameEntry := widget.NewEntry()
	textEntry := widget.NewEntry()
	textDescriptionEntry := widget.NewEntry()
	loginPasswordNameEntry := widget.NewEntry()
	loginEntry := widget.NewEntry()
	passwordEntry := widget.NewEntry()
	loginPasswordDescriptionEntry := widget.NewEntry()
	cardDescriptionEntry := widget.NewEntry()
	cardNameEntry := widget.NewEntry()
	paymentSystemEntry := widget.NewEntry()
	numberEntry := widget.NewEntry()
	holderEntry := widget.NewEntry()
	endDateEntry := widget.NewEntry()
	cvcEntry := widget.NewEntry()
	//---------------------------------------------------------------------- labels init
	labelAlertAuth := widget.NewLabel("")
	labelAlertLoginPassword := widget.NewLabel("")
	labelAlertText := widget.NewLabel("")
	labelAlertCard := widget.NewLabel("")
	labelAlertAuth.Hide()
	labelAlertLoginPassword.Hide()
	labelAlertText.Hide()
	labelAlertCard.Hide()
	//---------------------------------------------------------------------- forms init
	formLogin := component.GetFormLogin(usernameLoginEntry, passwordLoginEntry)
	formRegistration := component.GetFormRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry)
	formLoginPassword := component.GetFormLoginPassword(loginPasswordNameEntry, loginPasswordDescriptionEntry, loginEntry, passwordEntry)
	formText := component.GetFormText(textNameEntry, textDescriptionEntry, textEntry)
	formCard := component.GetFormCard(cardNameEntry, cardDescriptionEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)
	//---------------------------------------------------------------------- radio event
	radioAuth := widget.NewRadioGroup(radioOptions, func(value string) {
		log.Println("Radio set to ", value)
		if value == "Login" {
			window.SetContent(containerFormLogin)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
		if value == "Registration" {
			window.SetContent(containerFormRegistration)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
	})
	//---------------------------------------------------------------------- buttons event
	buttonTop = widget.NewButton(labels.BtnUpdateData, func() {
		dataTblText, dataTblCard, err = client.EventSynchronization(password, accessToken)
		if err != nil {
			labelAlertAuth.SetText(errors.ErrLogin)
		} else {
			tblText.Resize(fyne.NewSize(float32(len(dataTblText)), float32(len(dataTblText[0]))))
			tblText.Refresh()
			tblCard.Resize(fyne.NewSize(float32(len(dataTblCard)), float32(len(dataTblCard[0]))))
			tblCard.Refresh()
			window.SetContent(containerTabs)
		}
	})
	buttonLoginPassword = widget.NewButton(labels.BtnAddLoginPassword, func() {
		window.SetContent(containerFormLoginPassword)
		window.Show()
	})
	buttonText = widget.NewButton(labels.BtnAddText, func() {
		window.SetContent(containerFormText)
		window.Show()
	})
	buttonCard = widget.NewButton(labels.BtnAddCard, func() {
		window.SetContent(containerFormCard)
		window.Show()
	})
	//---------------------------------------------------------------------- table login password init
	tblLoginPassword = widget.NewTable(
		func() (int, int) {
			return len(dataTblLoginPassword), len(dataTblLoginPassword[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel(labels.TblLabel)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblLoginPassword[i.Row][i.Col])
		})
	form.SetDefaultColumnsWidthLoginPassword(tblLoginPassword)
	//---------------------------------------------------------------------- table text init
	tblText = widget.NewTable(
		func() (int, int) {
			return len(dataTblText), len(dataTblText[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel(labels.TblLabel)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblText[i.Row][i.Col])
		})
	form.SetDefaultColumnsWidthText(tblText)
	//---------------------------------------------------------------------- table card init
	tblCard = widget.NewTable(
		func() (int, int) {
			return len(dataTblCard), len(dataTblCard[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel(labels.TblLabel)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblCard[i.Row][i.Col])
		})
	form.SetDefaultColumnsWidthCard(tblCard)
	//---------------------------------------------------------------------- containerTabs
	tabLoginPassword = component.GetTabLoginPassword(tblLoginPassword, buttonTop, buttonLoginPassword)
	tabText = component.GetTabTexts(tblText, buttonTop, buttonText)
	tabCard = component.GetTabCards(tblCard, buttonTop, buttonCard)
	containerTabs = container.NewAppTabs(tabLoginPassword, tabText, tabCard)
	//---------------------------------------------------------------------- auth event
	buttonAuth = widget.NewButton("Submit", func() {
		labelAlertAuth.Show()
		valid = false
		if radioAuth.Selected == "Login" {
			valid = form.ValidateLogin(usernameLoginEntry, passwordLoginEntry, labelAlertAuth)
			if valid {
				accessToken, err = client.EventAuthentication(usernameLoginEntry.Text, passwordLoginEntry.Text)
				if err != nil {
					labelAlertAuth.SetText(errors.ErrLogin)
				} else {
					password = passwordLoginEntry.Text
					dataTblText, dataTblCard, err = client.EventSynchronization(password, accessToken)
					if err != nil {
						labelAlertAuth.SetText(errors.ErrLogin)
					} else {
						window.SetContent(containerTabs)
						window.Resize(fyne.NewSize(1250, 300))
						window.Show()
					}
				}
			}
		}
		if radioAuth.Selected == "Registration" {
			valid = form.ValidateRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry, labelAlertAuth)
			if valid {
				exist, err = client.EventUserExist(usernameRegistrationEntry.Text)
				if err != nil {
					labelAlertAuth.SetText(errors.ErrRegistration)
				}
				if exist {
					labelAlertAuth.SetText(errors.ErrUserExist)
				} else {
					accessToken, err = client.EventRegistration(usernameRegistrationEntry.Text, passwordRegistrationEntry.Text)
					if err != nil {
						labelAlertAuth.SetText(errors.ErrRegistration)
					} else {
						password = passwordRegistrationEntry.Text
						window.SetContent(containerTabs)
						window.Resize(fyne.NewSize(1250, 300))
						window.Show()
					}
				}
			}
		}
	})
	//---------------------------------------------------------------------- login password event
	buttonLoginPasswordAdd = widget.NewButton(labels.BtnAdd, func() {
		labelAlertLoginPassword.Show()
		valid = form.ValidateLoginPassword(false, loginPasswordNameEntry, loginPasswordDescriptionEntry, loginEntry,
			passwordEntry, labelAlertLoginPassword)
		if valid {
			err = client.EventCreateLoginPassword(loginPasswordNameEntry.Text, loginPasswordDescriptionEntry.Text, password,
				loginEntry.Text, passwordEntry.Text, accessToken)
			if err != nil {
				labelAlertLoginPassword.SetText(errors.ErrLoginPasswordAdd)
			} else {
				dataTblLoginPassword = append(dataTblLoginPassword, []string{loginPasswordNameEntry.Text, loginPasswordDescriptionEntry.Text,
					loginEntry.Text, passwordEntry.Text, time.Now().Format(layout), time.Now().Format(layout)})

				form.ClearLoginPassword(loginPasswordNameEntry, loginPasswordDescriptionEntry, loginEntry, passwordEntry)
				log.Info("Логин-пароль добавлен")

				formLoginPassword.Refresh()
				window.SetContent(containerTabs)
				window.Show()
			}
		}
		log.Debug(dataTblLoginPassword)
	})
	//---------------------------------------------------------------------- text event
	buttonTextAdd = widget.NewButton(labels.BtnAdd, func() {
		labelAlertText.Show()
		valid = false
		exist = table.SearchByColumn(dataTblText, 0, textNameEntry.Text) // search in map
		valid = form.ValidateText(exist, textNameEntry, textDescriptionEntry, textEntry, labelAlertText)
		if valid {
			err = client.EventCreateText(textNameEntry.Text, textDescriptionEntry.Text, password, textEntry.Text, accessToken)
			if err != nil {
				labelAlertText.SetText(errors.ErrTextAdd)
			} else {
				dataTblText = append(dataTblText, []string{textNameEntry.Text, textDescriptionEntry.Text, textEntry.Text,
					time.Now().Format(layout), time.Now().Format(layout)})

				form.ClearText(textNameEntry, textDescriptionEntry, textEntry)
				log.Info("Текст добавлен")

				formText.Refresh()
				window.SetContent(containerTabs)
				window.Show()
			}
		}
		log.Debug(dataTblText)
	})
	//---------------------------------------------------------------------- card event
	buttonCardAdd = widget.NewButton(labels.BtnAdd, func() {
		labelAlertCard.Show()
		valid = false
		exist = table.SearchByColumn(dataTblCard, 0, cardNameEntry.Text) // search in map
		valid = form.ValidateCard(exist, cardNameEntry, cardDescriptionEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry, labelAlertCard)
		if valid {
			err = client.EventCreateCard(cardNameEntry.Text, cardDescriptionEntry.Text, password, paymentSystemEntry.Text, numberEntry.Text, holderEntry.Text,
				endDateEntry.Text, cvcEntry.Text, accessToken)
			if err != nil {
				labelAlertCard.SetText(errors.ErrTextAdd)
			} else {
				layout := "01/02/2006 15:04:05"
				dataTblCard = append(dataTblCard, []string{cardNameEntry.Text, cardDescriptionEntry.Text, paymentSystemEntry.Text, numberEntry.Text, holderEntry.Text,
					cvcEntry.Text, endDateEntry.Text, time.Now().Format(layout), time.Now().Format(layout)})

				form.ClearCard(cardNameEntry, cardDescriptionEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)
				log.Info("Карта добавлена")

				formCard.Refresh()
				window.SetContent(containerTabs)
				window.Show()
			}

		}
		log.Debug(dataTblCard)
	})
	//---------------------------------------------------------------------- containers init
	containerRadio = container.NewVBox(radioAuth)
	containerFormLogin = container.NewVBox(formLogin, buttonAuth, labelAlertAuth, separator, radioAuth)
	containerFormRegistration = container.NewVBox(formRegistration, buttonAuth, labelAlertAuth, separator, radioAuth)
	containerFormLoginPassword = container.NewVBox(formLoginPassword, buttonLoginPasswordAdd, labelAlertLoginPassword)
	containerFormText = container.NewVBox(formText, buttonTextAdd, labelAlertText)
	containerFormCard = container.NewVBox(formCard, buttonCardAdd, labelAlertCard)
	//----------------------------------------------------------------------
	window.SetContent(containerRadio)
	window.ShowAndRun()
}
