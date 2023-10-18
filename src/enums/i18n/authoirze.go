package i18n

const AuthorizePrefix = "AUTHORIZE."

// EmptyAccount AuthorizeEmptyAccount 空账号
const EmptyAccount = AuthorizePrefix + "EMPTY_ACCOUNT"

// EmptyPassword AuthorizeEmptyPassword AuthorizeEmptyAccount 空密码
const EmptyPassword = AuthorizePrefix + "EMPTY_PASSWORD"
const ErrorPassword = AuthorizePrefix + "ERROR_PASSWORD"
const ErrorSingPassword = AuthorizePrefix + "ERROR_Sing_PASSWORD"

// NotExistsAccount 账号不存在
const NotExistsAccount = AuthorizePrefix + "NOT_EXISTS_ACCOUNT"

// ExistsAccount AuthorizeExistsAccount 账号不存在
const ExistsAccount = AuthorizePrefix + "EXISTS_ACCOUNT"
