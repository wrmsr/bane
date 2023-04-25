package schema

//

type ErrorCode string

//

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details any       `json:"details"`
}
