// ProjectAPI: API methods of the Project section.

package project

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
	"github.com/allymailer/allymailer-go-sdk/common"
)

// ProjectAPI — Проєкти дозволяють об'єднати кілька незалежних доменів із webhook, шаблонами та списками відписки.
// Це дуже корисно, якщо у вас є власні клієнти і ви керуєте їхніми даними з одного облікового запису
// AllyMailer. Кожен проєкт отримує власний ключ API, а виклики методів API з цим ключем обробляють
// лише дані цього проєкту. Таким чином, ви можете призначити окремий проєкт для кожного користувача і
// бути впевненими, що вони не конфліктуватимуть: наприклад, відписка від одного проєкту не вплине на
// інший.
//
// Методи проєктів за замовчуванням вимкнено через заходи захисту від спаму. Щоб їх увімкнути, вам
// необхідно звернутися до [служби підтримки](https://cp.allymailer.com/ua/support).
//
// Зверніть увагу: за замовчуванням домени для надсилання підтверджуються та використовуються незалежно
// для основного облікового запису й кожного з проєктів. Але є можливість увімкнути для проєктів
// використання доменів, підтверджених в основному обліковому записі. Просто попросіть про це
// [техпідтримку](https://cp.allymailer.com/ua/support).
type ProjectAPI struct {
	client *client.APIClient
}

// NewProjectAPI creates the Project API.
func NewProjectAPI(c *client.APIClient) *ProjectAPI {
	return &ProjectAPI{client: c}
}

// Create — Створити новий проєкт.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#project-create
func (a *ProjectAPI) Create(ctx context.Context, request *ProjectCreateRequest) (*ProjectCreateResponse, error) {
	response := &ProjectCreateResponse{}
	if err := a.client.Call(ctx, "/project/create.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Update — Змінити властивості проєкту.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#project-update
func (a *ProjectAPI) Update(ctx context.Context, request *ProjectUpdateRequest) (*ProjectUpdateResponse, error) {
	response := &ProjectUpdateResponse{}
	if err := a.client.Call(ctx, "/project/update.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Повертає список проєктів.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#project-list
func (a *ProjectAPI) List(ctx context.Context, request *ProjectListRequest) (*ProjectListResponse, error) {
	response := &ProjectListResponse{}
	if err := a.client.Call(ctx, "/project/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Видалити проєкт.
//
// Деякий час проєкт не видаляється остаточно, і його ще можна відновити — для цього необхідно
// звернутися в техпідтримку.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#project-delete
func (a *ProjectAPI) Delete(ctx context.Context, request *ProjectDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/project/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
