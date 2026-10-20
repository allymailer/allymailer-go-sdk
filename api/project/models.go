// Models of the Project section.

package project

import "github.com/allymailer/allymailer-go-sdk/common"

// ProjectCreateRequest
type ProjectCreateRequest struct {
	// Об'єкт із властивостями проєкту.
	Project *ProjectObject `json:"project"`
}

// ProjectObject — Об'єкт із властивостями проєкту.
type ProjectObject struct {
	// Назва проєкту, унікальна в межах облікового запису користувача.
	Name string `json:"name"`

	// Двобуквений код країни за правилами ISO-3166 alpha-2. Якщо задано, то AllyMailer обробляє
	// персональні дані проєкту відповідно до законів заданої країни — наприклад, відповідно до GDPR для
	// європейських країн.
	Country *string `json:"country,omitempty"`

	// Чи дозволено надсилання листів для проєкту.
	// За замовчуванням: `True`.
	SendEnabled *bool `json:"send_enabled,omitempty"`

	// Якщо false, то AllyMailer додає в кінець кожного листа системне посилання відписки. Значення true
	// вимикає автоматичне додавання системою посилання відписки в кінець листа й дозволяє використовувати
	// власне посилання відписки або надсилати листи зовсім без посилання відписки. Докладніше див.
	// [тут](https://docs.allymailer.com.ua/en/unsubscribe-link). Якщо ви не вказуєте значення
	// custom_unsubscribe_url_enabled під час створення проєкту — то це значення копіюється з головного
	// облікового запису (якщо в головному обліковому записі було встановлено примусове вимкнення посилання
	// відписки — то й у проєкту посилання відписки буде вимкнено).
	CustomUnsubscribeURLEnabled *bool `json:"custom_unsubscribe_url_enabled,omitempty"`

	// Унікальний ідентифікатор backend-домену, який визначає [домен посилань](/tracking-domains) або пул
	// [виділених IP-адрес](/dedicated-ip) за замовчуванням. Якщо відсутній — обліковому запису/проєкту
	// буде призначено дефолтний backend-домен за замовчуванням.
	BackendDomainID *int `json:"backend_domain_id,omitempty"`

	// Унікальний ідентифікатор, який визначає сторінку відписки за замовчуванням. Якщо відсутній —
	// обліковому запису/проєкту буде призначено системну сторінку відписки за замовчуванням.
	UnsubscribePageID *int `json:"unsubscribe_page_id,omitempty"`
}

// ProjectCreateResponse
type ProjectCreateResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Унікальний ідентифікатор проєкту, ASCII-рядок до 36 символів завдовжки.
	ProjectID string `json:"project_id,omitempty"`

	// API-ключ проєкту. Ви можете використовувати цей ключ замість користувацького api_key у всіх методах,
	// крім project-методів.
	ProjectAPIKey string `json:"project_api_key,omitempty"`
}

// ProjectUpdateRequest
type ProjectUpdateRequest struct {
	// Ідентифікатор проєкту, що оновлюється. Проєкт має бути ідентифіковано або цим полем project_id, або
	// полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше.
	ProjectID *string `json:"project_id,omitempty"`

	// Необов'язковий API-ключ проєкту, що оновлюється. Проєкт має бути ідентифіковано або полем
	// project_id, або полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше.
	ProjectAPIKey *string `json:"project_api_key,omitempty"`

	// Об'єкт із властивостями проєкту.
	Project *ProjectUpdateObject `json:"project"`
}

// ProjectUpdateObject — Об'єкт із властивостями проєкту.
type ProjectUpdateObject struct {
	// Назва проєкту, унікальна в межах облікового запису користувача.
	Name string `json:"name"`

	// Двобуквений код країни за правилами ISO-3166 alpha-2. Якщо задано, то AllyMailer обробляє
	// персональні дані проєкту відповідно до законів заданої країни — наприклад, відповідно до GDPR для
	// європейських країн.
	Country *string `json:"country,omitempty"`

	// Чи дозволено надсилання листів для проєкту.
	// За замовчуванням: `True`.
	SendEnabled *bool `json:"send_enabled,omitempty"`

	// Якщо false, то AllyMailer додає в кінець кожного листа системне посилання відписки. Значення true
	// вимикає автоматичне додавання системою посилання відписки в кінець листа й дозволяє використовувати
	// власне посилання відписки або надсилати листи зовсім без посилання відписки. Докладніше див.
	// [тут](https://docs.allymailer.com.ua/en/unsubscribe-link). Якщо ви не вказуєте значення
	// custom_unsubscribe_url_enabled під час створення проєкту — то це значення копіюється з головного
	// облікового запису (якщо в головному обліковому записі було встановлено примусове вимкнення посилання
	// відписки — то й у проєкту посилання відписки буде вимкнено).
	CustomUnsubscribeURLEnabled *bool `json:"custom_unsubscribe_url_enabled,omitempty"`

	// Унікальний ідентифікатор backend-домену, який визначає [домен посилань](/tracking-domains) або пул
	// [виділених IP-адрес](/dedicated-ip) за замовчуванням. Якщо відсутній — обліковому запису/проєкту
	// буде призначено дефолтний backend-домен за замовчуванням.
	BackendDomainID *int `json:"backend_domain_id,omitempty"`

	// Унікальний ідентифікатор, який визначає сторінку відписки за замовчуванням. Якщо відсутній —
	// обліковому запису/проєкту буде призначено системну сторінку відписки за замовчуванням.
	UnsubscribePageID *int `json:"unsubscribe_page_id,omitempty"`

	// Значення лічильника листів для проєкту, збільшується під час надсилання кожного листа. Приймає лише
	// значення 0 (тобто скидання лічильника в нуль) або оновлюється автоматично, залежно від режиму роботи
	// лічильника, заданого параметром 'email_counter_mode'. Якщо поле не передано в запиті або передано
	// null, то значення лічильника не змінюється.
	EmailCounter *int `json:"email_counter,omitempty"`

	// Дозволяє обмежити кількість надісланих проєктом листів. У разі досягнення цього ліміту надсилання
	// припиняється. Приймається будь-яке невід'ємне 64-бітне ціле число. Значення 0 означає відсутність
	// ліміту. Якщо поле не передано в запиті або передано null, то значення ліміту не змінюється.
	// За замовчуванням: `0`.
	EmailCounterLimit *int `json:"email_counter_limit,omitempty"`

	// Встановлює режим роботи лічильника листів. Дозволяє вибрати, як має поводитися лічильник:
	// автоматично скидатися чи залишатися постійним. Припустимі значення: default - лічильник автоматично
	// скидається в нуль на початку кожного облікового періоду, permanent - лічильник ніколи не скидається
	// автоматично. Скидання може бути виконане лише вручну користувачем, передавши значення 0 у полі
	// 'email_counter'. Якщо поле не передано в запиті або передано значення null, значення режиму не
	// змінюється.
	// За замовчуванням: `default`.
	EmailCounterMode common.EmailCounterMode `json:"email_counter_mode,omitempty"`
}

// ProjectUpdateResponse
type ProjectUpdateResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// API-ключ проєкту. Повертається лише в разі, якщо в параметрі запиту був project_api_key. Якщо ж
	// проєкт було ідентифіковано через project_id, то не повертається.
	ProjectAPIKey string `json:"project_api_key,omitempty"`
}

// ProjectListRequest
type ProjectListRequest struct {
	// Опціональна можливість указати project_id для включення до списку лише вказаного проєкту.
	ProjectID *string `json:"project_id,omitempty"`

	// Опціональна можливість указати project_api_key для включення до списку лише вказаного проєкту.
	ProjectAPIKey *string `json:"project_api_key,omitempty"`
}

// ProjectListResponse
type ProjectListResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Масив об'єктів, де кожен об'єкт відповідає окремому проєкту.
	Projects []ProjectListObject `json:"projects,omitempty"`
}

// ProjectListObject
type ProjectListObject struct {
	// Унікальний ідентифікатор проєкту, ASCII-рядок до 36 символів завдовжки.
	ID string `json:"id,omitempty"`

	// API-ключ проєкту. Ви можете використовувати цей ключ замість користувацького api_key у всіх методах,
	// крім project-методів.
	APIKey string `json:"api_key,omitempty"`

	// Назва проєкту, унікальна в межах облікового запису користувача.
	Name string `json:"name,omitempty"`

	// Двобуквений код країни за правилами ISO-3166 alpha-2. Якщо задано, то AllyMailer обробляє
	// персональні дані проєкту відповідно до законів заданої країни — наприклад, відповідно до GDPR для
	// європейських країн.
	Country string `json:"country,omitempty"`

	// Дата й час створення проєкту в часовому поясі UTC у форматі "РРРР-ММ-ДД гг:хх:сс".
	RegTime string `json:"reg_time,omitempty"`

	// Чи дозволено надсилання листів для проєкту.
	// За замовчуванням: `True`.
	SendEnabled bool `json:"send_enabled,omitempty"`

	// Якщо false, то AllyMailer додає в кінець кожного листа системне посилання відписки. Значення true
	// вимикає автоматичне додавання системою посилання відписки в кінець листа й дозволяє використовувати
	// власне посилання відписки або надсилати листи зовсім без посилання відписки. Докладніше див.
	// [тут](https://docs.allymailer.com.ua/en/unsubscribe-link). Якщо ви не вказуєте значення
	// custom_unsubscribe_url_enabled під час створення проєкту — то це значення копіюється з головного
	// облікового запису (якщо в головному обліковому записі було встановлено примусове вимкнення посилання
	// відписки — то й у проєкту посилання відписки буде вимкнено).
	CustomUnsubscribeURLEnabled bool `json:"custom_unsubscribe_url_enabled,omitempty"`

	// Унікальний ідентифікатор backend-домену, який визначає [домен посилань](/tracking-domains) або пул
	// [виділених IP-адрес](/dedicated-ip) за замовчуванням. Якщо відсутній — обліковому запису/проєкту
	// буде призначено дефолтний backend-домен за замовчуванням.
	BackendDomainID int `json:"backend_domain_id,omitempty"`

	// Значення лічильника листів для проєкту, збільшується під час надсилання кожного листа.
	EmailCounter int `json:"email_counter,omitempty"`

	// Дозволяє обмежити кількість надісланих проєктом листів. У разі досягнення цього ліміту надсилання
	// припиняється. Значення 0 (за замовчуванням) означає відсутність ліміту.
	EmailCounterLimit int `json:"email_counter_limit,omitempty"`

	// 'default' - лічильник автоматично скидається в нуль на початку кожного облікового періоду,
	// 'permanent' - лічильник ніколи не скидається автоматично.
	// За замовчуванням: `default`.
	EmailCounterMode common.EmailCounterMode `json:"email_counter_mode,omitempty"`

	// Унікальний ідентифікатор, який визначає сторінку відписки за замовчуванням. Якщо відсутній —
	// обліковому запису/проєкту буде призначено системну сторінку відписки за замовчуванням.
	UnsubscribePageID int `json:"unsubscribe_page_id,omitempty"`
}

// ProjectDeleteRequest
type ProjectDeleteRequest struct {
	// Ідентифікатор проєкту, що видаляється. Проєкт має бути ідентифіковано або цим полем project_id, або
	// полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше.
	ProjectID *string `json:"project_id,omitempty"`

	// Необов'язковий API-ключ проєкту, що видаляється. Проєкт має бути ідентифіковано або полем
	// project_id, або полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше.
	ProjectAPIKey *string `json:"project_api_key,omitempty"`
}
