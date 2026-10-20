// Models of the Template section.

package template

import "github.com/allymailer/allymailer-go-sdk/common"

// EditorType — Тип редактора шаблону у вебінтерфейсі, "html" або "visual". За замовчуванням — "html". Створення
// шаблонів із типом редактора "visual" можливе лише через вебінтерфейс.
type EditorType string

// EditorType values.
const (
	EditorTypeHTML   EditorType = "html"
	EditorTypeVisual EditorType = "visual"
)

// TemplateSetRequest — Об'єкт з усіма властивостями шаблону.
type TemplateSetRequest struct {
	// Об'єкт, що містить властивості шаблону.
	Template *TemplateObject `json:"template"`
}

// TemplateObject — Об'єкт, що містить властивості шаблону.
type TemplateObject struct {
	// Унікальний ідентифікатор шаблону. Якщо вказано, система спробує знайти шаблон із указаним id і
	// відредагувати його. Якщо не задано, буде створено новий шаблон.
	ID *string `json:"id,omitempty"`

	// Назва шаблону.
	Name string `json:"name"`

	// Тип редактора шаблону у вебінтерфейсі, "html" або "visual". За замовчуванням — "html". Створення
	// шаблонів із типом редактора "visual" можливе лише через вебінтерфейс.
	EditorType EditorType `json:"editor_type,omitempty"`

	// Параметр для вибору [шаблонізатора](https://docs.allymailer.com.ua/en/template-engines) — або
	// "simple", або "velocity", або "liquid", або "none". За замовчуванням — "simple".
	TemplateEngine common.TemplateEngine `json:"template_engine,omitempty"`

	// Об'єкт для передачі глобальних підстановок (наприклад, назва компанії). Якщо назви змінних
	// повторюються в об'єкті користувацьких підстановок "substitutions", значення змінних будуть взяті з
	// об'єкта "substitutions". Підстановки працюють у параметрах:
	//
	// - body.html
	// - body.plaintext
	// - body.amp
	// - subject
	// - from_name
	// - options.unsubscribe_url
	// Назви змінних можуть містити літери латинського алфавіту, цифри, символ "_". Першим символом може
	// бути лише літера, використання пробілів не допускається.
	GlobalSubstitutions map[string]string `json:"global_substitutions,omitempty"`

	// Об'єкт для передачі метаданих, спільних для всіх отримувачів, у форматі: "key": "value". Максимальна
	// кількість ключів: 10; максимальна довжина ключа: 64 символи; максимальна довжина значення ключа:
	// 1024 символи. Метадані передаються під час виклику
	// [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook). Для групування листів в одну
	// розсилку може використовуватися системний ключ "campaign_id" зі значенням у вигляді невід'ємного
	// десяткового цілого до 128 біт або UUID вигляду "c7703772-453e-4c77-a9ae-9a1b94051b5a". Неправильні
	// значення будуть прирівняні до '0'
	GlobalMetadata map[string]string `json:"global_metadata,omitempty"`

	// Об'єкт, що містить у собі html-, plaintext- і amp-частини листа. Або html-, або plaintext-частина
	// має бути присутня обов'язково.
	Body *common.BodyObject `json:"body,omitempty"`

	// Тема листа.
	Subject *string `json:"subject,omitempty"`

	// Email-адреса відправника.
	FromEmail string `json:"from_email"`

	// Ім'я відправника.
	FromName *string `json:"from_name,omitempty"`

	// Необов'язкова email-адреса для відповідей (на випадок, якщо відрізняється від адреси відправника)
	ReplyTo *string `json:"reply_to,omitempty"`

	// Необов'язкове ім'я для відповідей (якщо вказано email reply_to і ви хочете, щоб відображався не лише
	// цей email, але й ім'я)
	ReplyToName *string `json:"reply_to_name,omitempty"`

	// 1 = відстеження переходів за посиланнями увімкнено (значення за замовчуванням), 0=вимкнено. Для
	// використання track_links = 0 вам потрібно [попросити
	// техпідтримку](https://cp.allymailer.com/ua/support) увімкнути таку можливість.
	// Допустимі значення: `0`, `1`.
	TrackLinks *int `json:"track_links,omitempty"`

	// 1 = відстеження прочитань листа увімкнено (значення за замовчуванням), 0 = вимкнено. Для
	// використання track_read = 0 вам потрібно [попросити
	// техпідтримку](https://cp.allymailer.com/ua/support) увімкнути таку можливість.
	// Допустимі значення: `0`, `1`.
	TrackRead *int `json:"track_read,omitempty"`

	// Об'єкт, що містить додаткові заголовки листа, максимум 50 заголовків. Префікс “X-” в імені заголовка
	// обов'язковий, заголовки без нього будуть проігноровані, наприклад X-ALLYMAILER-Global-Language,
	// X-ALLYMAILER-Template-Engine. Стандартні заголовки "To", "CC", "BCC" передаються без префікса "X-",
	// але обробляються спеціальним чином, тому мають низку обмежень, докладніше про це можна прочитати
	// [тут](/cc-and-bcc). Також якщо [техпідтримка](https://cp.allymailer.com/ua/support) схвалила вам
	// вимкнення посилання відписки, то ви також можете передавати заголовки List-Unsubscribe,
	// List-Subscribe, List-Help, List-Owner, List-Archive, In-Reply-To і References.
	Headers *common.HeadersObject `json:"headers,omitempty"`

	// Необов'язковий масив вкладень.
	Attachments []common.AttachmentsArrayItem `json:"attachments,omitempty"`

	// Масив для передачі inline-вкладень, наприклад, для включення зображень у лист замість завантаження
	// їх за посиланням.
	InlineAttachments []common.InlineAttachmentsArrayItem `json:"inline_attachments,omitempty"`
}

// TemplateSetResponse
type TemplateSetResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Об'єкт, що містить властивості шаблону.
	Template *TemplateObject `json:"template,omitempty"`
}

// TemplateGetRequest
type TemplateGetRequest struct {
	// Ідентифікатор шаблону.
	ID string `json:"id"`
}

// TemplateGetResponse
type TemplateGetResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Об'єкт, що містить властивості шаблону.
	Template *TemplateObject `json:"template,omitempty"`
}

// TemplateListRequest
type TemplateListRequest struct {
	// Максимальна кількість шаблонів, що повертаються, за замовчуванням 50.
	// За замовчуванням: `50`.
	Limit *int `json:"limit,omitempty"`

	// Індекс першого шаблону зі списку, що повертається, рахуючи з нуля.
	// За замовчуванням: `0`.
	Offset *int `json:"offset,omitempty"`
}

// TemplateListResponse
type TemplateListResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Масив об'єктів-шаблонів.
	Templates []TemplateListObject `json:"templates,omitempty"`
}

// TemplateListObject
type TemplateListObject struct {
	// Унікальний ідентифікатор шаблону.
	ID string `json:"id,omitempty"`

	// Назва шаблону.
	Name string `json:"name,omitempty"`

	// Тип редактора шаблону у вебінтерфейсі, "html" або "visual". За замовчуванням — "html". Створення
	// шаблонів із типом редактора "visual" можливе лише через вебінтерфейс.
	EditorType EditorType `json:"editor_type,omitempty"`

	// Параметр для вибору [шаблонізатора](https://docs.allymailer.com.ua/en/template-engines) — або
	// "simple", або "velocity", або "liquid", або "none". За замовчуванням — "simple".
	TemplateEngine common.TemplateEngine `json:"template_engine,omitempty"`

	// Об'єкт для передачі глобальних підстановок (наприклад, назва компанії). Якщо назви змінних
	// повторюються в об'єкті користувацьких підстановок "substitutions", значення змінних будуть взяті з
	// об'єкта "substitutions". Підстановки працюють у параметрах:
	//
	// - body.html
	// - body.plaintext
	// - body.amp
	// - subject
	// - from_name
	// - options.unsubscribe_url
	// Назви змінних можуть містити літери латинського алфавіту, цифри, символ "_". Першим символом може
	// бути лише літера, використання пробілів не допускається.
	GlobalSubstitutions map[string]string `json:"global_substitutions,omitempty"`

	// Об'єкт для передачі метаданих, спільних для всіх отримувачів, у форматі: "key": "value". Максимальна
	// кількість ключів: 10; максимальна довжина ключа: 64 символи; максимальна довжина значення ключа:
	// 1024 символи. Метадані передаються під час виклику
	// [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook). Для групування листів в одну
	// розсилку може використовуватися системний ключ "campaign_id" зі значенням у вигляді невід'ємного
	// десяткового цілого до 128 біт або UUID вигляду "c7703772-453e-4c77-a9ae-9a1b94051b5a". Неправильні
	// значення будуть прирівняні до '0'
	GlobalMetadata map[string]string `json:"global_metadata,omitempty"`

	// Об'єкт, що містить у собі html-, plaintext- і amp-частини листа. Або html-, або plaintext-частина
	// має бути присутня обов'язково.
	Body *common.BodyObject `json:"body,omitempty"`

	// Тема листа.
	Subject string `json:"subject,omitempty"`

	// Email-адреса відправника.
	FromEmail string `json:"from_email,omitempty"`

	// Ім'я відправника.
	FromName string `json:"from_name,omitempty"`

	// Необов'язкова email-адреса для відповідей (на випадок, якщо відрізняється від адреси відправника)
	ReplyTo string `json:"reply_to,omitempty"`

	// Необов'язкове ім'я для відповідей (якщо вказано email reply_to і ви хочете, щоб відображався не лише
	// цей email, але й ім'я)
	ReplyToName string `json:"reply_to_name,omitempty"`

	// 1 = відстеження переходів за посиланнями увімкнено (значення за замовчуванням), 0=вимкнено. Для
	// використання track_links = 0 вам потрібно [попросити
	// техпідтримку](https://cp.allymailer.com/ua/support) увімкнути таку можливість.
	// Допустимі значення: `0`, `1`.
	TrackLinks int `json:"track_links,omitempty"`

	// 1 = відстеження прочитань листа увімкнено (значення за замовчуванням), 0 = вимкнено. Для
	// використання track_read = 0 вам потрібно [попросити
	// техпідтримку](https://cp.allymailer.com/ua/support) увімкнути таку можливість.
	// Допустимі значення: `0`, `1`.
	TrackRead int `json:"track_read,omitempty"`

	// Об'єкт, що містить додаткові заголовки листа, максимум 50 заголовків. Префікс “X-” в імені заголовка
	// обов'язковий, заголовки без нього будуть проігноровані, наприклад X-ALLYMAILER-Global-Language,
	// X-ALLYMAILER-Template-Engine. Стандартні заголовки "To", "CC", "BCC" передаються без префікса "X-",
	// але обробляються спеціальним чином, тому мають низку обмежень, докладніше про це можна прочитати
	// [тут](/cc-and-bcc). Також якщо [техпідтримка](https://cp.allymailer.com/ua/support) схвалила вам
	// вимкнення посилання відписки, то ви також можете передавати заголовки List-Unsubscribe,
	// List-Subscribe, List-Help, List-Owner, List-Archive, In-Reply-To і References.
	Headers *common.HeadersObject `json:"headers,omitempty"`

	// Необов'язковий масив вкладень.
	Attachments []common.AttachmentsArrayItem `json:"attachments,omitempty"`

	// Масив для передачі inline-вкладень, наприклад, для включення зображень у лист замість завантаження
	// їх за посиланням.
	InlineAttachments []common.InlineAttachmentsArrayItem `json:"inline_attachments,omitempty"`

	// Час створення шаблону в часовому поясі UTC у форматі "РРРР-ММ-ДД гг:хх:сс"
	Created string `json:"created,omitempty"`

	// Унікальний ідентифікатор користувача.
	UserID int `json:"user_id,omitempty"`

	// Унікальний ідентифікатор проєкту, ASCII-рядок до 36 символів завдовжки.
	ProjectID string `json:"project_id,omitempty"`

	// Назва проєкту, унікальна в межах облікового запису користувача.
	ProjectName string `json:"project_name,omitempty"`
}

// TemplateDeleteRequest
type TemplateDeleteRequest struct {
	// Ідентифікатор шаблону.
	ID string `json:"id"`
}
