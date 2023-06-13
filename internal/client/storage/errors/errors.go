package errors

var (
	ErrLogin              = "Неудачная попытка аутентификации"
	ErrRegistration       = "Неудачная попытка регистрации"
	ErrUserExist          = "Пользователь с таким username зарегистрирован"
	ErrUsernameIncorrect  = "Длинна username должна быть не менее шести символов"
	ErrPasswordIncorrect  = "Длинна password должна быть не менее шести символов"
	ErrPasswordDifferent  = "Пароли не совпали"
	ErrNameEmpty          = "Name не заполнен"
	ErrTextEmpty          = "Text не заполнен"
	ErrTextExist          = "Текст с таким name уже существует"
	ErrCartExist          = "Карта с таким name уже существует"
	ErrPaymentSystemEmpty = "Payment System не заполнен"
	ErrNumberEmpty        = "Number не заполнен"
	ErrHolderEmpty        = "Holder не заполнен"
	ErrEndDateEmpty       = "End date не заполнен"
	ErrEndDataIncorrect   = "End Date некорректный (пример: 01/02/2006)"
	ErrCvcEmpty           = "CVC не заполнен"
	ErrCvcIncorrect       = "CVC некорректный (пример: 123)"
)
