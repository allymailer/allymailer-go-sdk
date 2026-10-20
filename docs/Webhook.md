# Webhook

[До списку розділів](../README.md)

AllyMailer дозволяє вам отримувати сповіщення про зміну статусу доставки листа або спам-блокування. Ви надаєте URL, і в разі виникнення події, що вас цікавить, ми сповіщаємо вас, передаючи на цей URL JSON-дані про подію.

Ви можете налаштувати формат виклику webhook (просто JSON або JSON, запакований у gzip), які саме події вас цікавлять, чи дозволено групувати кілька подій в одному виклику, чи потрібні вам такі подробиці, як SMTP-відповідь сервера або user agent отримувача тощо.

Зверніть увагу, що ваш обробник має вміти приймати щонайменше 5 паралельних викликів, до 100 подій у кожному (але ви можете налаштувати максимально дозволену кількість паралельних викликів, збільшивши її до 100).

Якщо URL, на який надсилається webhook, недоступний (немає відповіді HTTP 200 OK протягом 3 секунд), спроби надсилання webhook на цей URL до отримання очікуваного 200 OK триватимуть протягом 24 годин з інтервалом у 10 хвилин з додатковим параметром retry_count, значення якого збільшуватиметься на 1 з кожним повторним надсиланням webhook. Якщо протягом 24 годин сталося щонайменше 10 подій і жоден із багаторазових викликів webhook про ці події не був успішним, то ми вважаємо обробник непрацездатним і автоматично встановлюємо його статус у неактивний та повідомляємо про це на email, [вказаний в особистому кабінеті](https://cp.allymailer.com/ua/user/alerting/technical/).

Формат webhook і структура кожної події описані [тут](https://docs.allymailer.com.ua/ua/web-api-ref#callback-format).

`webhook.WebhookAPI`, `client.Webhook`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`Set`](#set) | `POST /webhook/set.json` | Встановлює або змінює властивості [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook) — вашого обробника сповіщень про події. У момент виклику [webhook/set](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-set) ми очікуємо, що до вказаного вами URL можна звернутися GET-запитом без параметрів і у відповідь надійде 200 OK. |
| [`Get`](#get) | `POST /webhook/get.json` | Повертає властивості [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook) |
| [`List`](#list) | `POST /webhook/list.json` | Повернути список усіх або кількох [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook) (обробників сповіщень про події) користувача або проєкту. |
| [`Delete`](#delete) | `POST /webhook/delete.json` | Видалити webhook. |

## Set

`POST /webhook/set.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-set)

Встановлює або змінює властивості [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook) — вашого обробника сповіщень про події. У момент виклику [webhook/set](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-set) ми очікуємо, що до вказаного вами URL можна звернутися GET-запитом без параметрів і у відповідь надійде 200 OK.

### Запит: `WebhookSetRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `URL` | `string` | так | Рядок з URL, на який надсилатиметься запит у разі виникнення події (ідентифікатор обробника). В URL поки що підтримуються лише ASCII-символи. Якщо вам потрібно використати не-ASCII, сконвертуйте URL у [Punycode](https://en.wikipedia.org/wiki/Punycode). |
| `Status` | `webhook.WebhookStatus` |  | Статус webhook, за замовчуванням "active". Статус "disabled" означає, що webhook було вимкнено користувачем, "stopped" означає, що webhook було зупинено системою після 24 годин невдалих викликів (мінімум за 10 різними подіями). (`active`, `disabled`, `stopped`) |
| `EventFormat` | `webhook.WebhookEventFormat` |  | Формат сповіщень. "json_post"(за замовчуванням) або "json_post_gzip". (`json_post`, `json_post_gzip`) |
| `DeliveryInfo` | `int` |  | Повертати (1) чи ні (0) докладну інформацію про доставку. Якщо delivery_info=1, повертається додаткова інформація про відповідь сервера та внутрішній статус доставки в разі статусів доставки "hard_bounced" і "soft_bounced"; user agent, IP-адреса отримувача, геодані, визначена за user agent інформація про пристрій (тип і марка пристрою, назва ОС, браузера або поштового клієнта) і traffic_type (human, machine, protected або unknown) в разі статусів "opened" і "clicked"; URL для статусу "clicked". Для статусів "delivered", "soft_bounced" і "hard_bounced" AllyMailer також може повернути sender_ip — IP-адресу SMTP-сервера, з якого було надіслано лист, якщо він доступний. (`0`, `1`) |
| `SingleEvent` | `int` |  | 0=в одному виклику може надійти декілька подій, 1=лише одна подія на виклик (не рекомендується). (`0`, `1`) |
| `MaxParallel` | `int` |  | Максимальна кількість дозволених паралельних запитів до вашого сервера. Чим більше ваш сервер може обробити — тим краще. |
| `Events` | `webhook.WebhookEventsObject` |  | Об'єкт, що описує події, про які потрібно сповіщати. |

### Відповідь: `WebhookSetResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Object` | `webhook.WebhookObject` | Об'єкт з усіма властивостями webhook. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/webhook"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Webhook.Set(context.Background(), &webhook.WebhookSetRequest{
		URL: "https://yourhost.example.com/allymailer-webhook",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Object)
}
```

## Get

`POST /webhook/get.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-get)

Повертає властивості [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook)

### Запит: `WebhookGetRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `URL` | `string` | так | URL webhook. |

### Відповідь: `WebhookGetResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Object` | `webhook.WebhookObject` | Об'єкт з усіма властивостями webhook. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/webhook"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Webhook.Get(context.Background(), &webhook.WebhookGetRequest{
		URL: "https://example.com/webhook",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Object)
}
```

## List

`POST /webhook/list.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-list)

Повернути список усіх або кількох [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook) (обробників сповіщень про події) користувача або проєкту.

### Запит: `WebhookListRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Limit` | `int` |  | Максимальна кількість webhook, що повертаються за один виклик. Рекомендується 50. |
| `Offset` | `int` |  | Індекс першого webhook зі списку, що повертається, рахуючи з нуля. |

### Відповідь: `WebhookListResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Objects` | `[]webhook.WebhookObject` | Масив об'єктів з даними про webhook. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/webhook"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Webhook.List(context.Background(), &webhook.WebhookListRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Objects)
}
```

## Delete

`POST /webhook/delete.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-delete)

Видалити webhook.

Ви також можете тимчасово вимкнути webhook без його видалення, встановивши його поле status у "disabled" за допомогою методу [webhook/set](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-set).

### Запит: `WebhookDeleteRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `URL` | `string` | так | URL webhook. |

### Відповідь: `SuccessResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/webhook"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Webhook.Delete(context.Background(), &webhook.WebhookDeleteRequest{
		URL: "https://example.com/webhook",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
```

## Колбеки: `WebhookCallback`

Опис формату webhook.

Розбір тіла вхідного запиту вебхука:

```go
callback, err := webhook.ParseCallback(requestBody)
if err != nil {
	panic(err)
}
fmt.Printf("%+v\n", callback)
```

| Поле | Тип | Опис |
|---|---|---|
| `Auth` | `string` | MD5-хеш рядкового тіла повідомлення, у якому значення auth замінено на api_key користувача/проєкту, чий обробник викликається. За допомогою цього отримувач сповіщення може як провести автентифікацію, так і перевірити цілісність сповіщення. |
| `EventsByUser` | `[]webhook.WebhookCallbackEventsByUserItem` | Масив з одним елементом, що містить дані про події користувача/проєкту. |
