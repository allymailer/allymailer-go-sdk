// TagAPI: API methods of the Tag section.

package tag

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
	"github.com/allymailer/allymailer-go-sdk/common"
)

// TagAPI — Методи керування користувацькими мітками. За допомогою міток ви можете класифікувати листи за
// потрібними критеріями. Відокремити маркетингові листи від транзакційних або листи з підтвердженням
// реєстрації від сповіщень про замовлення. AllyMailer дозволяє позначати мітками кожне вихідне
// повідомлення, а потім сегментувати й аналізувати статистику на основі зазначених під час надсилання
// міток.
type TagAPI struct {
	client *client.APIClient
}

// NewTagAPI creates the Tag API.
func NewTagAPI(c *client.APIClient) *TagAPI {
	return &TagAPI{client: c}
}

// List — Метод без параметрів, повертає список усіх користувацьких міток.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#tag-list
func (a *TagAPI) List(ctx context.Context) (*TagListResponse, error) {
	response := &TagListResponse{}
	if err := a.client.Call(ctx, "/tag/list.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Видаляє мітку.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#tag-delete
func (a *TagAPI) Delete(ctx context.Context, request *TagDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/tag/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
