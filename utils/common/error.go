package common

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (err AppError) Error() string {
	return err.Message
}