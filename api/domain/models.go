// Models of the Domain section.

package domain

import "github.com/allymailer/allymailer-go-sdk/common"

// DkimStatus — Лише домени зі статусом DKIM-запису "active" можуть використовуватися для надсилання листів.
type DkimStatus string

// DkimStatus values.
const (
	DkimStatusActive   DkimStatus = "active"
	DkimStatusInactive DkimStatus = "inactive"
)

// DomainGetDnsRecordsRequest
type DomainGetDnsRecordsRequest struct {
	// Домен, для якого потрібно отримати інформацію про DNS-записи.
	Domain string `json:"domain"`
}

// DomainGetDnsRecordsResponse
type DomainGetDnsRecordsResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Домен, для якого потрібно отримати інформацію про DNS-записи.
	Domain string `json:"domain,omitempty"`

	// Запис, який потрібно внести "як є" для підтвердження володіння цим доменом.
	VerificationRecord string `json:"verification-record,omitempty"`

	// DKIM-підпис для домену. У цьому полі повертається лише ключ. Щоб TXT-запис був коректним, перед
	// ключем потрібно додати "v=DKIM1; k=rsa; p=".
	DKIM string `json:"dkim,omitempty"`

	// Значення CNAME-запису для хоста "_dmarc" домену.
	Dmarc string `json:"dmarc,omitempty"`
}

// DomainValidateVerificationRecordRequest
type DomainValidateVerificationRecordRequest struct {
	// Домен, для якого перевіряється запис про володіння.
	Domain string `json:"domain"`
}

// DomainValidateVerificationRecordResponse
type DomainValidateVerificationRecordResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Налагоджувальне повідомлення про хід перевірки.
	Message string `json:"message,omitempty"`
}

// DomainValidateDkimRequest
type DomainValidateDkimRequest struct {
	// Домен, для якого перевіряється DKIM-запис.
	Domain string `json:"domain"`
}

// DomainValidateDkimResponse
type DomainValidateDkimResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Налагоджувальне повідомлення про хід перевірки.
	Message string `json:"message,omitempty"`
}

// DomainListRequest
type DomainListRequest struct {
	// Опціональне ім'я домену, якщо вам потрібно отримати статуси лише за одним доменом.
	Domain *string `json:"domain,omitempty"`

	// Максимальна кількість доменів, що повертаються. За замовчуванням 50.
	// За замовчуванням: `50`.
	Limit *int `json:"limit,omitempty"`

	// Індекс першого домену зі списку, що повертається, рахуючи з нуля.
	// За замовчуванням: `0`.
	Offset *int `json:"offset,omitempty"`
}

// DomainListResponse
type DomainListResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Масив об'єктів, що описують домени.
	Domains []DomainListResponseDomainsItem `json:"domains,omitempty"`
}

// DomainListResponseDomainsItem
type DomainListResponseDomainsItem struct {
	// Ім'я домену.
	Domain string `json:"domain,omitempty"`

	// Об'єкт, що описує запис перевірки володіння та її статус.
	VerificationRecord *DomainListResponseDomainsItemVerificationRecord `json:"verification-record,omitempty"`

	// Об'єкт, що описує DKIM-запис та її статус.
	DKIM *DomainListResponseDomainsItemDkim `json:"dkim,omitempty"`
}

// DomainListResponseDomainsItemVerificationRecord — Об'єкт, що описує запис перевірки володіння та її статус.
type DomainListResponseDomainsItemVerificationRecord struct {
	// Запис, який потрібно внести "як є" для підтвердження володіння цим доменом.
	Value string `json:"value,omitempty"`

	// Лише домени зі статусом "confirmed" запису про володіння доменом можуть використовуватися для
	// надсилання листів.
	Status common.VerificationRecordStatus `json:"status,omitempty"`
}

// DomainListResponseDomainsItemDkim — Об'єкт, що описує DKIM-запис та її статус.
type DomainListResponseDomainsItemDkim struct {
	// DKIM-підпис для домену. У цьому полі повертається лише ключ. Щоб DNS-запис був коректним, перед
	// ключем потрібно додати "k=rsa, p=" (див.
	// [domain/get-dns-records](https://docs.allymailer.com.ua/ua/web-api-ref#domain-get-dns-records)).
	Key string `json:"key,omitempty"`

	// Лише домени зі статусом DKIM-запису "active" можуть використовуватися для надсилання листів.
	Status DkimStatus `json:"status,omitempty"`
}

// DomainGetSandboxDomainResponse
type DomainGetSandboxDomainResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Ім'я sandbox-домену.
	Domain string `json:"domain,omitempty"`
}

// DomainDeleteRequest
type DomainDeleteRequest struct {
	// Ім'я домену, який потрібно видалити.
	Domain string `json:"domain"`
}

// DomainDeleteResponse
type DomainDeleteResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Домен видалено
	Message string `json:"message,omitempty"`
}
