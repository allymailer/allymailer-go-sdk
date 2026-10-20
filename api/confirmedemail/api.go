// ConfirmedEmailAPI: API methods of the ConfirmedEmail section.

package confirmedemail

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
)

// ConfirmedEmailAPI — Керування email-адресами, підтвердженими для надсилання із sandbox-домену. Методи доступні лише з
// API-ключем основного облікового запису.
type ConfirmedEmailAPI struct {
	client *client.APIClient
}

// NewConfirmedEmailAPI creates the ConfirmedEmail API.
func NewConfirmedEmailAPI(c *client.APIClient) *ConfirmedEmailAPI {
	return &ConfirmedEmailAPI{client: c}
}

// Create — Adds an email address and sends a confirmation link.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#confirmed-email-create
func (a *ConfirmedEmailAPI) Create(ctx context.Context, request *ConfirmedEmailRequest) (*ConfirmedEmailCreateResponse, error) {
	response := &ConfirmedEmailCreateResponse{}
	if err := a.client.Call(ctx, "/confirmed-email/create.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Lists confirmed and unconfirmed email addresses for the account.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#confirmed-email-list
func (a *ConfirmedEmailAPI) List(ctx context.Context) (*ConfirmedEmailListResponse, error) {
	response := &ConfirmedEmailListResponse{}
	if err := a.client.Call(ctx, "/confirmed-email/list.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Deletes an email address from the account's confirmed email list.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#confirmed-email-delete
func (a *ConfirmedEmailAPI) Delete(ctx context.Context, request *ConfirmedEmailRequest) (*ConfirmedEmailDeleteResponse, error) {
	response := &ConfirmedEmailDeleteResponse{}
	if err := a.client.Call(ctx, "/confirmed-email/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
