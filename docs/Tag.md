# Tag

[До списку розділів](../README.md)

Методи керування користувацькими мітками. За допомогою міток ви можете класифікувати листи за потрібними критеріями. Відокремити маркетингові листи від транзакційних або листи з підтвердженням реєстрації від сповіщень про замовлення. AllyMailer дозволяє позначати мітками кожне вихідне повідомлення, а потім сегментувати й аналізувати статистику на основі зазначених під час надсилання міток.

`tag.TagAPI`, `client.Tag`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`List`](#list) | `POST /tag/list.json` | Метод без параметрів, повертає список усіх користувацьких міток. |
| [`Delete`](#delete) | `POST /tag/delete.json` | Видаляє мітку. |

## List

`POST /tag/list.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#tag-list)

Метод без параметрів, повертає список усіх користувацьких міток.

### Запит

Метод не має параметрів.

### Відповідь: `TagListResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Tags` | `[]tag.TagsResponseObject` | Масив об'єктів з даними користувацьких міток. |

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

	response, err := client.Tag.List(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Tags)
}
```

## Delete

`POST /tag/delete.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#tag-delete)

Видаляє мітку.

### Запит: `TagDeleteRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `TagID` | `int` | так | Унікальний id мітки. |

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
	"github.com/allymailer/allymailer-go-sdk/api/tag"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Tag.Delete(context.Background(), &tag.TagDeleteRequest{
		TagID: 54,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
```
