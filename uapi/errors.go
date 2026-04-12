package uapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

type RateLimitPolicyEntry struct {
	Name          string `json:"name"`
	Quota         *int64 `json:"quota,omitempty"`
	Unit          string `json:"unit,omitempty"`
	WindowSeconds *int   `json:"window_seconds,omitempty"`
}

type RateLimitStateEntry struct {
	Name              string `json:"name"`
	Remaining         *int64 `json:"remaining,omitempty"`
	Unit              string `json:"unit,omitempty"`
	ResetAfterSeconds *int   `json:"reset_after_seconds,omitempty"`
}

type ResponseMeta struct {
	RequestID                  string                           `json:"request_id,omitempty"`
	RetryAfterSeconds          *int                             `json:"retry_after_seconds,omitempty"`
	DebitStatus                string                           `json:"debit_status,omitempty"`
	CreditsRequested           *int64                           `json:"credits_requested,omitempty"`
	CreditsCharged             *int64                           `json:"credits_charged,omitempty"`
	CreditsPricing             string                           `json:"credits_pricing,omitempty"`
	ActiveQuotaBuckets         *int                             `json:"active_quota_buckets,omitempty"`
	StopOnEmpty                *bool                            `json:"stop_on_empty,omitempty"`
	RateLimitPolicyRaw         string                           `json:"rate_limit_policy_raw,omitempty"`
	RateLimitRaw               string                           `json:"rate_limit_raw,omitempty"`
	RateLimitPolicies          map[string]RateLimitPolicyEntry  `json:"rate_limit_policies,omitempty"`
	RateLimits                 map[string]RateLimitStateEntry   `json:"rate_limits,omitempty"`
	BalanceLimitCents          *int64                           `json:"balance_limit_cents,omitempty"`
	BalanceRemainingCents      *int64                           `json:"balance_remaining_cents,omitempty"`
	QuotaLimitCredits          *int64                           `json:"quota_limit_credits,omitempty"`
	QuotaRemainingCredits      *int64                           `json:"quota_remaining_credits,omitempty"`
	VisitorQuotaLimitCredits   *int64                           `json:"visitor_quota_limit_credits,omitempty"`
	VisitorQuotaRemainingCredits *int64                          `json:"visitor_quota_remaining_credits,omitempty"`
	RawHeaders                 map[string]string                `json:"raw_headers,omitempty"`
}

type UapiError struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Status  int             `json:"-"`
	Details json.RawMessage `json:"details,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
	Meta    *ResponseMeta   `json:"meta,omitempty"`
}

func (e *UapiError) Error() string { return fmt.Sprintf("[%d] %s: %s", e.Status, e.Code, e.Message) }
type ApiErrorError struct{ UapiError }
type AvatarNotFoundError struct{ UapiError }
type ConversionFailedError struct{ UapiError }
type FileOpenErrorError struct{ UapiError }
type FileRequiredError struct{ UapiError }
type InsufficientCreditsError struct{ UapiError }
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
type VisitorMonthlyQuotaExhaustedError struct{ UapiError }

func mapError(status int, body []byte, headers *fasthttp.ResponseHeader) error {
	var e struct{
		Code string `json:"code"`
		Error string `json:"error"`
		Message string `json:"message"`
		Details json.RawMessage `json:"details"`
		Quota json.RawMessage `json:"quota"`
		Docs json.RawMessage `json:"docs"`
	}
	_ = json.Unmarshal(body, &e)
	code := strings.TrimSpace(e.Code)
	if code == "" {
		code = strings.TrimSpace(e.Error)
	}
	if code == "" {
		code = defaultCode(status)
	}
	message := strings.TrimSpace(e.Message)
	if message == "" {
		message = strings.TrimSpace(string(body))
	}
	details := e.Details
	if len(details) == 0 {
		if len(e.Quota) > 0 {
			details = e.Quota
		} else if len(e.Docs) > 0 {
			details = e.Docs
		}
	}
	meta := extractMetaFromHeaders(headers)
	err := &UapiError{
		Code: code,
		Message: message,
		Status: status,
		Details: append(json.RawMessage(nil), details...),
		Payload: append(json.RawMessage(nil), body...),
		Meta: meta,
	}
	switch code {
	case "API_ERROR": return &ApiErrorError{ *err }
	case "AVATAR_NOT_FOUND": return &AvatarNotFoundError{ *err }
	case "CONVERSION_FAILED": return &ConversionFailedError{ *err }
	case "FILE_OPEN_ERROR": return &FileOpenErrorError{ *err }
	case "FILE_REQUIRED": return &FileRequiredError{ *err }
	case "INSUFFICIENT_CREDITS": return &InsufficientCreditsError{ *err }
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
	case "VISITOR_MONTHLY_QUOTA_EXHAUSTED": return &VisitorMonthlyQuotaExhaustedError{ *err }
	default: return err
	}
}

func defaultCode(status int) string {
	switch status {
	case 400: return "INVALID_PARAMETER"
	case 401: return "UNAUTHORIZED"
	case 402: return "INSUFFICIENT_CREDITS"
	case 404: return "NOT_FOUND"
	case 429: return "SERVICE_BUSY"
	case 500: return "INTERNAL_SERVER_ERROR"
	default: return "API_ERROR"
	}
}

func extractMetaFromHeaders(headers *fasthttp.ResponseHeader) *ResponseMeta {
	if headers == nil {
		return &ResponseMeta{}
	}
	raw := map[string]string{}
	headers.VisitAll(func(k, v []byte) {
		raw[strings.ToLower(string(k))] = string(v)
	})
	meta := &ResponseMeta{
		RequestID:          raw["x-request-id"],
		RetryAfterSeconds:  parseInt(raw["retry-after"]),
		DebitStatus:        raw["uapi-debit-status"],
		CreditsRequested:   parseInt64(raw["uapi-credits-requested"]),
		CreditsCharged:     parseInt64(raw["uapi-credits-charged"]),
		CreditsPricing:     raw["uapi-credits-pricing"],
		ActiveQuotaBuckets: parseInt(raw["uapi-quota-active-buckets"]),
		StopOnEmpty:        parseBool(raw["uapi-stop-on-empty"]),
		RateLimitPolicyRaw: raw["ratelimit-policy"],
		RateLimitRaw:       raw["ratelimit"],
		RateLimitPolicies:  map[string]RateLimitPolicyEntry{},
		RateLimits:         map[string]RateLimitStateEntry{},
		RawHeaders:         raw,
	}
	for _, item := range parseStructuredItems(raw["ratelimit-policy"]) {
		entry := RateLimitPolicyEntry{
			Name: item.Name,
			Quota: parseInt64(item.Params["q"]),
			Unit: item.Params["uapi-unit"],
			WindowSeconds: parseInt(item.Params["w"]),
		}
		meta.RateLimitPolicies[item.Name] = entry
	}
	for _, item := range parseStructuredItems(raw["ratelimit"]) {
		entry := RateLimitStateEntry{
			Name: item.Name,
			Remaining: parseInt64(item.Params["r"]),
			Unit: item.Params["uapi-unit"],
			ResetAfterSeconds: parseInt(item.Params["t"]),
		}
		meta.RateLimits[item.Name] = entry
	}
	if entry, ok := meta.RateLimitPolicies["billing-balance"]; ok {
		meta.BalanceLimitCents = entry.Quota
	}
	if entry, ok := meta.RateLimits["billing-balance"]; ok {
		meta.BalanceRemainingCents = entry.Remaining
	}
	if entry, ok := meta.RateLimitPolicies["billing-quota"]; ok {
		meta.QuotaLimitCredits = entry.Quota
	}
	if entry, ok := meta.RateLimits["billing-quota"]; ok {
		meta.QuotaRemainingCredits = entry.Remaining
	}
	if entry, ok := meta.RateLimitPolicies["visitor-quota"]; ok {
		meta.VisitorQuotaLimitCredits = entry.Quota
	}
	if entry, ok := meta.RateLimits["visitor-quota"]; ok {
		meta.VisitorQuotaRemainingCredits = entry.Remaining
	}
	return meta
}

type structuredItem struct {
	Name   string
	Params map[string]string
}

func parseStructuredItems(raw string) []structuredItem {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	items := make([]structuredItem, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		segments := strings.Split(part, ";")
		if len(segments) == 0 {
			continue
		}
		item := structuredItem{
			Name:   unquote(segments[0]),
			Params: map[string]string{},
		}
		for _, segment := range segments[1:] {
			segment = strings.TrimSpace(segment)
			if segment == "" {
				continue
			}
			key, value, ok := strings.Cut(segment, "=")
			if !ok {
				continue
			}
			item.Params[strings.TrimSpace(key)] = unquote(value)
		}
		items = append(items, item)
	}
	return items
}

func unquote(value string) string {
	value = strings.TrimSpace(value)
	if len(value) >= 2 && strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
		return value[1 : len(value)-1]
	}
	return value
}

func parseInt(value string) *int {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}
	return &n
}

func parseInt64(value string) *int64 {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}
	return &n
}

func parseBool(value string) *bool {
	value = strings.TrimSpace(strings.ToLower(value))
	switch value {
	case "true":
		v := true
		return &v
	case "false":
		v := false
		return &v
	default:
		return nil
	}
}
