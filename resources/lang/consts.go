package lang

// StateSuccess 成功状态
const StateSuccess = "STATE.SUCCESS"

// StateError 错误状态
const StateError = "STATE.ERROR"

// StateFail 失败状态
const StateFail = "STATE.FAIL"

// 重复请求问题
const TooManyRequests = "REQUESTS.TOO_MANY_REQUESTS"

const CaptchaPrefix = "CAPTCHA."
const CaptchaEmptyId = CaptchaPrefix + "EMPTY_ID"
const CaptchaErrorGenerateCode = CaptchaPrefix + "ERROR_GENERATE_CODE"
const CaptchaErrorStoreCode = CaptchaPrefix + "ERROR_STORE_CODE"
const CaptchaErrorCheck = CaptchaPrefix + "ERROR_CHECK"
