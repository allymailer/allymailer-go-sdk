// SystemAPI: API methods of the System section.

package system

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
)

// SystemAPI — Системні методи.
type SystemAPI struct {
	client *client.APIClient
}

// NewSystemAPI creates the System API.
func NewSystemAPI(c *client.APIClient) *SystemAPI {
	return &SystemAPI{client: c}
}

// Ping — Перевіряє валідність API-ключа та повертає ідентифікатор користувача.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#system-ping
func (a *SystemAPI) Ping(ctx context.Context) (*SystemPingResponse, error) {
	response := &SystemPingResponse{}
	if err := a.client.Call(ctx, "/system/ping.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Info — Повертає інформацію про користувача або проєкт за API-ключем.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#system-info
func (a *SystemAPI) Info(ctx context.Context) (*SystemInfoResponse, error) {
	response := &SystemInfoResponse{}
	if err := a.client.Call(ctx, "/system/info.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
