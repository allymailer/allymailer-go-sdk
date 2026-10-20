// Models of the System section.

package system

import "github.com/allymailer/allymailer-go-sdk/common"

// SystemPingResponse
type SystemPingResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Унікальний ідентифікатор користувача.
	UserID int `json:"user_id,omitempty"`
}

// SystemInfoResponse
type SystemInfoResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Унікальний ідентифікатор користувача.
	UserID int `json:"user_id,omitempty"`

	// Email користувача.
	Email string `json:"email,omitempty"`

	// Унікальний ідентифікатор проєкту, ASCII-рядок до 36 символів завдовжки. Присутній лише якщо
	// API-ключ, використаний для запиту, був API-ключем проєкту.
	ProjectID string `json:"project_id,omitempty"`

	// Назва проєкту, унікальна в межах облікового запису користувача. Присутня лише якщо API-ключ,
	// використаний для запиту, був API-ключем проєкту.
	ProjectName string `json:"project_name,omitempty"`

	// Об'єкт із властивостями лічильника листів. Присутній лише якщо API-ключ, використаний для запиту,
	// був API-ключем проєкту.
	ProjectAccounting *ProjectAccountingObject `json:"project_accounting,omitempty"`

	// Об'єкт з властивостями періоду обліку. Присутній лише якщо API-ключ, використаний для запиту, був
	// API-ключем користувача.
	Accounting *SystemInfoResponseAccounting `json:"accounting,omitempty"`
}

// ProjectAccountingObject — Об'єкт із властивостями лічильника листів. Присутній лише якщо API-ключ, використаний для запиту,
// був API-ключем проєкту.
type ProjectAccountingObject struct {
	// Значення лічильника листів для проєкту, збільшується під час надсилання кожного листа.
	EmailCounter int `json:"email_counter,omitempty"`

	// Дозволяє обмежити кількість надісланих проєктом листів. У разі досягнення цього ліміту надсилання
	// припиняється. Значення 0 (за замовчуванням) означає відсутність ліміту.
	EmailCounterLimit int `json:"email_counter_limit,omitempty"`

	// 'default' - лічильник автоматично скидається в нуль на початку кожного облікового періоду,
	// 'permanent' - лічильник ніколи не скидається автоматично.
	// За замовчуванням: `default`.
	EmailCounterMode common.EmailCounterMode `json:"email_counter_mode,omitempty"`
}

// SystemInfoResponseAccounting — Об'єкт з властивостями періоду обліку. Присутній лише якщо API-ключ, використаний для запиту, був
// API-ключем користувача.
type SystemInfoResponseAccounting struct {
	// Дата й час початку періоду обліку в часовому поясі UTC у форматі "РРРР-ММ-ДД гг:хх:сс".
	PeriodStart string `json:"period_start,omitempty"`

	// Дата й час кінця періоду обліку в часовому поясі UTC у форматі "РРРР-ММ-ДД гг:хх:сс".
	PeriodEnd string `json:"period_end,omitempty"`

	// Кількість листів, включених в обліковий період.
	EmailsIncluded int `json:"emails_included,omitempty"`

	// Кількість листів, надісланих за обліковий період. Може бути більшою за emails_included у разі
	// перевищення.
	EmailsSent int `json:"emails_sent,omitempty"`
}
