package gui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/api/events"
	"github.com/vasiliyantufev/gophkeeper/internal/client/component/form"
	"github.com/vasiliyantufev/gophkeeper/internal/client/component/function"
	"github.com/vasiliyantufev/gophkeeper/internal/client/component/tab"
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

	var indexTblLoginPassword = 0
	var selectedRowTblLoginPassword []string
	var indexTblText = 0
	var selectedRowTblText []string
	var indexTblCard = 0
	var selectedRowTblCard []string

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

	var containerFormLoginPasswordCreate *fyne.Container
	var containerFormTextCreate *fyne.Container
	var containerFormCardCreate *fyne.Container

	var containerFormLoginPasswordUpdate *fyne.Container
	var containerFormTextUpdate *fyne.Container
	var containerFormCardUpdate *fyne.Container
	//---------------------------------------------------------------------- buttons
	var buttonAuth *widget.Button

	var buttonTopBack *widget.Button
	var buttonTopSynchronization *widget.Button

	var buttonLoginPassword *widget.Button
	var buttonText *widget.Button
	var buttonCard *widget.Button

	var buttonLoginPasswordCreate *widget.Button
	var buttonLoginPasswordDelete *widget.Button
	var buttonLoginPasswordUpdate *widget.Button
	var buttonLoginPasswordFormUpdate *widget.Button

	var buttonTextCreate *widget.Button
	var buttonTextDelete *widget.Button
	var buttonTextUpdate *widget.Button
	var buttonTextFormUpdate *widget.Button

	var buttonCardCreate *widget.Button
	var buttonCardDelete *widget.Button
	var buttonCardUpdate *widget.Button
	var buttonCardFormUpdate *widget.Button
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

	textNameEntryCreate := widget.NewEntry()
	textDescriptionEntryCreate := widget.NewEntry()
	textEntryCreate := widget.NewEntry()

	textNameEntryUpdate := widget.NewEntry()
	textNameEntryUpdate.Disable()
	textDescriptionEntryUpdate := widget.NewEntry()
	textDescriptionEntryUpdate.Disable()
	textEntryUpdate := widget.NewEntry()

	loginPasswordNameEntryCreate := widget.NewEntry()
	loginPasswordDescriptionEntryCreate := widget.NewEntry()
	loginEntryCreate := widget.NewEntry()
	passwordEntryCreate := widget.NewEntry()

	loginPasswordNameEntryUpdate := widget.NewEntry()
	loginPasswordNameEntryUpdate.Disable()
	loginPasswordDescriptionEntryUpdate := widget.NewEntry()
	loginPasswordDescriptionEntryUpdate.Disable()
	loginEntryUpdate := widget.NewEntry()
	passwordEntryUpdate := widget.NewEntry()

	cardNameEntryCreate := widget.NewEntry()
	cardDescriptionEntryCreate := widget.NewEntry()
	paymentSystemEntryCreate := widget.NewEntry()
	numberEntryCreate := widget.NewEntry()
	holderEntryCreate := widget.NewEntry()
	endDateEntryCreate := widget.NewEntry()
	cvcEntryCreate := widget.NewEntry()

	cardNameEntryUpdate := widget.NewEntry()
	cardNameEntryUpdate.Disable()
	cardDescriptionEntryUpdate := widget.NewEntry()
	cardDescriptionEntryUpdate.Disable()
	paymentSystemEntryUpdate := widget.NewEntry()
	numberEntryUpdate := widget.NewEntry()
	holderEntryUpdate := widget.NewEntry()
	endDateEntryUpdate := widget.NewEntry()
	cvcEntryUpdate := widget.NewEntry()

	//---------------------------------------------------------------------- labels init
	labelAlertAuth := widget.NewLabel("")
	labelAlertLoginPassword := widget.NewLabel("")
	labelAlertLoginPasswordCreate := widget.NewLabel("")
	labelAlertLoginPasswordUpdate := widget.NewLabel("")
	labelAlertText := widget.NewLabel("")
	labelAlertTextCreate := widget.NewLabel("")
	labelAlertTextUpdate := widget.NewLabel("")
	labelAlertCard := widget.NewLabel("")
	labelAlertCardCreate := widget.NewLabel("")
	labelAlertCardUpdate := widget.NewLabel("")

	labelAlertAuth.Hide()
	labelAlertLoginPassword.Hide()
	labelAlertLoginPasswordCreate.Hide()
	labelAlertLoginPasswordUpdate.Hide()
	labelAlertText.Hide()
	labelAlertTextCreate.Hide()
	labelAlertTextUpdate.Hide()
	labelAlertCard.Hide()
	labelAlertCardCreate.Hide()
	labelAlertCardUpdate.Hide()
	//---------------------------------------------------------------------- forms init
	formLogin := form.GetFormLogin(usernameLoginEntry, passwordLoginEntry)
	formRegistration := form.GetFormRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry)

	formLoginPasswordCreate := form.GetFormLoginPassword(loginPasswordNameEntryCreate, loginPasswordDescriptionEntryCreate, loginEntryCreate, passwordEntryCreate)
	formLoginPasswordUpdate := form.GetFormLoginPassword(loginPasswordNameEntryUpdate, loginPasswordDescriptionEntryUpdate, loginEntryUpdate, passwordEntryUpdate)

	formTextCreate := form.GetFormText(textNameEntryCreate, textDescriptionEntryCreate, textEntryCreate)
	formTextUpdate := form.GetFormText(textNameEntryUpdate, textDescriptionEntryUpdate, textEntryUpdate)

	formCardCreate := form.GetFormCard(cardNameEntryCreate, cardDescriptionEntryCreate, paymentSystemEntryCreate, numberEntryCreate, holderEntryCreate, endDateEntryCreate, cvcEntryCreate)
	formCardUpdate := form.GetFormCard(cardNameEntryUpdate, cardDescriptionEntryUpdate, paymentSystemEntryUpdate, numberEntryUpdate, holderEntryUpdate, endDateEntryUpdate, cvcEntryUpdate)
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
	buttonTopSynchronization = widget.NewButton(labels.BtnUpdateData, func() {
		dataTblText, dataTblCard, dataTblLoginPassword, err = client.EventSynchronization(password, accessToken)
		if err != nil {
			labelAlertAuth.SetText(errors.ErrLogin)
		} else {
			tblLoginPassword.Resize(fyne.NewSize(float32(len(dataTblLoginPassword)), float32(len(dataTblLoginPassword[0]))))
			tblLoginPassword.Refresh()
			tblText.Resize(fyne.NewSize(float32(len(dataTblText)), float32(len(dataTblText[0]))))
			tblText.Refresh()
			tblCard.Resize(fyne.NewSize(float32(len(dataTblCard)), float32(len(dataTblCard[0]))))
			tblCard.Refresh()
			window.SetContent(containerTabs)
		}
	})
	buttonLoginPassword = widget.NewButton(labels.BtnAddLoginPassword, func() {
		window.SetContent(containerFormLoginPasswordCreate)
		window.Show()
	})
	buttonText = widget.NewButton(labels.BtnAddText, func() {
		window.SetContent(containerFormTextCreate)
		window.Show()
	})
	buttonCard = widget.NewButton(labels.BtnAddCard, func() {
		window.SetContent(containerFormCardCreate)
		window.Show()
	})
	buttonLoginPasswordDelete = widget.NewButton(labels.BtnDeleteLoginPassword, func() {
		function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
		if indexTblLoginPassword > 0 {
			client.EventDeleteLoginPassword(selectedRowTblLoginPassword, accessToken)
			// Удаляем строку с индексом indexTblLoginPassword
			dataTblLoginPassword = table.RemoveRow(dataTblLoginPassword, indexTblLoginPassword)
			indexTblLoginPassword = 0
		} else {
			logrus.Error(errors.ErrLoginPasswordTblIndexDelete)
			labelAlertLoginPassword.Show()
			labelAlertLoginPassword.SetText(errors.ErrLoginPasswordTblIndexDelete)
		}
	})
	buttonTextDelete = widget.NewButton(labels.BtnDeleteText, func() {
		function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
		if indexTblText > 0 {
			client.EventDeleteText(selectedRowTblText, accessToken)
			// Удаляем строку с индексом indexTblText
			dataTblText = table.RemoveRow(dataTblText, indexTblText)
			indexTblText = 0
		} else {
			logrus.Error(errors.ErrTextTblIndexDelete)
			labelAlertText.Show()
			labelAlertText.SetText(errors.ErrTextTblIndexDelete)
		}
	})
	buttonCardDelete = widget.NewButton(labels.BtnDeleteCard, func() {
		function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
		if indexTblCard > 0 {
			client.EventDeleteCard(selectedRowTblCard, accessToken)
			// Удаляем строку с индексом indexTblCard
			dataTblCard = table.RemoveRow(dataTblCard, indexTblCard)
			indexTblCard = 0
		} else {
			logrus.Error(errors.ErrCardTblIndexDelete)
			labelAlertCard.Show()
			labelAlertCard.SetText(errors.ErrCardTblIndexDelete)
		}
	})
	buttonLoginPasswordUpdate = widget.NewButton(labels.BtnUpdateLoginPassword, func() {
		if indexTblLoginPassword > 0 {
			function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
			function.SetLoginPasswordData(selectedRowTblLoginPassword, loginPasswordNameEntryUpdate, loginPasswordDescriptionEntryUpdate, loginEntryUpdate, passwordEntryUpdate)
			window.SetContent(containerFormLoginPasswordUpdate)
			window.Show()
		} else {
			logrus.Error(errors.ErrLoginPasswordTblIndexUpdate)
			labelAlertLoginPassword.Show()
			labelAlertLoginPassword.SetText(errors.ErrLoginPasswordTblIndexUpdate)
		}
	})
	buttonTextUpdate = widget.NewButton(labels.BtnUpdateText, func() {
		if indexTblText > 0 {
			function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
			function.SetTextData(selectedRowTblText, textNameEntryUpdate, textDescriptionEntryUpdate, textEntryUpdate)
			window.SetContent(containerFormTextUpdate)
			window.Show()
		} else {
			logrus.Error(errors.ErrTextTblIndexUpdate)
			labelAlertText.Show()
			labelAlertText.SetText(errors.ErrTextTblIndexUpdate)
		}
	})
	buttonCardUpdate = widget.NewButton(labels.BtnUpdateCard, func() {
		if indexTblCard > 0 {
			function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
			window.SetContent(containerFormCardUpdate)
			window.Show()
		} else {
			logrus.Error(errors.ErrCardTblIndexUpdate)
			labelAlertCard.Show()
			labelAlertCard.SetText(errors.ErrCardTblIndexUpdate)
		}
	})

	buttonLoginPasswordFormUpdate = widget.NewButton(labels.BtnUpdate, func() {
		logrus.Info(labels.BtnUpdate)
	})
	buttonTextFormUpdate = widget.NewButton(labels.BtnUpdate, func() {
		logrus.Info(labels.BtnUpdate)
	})
	buttonCardFormUpdate = widget.NewButton(labels.BtnUpdate, func() {
		logrus.Info(labels.BtnUpdate)
	})

	buttonTopBack = widget.NewButton(labels.BtnBack, func() {
		function.ClearLoginPassword(loginPasswordNameEntryCreate, loginPasswordDescriptionEntryCreate, loginEntryCreate, passwordEntryCreate)
		function.ClearText(textNameEntryCreate, textDescriptionEntryCreate, textEntryCreate)
		function.ClearCard(cardNameEntryCreate, cardDescriptionEntryCreate, paymentSystemEntryCreate, numberEntryCreate, holderEntryCreate, endDateEntryCreate, cvcEntryCreate)
		function.ClearLoginPassword(loginPasswordNameEntryUpdate, loginPasswordDescriptionEntryUpdate, loginEntryUpdate, passwordEntryUpdate)
		function.ClearText(textNameEntryUpdate, textDescriptionEntryUpdate, textEntryUpdate)
		function.ClearCard(cardNameEntryUpdate, cardDescriptionEntryUpdate, paymentSystemEntryUpdate, numberEntryUpdate, holderEntryUpdate, endDateEntryUpdate, cvcEntryUpdate)
		labelAlertLoginPasswordCreate.Hide()
		labelAlertLoginPasswordUpdate.Hide()
		labelAlertTextCreate.Hide()
		labelAlertTextUpdate.Hide()
		labelAlertCardCreate.Hide()
		labelAlertCardUpdate.Hide()
		window.SetContent(containerTabs)
		window.Resize(fyne.NewSize(1250, 300))
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
	function.SetDefaultColumnsWidthLoginPassword(tblLoginPassword)
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
	function.SetDefaultColumnsWidthText(tblText)
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
	function.SetDefaultColumnsWidthCard(tblCard)
	//---------------------------------------------------------------------- containerTabs
	tabLoginPassword = tab.GetTabLoginPassword(tblLoginPassword, buttonTopSynchronization, buttonLoginPassword, buttonLoginPasswordDelete, buttonLoginPasswordUpdate, labelAlertLoginPassword)
	tabText = tab.GetTabTexts(tblText, buttonTopSynchronization, buttonText, buttonTextDelete, buttonTextUpdate, labelAlertText)
	tabCard = tab.GetTabCards(tblCard, buttonTopSynchronization, buttonCard, buttonCardDelete, buttonCardUpdate, labelAlertCard)
	containerTabs = container.NewAppTabs(tabLoginPassword, tabText, tabCard)
	//----------------------------------------------------------------------
	// Get selected row data
	tblLoginPassword.OnSelected = func(id widget.TableCellID) {
		indexTblLoginPassword = id.Row
		selectedRowTblLoginPassword = dataTblLoginPassword[id.Row]
	}
	tblText.OnSelected = func(id widget.TableCellID) {
		indexTblText = id.Row
		selectedRowTblText = dataTblText[id.Row]
	}
	tblCard.OnSelected = func(id widget.TableCellID) {
		indexTblCard = id.Row
		selectedRowTblCard = dataTblCard[id.Row]
	}
	//---------------------------------------------------------------------- auth event
	buttonAuth = widget.NewButton("Submit", func() {
		labelAlertAuth.Show()
		valid = false
		if radioAuth.Selected == "Login" {
			valid = function.ValidateLogin(usernameLoginEntry, passwordLoginEntry, labelAlertAuth)
			if valid {
				accessToken, err = client.EventAuthentication(usernameLoginEntry.Text, passwordLoginEntry.Text)
				if err != nil {
					labelAlertAuth.SetText(errors.ErrLogin)
				} else {
					password = passwordLoginEntry.Text
					dataTblText, dataTblCard, dataTblLoginPassword, err = client.EventSynchronization(password, accessToken)
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
			valid = function.ValidateRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry, labelAlertAuth)
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
	buttonLoginPasswordCreate = widget.NewButton(labels.BtnAdd, func() {
		labelAlertLoginPasswordCreate.Show()
		function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
		exist = table.SearchByColumn(dataTblLoginPassword, 0, loginPasswordNameEntryCreate.Text) // search in map
		valid = false
		valid = function.ValidateLoginPassword(false, loginPasswordNameEntryCreate, loginPasswordDescriptionEntryCreate, loginEntryCreate,
			passwordEntryCreate, labelAlertLoginPasswordCreate)
		if valid {
			err = client.EventCreateLoginPassword(loginPasswordNameEntryCreate.Text, loginPasswordDescriptionEntryCreate.Text, password,
				loginEntryCreate.Text, passwordEntryCreate.Text, accessToken)
			if err != nil {
				labelAlertLoginPasswordCreate.SetText(errors.ErrLoginPasswordAdd)
			} else {
				dataTblLoginPassword = append(dataTblLoginPassword, []string{loginPasswordNameEntryCreate.Text, loginPasswordDescriptionEntryCreate.Text,
					loginEntryCreate.Text, passwordEntryCreate.Text, time.Now().Format(layout), time.Now().Format(layout)})

				function.ClearLoginPassword(loginPasswordNameEntryCreate, loginPasswordDescriptionEntryCreate, loginEntryCreate, passwordEntryCreate)
				log.Info("Логин-пароль добавлен")

				labelAlertLoginPasswordCreate.Hide()
				formLoginPasswordCreate.Refresh()
				window.SetContent(containerTabs)
				window.Show()
			}
		}
		log.Debug(dataTblLoginPassword)
	})
	//---------------------------------------------------------------------- text event
	buttonTextCreate = widget.NewButton(labels.BtnAdd, func() {
		labelAlertTextCreate.Show()
		function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
		exist = table.SearchByColumn(dataTblText, 0, textNameEntryCreate.Text) // search in map
		valid = false
		valid = function.ValidateText(exist, textNameEntryCreate, textDescriptionEntryCreate, textEntryCreate, labelAlertTextCreate)
		if valid {
			err = client.EventCreateText(textNameEntryCreate.Text, textDescriptionEntryCreate.Text, password, textEntryCreate.Text, accessToken)
			if err != nil {
				labelAlertTextCreate.SetText(errors.ErrTextAdd)
			} else {
				dataTblText = append(dataTblText, []string{textNameEntryCreate.Text, textDescriptionEntryCreate.Text, textEntryCreate.Text,
					time.Now().Format(layout), time.Now().Format(layout)})

				function.ClearText(textNameEntryCreate, textDescriptionEntryCreate, textEntryCreate)
				log.Info("Текст добавлен")

				labelAlertTextCreate.Hide()
				formTextCreate.Refresh()
				window.SetContent(containerTabs)
				window.Show()
			}
		}
		log.Debug(dataTblText)
	})
	//---------------------------------------------------------------------- card event
	buttonCardCreate = widget.NewButton(labels.BtnAdd, func() {
		labelAlertCardCreate.Show()
		function.HideLabelsTab(labelAlertLoginPassword, labelAlertText, labelAlertCard)
		exist = table.SearchByColumn(dataTblCard, 0, cardNameEntryCreate.Text) // search in map
		valid = false
		valid = function.ValidateCard(exist, cardNameEntryCreate, cardDescriptionEntryCreate, paymentSystemEntryCreate, numberEntryCreate, holderEntryCreate, endDateEntryCreate, cvcEntryCreate, labelAlertCardCreate)
		if valid {
			err = client.EventCreateCard(cardNameEntryCreate.Text, cardDescriptionEntryCreate.Text, password, paymentSystemEntryCreate.Text, numberEntryCreate.Text, holderEntryCreate.Text,
				endDateEntryCreate.Text, cvcEntryCreate.Text, accessToken)
			if err != nil {
				labelAlertCardCreate.SetText(errors.ErrCardAdd)
			} else {
				layout := "01/02/2006 15:04:05"
				dataTblCard = append(dataTblCard, []string{cardNameEntryCreate.Text, cardDescriptionEntryCreate.Text, paymentSystemEntryCreate.Text, numberEntryCreate.Text, holderEntryCreate.Text,
					cvcEntryCreate.Text, endDateEntryCreate.Text, time.Now().Format(layout), time.Now().Format(layout)})

				function.ClearCard(cardNameEntryCreate, cardDescriptionEntryCreate, paymentSystemEntryCreate, numberEntryCreate, holderEntryCreate, endDateEntryCreate, cvcEntryCreate)
				log.Info("Карта добавлена")

				labelAlertCardCreate.Hide()
				formCardCreate.Refresh()
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

	containerFormLoginPasswordCreate = container.NewVBox(buttonTopBack, formLoginPasswordCreate, buttonLoginPasswordCreate, labelAlertLoginPasswordCreate)
	containerFormLoginPasswordUpdate = container.NewVBox(buttonTopBack, formLoginPasswordUpdate, buttonLoginPasswordFormUpdate, labelAlertLoginPasswordUpdate)

	containerFormTextCreate = container.NewVBox(buttonTopBack, formTextCreate, buttonTextCreate, labelAlertTextCreate)
	containerFormTextUpdate = container.NewVBox(buttonTopBack, formTextUpdate, buttonTextFormUpdate, labelAlertTextUpdate)

	containerFormCardCreate = container.NewVBox(buttonTopBack, formCardCreate, buttonCardCreate, labelAlertCardCreate)
	containerFormCardUpdate = container.NewVBox(buttonTopBack, formCardUpdate, buttonCardFormUpdate, labelAlertCardUpdate)

	//----------------------------------------------------------------------
	window.SetContent(containerRadio)
	window.ShowAndRun()
}
