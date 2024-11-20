package defines

import (
	"time"
)

const (
	// Errors
	CannotParseJSONRequest            = "cannot parse JSON request"
	CannotHashPassword                = "cannot hash password"
	CannotGenerateSlugID              = "cannot generate slug id"
	CannotSaveUserOnDatabase          = "cannot save user on database"
	CannotSendEmail                   = "cannot send email"
	CannotGetUserByEmail              = "cannot get user by email"
	CannotGetUserByUsername           = "cannot get user by username"
	CannotGetUserBySlugID             = "cannot get user by slug id"
	CannotGetProfileByUserID          = "cannot get profile by user id"
	CannotUpdateUser                  = "cannot update user"
	CannotUpdateUserPassword          = "cannot update user password"
	CannotUpdateProfile               = "cannot update profile"
	CannotValidateLogin               = "cannot validate login"
	CannotValidateJwt                 = "cannot validate jwt"
	CannotCheckBrokenAccessControl    = "cannot check broken access control"
	CannotValidateRequest             = "cannot validate request"
	InvalidRequestBody                = "invalid request body"
	AccountAlreadyActive              = "account already active"
	WrongEmailOrPassword              = "wrong email or password"
	UsernameMustToBeProvided          = "username must to be provided"
	AccountNotActiveYet               = "account not active yet"
	AuthorizationMustToBeProvided     = "authorization must to be provided"
	UnauthorizedRequest               = "unauthorized request"
	CannotCreateProfileOnDatabase     = "cannot create profile on database"
	CannotGenerateJwt                 = "cannot generate jwt"
	CannotGetGender                   = "cannot get gender"
	CannotGetGenderID                 = "cannot get gender id"
	CannotGetSexualPreference         = "cannot get sexual preference"
	CannotGetSexualPreferenceID       = "cannot get sexual preference id"
	CannotDeleteUser                  = "cannot delete user"
	CannotEstablishDatabaseConnection = "cannot establish database connection"
	CannotDeleteProfile               = "cannot delete profile"

	// Regex
	EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	PendingStatus                    = "pending"
	ActiveStatus                     = "active"
	EmptyString                      = ""
	EmptyJson                        = "{}"
	JwtExpirationTime                = time.Hour * 1
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
