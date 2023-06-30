package function

import (
	"log"
	"strconv"
	"time"
	"unicode/utf8"

	"fyne.io/fyne/v2/widget"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/algorithm"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/errors"
)

func ValidateLoginForm(usernameLoginEntry *widget.Entry, passwordLoginEntry *widget.Entry, labelAlertAuth *widget.Label) (string, bool) {
	if utf8.RuneCountInString(usernameLoginEntry.Text) < 6 {
		return errors.ErrUsernameIncorrect, false
	}
	if !encryption.VerifyPassword(passwordLoginEntry.Text) {
		return errors.ErrPasswordIncorrect, false
	}
	return "", true
}

func ValidateRegistration(usernameRegistrationEntry *widget.Entry, passwordRegistrationEntry *widget.Entry,
	passwordConfirmationRegistrationEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if utf8.RuneCountInString(usernameRegistrationEntry.Text) < 6 {
		labelAlertAuth.SetText(errors.ErrUsernameIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	if !encryption.VerifyPassword(passwordRegistrationEntry.Text) {
		labelAlertAuth.SetText(errors.ErrPasswordIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	if passwordRegistrationEntry.Text != passwordConfirmationRegistrationEntry.Text {
		labelAlertAuth.SetText(errors.ErrPasswordDifferent)
		log.Print(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateLoginPasswordForm(loginPasswordNameEntry *widget.Entry, loginPasswordDescriptionEntry *widget.Entry,
	loginEntry *widget.Entry, passwordEntry *widget.Entry) (string, bool) {
	if loginPasswordNameEntry.Text == "" {
		return errors.ErrNameEmpty, false
	}
	if loginPasswordDescriptionEntry.Text == "" {
		return errors.ErrDescriptionEmpty, false
	}
	if loginEntry.Text == "" {
		return errors.ErrLoginEmpty, false
	}
	if passwordEntry.Text == "" {
		return errors.ErrPasswordEmpty, false
	}
	return "", true
}

func ValidateTextForm(textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry) (string, bool) {
	if textNameEntry.Text == "" {
		return errors.ErrNameEmpty, false
	}
	if textDescriptionEntry.Text == "" {
		return errors.ErrDescriptionEmpty, false
	}
	if textEntry.Text == "" {
		return errors.ErrTextEmpty, false
	}
	return "", true
}

func ValidateCardForm(cardNameEntry *widget.Entry, cardDescriptionEntry *widget.Entry, paymentSystemEntry *widget.Entry,
	numberEntry *widget.Entry, holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry) (string, bool) {
	var err error
	if cardNameEntry.Text == "" {
		return errors.ErrNameEmpty, false
	}
	if cardDescriptionEntry.Text == "" {
		return errors.ErrDescriptionEmpty, false
	}
	if paymentSystemEntry.Text == "" {
		return errors.ErrPaymentSystemEmpty, false
	}
	if numberEntry.Text == "" {
		return errors.ErrNumberEmpty, false
	}
	intNumber, err := strconv.Atoi(numberEntry.Text)
	if err != nil {
		return errors.ErrNumberIncorrect, false
	}
	if !algorithm.ValidLuhn(intNumber) {
		return errors.ErrNumberIncorrect, false
	}
	if holderEntry.Text == "" {
		return errors.ErrHolderEmpty, false
	}
	if endDateEntry.Text == "" {
		return errors.ErrEndDateEmpty, false
	} else {
		layout := "01/02/2006"
		_, err = time.Parse(layout, endDateEntry.Text)
		if err != nil {
			return errors.ErrEndDataIncorrect, false
		}
	}
	if cvcEntry.Text == "" {
		return errors.ErrCvcEmpty, false
	} else {
		_, err = strconv.Atoi(cvcEntry.Text)
		if err != nil {
			return errors.ErrCvcIncorrect, false
		}
	}
	return "", true
}
