package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/vasiliyantufev/gophkeeper/internal/client/api/data"
	"github.com/vasiliyantufev/gophkeeper/internal/client/component"
	form2 "github.com/vasiliyantufev/gophkeeper/internal/client/component/form"
	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/table"
)

func main() {
	application := app.New()
	application.Settings().SetTheme(theme.LightTheme())
	window := application.NewWindow("GophKeeper")
	window.Resize(fyne.NewSize(250, 80))
	//---------------------------------------------------------------------- variables
	var dataTblText = [][]string{{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"}}
	var dataTblCart = [][]string{{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED_AT", "UPDATED_AT"}}
	var radioOptions = []string{"Login", "Registration"}
	var user = model.User{}
	var exist bool
	var valid bool
	var layout string
	layout = "01/02/2006 15:04:05"
	//---------------------------------------------------------------------- containers
	var containerRadio *fyne.Container
	var containerFormLogin *fyne.Container
	var containerFormRegistration *fyne.Container
	var containerFormText *fyne.Container
	var containerFormCart *fyne.Container
	//---------------------------------------------------------------------- buttons
	var buttonAuth *widget.Button
	var buttonTop *widget.Button
	var buttonText *widget.Button
	var buttonCart *widget.Button
	var buttonTextAdd *widget.Button
	var buttonCartAdd *widget.Button
	//---------------------------------------------------------------------- tabs
	var containerTabs *container.AppTabs
	var tblText *widget.Table
	var tblCart *widget.Table
	var tabText *container.TabItem
	var tabCart *container.TabItem
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
	cartNameEntry := widget.NewEntry()
	paymentSystemEntry := widget.NewEntry()
	numberEntry := widget.NewEntry()
	holderEntry := widget.NewEntry()
	endDateEntry := widget.NewEntry()
	cvcEntry := widget.NewEntry()
	//---------------------------------------------------------------------- labels init
	labelAlertAuth := widget.NewLabel("")
	labelAlertText := widget.NewLabel("")
	labelAlertCart := widget.NewLabel("")
	labelAlertAuth.Hide()
	labelAlertText.Hide()
	labelAlertCart.Hide()
	//---------------------------------------------------------------------- forms init
	formLogin := component.GetFormLogin(usernameLoginEntry, passwordLoginEntry)
	formRegistration := component.GetFormRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry)
	formText := component.GetFormText(textNameEntry, textEntry, textDescriptionEntry)
	formCart := component.GetFormCart(cartNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)
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
	buttonTop = widget.NewButton("Обновить данные", func() {
		dataTblText, dataTblCart = data.Sync(user.ID)
		tblText.Resize(fyne.NewSize(float32(len(dataTblText)), float32(len(dataTblText[0]))))
		tblText.Refresh()
		tblCart.Resize(fyne.NewSize(float32(len(dataTblCart)), float32(len(dataTblCart[0]))))
		tblCart.Refresh()
		window.SetContent(containerTabs)
	})
	buttonText = widget.NewButton("Добавить текстовые данные", func() {
		window.SetContent(containerFormText)
		window.Show()
	})
	buttonCart = widget.NewButton("Добавить банковскую карту", func() {
		window.SetContent(containerFormCart)
		window.Show()
	})
	//---------------------------------------------------------------------- table text init
	tblText = widget.NewTable(
		func() (int, int) {
			return len(dataTblText), len(dataTblText[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblText[i.Row][i.Col])
		})
	form2.SetDefaultColumnsWidthText(tblText)
	//---------------------------------------------------------------------- table cart init
	tblCart = widget.NewTable(
		func() (int, int) {
			return len(dataTblCart), len(dataTblCart[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblCart[i.Row][i.Col])
		})
	form2.SetDefaultColumnsWidthCart(tblCart)
	//---------------------------------------------------------------------- containerTabs
	tabText = component.GetTabTexts(tblText, buttonTop, buttonText)
	tabCart = component.GetTabCarts(tblCart, buttonTop, buttonCart)
	containerTabs = container.NewAppTabs(tabText, tabCart)
	//---------------------------------------------------------------------- auth event
	buttonAuth = widget.NewButton("Submit", func() {
		labelAlertAuth.Show()
		valid = false
		if radioAuth.Selected == "Login" {
			valid = form2.ValidateLogin(usernameLoginEntry, passwordLoginEntry, labelAlertAuth)
			if valid {
				user, exist = user.Authentication(usernameLoginEntry.Text, passwordLoginEntry.Text) //ищем в бд
				if exist {
					dataTblText, dataTblCart = data.Sync(user.ID)
					window.SetContent(containerTabs)
					window.Resize(fyne.NewSize(1250, 300))
					window.Show()
				}
			}
		}
		if radioAuth.Selected == "Registration" {
			valid = form2.ValidateRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry, labelAlertAuth)
			if valid {
				exist = user.UserExist(usernameRegistrationEntry.Text) //ищем в бд
				if !exist {
					user = user.Registration(usernameRegistrationEntry.Text, passwordRegistrationEntry.Text)
					window.SetContent(containerTabs)
					window.Resize(fyne.NewSize(1250, 300))
					window.Show()
				}
			}
		}
	})
	//---------------------------------------------------------------------- text event
	buttonTextAdd = widget.NewButton("Добавить", func() {
		labelAlertText.Show()
		valid = false
		exist = table.SearchByColumn(dataTblText, 0, textNameEntry.Text) //ищем в мапке
		valid = form2.ValidateText(exist, textNameEntry, textEntry, textDescriptionEntry, labelAlertText)
		if valid {
			dataTblText = append(dataTblText, []string{textNameEntry.Text, textEntry.Text, textDescriptionEntry.Text,
				time.Now().Format(layout), time.Now().Format(layout)})

			form2.ClearText(textNameEntry, textEntry, textDescriptionEntry)
			log.Print("Текст добавлен")

			labelAlertText.Hide()
			formText.Refresh()
			window.SetContent(containerTabs)
			window.Show()
		}
		log.Print(dataTblText)
	})
	//---------------------------------------------------------------------- cart event
	buttonCartAdd = widget.NewButton("Добавить", func() {
		labelAlertCart.Show()
		valid = false
		exist = table.SearchByColumn(dataTblCart, 0, cartNameEntry.Text) //ищем в мапке
		valid = form2.ValidateCart(exist, cartNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry, labelAlertCart)
		if valid {
			layout := "01/02/2006 15:04:05"
			dataTblCart = append(dataTblCart, []string{cartNameEntry.Text, paymentSystemEntry.Text, numberEntry.Text, holderEntry.Text,
				cvcEntry.Text, endDateEntry.Text, time.Now().Format(layout), time.Now().Format(layout)})

			form2.ClearCart(cartNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)
			log.Print("Карта добавлена")

			labelAlertCart.Hide()
			formCart.Refresh()
			window.SetContent(containerTabs)
			window.Show()
		}
		log.Print(dataTblCart)
	})
	//---------------------------------------------------------------------- containers init
	containerRadio = container.NewVBox(radioAuth)
	containerFormLogin = container.NewVBox(formLogin, buttonAuth, labelAlertAuth, separator, radioAuth)
	containerFormRegistration = container.NewVBox(formRegistration, buttonAuth, labelAlertAuth, separator, radioAuth)
	containerFormText = container.NewVBox(formText, buttonTextAdd, labelAlertText)
	containerFormCart = container.NewVBox(formCart, buttonCartAdd, labelAlertCart)
	//----------------------------------------------------------------------
	window.SetContent(containerRadio)
	window.ShowAndRun()
}
