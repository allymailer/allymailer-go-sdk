// EmailValidationAPI: API methods of the EmailValidation section.

package emailvalidation

import (
	"context"

	"github.com/allymailer/allymailer-go-sdk/client"
)

// EmailValidationAPI — Методи перевірки email-адрес. Перевірка адреси електронної пошти надає докладну інформацію про
// працездатність email. Це дозволить виключити з надсилання неіснуючі email, знизити показник скарг,
// очистити список підписників і врешті-решт покращити вашу репутацію як відправника. Наприклад, ви
// можете інтегрувати процес перевірки у форму реєстрації вашого сайту і вже на етапі введення адреси
// повідомляти користувачам, що адреса, яку вони ввели у форму, недійсна.
type EmailValidationAPI struct {
	client *client.APIClient
}

// NewEmailValidationAPI creates the EmailValidation API.
func NewEmailValidationAPI(c *client.APIClient) *EmailValidationAPI {
	return &EmailValidationAPI{client: c}
}

// Single — Допускається не більше ніж два одночасні запити; у разі перевищення буде повернуто помилку. Оскільки
// перевірка має бути швидкою, ми обмежуємо час, що витрачається на перевірку адреси електронної пошти.
// Тому, якщо ми зіткнемося з повільним поштовим сервером, час відповіді якого перевищує 5 секунд,
// результат перевірки може бути невідомим.
//
// Документація: https://docs.allymailer.com.ua/ua/web-api-ref#email-validation-single
func (a *EmailValidationAPI) Single(ctx context.Context, request *EmailValidationSingleRequest) (*EmailValidationSingleResponse, error) {
	response := &EmailValidationSingleResponse{}
	if err := a.client.Call(ctx, "/email-validation/single.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
