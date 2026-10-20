// Models of the Tag section.

package tag

// TagListResponse
type TagListResponse struct {
	// Рядок "success".
	// Допустимі значення: `success`.
	Status string `json:"status,omitempty"`

	// Масив об'єктів з даними користувацьких міток.
	Tags []TagsResponseObject `json:"tags,omitempty"`
}

// TagsResponseObject — Об'єкт із властивостями користувацьких міток.
type TagsResponseObject struct {
	// Унікальний id мітки.
	TagID int `json:"tag_id,omitempty"`

	// Назва мітки.
	Tag string `json:"tag,omitempty"`
}

// TagDeleteRequest
type TagDeleteRequest struct {
	// Унікальний id мітки.
	TagID int `json:"tag_id"`
}
