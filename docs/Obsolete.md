# Obsolete

[До списку розділів](../README.md)

Застарілі методи роботи з відписками. Ми рекомендуємо користуватися замість них гнучкішими методами [suppression](https://docs.allymailer.com.ua/ua/web-api-ref#suppression). Зверніть увагу, що знову підписати відписаного адресата через WebAPI не можна. Можна лише надіслати йому email із посиланням повторної підписки за допомогою методу email/subscribe.

`obsolete.ObsoleteAPI`, `client.Obsolete`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`UnsubscribedSet`](#unsubscribedset) | `POST /unsubscribed/set.json` | Внести email до списку відписаних. |
| [`UnsubscribedCheck`](#unsubscribedcheck) | `POST /unsubscribed/check.json` | Перевіряє, чи відписано email. |
| [`UnsubscribedList`](#unsubscribedlist) | `POST /unsubscribed/list.json` | Повертає список адрес, відписаних починаючи з указаної дати. |

## UnsubscribedSet

`POST /unsubscribed/set.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#unsubscribed-set)

Внести email до списку відписаних.

Після внесення email до списку відписаних усі майбутні спроби надіслати на нього повідомлення будуть відхилятися. Зауважте, що видалити email зі списку відписаних безпосередньо за допомогою API не можна. Але можна надіслати йому повідомлення з посиланням повторної підписки за допомогою методу [email/subscribe](https://docs.allymailer.com.ua/ua/web-api-ref#email-subscribe) або видалити його зі списку відписаних вручну на сторінці [Пошук email](https://cp.allymailer.com/ua/user/email) у панелі керування (але кількість таких видалень на день сильно обмежена).

### Запит: `ObsoleteUnsubscribedSetRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Address` | `string` | так | Email, який має бути відписано. |

### Відповідь: `ObsoleteUnsubscribedSetResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Address` | `string` | Відписаний email. |
| `Message` | `obsolete.ObsoleteUnsubscribedSetMessage` | Подробиці відписки:<br><br>- "unsubscribed" якщо email було щойно відписано,<br>- "was inactive" якщо він уже був позначений як неактивний і немає сенсу відписувати,<br>- "was unsubscribed" якщо email уже було відписано раніше,<br>- "added and unsubscribed" якщо для цього email не зберігалося жодного більш раннього статусу, тому його було додано й одразу ж відписано. (`unsubscribed`, `was inactive`, `was unsubscribed`, `added and unsubscribed`) |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/obsolete"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Obsolete.UnsubscribedSet(context.Background(), &obsolete.ObsoleteUnsubscribedSetRequest{
		Address: "email@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Address)
}
```

## UnsubscribedCheck

`POST /unsubscribed/check.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#unsubscribed-check)

Перевіряє, чи відписано email.

### Запит: `ObsoleteUnsubscribedCheckRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Address` | `string` | так | Email, який необхідно перевірити. |

### Відповідь: `ObsoleteUnsubscribedCheckResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Address` | `string` | Email, який було перевірено. |
| `IsUnsubscribed` | `bool` | true якщо відписано, false якщо ні. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/obsolete"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Obsolete.UnsubscribedCheck(context.Background(), &obsolete.ObsoleteUnsubscribedCheckRequest{
		Address: "email@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Address)
}
```

## UnsubscribedList

`POST /unsubscribed/list.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#unsubscribed-list)

Повертає список адрес, відписаних починаючи з указаної дати.

### Запит: `ObsoleteUnsubscribedListRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `DateFrom` | `string` |  | Дата у форматі РРРР-ММ-ДД, щоб отримати всі відписані адреси з "date_from" по сьогоднішній день. Якщо параметр не вказати, період за замовчуванням буде з 00:00 сьогоднішнього дня за UTC. |

### Відповідь: `ObsoleteUnsubscribedListResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Unsubscribed` | `[]obsolete.ObsoleteUnsubscribedListResponseUnsubscribedItem` | Масив об'єктів з даними відписаних. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/obsolete"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Obsolete.UnsubscribedList(context.Background(), &obsolete.ObsoleteUnsubscribedListRequest{
		DateFrom: allymailer.String("2020-10-14"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Unsubscribed)
}
```
