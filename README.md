# AllyMailer Go SDK

[![Go](https://pkg.go.dev/badge/github.com/allymailer/allymailer-go-sdk.svg)](https://pkg.go.dev/github.com/allymailer/allymailer-go-sdk) [![License: MIT](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

Офіційна бібліотека Go для API транзакційних листів [AllyMailer](https://allymailer.com.ua).

[Довідник API](https://docs.allymailer.com.ua/ua/web-api-ref) · Документація для LLM: [llms.txt](https://docs.allymailer.com.ua/llms.txt)

> **Go net/http?** Встановіть [`github.com/allymailer/allymailer-nethttp`](https://github.com/allymailer/allymailer-nethttp): у ньому цей SDK, налаштування, поштовий транспорт і маршрут для вебхуків.

## Встановлення

```bash
go get github.com/allymailer/allymailer-go-sdk
```

Вимоги: Go 1.18+. Сторонніх залежностей немає.

Отримайте API-ключ у [панелі керування](https://cp.allymailer.com) і зберігайте його у змінній оточення, а не в коді:

```bash
export ALLYMAILER_API_KEY="your-api-key"
```

## Надсилання листа

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/email"
	"github.com/allymailer/allymailer-go-sdk/common"
)

func main() {
	client := allymailer.NewClient(os.Getenv("ALLYMAILER_API_KEY"))
	ctx := context.Background()

	response, err := client.Email.Send(ctx, &email.EmailSendRequest{
		Message: &email.MessageObject{
			Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("Олексій <alex@example.com>")},
			FromEmail:  allymailer.String("shop@example.com"),
			FromName:   allymailer.String("Магазин"),
			Subject:    "Ваше замовлення",
			Body:       common.HTMLBody("<p>Вітаємо, {{to_name}}! Ваше замовлення надіслано.</p>"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.JobID)
}
```

Адреса відправника має бути на домені, підтвердженому у вашому обліковому записі. Домену ще немає? Надішліть лист із [sandbox-домену](#без-свого-домену-тестовий-sandbox-домен) у три кроки.

## Зміст

- [Рецепти](#рецепти): персоналізація · вкладення · копія і прихована копія · відкладене надсилання · безпечний повтор · помилки · вебхуки
- [Без свого домену: тестовий sandbox-домен](#без-свого-домену-тестовий-sandbox-домен)
- [Запуск у продакшен](#запуск-у-продакшен)
- [Довідник](#довідник)
- [Допомога](#допомога)

## Рецепти

Кожен рецепт продовжує приклад вище: клієнт та імпорти вже є.

### Персоналізація для кожного одержувача

До 500 одержувачів в одному виклику, кожен отримує окрему копію зі своїми підстановками. Метадані одержувача повертаються в його подіях вебхука.

```go
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{
			email.RecipientFromAddress("Олексій <alex@example.com>", map[string]string{"order": "1001"}).
				WithMetadata(map[string]string{"user_id": "42"}),
			email.RecipientFromAddress("Марія <maria@example.com>", map[string]string{"order": "1002"}).
				WithMetadata(map[string]string{"user_id": "43"}),
		},
		FromEmail: allymailer.String("shop@example.com"),
		Subject:   "Замовлення №{{order}}",
		Body:      common.HTMLBody("<p>{{to_name}}, замовлення №{{order}} надіслано.</p>"),
	},
}); err != nil {
	log.Fatal(err)
}
```

### Вкладення

```go
import (
	"encoding/base64"
	"os"
)

content, err := os.ReadFile("invoice.pdf")
if err != nil {
	log.Fatal(err)
}
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("alex@example.com")},
		FromEmail:  allymailer.String("shop@example.com"),
		Subject:    "Рахунок",
		Body:       common.HTMLBody("<p>Рахунок у вкладенні.</p>"),
		Attachments: []common.AttachmentsArrayItem{
			{Type: "application/pdf", Name: "invoice.pdf", Content: base64.StdEncoding.EncodeToString(content)},
		},
	},
}); err != nil {
	log.Fatal(err)
}
```

### Копія і прихована копія

Усі адреси передаються в `recipients` (кожна отримує копію), видимі — ще й у заголовках. `audit@example.com` немає в жодному заголовку: це прихована копія. Правила й обмеження: [копія і прихована копія](https://docs.allymailer.com.ua/ua/cc-and-bcc).

```go
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{
			email.RecipientFromAddress("alex@example.com"),
			email.RecipientFromAddress("boss@example.com"),
			email.RecipientFromAddress("audit@example.com"),
		},
		FromEmail: allymailer.String("shop@example.com"),
		Subject:   "Звіт за тиждень",
		Body:      common.HTMLBody("<p>Звіт у вкладенні.</p>"),
		Headers: &common.HeadersObject{
			Extra: map[string]string{"To": "alex@example.com", "CC": "boss@example.com"},
		},
	},
}); err != nil {
	log.Fatal(err)
}
```

### Відкладене надсилання

`send_at` — в UTC, не далі ніж на 24 години вперед.

```go
import "time"

sendAt := time.Now().UTC().Add(time.Hour).Format("2006-01-02 15:04:05") // UTC, "YYYY-MM-DD hh:mm:ss"
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("alex@example.com")},
		FromEmail:  allymailer.String("shop@example.com"),
		Subject:    "Нагадування",
		Body:       common.TextBody("До зустрічі завтра!"),
		Options:    &email.MessageObjectOptions{SendAt: allymailer.String(sendAt)},
	},
}); err != nil {
	log.Fatal(err)
}
```

### Безпечний повтор

Повтор із тим самим `idempotence_key` протягом хвилини не надішлеться вдруге: з ним повторюйте мережеві помилки, таймаути, 429 і 5xx.

```go
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients:     []email.MessageObjectRecipientsItem{email.RecipientFromAddress("alex@example.com")},
		FromEmail:      allymailer.String("shop@example.com"),
		Subject:        "Скидання пароля",
		Body:           common.HTMLBody("<p>Посилання для скидання пароля.</p>"),
		IdempotenceKey: allymailer.String("password-reset-8f3a"),
	},
}); err != nil {
	log.Fatal(err)
}
```

### Обробка помилок

Одержувачі, яких API відкинув (відписані, неіснуючі, у стоп-листі), — не помилка: вони у `failed_emails` відповіді, решта лист отримає.

```go
import "errors"

if _, err := client.Email.Send(ctx, request); err != nil {
	var apiErr *allymailer.APIError
	if errors.As(err, &apiErr) {
		// HTTP-статус (0: мережева помилка або таймаут), код помилки API, повідомлення
		log.Println(apiErr.HTTPStatus, apiErr.ErrorCode(), apiErr.Message)
	}
}
```

### Статуси доставки

Зареєструйте свою HTTPS-адресу один раз (`webhook/set`). Потім перевіряйте підпис кожного сповіщення (`auth`: MD5 сирого тіла, у якому значення `auth` замінено на API-ключ; на неправильний — 401) і обробляйте події.

```go
import (
	"crypto/md5"
	"net/http"
	"os"
	"strings"

	"github.com/allymailer/allymailer-go-sdk/api/webhook"
)

// handleWebhook: тіло POST-запиту як надійшло; повертає HTTP-статус відповіді
func handleWebhook(body []byte) int {
	callback, err := webhook.ParseCallback(body)
	if err != nil {
		return http.StatusBadRequest
	}
	// auth: MD5 сирого тіла, у якому значення auth замінено на ваш API-ключ
	auth := callback.Auth
	signed := strings.Replace(string(body), auth, os.Getenv("ALLYMAILER_API_KEY"), 1)
	if auth == "" || fmt.Sprintf("%x", md5.Sum([]byte(signed))) != auth {
		return http.StatusUnauthorized
	}

	for _, user := range callback.EventsByUser {
		for _, event := range user.Events {
			// email, status (delivered, opened, hard_bounced, spam, ...), job_id, metadata
			log.Printf("%+v", event.EventData)
		}
	}
	return http.StatusOK
}
```

> **Відповідайте 200 протягом 3 секунд**, а події обробляйте у фоні (через чергу): інакше AllyMailer повторює запит кожні 10 хвилин протягом доби (зростає `retry_count`), а обробник, на якому за добу не пройшов жоден запит хоча б для 10 подій, вимикається. Обробник має витримувати **щонайменше 5 паралельних запитів до 100 подій у кожному**; повтор надсилає ті самі події ще раз — прибирайте дублікати.

## Без свого домену: тестовий sandbox-домен

Поки ваш домен не підтверджено, листи можна надсилати із **sandbox-домену** — тестового домену, який є в кожного облікового запису: він уже підтверджений, DKIM і SPF налаштовані. Листи з нього доставляються лише на підтверджені вами адреси (посилання для підтвердження приходить листом і діє 48 годин), до 100 на день. Через 31 день після створення sandbox-домен видаляється, але його можна отримати знову: наступний виклик `domain/get-sandbox-domain` створить новий (з іншою назвою, тому беріть її з відповіді, а не прописуйте в коді).

```go
import "github.com/allymailer/allymailer-go-sdk/api/confirmedemail"

// 1. Ваш sandbox-домен: створюється під час першого виклику, далі повертається той самий
sandbox, err := client.Domain.GetSandboxDomain(ctx)
if err != nil {
	log.Fatal(err)
}

// 2. Підтвердіть свою скриньку: AllyMailer надішле посилання, перейдіть за ним протягом 48 годин
confirm := &confirmedemail.ConfirmedEmailRequest{Email: "you@example.com"}
if _, err := client.ConfirmedEmail.Create(ctx, confirm); err != nil {
	log.Fatal(err)
}

// 3. Надішліть із будь-якої адреси sandbox-домену на підтверджену адресу
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("you@example.com")},
		FromEmail:  allymailer.String("noreply@" + sandbox.Domain),
		Subject:    "Привіт від AllyMailer",
		Body:       common.HTMLBody("<p>Працює!</p>"),
	},
}); err != nil {
	log.Fatal(err)
}
```

<details>
<summary>Обмеження sandbox</summary>

- Надсилання лише на підтверджені адреси: до 10 на обліковий запис.
- До 100 листів на день.
- Sandbox-домен спливає через 31 день; знову викличте `domain/get-sandbox-domain`, щоб отримати новий.
- Працює лише з API-ключем облікового запису, не з ключем проєкту.
- Не показується в `domain/list`; інші методи `domain/*` до нього не застосовуються.

Sandbox — для розробки та тестування. Не використовуйте sandbox-домен як відправника на проді.

</details>

<details>
<summary>Те саме через curl</summary>

```bash
BASE=https://cp.allymailer.com/ua/transactional/api/v1

curl -s -X POST $BASE/domain/get-sandbox-domain.json \
  -H "Content-Type: application/json" -H "X-API-KEY: $ALLYMAILER_API_KEY" -d '{}'

curl -s -X POST $BASE/confirmed-email/create.json \
  -H "Content-Type: application/json" -H "X-API-KEY: $ALLYMAILER_API_KEY" \
  -d '{"email": "you@example.com"}'

curl -s -X POST $BASE/email/send.json \
  -H "Content-Type: application/json" -H "X-API-KEY: $ALLYMAILER_API_KEY" \
  -d '{
        "message": {
          "recipients": [
            {
              "email": "you@example.com"
            }
          ],
          "from_email": "noreply@<sandbox-domain>",
          "subject": "Привіт від AllyMailer",
          "body": {
            "html": "<p>Працює!</p>"
          }
        }
      }'
```

</details>

## Запуск у продакшен

**1. Підтвердіть домен надсилання.** `domain/get-dns-records` повертає запис підтвердження, DKIM-ключ і адресу для DMARC; додайте записи у свого DNS-провайдера, потім `domain/validate-verification-record` і `domain/validate-dkim`.

| Тип | Хост | Значення | Звідки взяти |
|---|---|---|---|
| TXT | `@` | `allymailer-validate-hash=…` | поле `verification-record` відповіді, як є |
| TXT | `<селектор>._domainkey` (селектор — на сторінці домену в панелі керування, наприклад `us`) | `v=DKIM1; k=rsa; p=<dkim>` | поле `dkim` відповіді після `v=DKIM1; k=rsa; p=` |
| TXT | `@` | `v=spf1 include:spf.allymailer.com ~all` | у відповіді немає; якщо SPF-запис уже є, додайте include до нього |
| CNAME | `_dmarc` | `<домен>.dmarc.…` | поле `dmarc` відповіді (DMARC-політику публікує AllyMailer); якщо в домену вже є свій TXT-запис DMARC, залиште його |

**2. Надійне надсилання.**

| | |
|---|---|
| Повтори | `429` і `5xx` з backoff і тим самим `idempotence_key`; інші помилки виправляйте, а не повторюйте |
| Ліміти | 500 одержувачів і 10 МБ на запит разом із вкладеннями: у base64 файл стає приблизно на третину більшим |
| Тіло | завжди `html` або `plaintext` (або із шаблону) |
| Статус доставки | вебхук, а не опитування |
| Черга | надсилайте з фонового завдання, а не всередині веб-запиту |
| Налаштування | `from_email` з конфігурації: перехід із sandbox на продакшен без змін у коді |
| Типові налаштування | не прибирайте блок відписки та трекінг; змінити їх можна через підтримку AllyMailer |

## Довідник

### Розділи

| Розділ | Методи | Опис |
|---|---|---|
| [Email](docs/Email.md) | `Send`, `Subscribe` | Ця група методів пов'язана з надсиланням email. |
| [EmailValidation](docs/EmailValidation.md) | `Single` | Методи перевірки email-адрес. Перевірка адреси електронної пошти надає докладну інформацію про працездатність email. Це дозволить виключити з надсилання неіснуючі email, знизити показник скарг, очистити список підписників і врешті-решт покращити вашу репутацію як відправника. Наприклад, ви можете інтегрувати процес перевірки у форму реєстрації вашого сайту і вже на етапі введення адреси повідомляти користувачам, що адреса, яку вони ввели у форму, недійсна. |
| [ConfirmedEmail](docs/ConfirmedEmail.md) | `Create`, `List`, `Delete` | Керування email-адресами, підтвердженими для надсилання із sandbox-домену. Методи доступні лише з API-ключем основного облікового запису. |
| [Template](docs/Template.md) | `Set`, `Get`, `List`, `Delete` | Ви можете зберегти шаблон на боці AllyMailer і потім посилатися на нього за ідентифікатором. Це не обов'язково, але вкрай рекомендовано у разі повторюваних email-повідомлень або за великих обсягів надсилання. AllyMailer підтримує два різні синтаксиси для шаблонів, докладніше читайте [тут](https://docs.allymailer.com.ua/ua/template-engines). |
| [Webhook](docs/Webhook.md) | `Set`, `Get`, `List`, `Delete` | AllyMailer дозволяє вам отримувати сповіщення про зміну статусу доставки листа або спам-блокування. Ви надаєте URL, і в разі виникнення події, що вас цікавить, ми сповіщаємо вас, передаючи на цей URL JSON-дані про подію. |
| [Suppression](docs/Suppression.md) | `Set`, `Get`, `List`, `Delete` | Методи для керування адресами, надсилання на які заблоковано з різних причин — відписка, недоступність, скарги. |
| [Domain](docs/Domain.md) | `GetDNSRecords`, `ValidateVerificationRecord`, `ValidateDKIM`, `List`, `GetSandboxDomain`, `Delete` | Методи для налаштування доменів надсилання та перевірки їх статусів. Корисні, якщо вам недостатньо додавання доменів через вебінтерфейс AllyMailer і, наприклад, хочеться автоматизувати додавання доменів для ваших власних клієнтів. |
| [EventDump](docs/EventDump.md) | `Create`, `Get`, `List`, `Delete` | Методи для отримання списку подій, що сталися, за певний період. Дані зберігаються до 45 днів, залежно від вашого тарифу. Це означає, що можна запросити події, починаючи з 00:00:00 45-го дня тому і до поточного моменту часу. Для індивідуальних тарифів з великими обсягами за замовчуванням події доступні лише за останні 7 днів. |
| [Tag](docs/Tag.md) | `List`, `Delete` | Методи керування користувацькими мітками. За допомогою міток ви можете класифікувати листи за потрібними критеріями. Відокремити маркетингові листи від транзакційних або листи з підтвердженням реєстрації від сповіщень про замовлення. AllyMailer дозволяє позначати мітками кожне вихідне повідомлення, а потім сегментувати й аналізувати статистику на основі зазначених під час надсилання міток. |
| [Project](docs/Project.md) | `Create`, `Update`, `List`, `Delete` | Проєкти дозволяють об'єднати кілька незалежних доменів із webhook, шаблонами та списками відписки. Це дуже корисно, якщо у вас є власні клієнти і ви керуєте їхніми даними з одного облікового запису AllyMailer. Кожен проєкт отримує власний ключ API, а виклики методів API з цим ключем обробляють лише дані цього проєкту. Таким чином, ви можете призначити окремий проєкт для кожного користувача і бути впевненими, що вони не конфліктуватимуть: наприклад, відписка від одного проєкту не вплине на інший. |
| [System](docs/System.md) | `Ping`, `Info` | Системні методи. |
| [Obsolete](docs/Obsolete.md) | `UnsubscribedSet`, `UnsubscribedCheck`, `UnsubscribedList` | Застарілі методи роботи з відписками. Ми рекомендуємо користуватися замість них гнучкішими методами [suppression](https://docs.allymailer.com.ua/ua/web-api-ref#suppression). Зверніть увагу, що знову підписати відписаного адресата через WebAPI не можна. Можна лише надіслати йому email із посиланням повторної підписки за допомогою методу email/subscribe. |

### Обробка помилок

Кожен метод кидає `*client.APIError`, якщо API повернув помилку; у ньому є [код помилки API](https://docs.allymailer.com.ua/ua/web-api-ref#api-errors) і HTTP-статус.

Транспортна помилка (таймаут, DNS, TLS, скасований контекст) обгортається: `errors.Is(err, context.DeadlineExceeded)` і `errors.As(err, &netErr)` працюють через `*client.APIError` (`Err`, `Unwrap()`).

```go
response, err := client.Email.Send(ctx, request)
var apiErr *client.APIError
var netErr net.Error
switch {
case errors.Is(err, context.DeadlineExceeded):
	// the timeout of ctx or of the HTTP client
case errors.As(err, &netErr):
	// network error
case errors.As(err, &apiErr) && apiErr.Code != nil:
	fmt.Println(apiErr.HTTPStatus, *apiErr.Code, apiErr.Message)
}
```

### HTTP-клієнт, заголовки, повтори

Запити надсилає `Configuration.HTTPClient`: його `Transport` (`http.RoundTripper`) — місце для повторів, логування, метрик і проксі. `Configuration.Headers` додає HTTP-заголовки до кожного запиту.

```go
type retryTransport struct{ next http.RoundTripper }

func (t retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Println(req.Method, req.URL.Path)
	// requests are POST with a JSON body: GetBody gives it again for a retry
	return t.next.RoundTrip(req)
}

config := client.NewConfiguration("YOUR_API_KEY")
config.HTTPClient = &http.Client{Timeout: 30 * time.Second, Transport: retryTransport{http.DefaultTransport}}
config.Headers = map[string]string{"X-Request-Source": "billing"}
c := allymailer.NewClientWithConfig(config)
```

### Додаткові ключі моделі

`HeadersObject` приймає ключі окрім іменованих полів (будь-який заголовок `X-*`, `CC`, `List-Unsubscribe`, ...) у мапі `Extra`: вони надсилаються поруч з іменованими полями, за збігу імен перемагає іменоване поле; невідомі ключі відповіді потрапляють в `Extra`.

```go
headers := &common.HeadersObject{
	XALLYMAILERGlobalLanguage: allymailer.String("en"),
	Extra: map[string]string{
		"X-MyHeader":       "some data",
		"CC":               "copy@example.com",
		"List-Unsubscribe": "<mailto:unsubscribe@example.com?subject=unsubscribe>",
	},
}
```

## Допомога

- Для ШІ-агентів: [llms.txt](https://docs.allymailer.com.ua/llms.txt), скіли: `npx skills add https://github.com/allymailer/allymailer-skills.git`
- [Підтримка](https://cp.allymailer.com/ua/support) · [Задачі](https://github.com/allymailer/allymailer-go-sdk/issues)
- Розробникам: [AGENTS.md](AGENTS.md) — SDK генерується з OpenAPI-специфікації.
