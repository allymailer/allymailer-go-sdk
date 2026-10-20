# ConfirmedEmail

[До списку розділів](../README.md)

Керування email-адресами, підтвердженими для надсилання із sandbox-домену. Методи доступні лише з API-ключем основного облікового запису.

`confirmedemail.ConfirmedEmailAPI`, `client.ConfirmedEmail`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`Create`](#create) | `POST /confirmed-email/create.json` | Adds an email address and sends a confirmation link. |
| [`List`](#list) | `POST /confirmed-email/list.json` | Lists confirmed and unconfirmed email addresses for the account. |
| [`Delete`](#delete) | `POST /confirmed-email/delete.json` | Deletes an email address from the account's confirmed email list. |

## Create

`POST /confirmed-email/create.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#confirmed-email-create)

Adds an email address and sends a confirmation link.

### Запит: `ConfirmedEmailRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Email` | `string` | так | Email address to add or delete. |

### Відповідь: `ConfirmedEmailCreateResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `ConfirmedEmail` | `confirmedemail.ConfirmedEmail` |  |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/confirmedemail"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.ConfirmedEmail.Create(context.Background(), &confirmedemail.ConfirmedEmailRequest{
		Email: "dev@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ConfirmedEmail)
}
```

## List

`POST /confirmed-email/list.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#confirmed-email-list)

Lists confirmed and unconfirmed email addresses for the account.

### Запит

Метод не має параметрів.

### Відповідь: `ConfirmedEmailListResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `ConfirmedEmails` | `[]confirmedemail.ConfirmedEmail` |  |

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

	response, err := client.ConfirmedEmail.List(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ConfirmedEmails)
}
```

## Delete

`POST /confirmed-email/delete.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#confirmed-email-delete)

Deletes an email address from the account's confirmed email list.

### Запит: `ConfirmedEmailRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Email` | `string` | так | Email address to add or delete. |

### Відповідь: `ConfirmedEmailDeleteResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Message` | `string` |  |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/confirmedemail"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.ConfirmedEmail.Delete(context.Background(), &confirmedemail.ConfirmedEmailRequest{
		Email: "dev@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Message)
}
```
