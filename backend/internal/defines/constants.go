package defines

const (
	EmptyString = ""
	EmptyJson   = "{}"

	// Errors
	CannotParseJSONRequest   = "cannot parse JSON request"
	InvalidRequestBody       = "invalid request body"
	CannotHashPassword       = "cannot hash password"
	CannotGenerateSlugID     = "cannot generate slug id"
	CannotSaveUserOnDatabase = "cannot save user on database"
	CannotSendEmail          = "cannot send email"
	CannotGetUserByEmail     = "cannot get user by email"
	CannotGetUserBySlugID    = "cannot get user by slug id"
	CannotGetProfileByUserID = "cannot get profile by user id"
	CannotUpdateUser         = "cannot update user"
	CannotUpdateProfile      = "cannot update profile"
	AccountAlreadyActive     = "account already active"

	// Regex
	EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	PendingStatus = "pending"
	ActiveStatus  = "active"
)

var (
	SqlKeywords       = []string{"select", "insert", "update", "delete", "drop", "union", "or", "and"}
	CommonSqlPatterns = []string{
		"(?i)\\bselect\\b", "(?i)\\binsert\\b", "(?i)\\bupdate\\b", "(?i)\\bdelete\\b",
		"(?i)\\bdrop\\b", "(?i)\\bunion\\b", "(?i)\\b--\\b", "(?i)\\b;\\b",
		"(?i)\\b'\\b", "(?i)\\b\"\\b", "(?i)\\b#\\b", "(?i)\\b--\\b",
	}

	userAccountStatus = []string{PendingStatus, ActiveStatus}
)
