package lang

// StateSuccess 成功状态
const StateSuccess = "STATE.SUCCESS"

// StateError 错误状态
const StateError = "STATE.ERROR"

// StateFail 失败状态
const StateFail = "STATE.FAIL"

// TooManyRequests 重复请求问题
const TooManyRequests = "REQUESTS.TOO_MANY_REQUESTS"

const CaptchaEmpty = "CAPTCHA.EMPTY"
const CaptchaGenerate = "CAPTCHA.GENERATE"
const CaptchaEmptyId = "CAPTCHA.EMPTY_ID"
const CaptchaErrorGenerateCode = "CAPTCHA.ERROR_GENERATE_CODE"
const CaptchaErrorStoreCode = "CAPTCHA.ERROR_STORE_CODE"
const CaptchaErrorCheck = "CAPTCHA.ERROR_CHECK"

// EmptyEmail   空账号
const EmptyEmail = "AUTHORIZE.EMPTY_EMAIL"
const ErrorPassword = "AUTHORIZE.ERROR_PASSWORD"
const ErrorSingPassword = "AUTHORIZE.ERROR_Sing_PASSWORD"

// NotExistsEmail 账号不存在
const NotExistsEmail = "AUTHORIZE.NOT_EXISTS_EMAIL"

// ExistsEmail  账号不存在
const ExistsEmail = "AUTHORIZE.EXISTS_EMAIL"

// ErrorFormatEmail 格式错误
const ErrorFormatEmail = "AUTHORIZE.ERROR_FORMAT_EMAIL"
