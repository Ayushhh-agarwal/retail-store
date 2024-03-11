package errors

type ErrorData struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

type Public struct {
	ErrorMessage string `json:"error_message"`
}

func (err ErrorData) GetHttpCode() int {
	return err.Code
}

func (err ErrorData) Public() Public {
	return Public{
		ErrorMessage: err.Message,
	}
}

const (
	RedisLockError = "Error occurred with code : 210 and message : redislock: not obtained"
)
