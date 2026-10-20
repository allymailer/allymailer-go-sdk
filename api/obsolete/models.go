// Models of the Obsolete section.

package obsolete

// ObsoleteUnsubscribedSetMessage — Подробиці відписки:
//
// - "unsubscribed" якщо email було щойно відписано,
// - "was inactive" якщо він уже був позначений як неактивний і немає сенсу відписувати,
// - "was unsubscribed" якщо email уже було відписано раніше,
// - "added and unsubscribed" якщо для цього email не зберігалося жодного більш раннього статусу, тому його було додано й одразу ж відписано.
type ObsoleteUnsubscribedSetMessage string

// ObsoleteUnsubscribedSetMessage values.
const (
	ObsoleteUnsubscribedSetMessageUnsubscribed         ObsoleteUnsubscribedSetMessage = "unsubscribed"
	ObsoleteUnsubscribedSetMessageWasInactive          ObsoleteUnsubscribedSetMessage = "was inactive"
	ObsoleteUnsubscribedSetMessageWasUnsubscribed      ObsoleteUnsubscribedSetMessage = "was unsubscribed"
	ObsoleteUnsubscribedSetMessageAddedAndUnsubscribed ObsoleteUnsubscribedSetMessage = "added and unsubscribed"
)

// ObsoleteUnsubscribedSetRequest
type ObsoleteUnsubscribedSetRequest struct {
	// Email, який має бути відписано.
	Address string `json:"address"`
}

// ObsoleteUnsubscribedSetResponse
type ObsoleteUnsubscribedSetResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Відписаний email.
	Address string `json:"address,omitempty"`

	// Подробиці відписки:
	//
	// - "unsubscribed" якщо email було щойно відписано,
	// - "was inactive" якщо він уже був позначений як неактивний і немає сенсу відписувати,
	// - "was unsubscribed" якщо email уже було відписано раніше,
	// - "added and unsubscribed" якщо для цього email не зберігалося жодного більш раннього статусу, тому його було додано й одразу ж відписано.
	Message ObsoleteUnsubscribedSetMessage `json:"message,omitempty"`
}

// ObsoleteUnsubscribedCheckRequest
type ObsoleteUnsubscribedCheckRequest struct {
	// Email, який необхідно перевірити.
	Address string `json:"address"`
}

// ObsoleteUnsubscribedCheckResponse
type ObsoleteUnsubscribedCheckResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Email, який було перевірено.
	Address string `json:"address,omitempty"`

	// true якщо відписано, false якщо ні.
	IsUnsubscribed bool `json:"is_unsubscribed,omitempty"`
}

// ObsoleteUnsubscribedListRequest
type ObsoleteUnsubscribedListRequest struct {
	// Дата у форматі РРРР-ММ-ДД, щоб отримати всі відписані адреси з "date_from" по сьогоднішній день.
	// Якщо параметр не вказати, період за замовчуванням буде з 00:00 сьогоднішнього дня за UTC.
	DateFrom *string `json:"date_from,omitempty"`
}

// ObsoleteUnsubscribedListResponse
type ObsoleteUnsubscribedListResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Масив об'єктів з даними відписаних.
	Unsubscribed []ObsoleteUnsubscribedListResponseUnsubscribedItem `json:"unsubscribed,omitempty"`
}

// ObsoleteUnsubscribedListResponseUnsubscribedItem
type ObsoleteUnsubscribedListResponseUnsubscribedItem struct {
	// Відписаний email.
	Address string `json:"address,omitempty"`

	// Дата й час відписки адреси у форматі РРРР-ММ-ДД гг:хх:сс у часовому поясі UTC.
	UnsubscribedOn string `json:"unsubscribed_on,omitempty"`
}
