# Domain

[До списку розділів](../README.md)

Методи для налаштування доменів надсилання та перевірки їх статусів. Корисні, якщо вам недостатньо додавання доменів через вебінтерфейс AllyMailer і, наприклад, хочеться автоматизувати додавання доменів для ваших власних клієнтів.

Щоб додати новий домен надсилання через API, ви маєте зробити таке:
- Підтвердіть домен, з якого надсилатимуться листи, внісши необхідні DNS-записи, отримавши їх за допомогою методу [domain/get-dns-records](https://docs.allymailer.com.ua/ua/web-api-ref#domain-get-dns-records).
- Після додавання раз на 5–10 хвилин викликайте [domain/list](https://docs.allymailer.com.ua/ua/web-api-ref#domain-list) із зазначенням доданого домену (або кількох доменів одразу, якщо перевірити треба декілька) і перевіряйте verification-record.status та dkim.status у відповіді domain/list. Можна робити запит і частіше — але точно не частіше за раз на хвилину.
- Якщо verification-record.status="unconfirmed", то ініціюйте перевірку правильності запису володіння доменом за допомогою виклику [domain/validate-verification-record](https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-verification-record).
- Якщо dkim.status="inactive", то ініціюйте перевірку правильності DKIM-запису за допомогою виклику [domain/validate-dkim](https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-dkim) (втім, DKIM-записи AllyMailer і сам перевіряє раз на пів години).
- І лише коли domain/list повернув verification-record.status="confirmed" та dkim.status="active" для домену, можна використовувати домен для надсилання та припинити його очікування й виклики для нього зазначених вище методів.

`domain.DomainAPI`, `client.Domain`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`GetDNSRecords`](#getdnsrecords) | `POST /domain/get-dns-records.json` | Повертає інформацію про записи, які потрібно внести в DNS для використання домену в AllyMailer. |
| [`ValidateVerificationRecord`](#validateverificationrecord) | `POST /domain/validate-verification-record.json` | Ініціює перевірку запису про володіння доменом. |
| [`ValidateDKIM`](#validatedkim) | `POST /domain/validate-dkim.json` | Перевіряє DKIM-запис домену. |
| [`List`](#list) | `POST /domain/list.json` | Повертає список усіх зареєстрованих доменів користувача або проєкту разом зі статусами останніх перевірок. |
| [`GetSandboxDomain`](#getsandboxdomain) | `POST /domain/get-sandbox-domain.json` | Створює sandbox-домен облікового запису або повертає наявний. |
| [`Delete`](#delete) | `POST /domain/delete.json` | Видаляє домен. Увага: для видалення домену, що належить проєкту, вам потрібно використати API-ключ проєкту. |

## GetDNSRecords

`POST /domain/get-dns-records.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#domain-get-dns-records)

Повертає інформацію про записи, які потрібно внести в DNS для використання домену в AllyMailer.

Щоб надсилати листи, вам потрібно підтвердити володіння доменом, зі зворотною адресою якого ви збираєтеся їх надсилати, а також налаштувати DKIM-підпис для цього домену. Метод дозволяє вам зареєструвати домен у AllyMailer і згенерувати запис для перевірки володіння та DKIM-ключ, або отримати їх, якщо вони були згенеровані раніше.

Після отримання успішної відповіді вам потрібно додати DNS-записи на кшталт цих:

`@             IN TXT "allymailer-validate-hash=483bb362ebdbeedd755cfb1d4d661"`

`us._domainkey IN TXT "k=rsa; p=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDo7"`

`@             IN TXT  "v=spf1 include:spf.allymailer.com.ua ~all"`

- Перший запис представлено вмістом поля "verification-record".
- Другий запис має включати вміст поля "dkim" одразу після "p=". Зверніть увагу на ім'я запису! Префікс "us._domainkey." дуже важливий для того, щоб DKIM-селектор AllyMailer працював правильно.
- Третій запис необов'язковий, але ми дуже рекомендуємо його додавати — це SPF-запис, який підвищує доставлюваність. Можливо, вона у вас уже є — у цьому разі просто додайте include:spf.allymailer.com.ua після v=spf1.

Вимоги до синтаксису записів у різних DNS-провайдерів можуть відрізнятися:
- іноді потрібно вказувати повне ім'я домену із завершальною крапкою, на кшталт "example.com." замість символа @.
- деякі провайдери вимагають указувати час життя (TTL, наприклад 3600)
- для деяких провайдерів буде зайвим поміщення значень у подвійні лапки
Будь ласка, перевірте вимоги свого провайдера й підкажіть це своїм користувачам.

Для реалізації white-label рішень є можливість змінити стандартні DNS-записи на ті, що не містять підрядка "allymailer". Щоб дізнатися докладніше — напишіть у техпідтримку.

### Запит: `DomainGetDnsRecordsRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Domain` | `string` | так | Домен, для якого потрібно отримати інформацію про DNS-записи. |

### Відповідь: `DomainGetDnsRecordsResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Domain` | `string` | Домен, для якого потрібно отримати інформацію про DNS-записи. |
| `VerificationRecord` | `string` | Запис, який потрібно внести "як є" для підтвердження володіння цим доменом. |
| `DKIM` | `string` | DKIM-підпис для домену. У цьому полі повертається лише ключ. Щоб TXT-запис був коректним, перед ключем потрібно додати "v=DKIM1; k=rsa; p=". |
| `Dmarc` | `string` | Значення CNAME-запису для хоста "_dmarc" домену. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/domain"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Domain.GetDNSRecords(context.Background(), &domain.DomainGetDnsRecordsRequest{
		Domain: "example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Domain)
}
```

## ValidateVerificationRecord

`POST /domain/validate-verification-record.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-verification-record)

Ініціює перевірку запису про володіння доменом.

Якщо результат виклику цього методу успішний (200 OK) — це означає, що перевірку запущено. Для отримання результату потрібно пізніше перевірити поле verification-record.status, що повертається методом [domain/list](https://docs.allymailer.com.ua/ua/web-api-ref#domain-list). Докладніше читайте [тут](https://docs.allymailer.com.ua/ua/web-api-ref#domain).

### Запит: `DomainValidateVerificationRecordRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Domain` | `string` | так | Домен, для якого перевіряється запис про володіння. |

### Відповідь: `DomainValidateVerificationRecordResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Message` | `string` | Налагоджувальне повідомлення про хід перевірки. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/domain"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Domain.ValidateVerificationRecord(context.Background(), &domain.DomainValidateVerificationRecordRequest{
		Domain: "example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Message)
}
```

## ValidateDKIM

`POST /domain/validate-dkim.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#domain-validate-dkim)

Перевіряє DKIM-запис домену.

Якщо результат виклику цього методу успішний (200 OK) — це означає, що перевірку запущено. Для отримання результату потрібно пізніше перевірити поле dkim.status, що повертається методом [domain/list](https://docs.allymailer.com.ua/ua/web-api-ref#domain-list). Докладніше читайте [тут](https://docs.allymailer.com.ua/ua/web-api-ref#domain).

### Запит: `DomainValidateDkimRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Domain` | `string` | так | Домен, для якого перевіряється DKIM-запис. |

### Відповідь: `DomainValidateDkimResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Message` | `string` | Налагоджувальне повідомлення про хід перевірки. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/domain"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Domain.ValidateDKIM(context.Background(), &domain.DomainValidateDkimRequest{
		Domain: "example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Message)
}
```

## List

`POST /domain/list.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#domain-list)

Повертає список усіх зареєстрованих доменів користувача або проєкту разом зі статусами останніх перевірок.

Повертає статус останньої перевірки DNS-запису володіння доменом і DKIM-запису для кожного домену. Докладніше [тут](https://docs.allymailer.com.ua/ua/web-api-ref#domain).

### Запит: `DomainListRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Domain` | `string` |  | Опціональне ім'я домену, якщо вам потрібно отримати статуси лише за одним доменом. |
| `Limit` | `int` |  | Максимальна кількість доменів, що повертаються. За замовчуванням 50. |
| `Offset` | `int` |  | Індекс першого домену зі списку, що повертається, рахуючи з нуля. |

### Відповідь: `DomainListResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Domains` | `[]domain.DomainListResponseDomainsItem` | Масив об'єктів, що описують домени. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/domain"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Domain.List(context.Background(), &domain.DomainListRequest{
		Domain: allymailer.String("example.com"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Domains)
}
```

## GetSandboxDomain

`POST /domain/get-sandbox-domain.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#domain-get-sandbox-domain)

Створює sandbox-домен облікового запису або повертає наявний.

Повертає sandbox-домен облікового запису, готовий до використання в полі `from_email` методу [email/send](https://docs.allymailer.com.ua/ua/web-api-ref#email-send). Якщо sandbox-домену немає, AllyMailer створює його та повертає. Повторні виклики повертають той самий домен.

Метод доступний лише з API-ключем облікового запису. API-ключі проєктів не можуть створювати sandbox-домени.

### Запит

Метод не має параметрів.

### Відповідь: `DomainGetSandboxDomainResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Domain` | `string` | Ім'я sandbox-домену. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Domain.GetSandboxDomain(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Domain)
}
```

## Delete

`POST /domain/delete.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#domain-delete)

Видаляє домен. Увага: для видалення домену, що належить проєкту, вам потрібно використати API-ключ проєкту.

Видаляє домен.

### Запит: `DomainDeleteRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Domain` | `string` | так | Ім'я домену, який потрібно видалити. |

### Відповідь: `DomainDeleteResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Message` | `string` | Домен видалено |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/domain"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Domain.Delete(context.Background(), &domain.DomainDeleteRequest{
		Domain: "example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Message)
}
```
