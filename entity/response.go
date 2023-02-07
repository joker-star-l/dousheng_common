package common

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

const StatusSuccess = 0
const StatusError = 1

func SuccessResponse() Response {
	return Response{StatusSuccess, "success"}
}

func ErrorResponse(msg string) Response {
	return Response{StatusError, msg}
}
