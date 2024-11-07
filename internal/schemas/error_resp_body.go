package schemas



type ErrorResponseBody struct {
	Message  string      `json:"message"`
	MetaData interface{} `json:"metaData"`
}

type AuthResponse struct {
	AccessToken                 *string `json:"accessToken,omitempty"`
	RefreshToken                *string `json:"refreshToken,omitempty"`
	AccessExpiresAtEpoach       int64   `json:"atExpiresIn,omitempty"`
	RefreshTokenExpiresAtEpoach int64   `json:"rtExpiresIn"`
	Message                     string  `json:"message"`
}
