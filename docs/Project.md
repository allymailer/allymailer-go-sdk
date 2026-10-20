# Project

[До списку розділів](../README.md)

Проєкти дозволяють об'єднати кілька незалежних доменів із webhook, шаблонами та списками відписки. Це дуже корисно, якщо у вас є власні клієнти і ви керуєте їхніми даними з одного облікового запису AllyMailer. Кожен проєкт отримує власний ключ API, а виклики методів API з цим ключем обробляють лише дані цього проєкту. Таким чином, ви можете призначити окремий проєкт для кожного користувача і бути впевненими, що вони не конфліктуватимуть: наприклад, відписка від одного проєкту не вплине на інший.

Методи проєктів за замовчуванням вимкнено через заходи захисту від спаму. Щоб їх увімкнути, вам необхідно звернутися до [служби підтримки](https://cp.allymailer.com/ua/support).

Зверніть увагу: за замовчуванням домени для надсилання підтверджуються та використовуються незалежно для основного облікового запису й кожного з проєктів. Але є можливість увімкнути для проєктів використання доменів, підтверджених в основному обліковому записі. Просто попросіть про це [техпідтримку](https://cp.allymailer.com/ua/support).

`project.ProjectAPI`, `client.Project`

| Метод | Ендпоінт | Опис |
|---|---|---|
| [`Create`](#create) | `POST /project/create.json` | Створити новий проєкт. |
| [`Update`](#update) | `POST /project/update.json` | Змінити властивості проєкту. |
| [`List`](#list) | `POST /project/list.json` | Повертає список проєктів. |
| [`Delete`](#delete) | `POST /project/delete.json` | Видалити проєкт. |

## Create

`POST /project/create.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#project-create)

Створити новий проєкт.

### Запит: `ProjectCreateRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `Project` | `project.ProjectObject` | так | Об'єкт із властивостями проєкту. |

### Відповідь: `ProjectCreateResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `ProjectID` | `string` | Унікальний ідентифікатор проєкту, ASCII-рядок до 36 символів завдовжки. |
| `ProjectAPIKey` | `string` | API-ключ проєкту. Ви можете використовувати цей ключ замість користувацького api_key у всіх методах, крім project-методів. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/project"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Project.Create(context.Background(), &project.ProjectCreateRequest{
		Project: &project.ProjectObject{
			Name:              "Project 1A",
			Country:           allymailer.String("ES"),
			BackendDomainID:   allymailer.Int(1234),
			UnsubscribePageID: allymailer.Int(1),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ProjectID)
}
```

## Update

`POST /project/update.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#project-update)

Змінити властивості проєкту.

### Запит: `ProjectUpdateRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `ProjectID` | `string` |  | Ідентифікатор проєкту, що оновлюється. Проєкт має бути ідентифіковано або цим полем project_id, або полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше. |
| `ProjectAPIKey` | `string` |  | Необов'язковий API-ключ проєкту, що оновлюється. Проєкт має бути ідентифіковано або полем project_id, або полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше. |
| `Project` | `project.ProjectUpdateObject` | так | Об'єкт із властивостями проєкту. |

### Відповідь: `ProjectUpdateResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `ProjectAPIKey` | `string` | API-ключ проєкту. Повертається лише в разі, якщо в параметрі запиту був project_api_key. Якщо ж проєкт було ідентифіковано через project_id, то не повертається. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/project"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Project.Update(context.Background(), &project.ProjectUpdateRequest{
		ProjectID:     allymailer.String("1234567890"),
		ProjectAPIKey: allymailer.String("your api_key"),
		Project: &project.ProjectUpdateObject{
			Name:              "Project 1A",
			Country:           allymailer.String("ES"),
			BackendDomainID:   allymailer.Int(1234),
			UnsubscribePageID: allymailer.Int(1),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ProjectAPIKey)
}
```

## List

`POST /project/list.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#project-list)

Повертає список проєктів.

### Запит: `ProjectListRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `ProjectID` | `string` |  | Опціональна можливість указати project_id для включення до списку лише вказаного проєкту. |
| `ProjectAPIKey` | `string` |  | Опціональна можливість указати project_api_key для включення до списку лише вказаного проєкту. |

### Відповідь: `ProjectListResponse`

| Поле | Тип | Опис |
|---|---|---|
| `Status` | `string` | Рядок "success". (`success`) |
| `Projects` | `[]project.ProjectListObject` | Масив об'єктів, де кожен об'єкт відповідає окремому проєкту. |

### Приклад

```go
package main

import (
	"context"
	"fmt"

	allymailer "github.com/allymailer/allymailer-go-sdk"
	"github.com/allymailer/allymailer-go-sdk/api/project"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Project.List(context.Background(), &project.ProjectListRequest{
		ProjectID:     allymailer.String("1234567890"),
		ProjectAPIKey: allymailer.String("your api_key"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Projects)
}
```

## Delete

`POST /project/delete.json` · [Документація](https://docs.allymailer.com.ua/ua/web-api-ref#project-delete)

Видалити проєкт.

Деякий час проєкт не видаляється остаточно, і його ще можна відновити — для цього необхідно звернутися в техпідтримку.

### Запит: `ProjectDeleteRequest`

| Поле | Тип | Обов'язкове | Опис |
|---|---|---|---|
| `ProjectID` | `string` |  | Ідентифікатор проєкту, що видаляється. Проєкт має бути ідентифіковано або цим полем project_id, або полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше. |
| `ProjectAPIKey` | `string` |  | Необов'язковий API-ключ проєкту, що видаляється. Проєкт має бути ідентифіковано або полем project_id, або полем project_api_key. Ми рекомендуємо використовувати project_id, це безпечніше. |

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
	"github.com/allymailer/allymailer-go-sdk/api/project"
)

func main() {
	client := allymailer.NewClient("YOUR_API_KEY")

	response, err := client.Project.Delete(context.Background(), &project.ProjectDeleteRequest{
		ProjectID:     allymailer.String("1234567890"),
		ProjectAPIKey: allymailer.String("your api_key"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
```
