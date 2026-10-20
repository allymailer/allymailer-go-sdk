// ObsoleteAPI: API methods of the Obsolete section.

package obsolete

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
)

// ObsoleteAPI — Застарілі методи роботи з відписками. Ми рекомендуємо користуватися замість них гнучкішими методами
// [suppression](https://docs.allymailer.com.ua/ua/web-api-ref#suppression). Зверніть увагу, що знову
// підписати відписаного адресата через WebAPI не можна. Можна лише надіслати йому email із посиланням
// повторної підписки за допомогою методу email/subscribe.
type ObsoleteAPI struct {
	client *client.APIClient
}

// NewObsoleteAPI creates the Obsolete API.
func NewObsoleteAPI(c *client.APIClient) *ObsoleteAPI {
	return &ObsoleteAPI{client: c}
}

// UnsubscribedSet — Внести email до списку відписаних.
//
// Після внесення email до списку відписаних усі майбутні спроби надіслати на нього повідомлення будуть
// відхилятися. Зауважте, що видалити email зі списку відписаних безпосередньо за допомогою API не
// можна. Але можна надіслати йому повідомлення з посиланням повторної підписки за допомогою методу
// [email/subscribe](https://docs.allymailer.com.ua/ua/web-api-ref#email-subscribe) або видалити його
// зі списку відписаних вручну на сторінці [Пошук email](https://cp.allymailer.com/ua/user/email) у
// панелі керування (але кількість таких видалень на день сильно обмежена).
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#unsubscribed-set
func (a *ObsoleteAPI) UnsubscribedSet(ctx context.Context, request *ObsoleteUnsubscribedSetRequest) (*ObsoleteUnsubscribedSetResponse, error) {
	response := &ObsoleteUnsubscribedSetResponse{}
	if err := a.client.Call(ctx, "/unsubscribed/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UnsubscribedCheck — Перевіряє, чи відписано email.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#unsubscribed-check
func (a *ObsoleteAPI) UnsubscribedCheck(ctx context.Context, request *ObsoleteUnsubscribedCheckRequest) (*ObsoleteUnsubscribedCheckResponse, error) {
	response := &ObsoleteUnsubscribedCheckResponse{}
	if err := a.client.Call(ctx, "/unsubscribed/check.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UnsubscribedList — Повертає список адрес, відписаних починаючи з указаної дати.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#unsubscribed-list
func (a *ObsoleteAPI) UnsubscribedList(ctx context.Context, request *ObsoleteUnsubscribedListRequest) (*ObsoleteUnsubscribedListResponse, error) {
	response := &ObsoleteUnsubscribedListResponse{}
	if err := a.client.Call(ctx, "/unsubscribed/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
