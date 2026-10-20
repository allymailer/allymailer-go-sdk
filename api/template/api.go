// TemplateAPI: API methods of the Template section.

package template

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
	"github.com/allymailer/allymailer-go-sdk/common"
)

// TemplateAPI — Ви можете зберегти шаблон на боці AllyMailer і потім посилатися на нього за ідентифікатором. Це не
// обов'язково, але вкрай рекомендовано у разі повторюваних email-повідомлень або за великих обсягів
// надсилання. AllyMailer підтримує два різні синтаксиси для шаблонів, докладніше читайте
// [тут](https://docs.allymailer.com.ua/en/template-engines).
type TemplateAPI struct {
	client *client.APIClient
}

// NewTemplateAPI creates the Template API.
func NewTemplateAPI(c *client.APIClient) *TemplateAPI {
	return &TemplateAPI{client: c}
}

// Set — Створити або оновити шаблон.
//
// Обмеження:
// * Максимальний розмір запиту — 10 MB.
// * Кількість шаблонів обмежено 10000.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#template-set
func (a *TemplateAPI) Set(ctx context.Context, request *TemplateSetRequest) (*TemplateSetResponse, error) {
	response := &TemplateSetResponse{}
	if err := a.client.Call(ctx, "/template/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Get — Повертає властивості шаблону за його ідентифікатором.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#template-get
func (a *TemplateAPI) Get(ctx context.Context, request *TemplateGetRequest) (*TemplateGetResponse, error) {
	response := &TemplateGetResponse{}
	if err := a.client.Call(ctx, "/template/get.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Повертає повний список або частину списку всіх шаблонів.
//
// Ви можете переглядати великий список частинами, використовуючи параметри "offset" і "limit". Якщо
// повернуто менше ніж "limit" шаблонів — значить, ви досягли кінця списку.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#template-list
func (a *TemplateAPI) List(ctx context.Context, request *TemplateListRequest) (*TemplateListResponse, error) {
	response := &TemplateListResponse{}
	if err := a.client.Call(ctx, "/template/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Видаляє шаблон за ідентифікатором.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#template-delete
func (a *TemplateAPI) Delete(ctx context.Context, request *TemplateDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/template/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
