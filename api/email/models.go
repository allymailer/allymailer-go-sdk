// Models of the Email section.

package email

import "github.com/allymailer/allymailer-go-sdk/common"

// EmailSendRequest
type EmailSendRequest struct {
	// Об'єкт, що містить усі властивості повідомлення, яке надсилається.
	Message *MessageObject `json:"message"`
}

// MessageObject — Об'єкт, що містить усі властивості повідомлення, яке надсилається.
type MessageObject struct {
	// Масив об'єктів з email-адресами, підстановками та метаданими отримувачів.
	Recipients []MessageObjectRecipientsItem `json:"recipients"`

	// Унікальний ідентифікатор шаблону, створеного раніше методом
	// [template/set](https://docs.allymailer.com.ua/ua/web-api-ref#template-set). Якщо вказано, то поля
	// шаблону підставляються замість пропущених полів email/send. Наприклад, якщо в email/send не вказано
	// body — береться тіло листа із шаблону, а якщо не вказано subject — береться тема із шаблону.
	TemplateID *string `json:"template_id,omitempty"`

	// Масив із рядків, до 100 символів у кожному. Максимум 4 рядки в масиві, рядки мають бути унікальними.
	// Для кожного облікового запису/проєкту дозволено не більше ніж 10 000 міток, у разі перевищення метод
	// email/send повертатиме помилку. Мітки можуть використовуватися для категоризації листів за обраними
	// вами критеріями. Мітки передаються методами event-dump.
	Tags []string `json:"tags,omitempty"`

	// Пропустити чи не пропускати додавання стандартного блока з посиланням відписки до HTML-частини
	// листа. 1=пропустити, 0=додати, за замовчуванням 0. Для использования skip_unsubscribe=1 вам надо
	// [попросить техподдержку](https://cp.allymailer.com/ru/support) включить такую возможность.
	// Допустимі значення: `0`, `1`.
	SkipUnsubscribe *int `json:"skip_unsubscribe,omitempty"`

	// Заголовок для вибору мови посилання та сторінки відписки. Припустимі значення "be", "de", "en",
	// "es", "fr", "it", "pl", "pt", "ru", "ua", "kz".
	GlobalLanguage *string `json:"global_language,omitempty"`

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
	Body *common.BodyObject `json:"body"`

	// Тема листа.
	Subject string `json:"subject"`

	// Email-адреса відправника. Обов'язковий лише за порожнього template_id.
	FromEmail *string `json:"from_email,omitempty"`

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

	// (0 або 1, за замовчуванням 0) Значення 1=ігнорувати список глобальної недоступності, під час
	// надсилання не враховуватимуться дані про недоступність адрес або скарги, отримані під час надсилання
	// іншими користувачами AllyMailer. Для деяких системних записів може бути заборонено
	// Допустимі значення: `0`, `1`.
	BypassGlobal *int `json:"bypass_global,omitempty"`

	// (0 або 1, за замовчуванням 0) Значення 1=ігнорувати список недоступних адрес поточного облікового
	// запису/проєкту. Указання 1 можливе лише за одночасного вказання bypass_global=1.
	// Допустимі значення: `0`, `1`.
	BypassUnavailable *int `json:"bypass_unavailable,omitempty"`

	// (0 або 1, за замовчуванням 0) Значення 1=ігнорувати список відписаних від розсилок поточного
	// облікового запису/проєкту. Указання 1 можливе лише за одночасного вказання bypass_global=1. Параметр
	// доступний лише тим користувачам, яким [техпідтримка](https://cp.allymailer.com/ua/support) схвалила
	// вимкнення посилання відписки.
	// Допустимі значення: `0`, `1`.
	BypassUnsubscribed *int `json:"bypass_unsubscribed,omitempty"`

	// (0 або 1, за замовчуванням 0) Значення 1=ігнорувати список тих, хто поскаржився на обліковий
	// запис/проєкт. Указання 1 можливе лише за одночасного вказання bypass_global=1. Параметр доступний
	// лише тим користувачам, яким [техпідтримка](https://cp.allymailer.com/ua/support) схвалила вимкнення
	// посилання відписки.
	// Допустимі значення: `0`, `1`.
	BypassComplained *int `json:"bypass_complained,omitempty"`

	// Рядок до 64 символів завдовжки з унікальним ключем запиту. Може використовуватися для запобігання
	// випадковому дублюванню листів. У разі повторного надсилання запиту з тим самим ключем протягом
	// хвилини запит буде відхилено. Може створюватися автоматично для кожного листа; щоб увімкнути
	// автоматичну генерацію ключа, зв'яжіться з [техпідтримкою](https://cp.allymailer.com/ua/support).
	IdempotenceKey *string `json:"idempotence_key,omitempty"`

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

	// Додаткові опції.
	Options *MessageObjectOptions `json:"options,omitempty"`
}

// MessageObjectRecipientsItem
type MessageObjectRecipientsItem struct {
	// Email-адреса отримувача
	Email string `json:"email"`

	// Об'єкт, що описує підстановки для конкретного отримувача (наприклад, ім'я, товар для показу саме
	// цьому отримувачу — див. [Шаблонізатори](https://docs.allymailer.com.ua/en/template-engines)).
	// Підстановки працюють у параметрах:
	//
	// - body.html
	// - body.plaintext
	// - body.amp
	// - subject
	// - from_name
	// - headers["List-Unsubscribe"]
	// - options.unsubscribe_url
	// Назви змінних можуть містити літери латинського алфавіту, цифри, символ "_". Першим символом може
	// бути лише літера, використання пробілів не допускається. Є спеціальна змінна "to_name", що дозволяє
	// задати ім'я отримувача. У разі підстановки туди "Ім'я Прізвище" SMTP-заголовок “To” виглядатиме як
	// "Ім'я Прізвище <email@example.com>". Довжина змінної "to_name" обмежена 78 символами.
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Об'єкт для передачі метаданих, специфічних для окремого отримувача, де "key" — рядок до 64 символів
	// завдовжки, "value" — рядок до 1024 символів завдовжки, максимальна кількість ключів — 10. Метадані
	// передаються під час виклику [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook) і
	// методів [event-dump](https://docs.allymailer.com.ua/ua/web-api-ref#event-dump). Для групування
	// листів в одну розсилку може використовуватися системний ключ "campaign_id" зі значенням у вигляді
	// невід'ємного десяткового цілого або UUID до 128 біт.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// MessageObjectOptions — Додаткові опції.
type MessageObjectOptions struct {
	// Дата й час у форматі "РРРР-ММ-ДД гг:хх:сс" у часовому поясі UTC. Дозволяє запланувати час надсилання
	// на майбутнє, у межах 24 годин від поточного часу.
	SendAt *string `json:"send_at,omitempty"`

	// Користувацьке посилання відписки. Докладніше [тут](https://docs.allymailer.com.ua/unsubscribe-link).
	UnsubscribeURL *string `json:"unsubscribe_url,omitempty"`

	// Ідентифікатор backend-домену, через який надсилається повідомлення. Якщо відсутній — надсилання буде
	// від імені backend-домену, призначеного вашому обліковому запису/проєкту за замовчуванням. Ви можете
	// придбати один або декілька виділених IP-адрес і використовувати цю опцію для вибору.
	CustomBackendID *int `json:"custom_backend_id,omitempty"`

	// Ідентифікатор пулу SMTP, через який надсилається повідомлення. Зазвичай немає необхідності вказувати
	// пул явно, оскільки він обирається автоматично або як пул за замовчуванням для вашого облікового
	// запису/проєкту, або відповідно до переданого параметра custom_backend_id.
	SMTPPoolID *string `json:"smtp_pool_id,omitempty"`
}

// EmailSendResponse
type EmailSendResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Ідентифікатор завдання з надсилання, може виявитися корисним під час з'ясування причин збоїв.
	JobID string `json:"job_id,omitempty"`

	// Масив email-адрес, успішно прийнятих до надсилання.
	Emails []string `json:"emails,omitempty"`

	// Об'єкт із email-адресами, не прийнятими до надсилання. Email представлено іменем властивості, а його
	// статус — значенням властивості, наприклад: {"email1@gmail.com": "temporary_unavailable"}. Можливі
	// статуси адрес, на які не вдалося здійснити надсилання:
	//
	// - __unsubscribed__ — зазначену адресу відписано;
	// - __invalid__ — адреса не існує або введена некоректно;
	// - __duplicate__ — адреса повторюється в запиті, повторне надсилання на той самий email запобігнуто;
	// - __temporary_unavailable__ — адреса тимчасово недоступна. Це означає, що протягом наступних трьох діб надсилання на цю адресу повертатиме помилку. Email може бути тимчасово недоступним з різних причин, наприклад:
	//
	// - попереднє надсилання було відхилено сервером отримувача як спам;
	// - поштова скринька отримувача переповнена або не використовується;
	// - домен не приймає пошту через неправильне налаштування на боці отримувача або з інших причин;
	// - сервер відправника було відхилено через блеклистинг;
	//
	// - __permanent_unavailable__ — адреса перманентно недоступна через багаторазові недоставки;
	// - __complained__ — в одному з попередніх листів адресат натиснув “Це спам”;
	// - __blocked__ — надсилання на цю адресу заборонено адміністрацією AllyMailer.
	// У майбутньому можуть з'явитися й інші статуси.
	FailedEmails map[string]string `json:"failed_emails,omitempty"`
}

// EmailSubscribeRequest
type EmailSubscribeRequest struct {
	// Адреса відправника на підтвердженому домені. (Також підтримується іменування поля як
	// "email_address_from").
	FromEmail string `json:"from_email"`

	// Ім'я відправника. (Також підтримується іменування поля як "name_from").
	FromName string `json:"from_name"`

	// Адреса отримувача. (Також підтримується іменування поля як "email_address_to").
	ToEmail string `json:"to_email"`
}
