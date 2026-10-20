// SuppressionAPI: API methods of the Suppression section.

package suppression

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
	"github.com/allymailer/allymailer-go-sdk/common"
)

// SuppressionAPI — Методи для керування адресами, надсилання на які заблоковано з різних причин — відписка,
// недоступність, скарги.
type SuppressionAPI struct {
	client *client.APIClient
}

// NewSuppressionAPI creates the Suppression API.
func NewSuppressionAPI(c *client.APIClient) *SuppressionAPI {
	return &SuppressionAPI{client: c}
}

// Set — Додає email-адресу до списку недоступних. Ви завжди зможете пізніше видалити цю адресу зі списку
// недоступних методом
// [suppression/delete](https://docs.allymailer.com.ua/ua/web-api-ref#suppression-delete).
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#suppression-set
func (a *SuppressionAPI) Set(ctx context.Context, request *SuppressionSetRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/suppression/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Get — Повертає причину й дату блокування email-адреси.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#suppression-get
func (a *SuppressionAPI) Get(ctx context.Context, request *SuppressionGetRequest) (*SuppressionGetResponse, error) {
	response := &SuppressionGetResponse{}
	if err := a.client.Call(ctx, "/suppression/get.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Повертає список недоступних, починаючи з указаної дати.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#suppression-list
func (a *SuppressionAPI) List(ctx context.Context, request *SuppressionListRequest) (*SuppressionListResponse, error) {
	response := &SuppressionListResponse{}
	if err := a.client.Call(ctx, "/suppression/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Видаляє email-адресу зі списку недоступних (лише записи з is_deletable=true). Якщо жодного запису
// для email не знайдено, повертається помилка API 3003. Якщо записи про недоступність знайдено, але
// жодну з них видалити не можна, повертається помилка API 3004. Якщо перевищено ліміт видалень на
// добу, повертається помилка API 906.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#suppression-delete
func (a *SuppressionAPI) Delete(ctx context.Context, request *SuppressionDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/suppression/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
