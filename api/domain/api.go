// DomainAPI: API methods of the Domain section.

package domain

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
)

// DomainAPI — Методи для налаштування доменів надсилання та перевірки їх статусів. Корисні, якщо вам недостатньо
// додавання доменів через вебінтерфейс AllyMailer і, наприклад, хочеться автоматизувати додавання
// доменів для ваших власних клієнтів.
//
// Щоб додати новий домен надсилання через API, ви маєте зробити таке:
// - Підтвердіть домен, з якого надсилатимуться листи, внісши необхідні DNS-записи, отримавши їх за допомогою методу [domain/get-dns-records](https://docs.allymailer.com.ua/ua/web-api-ref#domain-get-dns-records).
// - Після додавання раз на 5–10 хвилин викликайте [domain/list](https://docs.allymailer.com.ua/ua/web-api-ref#domain-list) із зазначенням доданого домену (або кількох доменів одразу, якщо перевірити треба декілька) і перевіряйте verification-record.status та dkim.status у відповіді domain/list. Можна робити запит і частіше — але точно не частіше за раз на хвилину.
// - Якщо verification-record.status="unconfirmed", то ініціюйте перевірку правильності запису володіння доменом за допомогою виклику [domain/validate-verification-record](https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-verification-record).
// - Якщо dkim.status="inactive", то ініціюйте перевірку правильності DKIM-запису за допомогою виклику [domain/validate-dkim](https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-dkim) (втім, DKIM-записи AllyMailer і сам перевіряє раз на пів години).
// - І лише коли domain/list повернув verification-record.status="confirmed" та dkim.status="active" для домену, можна використовувати домен для надсилання та припинити його очікування й виклики для нього зазначених вище методів.
type DomainAPI struct {
	client *client.APIClient
}

// NewDomainAPI creates the Domain API.
func NewDomainAPI(c *client.APIClient) *DomainAPI {
	return &DomainAPI{client: c}
}

// GetDNSRecords — Повертає інформацію про записи, які потрібно внести в DNS для використання домену в AllyMailer.
//
// Щоб надсилати листи, вам потрібно підтвердити володіння доменом, зі зворотною адресою якого ви
// збираєтеся їх надсилати, а також налаштувати DKIM-підпис для цього домену. Метод дозволяє вам
// зареєструвати домен у AllyMailer і згенерувати запис для перевірки володіння та DKIM-ключ, або
// отримати їх, якщо вони були згенеровані раніше.
//
// Після отримання успішної відповіді вам потрібно додати DNS-записи на кшталт цих:
//
// `@             IN TXT "allymailer-validate-hash=483bb362ebdbeedd755cfb1d4d661"`
//
// `us._domainkey IN TXT "v=DKIM1; k=rsa; p=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDo7"`
//
// `@             IN TXT  "v=spf1 include:spf.allymailer.com.ua ~all"`
//
// `_dmarc        IN CNAME "example.com.dmarc.allymailer.com.ua."`
//
// - Перший запис представлено вмістом поля "verification-record".
// - Другий запис має включати вміст поля "dkim" одразу після "v=DKIM1; k=rsa; p=". Зверніть увагу на ім'я запису! Префікс "us._domainkey." дуже важливий для того, щоб DKIM-селектор AllyMailer працював правильно.
// - Третій запис необов'язковий, але ми дуже рекомендуємо його додавати — це SPF-запис, який підвищує доставлюваність. Метод його не повертає. Можливо, вона у вас уже є — у цьому разі просто додайте include:spf.allymailer.com.ua після v=spf1.
// - Четвертий запис — CNAME хоста "_dmarc" на значення поля "dmarc": тоді DMARC-політику домену публікує AllyMailer. Якщо в домену вже є свій TXT-запис DMARC, залиште його.
//
// Вимоги до синтаксису записів у різних DNS-провайдерів можуть відрізнятися:
// - іноді потрібно вказувати повне ім'я домену із завершальною крапкою, на кшталт "example.com." замість символа @.
// - деякі провайдери вимагають указувати час життя (TTL, наприклад 3600)
// - для деяких провайдерів буде зайвим поміщення значень у подвійні лапки
// Будь ласка, перевірте вимоги свого провайдера й підкажіть це своїм користувачам.
//
// Для реалізації white-label рішень є можливість змінити стандартні DNS-записи на ті, що не містять
// підрядка "allymailer". Щоб дізнатися докладніше — напишіть у техпідтримку.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#domain-get-dns-records
func (a *DomainAPI) GetDNSRecords(ctx context.Context, request *DomainGetDnsRecordsRequest) (*DomainGetDnsRecordsResponse, error) {
	response := &DomainGetDnsRecordsResponse{}
	if err := a.client.Call(ctx, "/domain/get-dns-records.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ValidateVerificationRecord — Ініціює перевірку запису про володіння доменом.
//
// Якщо результат виклику цього методу успішний (200 OK) — це означає, що перевірку запущено. Для
// отримання результату потрібно пізніше перевірити поле verification-record.status, що повертається
// методом [domain/list](https://docs.allymailer.com.ua/ua/web-api-ref#domain-list). Докладніше читайте
// [тут](https://docs.allymailer.com.ua/ua/web-api-ref#domain).
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-verification-record
func (a *DomainAPI) ValidateVerificationRecord(ctx context.Context, request *DomainValidateVerificationRecordRequest) (*DomainValidateVerificationRecordResponse, error) {
	response := &DomainValidateVerificationRecordResponse{}
	if err := a.client.Call(ctx, "/domain/validate-verification-record.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ValidateDKIM — Перевіряє DKIM-запис домену.
//
// Якщо результат виклику цього методу успішний (200 OK) — це означає, що перевірку запущено. Для
// отримання результату потрібно пізніше перевірити поле dkim.status, що повертається методом
// [domain/list](https://docs.allymailer.com.ua/ua/web-api-ref#domain-list). Докладніше читайте
// [тут](https://docs.allymailer.com.ua/ua/web-api-ref#domain).
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-dkim
func (a *DomainAPI) ValidateDKIM(ctx context.Context, request *DomainValidateDkimRequest) (*DomainValidateDkimResponse, error) {
	response := &DomainValidateDkimResponse{}
	if err := a.client.Call(ctx, "/domain/validate-dkim.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Повертає список усіх зареєстрованих доменів користувача або проєкту разом зі статусами останніх
// перевірок.
//
// Повертає статус останньої перевірки DNS-запису володіння доменом і DKIM-запису для кожного домену.
// Докладніше [тут](https://docs.allymailer.com.ua/ua/web-api-ref#domain).
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#domain-list
func (a *DomainAPI) List(ctx context.Context, request *DomainListRequest) (*DomainListResponse, error) {
	response := &DomainListResponse{}
	if err := a.client.Call(ctx, "/domain/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetSandboxDomain — Створює sandbox-домен облікового запису або повертає наявний.
//
// Повертає sandbox-домен облікового запису, готовий до використання в полі `from_email` методу
// [email/send](https://docs.allymailer.com.ua/ua/web-api-ref#email-send). Якщо sandbox-домену немає,
// AllyMailer створює його та повертає. Повторні виклики повертають той самий домен.
//
// Sandbox-домен видаляється через 31 день після створення; наступний виклик створить новий. З нього
// можна надіслати до 100 листів на день (ліміт облікового запису за замовчуванням).
//
// Метод доступний лише з API-ключем облікового запису. API-ключі проєктів не можуть створювати
// sandbox-домени.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#domain-get-sandbox-domain
func (a *DomainAPI) GetSandboxDomain(ctx context.Context) (*DomainGetSandboxDomainResponse, error) {
	response := &DomainGetSandboxDomainResponse{}
	if err := a.client.Call(ctx, "/domain/get-sandbox-domain.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Видаляє домен. Увага: для видалення домену, що належить проєкту, вам потрібно використати API-ключ
// проєкту.
//
// Видаляє домен.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#domain-delete
func (a *DomainAPI) Delete(ctx context.Context, request *DomainDeleteRequest) (*DomainDeleteResponse, error) {
	response := &DomainDeleteResponse{}
	if err := a.client.Call(ctx, "/domain/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
