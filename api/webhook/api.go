// WebhookAPI: API methods of the Webhook section.

package webhook

import (
	"context"
	"encoding/json"

	"github.com/allymailer/allymailer-go-sdk/client"
	"github.com/allymailer/allymailer-go-sdk/common"
)

// WebhookAPI — AllyMailer дозволяє вам отримувати сповіщення про зміну статусу доставки листа або спам-блокування.
// Ви надаєте URL, і в разі виникнення події, що вас цікавить, ми сповіщаємо вас, передаючи на цей URL
// JSON-дані про подію.
//
// Ви можете налаштувати формат виклику webhook (просто JSON або JSON, запакований у gzip), які саме
// події вас цікавлять, чи дозволено групувати кілька подій в одному виклику, чи потрібні вам такі
// подробиці, як SMTP-відповідь сервера або user agent отримувача тощо.
//
// Зверніть увагу, що ваш обробник має вміти приймати щонайменше 5 паралельних викликів, до 100 подій у
// кожному (але ви можете налаштувати максимально дозволену кількість паралельних викликів, збільшивши
// її до 100).
//
// Якщо URL, на який надсилається webhook, недоступний (немає відповіді HTTP 200 OK протягом 3 секунд),
// спроби надсилання webhook на цей URL до отримання очікуваного 200 OK триватимуть протягом 24 годин з
// інтервалом у 10 хвилин з додатковим параметром retry_count, значення якого збільшуватиметься на 1 з
// кожним повторним надсиланням webhook. Якщо протягом 24 годин сталося щонайменше 10 подій і жоден із
// багаторазових викликів webhook про ці події не був успішним, то ми вважаємо обробник непрацездатним
// і автоматично встановлюємо його статус у неактивний та повідомляємо про це на email, [вказаний в
// особистому кабінеті](https://cp.allymailer.com/ua/user/alerting/technical/).
//
// Формат webhook і структура кожної події описані
// [тут](https://docs.allymailer.com.ua/ua/web-api-ref#callback-format).
type WebhookAPI struct {
	client *client.APIClient
}

// NewWebhookAPI creates the Webhook API.
func NewWebhookAPI(c *client.APIClient) *WebhookAPI {
	return &WebhookAPI{client: c}
}

// Set — Встановлює або змінює властивості [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook) —
// вашого обробника сповіщень про події. У момент виклику
// [webhook/set](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-set) ми очікуємо, що до
// вказаного вами URL можна звернутися GET-запитом без параметрів і у відповідь надійде 200 OK.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#webhook-set
func (a *WebhookAPI) Set(ctx context.Context, request *WebhookSetRequest) (*WebhookSetResponse, error) {
	response := &WebhookSetResponse{}
	if err := a.client.Call(ctx, "/webhook/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Get — Повертає властивості [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook)
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#webhook-get
func (a *WebhookAPI) Get(ctx context.Context, request *WebhookGetRequest) (*WebhookGetResponse, error) {
	response := &WebhookGetResponse{}
	if err := a.client.Call(ctx, "/webhook/get.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Повернути список усіх або кількох [webhook](https://docs.allymailer.com.ua/ua/web-api-ref#webhook)
// (обробників сповіщень про події) користувача або проєкту.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#webhook-list
func (a *WebhookAPI) List(ctx context.Context, request *WebhookListRequest) (*WebhookListResponse, error) {
	response := &WebhookListResponse{}
	if err := a.client.Call(ctx, "/webhook/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Видалити webhook.
//
// Ви також можете тимчасово вимкнути webhook без його видалення, встановивши його поле status у
// "disabled" за допомогою методу
// [webhook/set](https://docs.allymailer.com.ua/ua/web-api-ref#webhook-set).
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#webhook-delete
func (a *WebhookAPI) Delete(ctx context.Context, request *WebhookDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/webhook/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ParseCallback — Опис формату webhook.
func ParseCallback(body []byte) (*WebhookCallback, error) {
	callback := &WebhookCallback{}
	if err := json.Unmarshal(body, callback); err != nil {
		return nil, err
	}
	return callback, nil
}
