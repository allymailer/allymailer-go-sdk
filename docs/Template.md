# Template

[До списку розділів](../README.md)

Ви можете зберегти шаблон на боці AllyMailer і потім посилатися на нього за ідентифікатором. Це не обов'язково, але вкрай рекомендовано у разі повторюваних email-повідомлень або за великих обсягів надсилання. AllyMailer підтримує два різні синтаксиси для шаблонів, докладніше читайте [тут](https://docs.allymailer.com.ua/en/template-engines).

`template.TemplateAPI`, `client.Template`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`Set`](#set) | `POST /template/set.json` | Створити або оновити шаблон. |
| [`Get`](#get) | `POST /template/get.json` | Повертає властивості шаблону за його ідентифікатором. |
| [`List`](#list) | `POST /template/list.json` | Повертає повний список або частину списку всіх шаблонів. |
| [`Delete`](#delete) | `POST /template/delete.json` | Видаляє шаблон за ідентифікатором. |

## Set

`POST /template/set.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#template-set)

Створити або оновити шаблон.

Обмеження:
* Максимальний розмір запиту — 10 MB.
* Кількість шаблонів обмежено 10000.

### Запит: `TemplateSetRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Template` | `template.TemplateObject` | так | Об'єкт, що містить властивості шаблону. |

### Відповідь: `TemplateSetResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Template` | `template.TemplateObject` | Об'єкт, що містить властивості шаблону. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/template"
	"github.com/allymailer/allymailer-go-sdk/common"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Template.Set(context.Background(), &template.TemplateSetRequest{
		Template: &template.TemplateObject{
			Name:           "name",
			EditorType:     template.EditorTypeHTML,
			TemplateEngine: common.TemplateEngineVelocity,
			Subject:        allymailer.String("AllyMailer test email "),
			FromEmail:      "user@example.com",
			FromName:       allymailer.String("John Smith"),
			ReplyTo:        allymailer.String("email@example.com"),
			ReplyToName:    allymailer.String("John Smith"),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Template)
}
```

## Get

`POST /template/get.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#template-get)

Повертає властивості шаблону за його ідентифікатором.

### Запит: `TemplateGetRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `ID` | `string` | так | Ідентифікатор шаблону. |

### Відповідь: `TemplateGetResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Template` | `template.TemplateObject` | Об'єкт, що містить властивості шаблону. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/template"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Template.Get(context.Background(), &template.TemplateGetRequest{
		ID: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Template)
}
```

## List

`POST /template/list.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#template-list)

Повертає повний список або частину списку всіх шаблонів.

Ви можете переглядати великий список частинами, використовуючи параметри "offset" і "limit". Якщо повернуто менше ніж "limit" шаблонів — значить, ви досягли кінця списку.

### Запит: `TemplateListRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Limit` | `int` |  | Максимальна кількість шаблонів, що повертаються, за замовчуванням 50. |
| `Offset` | `int` |  | Індекс першого шаблону зі списку, що повертається, рахуючи з нуля. |

### Відповідь: `TemplateListResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Templates` | `[]template.TemplateListObject` | Масив об'єктів-шаблонів. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/template"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Template.List(context.Background(), &template.TemplateListRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Templates)
}
```

## Delete

`POST /template/delete.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#template-delete)

Видаляє шаблон за ідентифікатором.

### Запит: `TemplateDeleteRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `ID` | `string` | так | Ідентифікатор шаблону. |

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
	"github.com/allymailer/allymailer-go-sdk/api/template"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Template.Delete(context.Background(), &template.TemplateDeleteRequest{
		ID: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
```
