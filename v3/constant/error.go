package constant

var (
	ERROR_USER_ALREADY_EXISTS       = "user with this email address already exists"
	ERROR_CREATE_USER_FAILED        = "failed to create user"
	ERROR_GENERATE_TOKEN            = "failed to generate token"
	ERROR_GENERATE_OTP_CODE         = "failed to generate otp code"
	ERROR_USER_NOT_FOUND            = "user not found with given email"
	ERROR_GET_USER_FAILED           = "failed to get user"
	ERROR_PASSWORD_NOT_MATCH        = "password does not match"
	ERROR_ACCOUNT_NOT_ACTIVATED     = "account has not been activated yet"
	ERROR_ACTIVATE_USER_FAILED      = "failed to activate user"
	ERROR_OTP_CODE_EXPIRED          = "your activation code has expired"
	ERROR_CHANGE_PASSWORD_FAILED    = "failed to change user password"
	ERROR_ACCOUNT_ALREADY_ACTIVATED = "account has already been activated"
	ERROR_RESEND_LIMIT_REACHED      = "resend limit reached, please try again after a while"
	ERROR_UPDATE_USER_FAILED        = "failed to update user"
	ERROR_SAVE_PROFILE_PHOTO_FAILED = "failed to save user profile photo"
)

var (
	ERROR_FAILED_TO_CONVERT_HOUR = "failed to convert hour"
	ERROR_MOOD_NOT_LISTED        = "mood given is not listed in our database"
	ERROR_CREATE_EMOTION_FAILED  = "failed to create new emotion"
	ERROR_EMOTION_ALREADY_EXISTS = "this patient emotion for this time frame today is already exists"
	ERROR_GET_ALL_EMOTION_FAILED = "failed to get all emotions"
)
