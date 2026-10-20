# System

[До списку розділів](../README.md)

Системні методи.

`system.SystemAPI`, `client.System`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`Ping`](#ping) | `POST /system/ping.json` | Перевіряє валідність API-ключа та повертає ідентифікатор користувача. |
| [`Info`](#info) | `POST /system/info.json` | Повертає інформацію про користувача або проєкт за API-ключем. |

## Ping

`POST /system/ping.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#system-ping)

Перевіряє валідність API-ключа та повертає ідентифікатор користувача.

### Запит

Метод не має параметрів.

### Відповідь: `SystemPingResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `UserID` | `int` | Унікальний ідентифікатор користувача. |

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

	response, err := client.System.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.UserID)
}
```

## Info

`POST /system/info.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#system-info)

Повертає інформацію про користувача або проєкт за API-ключем.

### Запит

Метод не має параметрів.

### Відповідь: `SystemInfoResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `UserID` | `int` | Унікальний ідентифікатор користувача. |
| `Email` | `string` | Email користувача. |
| `ProjectID` | `string` | Унікальний ідентифікатор проєкту, ASCII-рядок до 36 символів завдовжки. Присутній лише якщо API-ключ, використаний для запиту, був API-ключем проєкту. |
| `ProjectName` | `string` | Назва проєкту, унікальна в межах облікового запису користувача. Присутня лише якщо API-ключ, використаний для запиту, був API-ключем проєкту. |
| `ProjectAccounting` | `system.ProjectAccountingObject` | Об'єкт із властивостями лічильника листів. Присутній лише якщо API-ключ, використаний для запиту, був API-ключем проєкту. |
| `Accounting` | `system.SystemInfoResponseAccounting` | Об'єкт з властивостями періоду обліку. Присутній лише якщо API-ключ, використаний для запиту, був API-ключем користувача. |

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

	response, err := client.System.Info(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.UserID)
}
```
