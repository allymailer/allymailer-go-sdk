// Models of the EmailValidation section.

package emailvalidation

// EmailValidationSingleRequest
type EmailValidationSingleRequest struct {
	// Email-адреса, яку необхідно перевірити.
	Email string `json:"email"`
}

// EmailValidationSingleResponse
type EmailValidationSingleResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Email, що перевіряється.
	Email string `json:"email,omitempty"`

	// Можливі статуси:
	//
	// - "valid" - email-адреса дійсна,
	// - "invalid" - email-адреса недійсна,
	// - "suspicious" - email-адреса підозріла,
	// - "unknown" - адресу не вдалося перевірити. Приймаючий сервер не відповів у межах припустимого часу перевірки.
	Result string `json:"result,omitempty"`

	// Можливі статуси:
	//
	// - "no_mx_record" - у домену відсутні MX-записи,
	// - "syntax_error" - адреса не відповідає стандартам email,
	// - "possible_typo" - можливо, адресу введено з помилкою,
	// - "mailbox_not_found" - адреса не існує,
	// - "global_suppression" - адреса перманентно недоступна в сервісі через багаторазові недоставки,
	// - "disposable" - одноразовий email, зазвичай такі адреси приймають пошту лише декілька хвилин,
	// - "role" - загальна адреса електронної пошти, найчастіше використовується групою людей і не належить конкретній особі,
	// - "abuse" - адреси, з яких надходить велика кількість скарг і натискань на "спам", іноді навіть в автоматичному режимі,
	// - "spamtrap" - адреси-спам-пастки, зазвичай розміщуються поштовими провайдерами у відкритому доступі, але ніколи не використовуються для підписки. Надсилання листів на такі адреси сильно шкодить репутації відправника,
	// - "smtp_connection_failed" - SMTP-сервер недоступний — можливо, адресу введено з помилкою.
	Cause string `json:"cause,omitempty"`

	// Імовірність того, що адреса дійсна. Від 0 до 100 (100 - дійсна, 0 - недійсна).
	Validity int `json:"validity,omitempty"`

	// Локальна частина email-адреси (усе, що перед знаком "@").
	LocalPart string `json:"local_part,omitempty"`

	// Доменна частина email-адреси.
	Domain string `json:"domain,omitempty"`

	// Чи є у домену MX-запис? true якщо є, false якщо відсутній.
	MxFound bool `json:"mx_found,omitempty"`

	// Бажана MX-запис домену.
	MxRecord string `json:"mx_record,omitempty"`

	// Можлива адреса, отримана після виправлення помилки (для cause=possible_typo).
	DidYouMean string `json:"did_you_mean,omitempty"`

	// Дата й час перевірки адреси у форматі РРРР-ММ-ДД гг:хх:сс у часовому поясі UTC.
	ProcessedAt string `json:"processed_at,omitempty"`
}
