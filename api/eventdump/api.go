// EventDumpAPI: API methods of the EventDump section.

package eventdump

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
	"github.com/allymailer/allymailer-go-sdk/common"
)

// EventDumpAPI — Методи для отримання списку подій, що сталися, за певний період. Дані зберігаються до 45 днів,
// залежно від вашого тарифу. Це означає, що можна запросити події, починаючи з 00:00:00 45-го дня тому
// і до поточного моменту часу. Для індивідуальних тарифів з великими обсягами за замовчуванням події
// доступні лише за останні 7 днів.
type EventDumpAPI struct {
	client *client.APIClient
}

// NewEventDumpAPI creates the EventDump API.
func NewEventDumpAPI(c *client.APIClient) *EventDumpAPI {
	return &EventDumpAPI{client: c}
}

// Create — Асинхронний метод, який ініціює підготовку для завантаження файлів дампа у форматі CSV. Максимальна
// кількість збережених/створюваних дампів обмежена десятьма. У разі досягнення ліміту для запиту на
// підготовку нового дампа потрібно видалити старі. Час зберігання сформованого дампа — 8 годин, після
// чого він автоматично видаляється. Якщо період запиту даних більший за добу або подій багато, буде
// сформовано декілька файлів, у кожному з яких буде не більше ніж 100 000 подій. Файл містить колонки
// з такою інформацією — 'event_time', 'job_id', 'from_email', 'email', 'status', 'delivery_status',
// 'metadata', 'destination_response', 'user_agent', 'url', 'ip', 'tags' і 'project_id' (якщо у
// користувача увімкнено проєкти). Ви можете відфільтрувати поля та змінити їх порядок, а також додати
// поле 'subject' за допомогою параметра 'dump_fileds'.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-create
func (a *EventDumpAPI) Create(ctx context.Context, request *EventDumpCreateRequest) (*EventDumpCreateResponse, error) {
	response := &EventDumpCreateResponse{}
	if err := a.client.Call(ctx, "/event-dump/create.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Get — Повертає властивості дампа за його ідентифікатором.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-get
func (a *EventDumpAPI) Get(ctx context.Context, request *EventDumpGetRequest) (*EventDumpGetResponse, error) {
	response := &EventDumpGetResponse{}
	if err := a.client.Call(ctx, "/event-dump/get.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Метод без параметрів, повертає список усіх дампів подій.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-list
func (a *EventDumpAPI) List(ctx context.Context) (*EventDumpListResponse, error) {
	response := &EventDumpListResponse{}
	if err := a.client.Call(ctx, "/event-dump/list.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Видаляє дамп із черги та/або зі сховища.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#event-dump-delete
func (a *EventDumpAPI) Delete(ctx context.Context, request *EventDumpDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/event-dump/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
