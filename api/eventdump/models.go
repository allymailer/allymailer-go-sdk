// Models of the EventDump section.

package eventdump

import "github.com/allymailer/allymailer-go-sdk/common"

// EventDumpCreateRequest
type EventDumpCreateRequest struct {
	// Дата й час у форматі РРРР-ММ-ДД гг:хх:сс, починаючи з якої необхідно отримати події. Дані
	// зберігаються до 45 днів, залежно від вашого тарифу.
	StartTime string `json:"start_time"`

	// Дата й час у форматі РРРР-ММ-ДД гг:хх:сс, до якої (не включно) запитуються події.
	EndTime *string `json:"end_time,omitempty"`

	// Максимальна кількість подій, що повертаються, за замовчуванням 50. Якщо це значення перевищує 100
	// 000, буде створено декілька файлів, кожен з яких містить максимум 100 000 подій.
	Limit *int `json:"limit,omitempty"`

	// Якщо у користувача увімкнено проєкти, all_projects=true дозволить отримати дані всіх проєктів.
	// За замовчуванням: `False`.
	AllProjects *bool `json:"all_projects,omitempty"`

	// Об'єкт із властивостями фільтра дампа подій.
	Filter *FilterObject `json:"filter,omitempty"`

	// Масив із рядків з переліком назв стовпців, які необхідно отримати у вивантаженні (у порядку
	// переліку). Якщо 'dump_fields' відсутній або порожній — застосовується набір і порядок полів за
	// замовчуванням. Припустимі значення (у порядку за замовчуванням): 'project_id', 'event_time',
	// 'job_id', 'from_email', 'email', 'status', 'delivery_status', 'metadata', 'destination_response',
	// 'user_agent', 'url', 'ip', 'tags'. З метою більшої конфіденційності поле 'subject' не входить до
	// набору полів за замовчуванням, але ви можете додати його, вказавши в параметрі 'dump_fields'.
	DumpFields []string `json:"dump_fields,omitempty"`

	// Дозволяє отримати агреговану статистику за вказаний період. Може набувати одного з двох значень:
	//
	// - __day_status__ - отримати агреговану статистику; за такого значення параметр filter має бути відсутнім або порожнім об'єктом (фільтри неприпустимі);
	// - __порожній рядок__ - значення за замовчуванням, еквівалентне відсутності параметра aggregate, і метод event_dump/create у цьому разі працює як раніше.
	Aggregate *string `json:"aggregate,omitempty"`

	// Роздільник полів, за замовчуванням ','. Припустимі значення ',' і ';'.
	Delimiter *string `json:"delimiter,omitempty"`

	// Формат отримуваного файлу “csv” (за замовчуванням) або “csv_gzip”.
	// За замовчуванням: `csv`.
	Format *string `json:"format,omitempty"`
}

// FilterObject — Об'єкт із властивостями фільтра дампа подій.
type FilterObject struct {
	// Ідентифікатор завдання, раніше повернутий методом
	// [email/send](https://docs.allymailer.com.ua/ua/web-api-ref#email-send).
	JobID *string `json:"job_id,omitempty"`

	// Статус email. Можна вказати одне з таких значень:
	//
	// __accepted__ — повідомлення прийнято, але ще не надіслано.
	//
	// __sent__ — повідомлення надіслано, але ще не доставлено.
	//
	// __delivered__ — повідомлення доставлено.
	//
	// __opened__ — повідомлення доставлено й зареєстровано його прочитання.
	//
	// __clicked__ — повідомлення доставлено, прочитано, було зареєстровано перехід за одним з посилань у
	// листі.
	//
	// __unsubscribed__ — повідомлення було доставлено отримувачу й прочитано ним, але користувач
	// відписався за посиланням у листі.
	//
	// __subscribed__ — повідомлення було доставлено отримувачу й прочитано ним, але користувач відписався,
	// а потім знову підписався за посиланням у листі.
	//
	// __soft_bounced__ — тимчасова недоставка повідомлення. AllyMailer продовжує спроби доставки протягом
	// 2 діб.
	//
	// __hard_bounced__ — доставка не вдалася, спроб доставки більше не буде.
	//
	// __spam__ — повідомлення доставлено, але позначено як “спам” отримувачем.
	Status common.EmailStatus `json:"status,omitempty"`

	// Внутрішній детальний статус доставки AllyMailer. Можна вказати декілька статусів через кому. Ось
	// деякі з можливих статусів:
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
	DeliveryStatus *string `json:"delivery_status,omitempty"`

	// Адреса отримувача
	Email *string `json:"email,omitempty"`

	// Адреса відправника
	EmailFrom *string `json:"email_from,omitempty"`

	// Домен отримувача
	Domain *string `json:"domain,omitempty"`

	// Ідентифікатор розсилки, невід'ємне десяткове ціле або UUID до 128 біт, передане в метаданих у полі з
	// іменем "campaign_id" (ім'я можна налаштувати через техпідтримку).
	CampaignID *string `json:"campaign_id,omitempty"`
}

// EventDumpCreateResponse
type EventDumpCreateResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Ідентифікатор дампа. Використовується для виклику методу
	// [event-dump/get](https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-get).
	DumpID string `json:"dump_id,omitempty"`
}

// EventDumpGetRequest
type EventDumpGetRequest struct {
	// Ідентифікатор дампа, отриманий методом
	// [event-dump/create](https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-create).
	DumpID string `json:"dump_id"`
}

// EventDumpGetResponse
type EventDumpGetResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Об'єкт із властивостями дампа подій.
	EventDump *EventDumpObject `json:"event_dump,omitempty"`
}

// EventDumpObject — Об'єкт із властивостями дампа подій.
type EventDumpObject struct {
	// Ідентифікатор дампа, отриманий методом
	// [event-dump/create](https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-create).
	DumpID string `json:"dump_id,omitempty"`

	// Статус задачі. Одне з чотирьох значень:
	//
	// - __queued__ - очікує в черзі, обробка ще не починалася;
	// - __in_process__ - триває обробка;
	// - __ready__ - обробку завершено, файли доступні для завантаження;
	// - __failed__ - обробку перервано з помилкою.
	DumpStatus string `json:"dump_status,omitempty"`

	// Масив об'єктів, кожен з яких позначає готовий до завантаження файл. Навіть якщо
	// "dump_status"="in_process", уже можна завантажувати ті файли, які вказані в "files". У разі
	// відсутності подій за заданими параметрами буде повернуто порожній масив.
	Files []EventDumpFilesObject `json:"files,omitempty"`
}

// EventDumpFilesObject — Об'єкт із властивостями файлу дампа подій.
type EventDumpFilesObject struct {
	// Унікальний URL для завантаження файлу з розширенням .csv або .csv.gz, залежно від запитаного
	// формату.
	URL string `json:"url,omitempty"`

	// Розмір файлу.
	Size int `json:"size,omitempty"`
}

// EventDumpListResponse
type EventDumpListResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Масив об'єктів, що описують дампи подій. Кожен об'єкт містить ті самі властивості, що й об'єкт
	// 'event_dump', який повертає метод
	// [event-dump/get](https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-get).
	EventDumps []EventDumpObject `json:"event_dumps,omitempty"`
}

// EventDumpDeleteRequest
type EventDumpDeleteRequest struct {
	// Ідентифікатор дампа, отриманий методом
	// [event-dump/create](https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-create).
	DumpID string `json:"dump_id"`
}
