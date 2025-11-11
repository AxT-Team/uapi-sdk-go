package uapi

import (
	"encoding/json"
	"fmt"
)

type UapiError struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Status  int             `json:"-"`
	Details json.RawMessage `json:"details,omitempty"`
}

func (e *UapiError) Error() string { return fmt.Sprintf("[%d] %s: %s", e.Status, e.Code, e.Message) }
type ApiErrorError struct{ UapiError }
type AvatarNotFoundError struct{ UapiError }
type ConversionFailedError struct{ UapiError }
type FileOpenErrorError struct{ UapiError }
type FileRequiredError struct{ UapiError }
type InternalServerErrorError struct{ UapiError }
type InvalidParameterError struct{ UapiError }
type InvalidParamsError struct{ UapiError }
type NotFoundError struct{ UapiError }
type NoMatchError struct{ UapiError }
type NoTrackingDataError struct{ UapiError }
type PhoneInfoFailedError struct{ UapiError }
type RecognitionFailedError struct{ UapiError }
type RequestEntityTooLargeError struct{ UapiError }
type ServiceBusyError struct{ UapiError }
type TimezoneNotFoundError struct{ UapiError }
type UnauthorizedError struct{ UapiError }
type UnsupportedCarrierError struct{ UapiError }
type UnsupportedFormatError struct{ UapiError }

func mapError(status int, body []byte) error {
	var e struct{
		Code string `json:"code"`
		Message string `json:"message"`
		Details json.RawMessage `json:"details"`
	}
	_ = json.Unmarshal(body, &e)
	if e.Code == "" { e.Code = defaultCode(status) }
	err := &UapiError{ Code: e.Code, Message: e.Message, Status: status, Details: e.Details }
	switch e.Code {
	case "API_ERROR": return &ApiErrorError{ *err }
	case "AVATAR_NOT_FOUND": return &AvatarNotFoundError{ *err }
	case "CONVERSION_FAILED": return &ConversionFailedError{ *err }
	case "FILE_OPEN_ERROR": return &FileOpenErrorError{ *err }
	case "FILE_REQUIRED": return &FileRequiredError{ *err }
	case "INTERNAL_SERVER_ERROR": return &InternalServerErrorError{ *err }
	case "INVALID_PARAMETER": return &InvalidParameterError{ *err }
	case "INVALID_PARAMS": return &InvalidParamsError{ *err }
	case "NOT_FOUND": return &NotFoundError{ *err }
	case "NO_MATCH": return &NoMatchError{ *err }
	case "NO_TRACKING_DATA": return &NoTrackingDataError{ *err }
	case "PHONE_INFO_FAILED": return &PhoneInfoFailedError{ *err }
	case "RECOGNITION_FAILED": return &RecognitionFailedError{ *err }
	case "REQUEST_ENTITY_TOO_LARGE": return &RequestEntityTooLargeError{ *err }
	case "SERVICE_BUSY": return &ServiceBusyError{ *err }
	case "TIMEZONE_NOT_FOUND": return &TimezoneNotFoundError{ *err }
	case "UNAUTHORIZED": return &UnauthorizedError{ *err }
	case "UNSUPPORTED_CARRIER": return &UnsupportedCarrierError{ *err }
	case "UNSUPPORTED_FORMAT": return &UnsupportedFormatError{ *err }
	default: return err
	}
}

func defaultCode(status int) string {
	switch status {
	case 400: return "INVALID_PARAMETER"
	case 401: return "UNAUTHORIZED"
	case 404: return "NOT_FOUND"
	case 429: return "SERVICE_BUSY"
	case 500: return "INTERNAL_SERVER_ERROR"
	default: return "API_ERROR"
	}
}
