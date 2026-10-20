// Models of the Common section.

package common

import (
	"encoding/json"
	"strings"
)

// TemplateEngine — Параметр для вибору [шаблонізатора](https://docs.allymailer.com.ua/en/template-engines) — або
// "simple", або "velocity", або "liquid", або "none". За замовчуванням — "simple".
type TemplateEngine string

// TemplateEngine values.
const (
	TemplateEngineSimple   TemplateEngine = "simple"
	TemplateEngineVelocity TemplateEngine = "velocity"
	TemplateEngineLiquid   TemplateEngine = "liquid"
	TemplateEngineNone     TemplateEngine = "none"
)

// EmailStatus
type EmailStatus string

// EmailStatus values.
const (
	EmailStatusSent         EmailStatus = "sent"
	EmailStatusDelivered    EmailStatus = "delivered"
	EmailStatusOpened       EmailStatus = "opened"
	EmailStatusClicked      EmailStatus = "clicked"
	EmailStatusUnsubscribed EmailStatus = "unsubscribed"
	EmailStatusSubscribed   EmailStatus = "subscribed"
	EmailStatusSoftBounced  EmailStatus = "soft_bounced"
	EmailStatusHardBounced  EmailStatus = "hard_bounced"
	EmailStatusSpam         EmailStatus = "spam"
)

// VerificationRecordStatus — Лише домени зі статусом "confirmed" запису про володіння доменом можуть використовуватися для
// надсилання листів.
type VerificationRecordStatus string

// VerificationRecordStatus values.
const (
	VerificationRecordStatusConfirmed   VerificationRecordStatus = "confirmed"
	VerificationRecordStatusUnconfirmed VerificationRecordStatus = "unconfirmed"
)

// EmailCounterMode — Встановлює режим роботи лічильника листів. Дозволяє вибрати, як має поводитися лічильник:
// автоматично скидатися чи залишатися постійним. Припустимі значення: default - лічильник автоматично
// скидається в нуль на початку кожного облікового періоду, permanent - лічильник ніколи не скидається
// автоматично. Скидання може бути виконане лише вручну користувачем, передавши значення 0 у полі
// 'email_counter'. Якщо поле не передано в запиті або передано значення null, значення режиму не
// змінюється.
type EmailCounterMode string

// EmailCounterMode values.
const (
	EmailCounterModeDefault   EmailCounterMode = "default"
	EmailCounterModePermanent EmailCounterMode = "permanent"
)

// BodyObject — Об'єкт, що містить у собі html-, plaintext- і amp-частини листа. Або html-, або plaintext-частина
// має бути присутня обов'язково.
type BodyObject struct {
	// HTML-частина листа
	HTML *string `json:"html,omitempty"`

	// Plaintext-частина листа.
	Plaintext *string `json:"plaintext,omitempty"`

	// Необов'язкова
	// [AMP](http://amp.dev/documentation/guides-and-tutorials/learn/email_fundamentals/)-частина листа.
	Amp *string `json:"amp,omitempty"`
}

// HeadersObject — Об'єкт, що містить додаткові заголовки листа, максимум 50 заголовків. Префікс “X-” в імені заголовка
// обов'язковий, заголовки без нього будуть проігноровані, наприклад X-ALLYMAILER-Global-Language,
// X-ALLYMAILER-Template-Engine. Стандартні заголовки "To", "CC", "BCC" передаються без префікса "X-",
// але обробляються спеціальним чином, тому мають низку обмежень, докладніше про це можна прочитати
// [тут](/cc-and-bcc). Також якщо [техпідтримка](https://cp.allymailer.com/ua/support) схвалила вам
// вимкнення посилання відписки, то ви також можете передавати заголовки List-Unsubscribe,
// List-Subscribe, List-Help, List-Owner, List-Archive, In-Reply-To і References.
//
// Keys beyond the named fields (any header, e.g. X-MyHeader, CC, List-Unsubscribe) go to Extra: they
// are sent next to the named fields, a named field wins over an Extra key with the same name; unknown
// keys of a response land in Extra.
type HeadersObject struct {
	// Заголовок для вибору мови посилання та сторінки відписки. Припустимі значення "be", "de", "en",
	// "es", "fr", "it", "pl", "pt", "ru", "ua", "kz".
	XALLYMAILERGlobalLanguage *string `json:"X-ALLYMAILER-Global-Language,omitempty"`

	// Заголовок для вибору шаблонізатора — або “simple”, або “velocity”, або “liquid”, або “none”. За
	// замовчуванням — “simple”. Має пріоритет над значенням, що передається в параметрі “template_engine”.
	XALLYMAILERTemplateEngine *string `json:"X-ALLYMAILER-Template-Engine,omitempty"`

	// Extra holds the keys beyond the named fields, e.g. {"X-MyHeader": "some data"}.
	Extra map[string]string `json:"-"`
}

// headersObjectNamed are the JSON keys of the named fields of HeadersObject, lower case (encoding/json matches them so).
var headersObjectNamed = map[string]bool{
	"x-allymailer-global-language": true,
	"x-allymailer-template-engine": true,
}

// MarshalJSON writes the named fields and the Extra keys; a named field wins over an Extra key.
func (m HeadersObject) MarshalJSON() ([]byte, error) {
	type plain HeadersObject
	data, err := json.Marshal(plain(m))
	if err != nil || len(m.Extra) == 0 {
		return data, err
	}
	merged := map[string]json.RawMessage{}
	for key, value := range m.Extra {
		raw, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		merged[key] = raw
	}
	named := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &named); err != nil {
		return nil, err
	}
	for key, value := range named {
		merged[key] = value
	}
	return json.Marshal(merged)
}

// UnmarshalJSON reads the named fields; the other keys go to Extra.
func (m *HeadersObject) UnmarshalJSON(data []byte) error {
	type plain HeadersObject
	var fields plain
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	all := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}
	*m = HeadersObject(fields)
	m.Extra = nil
	for key, raw := range all {
		if headersObjectNamed[strings.ToLower(key)] {
			continue
		}
		var value string
		if json.Unmarshal(raw, &value) != nil {
			value = string(raw) // a number or an object: its JSON text
		}
		if m.Extra == nil {
			m.Extra = map[string]string{}
		}
		m.Extra[key] = value
	}
	return nil
}

// AttachmentsArrayItem
type AttachmentsArrayItem struct {
	// Тип вкладення, див. [MIME](https://en.wikipedia.org/wiki/MIME). Якщо не впевнені, використовуйте
	// "application/octet-stream".
	Type string `json:"type"`

	// Назва вкладення у форматі: "ім'я.розширення". У разі передачі кількох вкладень їхні назви мають бути
	// унікальними. У назві вкладень заборонено символ ‘/’.
	Name string `json:"name"`

	// Вміст файлу у [base64](https://en.wikipedia.org/wiki/Base64). Максимальний розмір файлу 7 Мб
	// (9786710 байт у base64).
	Content string `json:"content"`
}

// InlineAttachmentsArrayItem
type InlineAttachmentsArrayItem struct {
	// Тип вкладення, див. [MIME](https://en.wikipedia.org/wiki/MIME). Якщо не впевнені, використовуйте
	// "application/octet-stream".
	Type string `json:"type"`

	// Ідентифікатор вкладення. Якщо, наприклад, "name"="IMAGECID", то для виведення картинки з цього
	// вкладення в HTML-тексті листа потрібно вказати <img src="cid:IMAGECID">.
	Name string `json:"name"`

	// Вміст файлу у [base64](https://en.wikipedia.org/wiki/Base64). Максимальний розмір файлу 7 Мб
	// (9786710 байт у base64).
	Content string `json:"content"`
}

// SuccessResponse
type SuccessResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`
}
