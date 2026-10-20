// Models of the Webhook section.

package webhook

import "github.com/allymailer/allymailer-go-sdk/common"

// WebhookStatus — Статус webhook, за замовчуванням "active". Статус "disabled" означає, що webhook було вимкнено
// користувачем, "stopped" означає, що webhook було зупинено системою після 24 годин невдалих викликів
// (мінімум за 10 різними подіями).
type WebhookStatus string

// WebhookStatus values.
const (
	WebhookStatusActive   WebhookStatus = "active"
	WebhookStatusDisabled WebhookStatus = "disabled"
	WebhookStatusStopped  WebhookStatus = "stopped"
)

// WebhookEventFormat — Формат сповіщень. "json_post"(за замовчуванням) або "json_post_gzip".
type WebhookEventFormat string

// WebhookEventFormat values.
const (
	WebhookEventFormatJSONPost     WebhookEventFormat = "json_post"
	WebhookEventFormatJSONPostGzip WebhookEventFormat = "json_post_gzip"
)

// EventName — Тип даних, що містяться в об'єкті event_data. Або "transactional_email_status", або
// "transactional_spam_block".
type EventName string

// EventName values.
const (
	EventNameTransactionalEmailStatus EventName = "transactional_email_status"
	EventNameTransactionalSpamBlock   EventName = "transactional_spam_block"
)

// DeviceType — Тип пристрою отримувача, визначений за user agent. Присутній для статусів "clicked" і "opened";
// дорівнює "unknown", якщо тип пристрою визначити не вдалося.
type DeviceType string

// DeviceType values.
const (
	DeviceTypeDesktop DeviceType = "desktop"
	DeviceTypeMobile  DeviceType = "mobile"
	DeviceTypeTablet  DeviceType = "tablet"
	DeviceTypeOther   DeviceType = "other"
	DeviceTypeUnknown DeviceType = "unknown"
)

// TrafficType — Якість трафіку для статусів "opened" і "clicked". "human" — розпізнано звичайний браузер або
// поштовий клієнт без машинних ознак; "machine" — бот, сканер посилань або проксі безпеки; "protected"
// — механізм приватності скриньки/ОС, участь людини підтвердити не можна (Apple Mail Privacy
// Protection / iCloud Private Relay, Gmail Image Proxy; Apple MPP — це protected, не machine);
// "unknown" — User-Agent відсутній, пошкоджений або не розпізнаний.
type TrafficType string

// TrafficType values.
const (
	TrafficTypeHuman     TrafficType = "human"
	TrafficTypeMachine   TrafficType = "machine"
	TrafficTypeProtected TrafficType = "protected"
	TrafficTypeUnknown   TrafficType = "unknown"
)

// BlockType — Тип спам-блока — заблоковано або один відправляючий SMTP, або одразу багато. У разі блокування
// одиночного SMTP із загального пулу AllyMailer пробуватиме ще декілька. Властивість присутня лише
// якщо event_name="transactional_spam_block".
type BlockType string

// BlockType values.
const (
	BlockTypeOneSMTP BlockType = "one_smtp"
	BlockTypeAllSMTP BlockType = "all_smtp"
)

// DomainStatus — Це подія блокування чи розблокування. Властивість присутня лише якщо
// event_name="transactional_spam_block".
type DomainStatus string

// DomainStatus values.
const (
	DomainStatusBlocked   DomainStatus = "blocked"
	DomainStatusUnblocked DomainStatus = "unblocked"
)

// WebhookSetRequest
type WebhookSetRequest struct {
	// Рядок з URL, на який надсилатиметься запит у разі виникнення події (ідентифікатор обробника). В URL
	// поки що підтримуються лише ASCII-символи. Якщо вам потрібно використати не-ASCII, сконвертуйте URL у
	// [Punycode](https://en.wikipedia.org/wiki/Punycode).
	URL string `json:"url"`

	// Статус webhook, за замовчуванням "active". Статус "disabled" означає, що webhook було вимкнено
	// користувачем, "stopped" означає, що webhook було зупинено системою після 24 годин невдалих викликів
	// (мінімум за 10 різними подіями).
	// За замовчуванням: `active`.
	Status WebhookStatus `json:"status,omitempty"`

	// Формат сповіщень. "json_post"(за замовчуванням) або "json_post_gzip".
	// За замовчуванням: `json_post`.
	EventFormat WebhookEventFormat `json:"event_format,omitempty"`

	// Повертати (1) чи ні (0) докладну інформацію про доставку. Якщо delivery_info=1, повертається
	// додаткова інформація про відповідь сервера та внутрішній статус доставки в разі статусів доставки
	// "hard_bounced" і "soft_bounced"; user agent, IP-адреса отримувача, геодані, визначена за user agent
	// інформація про пристрій (тип і марка пристрою, назва ОС, браузера або поштового клієнта) і
	// traffic_type (human, machine, protected або unknown) в разі статусів "opened" і "clicked"; URL для
	// статусу "clicked". Для статусів "delivered", "soft_bounced" і "hard_bounced" AllyMailer також може
	// повернути sender_ip — IP-адресу SMTP-сервера, з якого було надіслано лист, якщо він доступний.
	// Допустимі значення: `0`, `1`.
	// За замовчуванням: `0`.
	DeliveryInfo *int `json:"delivery_info,omitempty"`

	// 0=в одному виклику може надійти декілька подій, 1=лише одна подія на виклик (не рекомендується).
	// Допустимі значення: `0`, `1`.
	// За замовчуванням: `0`.
	SingleEvent *int `json:"single_event,omitempty"`

	// Максимальна кількість дозволених паралельних запитів до вашого сервера. Чим більше ваш сервер може
	// обробити — тим краще.
	// За замовчуванням: `10`.
	MaxParallel *int `json:"max_parallel,omitempty"`

	// Об'єкт, що описує події, про які потрібно сповіщати.
	Events *WebhookEventsObject `json:"events,omitempty"`
}

// WebhookEventsObject — Об'єкт, що описує події, про які потрібно сповіщати.
type WebhookEventsObject struct {
	// Якщо присутній, то надсилатимуться сповіщення про спам-блокування. У масиві має бути єдиний елемент
	// із рядком "*".
	SpamBlock []string `json:"spam_block,omitempty"`

	// Якщо присутній, то надсилатимуться сповіщення про зміну статусу email. У масиві перелічуються ті
	// статуси, про які надсилатимуться сповіщення.
	EmailStatus []common.EmailStatus `json:"email_status,omitempty"`
}

// WebhookSetResponse
type WebhookSetResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Об'єкт з усіма властивостями webhook.
	Object *WebhookObject `json:"object,omitempty"`
}

// WebhookObject — Об'єкт з усіма властивостями webhook.
type WebhookObject struct {
	// Унікальний ідентифікатор webhook.
	ID int `json:"id,omitempty"`

	// URL webhook.
	URL string `json:"url,omitempty"`

	// Статус webhook, за замовчуванням "active". Статус "disabled" означає, що webhook було вимкнено
	// користувачем, "stopped" означає, що webhook було зупинено системою після 24 годин невдалих викликів
	// (мінімум за 10 різними подіями).
	// За замовчуванням: `active`.
	Status WebhookStatus `json:"status,omitempty"`

	// Формат сповіщень. "json_post"(за замовчуванням) або "json_post_gzip".
	// За замовчуванням: `json_post`.
	EventFormat WebhookEventFormat `json:"event_format,omitempty"`

	// Повертати (1) чи ні (0) докладну інформацію про доставку. Якщо delivery_info=1, повертається
	// додаткова інформація про відповідь сервера та внутрішній статус доставки в разі статусів доставки
	// "hard_bounced" і "soft_bounced"; user agent, IP-адреса отримувача, геодані, визначена за user agent
	// інформація про пристрій (тип і марка пристрою, назва ОС, браузера або поштового клієнта) і
	// traffic_type (human, machine, protected або unknown) в разі статусів "opened" і "clicked"; URL для
	// статусу "clicked". Для статусів "delivered", "soft_bounced" і "hard_bounced" AllyMailer також може
	// повернути sender_ip — IP-адресу SMTP-сервера, з якого було надіслано лист, якщо він доступний.
	// Допустимі значення: `0`, `1`.
	// За замовчуванням: `0`.
	DeliveryInfo int `json:"delivery_info,omitempty"`

	// 0=в одному виклику може надійти декілька подій, 1=лише одна подія на виклик (не рекомендується).
	// Допустимі значення: `0`, `1`.
	// За замовчуванням: `0`.
	SingleEvent int `json:"single_event,omitempty"`

	// Максимальна кількість дозволених паралельних запитів до вашого сервера. Чим більше ваш сервер може
	// обробити — тим краще.
	// За замовчуванням: `10`.
	MaxParallel int `json:"max_parallel,omitempty"`

	// Час останнього оновлення властивостей webhook у часовому поясі UTC у форматі "РРРР-ММ-ДД гг:хх:сс".
	UpdatedAt string `json:"updated_at,omitempty"`

	// Об'єкт, що описує події, про які потрібно сповіщати.
	Events *WebhookEventsObject `json:"events,omitempty"`
}

// WebhookCallback — Опис формату webhook.
type WebhookCallback struct {
	// MD5-хеш рядкового тіла повідомлення, у якому значення auth замінено на api_key користувача/проєкту,
	// чий обробник викликається. За допомогою цього отримувач сповіщення може як провести автентифікацію,
	// так і перевірити цілісність сповіщення.
	Auth string `json:"auth,omitempty"`

	// Масив з одним елементом, що містить дані про події користувача/проєкту.
	EventsByUser []WebhookCallbackEventsByUserItem `json:"events_by_user,omitempty"`
}

// WebhookCallbackEventsByUserItem
type WebhookCallbackEventsByUserItem struct {
	// Унікальний ідентифікатор користувача.
	UserID int `json:"user_id,omitempty"`

	// Ідентифікатор проєкту, присутній, якщо webhook було зареєстровано за допомогою API-ключа проєкту.
	ProjectID string `json:"project_id,omitempty"`

	// Назва проєкту, присутня, якщо webhook було зареєстровано за допомогою API-ключа проєкту.
	ProjectName string `json:"project_name,omitempty"`

	// Масив подій, про які сповіщає виклик webhook.
	Events []WebhookCallbackEventsByUserItemEventsItem `json:"events,omitempty"`
}

// WebhookCallbackEventsByUserItemEventsItem
type WebhookCallbackEventsByUserItemEventsItem struct {
	// Тип даних, що містяться в об'єкті event_data. Або "transactional_email_status", або
	// "transactional_spam_block".
	EventName EventName `json:"event_name,omitempty"`

	// Об'єкт з різними властивостями події, що залежать від поля "event_name". Нижче перелічено всі
	// можливі властивості, спочатку для події "transactional_email_status", потім для події
	// "transactional_spam_block".
	EventData *WebhookCallbackEventsByUserItemEventsItemEventData `json:"event_data,omitempty"`
}

// WebhookCallbackEventsByUserItemEventsItemEventData — Об'єкт з різними властивостями події, що залежать від поля "event_name". Нижче перелічено всі
// можливі властивості, спочатку для події "transactional_email_status", потім для події
// "transactional_spam_block".
type WebhookCallbackEventsByUserItemEventsItemEventData struct {
	// Ідентифікатор завдання, раніше повернутий методом
	// [email/send](https://docs.allymailer.com.ua/ua/web-api-ref#email-send). Присутній лише якщо
	// event_name="transactional_email_status".
	JobID string `json:"job_id,omitempty"`

	// Метадані, раніше передані методом
	// [email/send](https://docs.allymailer.com.ua/ua/web-api-ref#email-send) у полі recipients.metadata
	// або global_metadata. Присутні лише якщо event_name="transactional_email_status".
	Metadata map[string]string `json:"metadata,omitempty"`

	// Email, на який було здійснено надсилання. Присутній лише якщо
	// event_name="transactional_email_status".
	Email string `json:"email,omitempty"`

	// Статус доставки листа. Присутній лише якщо event_name="transactional_email_status". Набуває одного з
	// таких значень:
	//
	// __accepted__ — повідомлення прийнято, але ще не надіслано.
	//
	// __sent__ — повідомлення надіслано, але ще не доставлено.
	//
	// __delivered__ — повідомлення доставлено. Може змінитися на “opened”, “clicked”, “unsubscribed”,
	// “subscribed” або “spam”.
	//
	// __opened__ — повідомлення доставлено й зареєстровано його прочитання. Може змінитися на “clicked”,
	// “unsubscribed”, “subscribed” або “spam”.
	//
	// __clicked__ — повідомлення доставлено, прочитано, було зареєстровано перехід за одним з посилань у
	// листі. Може змінитися на “unsubscribed”, “subscribed” або “spam”.
	//
	// __unsubscribed__ — повідомлення було доставлено отримувачу й прочитано ним, але користувач
	// відписався за посиланням у листі. Може змінитися на “subscribed”.
	//
	// __subscribed__ — повідомлення було доставлено отримувачу й прочитано ним, але користувач відписався,
	// а потім знову підписався за посиланням у листі. Може змінитися на “unsubscribed”.
	//
	// __soft_bounced__ — тимчасова недоставка повідомлення. AllyMailer продовжує спроби доставки протягом
	// 2 діб. У разі успішної доставки лист перейде у статус “delivered”, у разі неуспішної — у
	// “hard_bounced”.
	//
	// __hard_bounced__ — доставка не вдалася, спроб доставки більше не буде. Статус остаточний. Можливих
	// причин невдалої доставки безліч, для їх аналізу можна використати текст відповіді SMTP-сервера з
	// поля delivery_info.destination_response або нашу внутрішню класифікацію помилок з поля
	// delivery_info.delivery_status.
	//
	// __spam__ — повідомлення доставлено, але позначено як “спам” отримувачем. Статус остаточний. Ось
	// деякі домени, від яких AllyMailer може отримати й обробити скаргу про спам (використовуючи
	// технологію FBL): mail.ru, msn.com, outlook.com, hotmail.com, live.com, ukr.net, yahoo.com, aol.com.
	Status common.EmailStatus `json:"status,omitempty"`

	// Дата й час події в часовому поясі UTC у форматі "РРРР-ММ-ДД гг:хх:сс". Присутній лише якщо
	// event_name="transactional_email_status".
	EventTime string `json:"event_time,omitempty"`

	// URL у разі статусів "opened" і "clicked". Присутній лише якщо
	// event_name="transactional_email_status".
	URL string `json:"url,omitempty"`

	// Об'єкт з додатковою інформацією про доставку. Присутній лише якщо властивість delivery_info під час
	// встановлення webhook дорівнювала 1 і якщо event_name="transactional_email_status".
	DeliveryInfo *WebhookCallbackEventsByUserItemEventsItemEventDataDeliveryInfo `json:"delivery_info,omitempty"`

	// Дата й час спам-блока в часовому поясі UTC у форматі "РРРР-ММ-ДД гг:хх:сс". Присутній лише якщо
	// event_name="transactional_spam_block".
	BlockTime string `json:"block_time,omitempty"`

	// Тип спам-блока — заблоковано або один відправляючий SMTP, або одразу багато. У разі блокування
	// одиночного SMTP із загального пулу AllyMailer пробуватиме ще декілька. Властивість присутня лише
	// якщо event_name="transactional_spam_block".
	BlockType BlockType `json:"block_type,omitempty"`

	// Домен, що заблокував надсилання. Властивість присутня лише якщо
	// event_name="transactional_spam_block".
	Domain string `json:"domain,omitempty"`

	// Кількість заблокованих вихідних SMTP. Властивість присутня лише якщо
	// event_name="transactional_spam_block".
	SMTPBlocksCount int `json:"SMTP_blocks_count,omitempty"`

	// Це подія блокування чи розблокування. Властивість присутня лише якщо
	// event_name="transactional_spam_block".
	DomainStatus DomainStatus `json:"domain_status,omitempty"`
}

// WebhookCallbackEventsByUserItemEventsItemEventDataDeliveryInfo — Об'єкт з додатковою інформацією про доставку. Присутній лише якщо властивість delivery_info під час
// встановлення webhook дорівнювала 1 і якщо event_name="transactional_email_status".
type WebhookCallbackEventsByUserItemEventsItemEventDataDeliveryInfo struct {
	// Внутрішній детальний статус доставки AllyMailer. Ось деякі з можливих статусів:
	//
	// __err_user_unknown__ -- адреса не існує (і невідомо, чи існувала колись);
	//
	// __err_user_inactive__ -- адреса не використовується;
	//
	// __err_will_retry__ -- лист було тимчасово відхилено, пізніше буде здійснено повторну спробу
	// доставки;
	//
	// __err_mailbox_discarded__ -- адресу видалено (тобто колись існувала, але тепер її немає);
	//
	// __err_mailbox_full__ -- поштова скринька переповнена;
	//
	// __err_spam_rejected__ -- лист відхилено як спам;
	//
	// __err_blacklisted__ -- лист було відхилено через наявність IP-адреси або домену відправника в
	// якомусь із чорних списків;
	//
	// __err_too_large__ -- на думку сервера отримувача, лист перевищує допустимий розмір;
	//
	// __err_unsubscribed__ -- адресат раніше відписався від ваших розсилок;
	//
	// __err_unreachable__ -- адресу позначено як постійно недоступну через багаторазові помилки доставки
	// на цю адресу;
	//
	// __err_skip_letter__ -- надсилання скасовано, оскільки email-адреса тимчасово недоступна;
	//
	// __err_domain_inactive__ -- домен не приймає пошту або не існує;
	//
	// __err_destination_misconfigured__ -- домен не приймає пошту через неправильне налаштування на боці
	// отримувача, причому відповідь сервера містить інформацію про усувну причину (наприклад,
	// використовується непрацюючий блеклист тощо);
	//
	// __err_delivery_failed__ -- доставка не вдалася з інших причин;
	//
	// __err_spam_skipped__ -- надсилання скасовано через блокування розсилки як спаму;
	//
	// __err_lost__ -- лист не було надіслано через неузгодженість його частин, або його було втрачено
	// через збій на нашому боці. Відправник має повторно надіслати лист самостійно, оскільки оригінал не
	// зберігся;
	//
	// Ми залишаємо за собою право змінювати й додавати внутрішні статуси, тому решту
	// використовуваних наразі статусів перелічимо лише довідково, без пояснень:
	// ok_accepted, ok_sent, ok_delivered, ok_read, ok_link_visited, ok_unsubscribed, ok_resubscribed,
	// ok_spam_folder, ok_fbl,
	// not_sent, skip_dup_unreachable, skip_dup_temp_unreachable, skip_dup_mailbox_full,
	// err_spam_removed, err_resend, err_unknown, err_retry_letter, err_src_invalid,
	// err_dest_invalid, err_not_available, err_internal, err_no_dns, err_no_smtp, err_giveup.
	DeliveryStatus string `json:"delivery_status,omitempty"`

	// Відповідь SMTP-сервера.
	DestinationResponse string `json:"destination_response,omitempty"`

	// Рядок user agent отримувача. Присутній лише якщо вдалося його визначити, для статусів "clicked" і
	// "opened".
	UserAgent string `json:"user_agent,omitempty"`

	// Тип пристрою отримувача, визначений за user agent. Присутній для статусів "clicked" і "opened";
	// дорівнює "unknown", якщо тип пристрою визначити не вдалося.
	DeviceType DeviceType `json:"device_type,omitempty"`

	// Марка пристрою отримувача, визначена за user agent. Присутня для статусів "clicked" і "opened", якщо
	// її вдалося визначити. Для device_type "mobile" і "tablet" присутня завжди та дорівнює "unknown",
	// якщо марку визначити не вдалося.
	DeviceBrand string `json:"device_brand,omitempty"`

	// Назва операційної системи пристрою отримувача, визначена за user agent. Присутня для статусів
	// "clicked" і "opened"; дорівнює "unknown", якщо ОС визначити не вдалося.
	OsName string `json:"os_name,omitempty"`

	// Назва браузера, у якому було відкрито лист, визначена за user agent. Присутня для статусів "clicked"
	// і "opened", лише якщо було визначено браузер, тому ніколи не надходить разом з email_client_name.
	BrowserName string `json:"browser_name,omitempty"`

	// Назва поштового клієнта, у якому було відкрито лист, визначена за user agent. Присутня для статусів
	// "clicked" і "opened", лише якщо було визначено поштовий клієнт, тому ніколи не надходить разом з
	// browser_name.
	EmailClientName string `json:"email_client_name,omitempty"`

	// Якість трафіку для статусів "opened" і "clicked". "human" — розпізнано звичайний браузер або
	// поштовий клієнт без машинних ознак; "machine" — бот, сканер посилань або проксі безпеки; "protected"
	// — механізм приватності скриньки/ОС, участь людини підтвердити не можна (Apple Mail Privacy
	// Protection / iCloud Private Relay, Gmail Image Proxy; Apple MPP — це protected, не machine);
	// "unknown" — User-Agent відсутній, пошкоджений або не розпізнаний.
	TrafficType TrafficType `json:"traffic_type,omitempty"`

	// IP-адреса отримувача. Присутня лише якщо вдалося її визначити, для статусів "clicked" і "opened".
	IP string `json:"ip,omitempty"`

	// ISO 3166-1 alpha-2 код країни, визначений за IP-адресою отримувача. Присутній для статусів "clicked"
	// і "opened".
	CountryCode string `json:"country_code,omitempty"`

	// Країна, визначена за IP-адресою отримувача. Присутня для статусів "clicked" і "opened".
	Country string `json:"country,omitempty"`

	// Місто, визначене за IP-адресою отримувача. Присутнє для статусів "clicked" і "opened".
	City string `json:"city,omitempty"`

	// IP-адреса SMTP-сервера, з якого було надіслано лист. Присутня лише для статусів "delivered",
	// "soft_bounced" і "hard_bounced", якщо доступна.
	SenderIP string `json:"sender_ip,omitempty"`
}

// WebhookGetRequest
type WebhookGetRequest struct {
	// URL webhook.
	URL string `json:"url"`
}

// WebhookGetResponse
type WebhookGetResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Об'єкт з усіма властивостями webhook.
	Object *WebhookObject `json:"object,omitempty"`
}

// WebhookListRequest
type WebhookListRequest struct {
	// Максимальна кількість webhook, що повертаються за один виклик. Рекомендується 50.
	Limit *int `json:"limit,omitempty"`

	// Індекс першого webhook зі списку, що повертається, рахуючи з нуля.
	// За замовчуванням: `0`.
	Offset *int `json:"offset,omitempty"`
}

// WebhookListResponse
type WebhookListResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Масив об'єктів з даними про webhook.
	Objects []WebhookObject `json:"objects,omitempty"`
}

// WebhookDeleteRequest
type WebhookDeleteRequest struct {
	// URL webhook.
	URL string `json:"url"`
}
