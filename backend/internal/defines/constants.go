package defines

import (
	"time"
)

const (
	// Errors
	CannotParseJSONRequest         = "cannot parse JSON request"
	CannotHashPassword             = "cannot hash password"
	CannotGenerateSlugID           = "cannot generate slug id"
	CannotSaveUserOnDatabase       = "cannot save user on database"
	CannotSendEmail                = "cannot send email"
	CannotGetUserByEmail           = "cannot get user by email"
	CannotGetUserByUsername        = "cannot get user by username"
	CannotGetUserBySlugID          = "cannot get user by slug id"
	CannotGetProfileByUserID       = "cannot get profile by user id"
	CannotUpdateUser               = "cannot update user"
	CannotUpdateProfile            = "cannot update profile"
	CannotValidateLogin            = "cannot validate login"
	CannotValidateJwt              = "cannot validate jwt"
	CannotCheckBrokenAccessControl = "cannot check broken access control"
	CannotValidateRequest          = "cannot validate request"
	CannotCreateSexualPreference   = "cannot create sexual preference"
	InvalidRequestBody             = "invalid request body"
	AccountAlreadyActive           = "account already active"
	InvalidPassword                = "invalid password"
	UsernameMustToBeProvided       = "username must to be provided"
	AccountNotActiveYet            = "account not active yet"
	AuthorizationMustToBeProvided  = "authorization must to be provided"
	UnauthorizedRequest            = "unauthorized request"

	// Regex
	EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	PendingStatus                    = "pending"
	ActiveStatus                     = "active"
	EmptyString                      = ""
	EmptyJson                        = "{}"
	JwtExpirationTime                = time.Hour * 24
	RecoverPasswordJwtExpirationTime = time.Minute * 15
	SexualPreferenceSeparator        = "_"

	ActiveAccountEmailSubject   = "Welcome to Matcha!"
	ActiveAccountEmailContent   = "Welcome aboard!\n\nClick here to activate your account: "
	RecoverPasswordEmailSubject = "Recover your Matcha password"
	RecoverPasswordEmailContent = "Click here to recover your password: "
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
