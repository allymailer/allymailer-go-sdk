// Package allymailer is the AllyMailer Web API client.
package allymailer

import (
	"github.com/allymailer/allymailer-go-sdk/api/confirmedemail"
	"github.com/allymailer/allymailer-go-sdk/api/domain"
	"github.com/allymailer/allymailer-go-sdk/api/email"
	"github.com/allymailer/allymailer-go-sdk/api/emailvalidation"
	"github.com/allymailer/allymailer-go-sdk/api/eventdump"
	"github.com/allymailer/allymailer-go-sdk/api/obsolete"
	"github.com/allymailer/allymailer-go-sdk/api/project"
	"github.com/allymailer/allymailer-go-sdk/api/suppression"
	"github.com/allymailer/allymailer-go-sdk/api/system"
	"github.com/allymailer/allymailer-go-sdk/api/tag"
	"github.com/allymailer/allymailer-go-sdk/api/template"
	"github.com/allymailer/allymailer-go-sdk/api/webhook"
	"github.com/allymailer/allymailer-go-sdk/client"
)

// Client is the entry point of the AllyMailer API: one field per section.
type Client struct {
	// APIClient is the low level transport.
	APIClient *client.APIClient

	// Email — Ця група методів пов'язана з надсиланням email.
	Email *email.EmailAPI

	// EmailValidation — Методи перевірки email-адрес. Перевірка адреси електронної пошти надає докладну інформацію про
	// працездатність email. Це дозволить виключити з надсилання неіснуючі email, знизити показник скарг,
	// очистити список підписників і врешті-решт покращити вашу репутацію як відправника. Наприклад, ви
	// можете інтегрувати процес перевірки у форму реєстрації вашого сайту і вже на етапі введення адреси
	// повідомляти користувачам, що адреса, яку вони ввели у форму, недійсна.
	EmailValidation *emailvalidation.EmailValidationAPI

	// ConfirmedEmail — Керування email-адресами, підтвердженими для надсилання із sandbox-домену. Методи доступні лише з
	// API-ключем основного облікового запису.
	ConfirmedEmail *confirmedemail.ConfirmedEmailAPI

	// Template — Ви можете зберегти шаблон на боці AllyMailer і потім посилатися на нього за ідентифікатором. Це не
	// обов'язково, але вкрай рекомендовано у разі повторюваних email-повідомлень або за великих обсягів
	// надсилання. AllyMailer підтримує два різні синтаксиси для шаблонів, докладніше читайте
	// [тут](https://docs.allymailer.com.ua/en/template-engines).
	Template *template.TemplateAPI

	// Webhook — AllyMailer дозволяє вам отримувати сповіщення про зміну статусу доставки листа або спам-блокування.
	// Ви надаєте URL, і в разі виникнення події, що вас цікавить, ми сповіщаємо вас, передаючи на цей URL
	// JSON-дані про подію.
	Webhook *webhook.WebhookAPI

	// Suppression — Методи для керування адресами, надсилання на які заблоковано з різних причин — відписка,
	// недоступність, скарги.
	Suppression *suppression.SuppressionAPI

	// Domain — Методи для налаштування доменів надсилання та перевірки їх статусів. Корисні, якщо вам недостатньо
	// додавання доменів через вебінтерфейс AllyMailer і, наприклад, хочеться автоматизувати додавання
	// доменів для ваших власних клієнтів.
	Domain *domain.DomainAPI

	// EventDump — Методи для отримання списку подій, що сталися, за певний період. Дані зберігаються до 45 днів,
	// залежно від вашого тарифу. Це означає, що можна запросити події, починаючи з 00:00:00 45-го дня тому
	// і до поточного моменту часу. Для індивідуальних тарифів з великими обсягами за замовчуванням події
	// доступні лише за останні 7 днів.
	EventDump *eventdump.EventDumpAPI

	// Tag — Методи керування користувацькими мітками. За допомогою міток ви можете класифікувати листи за
	// потрібними критеріями. Відокремити маркетингові листи від транзакційних або листи з підтвердженням
	// реєстрації від сповіщень про замовлення. AllyMailer дозволяє позначати мітками кожне вихідне
	// повідомлення, а потім сегментувати й аналізувати статистику на основі зазначених під час надсилання
	// міток.
	Tag *tag.TagAPI

	// Project — Проєкти дозволяють об'єднати кілька незалежних доменів із webhook, шаблонами та списками відписки.
	// Це дуже корисно, якщо у вас є власні клієнти і ви керуєте їхніми даними з одного облікового запису
	// AllyMailer. Кожен проєкт отримує власний ключ API, а виклики методів API з цим ключем обробляють
	// лише дані цього проєкту. Таким чином, ви можете призначити окремий проєкт для кожного користувача і
	// бути впевненими, що вони не конфліктуватимуть: наприклад, відписка від одного проєкту не вплине на
	// інший.
	Project *project.ProjectAPI

	// System — Системні методи.
	System *system.SystemAPI

	// Obsolete — Застарілі методи роботи з відписками. Ми рекомендуємо користуватися замість них гнучкішими методами
	// [suppression](https://docs.allymailer.com.ua/ua/web-api-ref#suppression). Зверніть увагу, що знову
	// підписати відписаного адресата через WebAPI не можна. Можна лише надіслати йому email із посиланням
	// повторної підписки за допомогою методу email/subscribe.
	Obsolete *obsolete.ObsoleteAPI
}

// API hosts for NewClient, e.g. NewClient("YOUR_API_KEY", HostCP).
const (
	HostCP = client.HostCP
)

// NewClient returns a client with the default configuration. The optional host is an API host
// (HostCP, "cp.allymailer.com") or a full base URL.
func NewClient(apiKey string, host ...string) *Client {
	config := client.NewConfiguration(apiKey)
	if len(host) > 0 && host[0] != "" {
		config.BaseURL = client.BaseURLFor(host[0])
	}
	return NewClientWithConfig(config)
}

// NewClientWithConfig returns a client with a custom configuration.
func NewClientWithConfig(config *client.Configuration) *Client {
	c := client.NewAPIClient(config)
	return &Client{
		APIClient:       c,
		Email:           email.NewEmailAPI(c),
		EmailValidation: emailvalidation.NewEmailValidationAPI(c),
		ConfirmedEmail:  confirmedemail.NewConfirmedEmailAPI(c),
		Template:        template.NewTemplateAPI(c),
		Webhook:         webhook.NewWebhookAPI(c),
		Suppression:     suppression.NewSuppressionAPI(c),
		Domain:          domain.NewDomainAPI(c),
		EventDump:       eventdump.NewEventDumpAPI(c),
		Tag:             tag.NewTagAPI(c),
		Project:         project.NewProjectAPI(c),
		System:          system.NewSystemAPI(c),
		Obsolete:        obsolete.NewObsoleteAPI(c),
	}
}

// String returns a pointer to v, for optional fields of the requests.
func String(v string) *string {
	return &v
}

// Int returns a pointer to v, for optional fields of the requests.
func Int(v int) *int {
	return &v
}

// Float64 returns a pointer to v, for optional fields of the requests.
func Float64(v float64) *float64 {
	return &v
}

// Bool returns a pointer to v, for optional fields of the requests.
func Bool(v bool) *bool {
	return &v
}
