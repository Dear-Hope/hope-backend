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
	ERROR_BLOCK_USER_FAILED         = "failed to block user"
	ERROR_GET_BLOCKED_USER_FAILED   = "failed to get blocked user"
)

var (
	ERROR_FAILED_TO_CONVERT_HOUR = "failed to convert hour"
	ERROR_MOOD_NOT_LISTED        = "mood given is not listed in our database"
	ERROR_CREATE_EMOTION_FAILED  = "failed to create new emotion"
	ERROR_EMOTION_ALREADY_EXISTS = "this patient emotion for this time frame today is already exists"
	ERROR_GET_ALL_EMOTION_FAILED = "failed to get all emotions"
)

var (
	ERROR_INVALID_EMAIL      = "invalid email given"
	ERROR_SUBSCRIBE_FAILED   = "failed to subscribe"
	ERROR_UNSUBSCRIBE_FAILED = "failed to unsubscribe"
)

var (
	ERROR_MOVIE_MINIMAL_ONE_MOOD    = "movie must be associated with minimal 1 mood"
	ERROR_MOVIE_ALREADY_EXISTS      = "movie already exists"
	ERROR_CREATE_MOVIE_FAILED       = "failed to create new movie"
	ERROR_STORE_MOVIE_DETAIL_FAILED = "failed to store new movie %s detail"
	ERROR_GET_LIST_MOVIE_FAILED     = "failed to get list movies"
	ERROR_GET_MOVIE_FAILED          = "failed to get detail movie"
	ERROR_MOVIE_NOT_FOUND           = "movie not found with given id"
)

var (
	ERROR_GET_LIST_CATEGORIES_FAILED    = "failed to get list categories"
	ERROR_GET_TOTAL_PER_CATEGORY_FAILED = "failed to get total per category"
)

var (
	ERROR_GET_LIST_BREATHING_EXERCISES_FAILED = "failed to get list breathing exercises"
	ERROR_SET_LAST_PLAYED_FAILED              = "failed to set last played %s"
	ERROR_GET_LAST_PLAYED_FAILED              = "failed to get last played %s"
)

var (
	ERROR_GET_LIST_SELF_HEALING_AUDIO_THEMES_FAILED = "failed to get list self healing audio themes"
	ERROR_GET_SELF_HEALING_AUDIO_THEME_FAILED       = "failed to get self healing audio theme"
	ERROR_SELF_HEALING_AUDIO_THEME_NOT_FOUND        = "self healing audio theme not found with given id"
	ERROR_GET_SELF_HEALING_AUDIO_FAILED             = "failed to get self healing audio"
	ERROR_SELF_HEALING_AUDIO_NOT_FOUND              = "self healing audio not found with given id"
	ERROR_GET_LIST_SELF_HEALING_AUDIO_FAILED        = "failed to get list self healing audios"
)

var (
	ERROR_POST_MINIMAL_ONE_CATEGORY         = "post must be associated with minimal 1 category"
	ERROR_POST_ALREADY_EXISTS               = "post already exists"
	ERROR_COMMENT_ALREADY_EXISTS            = "comment already exists"
	ERROR_CREATE_POST_FAILED                = "failed to create new post"
	ERROR_CREATE_COMMENT_FAILED             = "failed to create new comment on post id %d"
	ERROR_CREATE_CATEGORY_POST_FAILED       = "failed to associate category on post id %d"
	ERROR_GET_LIST_POST_FAILED              = "failed to get list posts"
	ERROR_GET_CATEGORY_BY_POST_ID_FAILED    = "failed to get category on post id %d"
	ERROR_GET_COMMENT_BY_POST_ID_FAILED     = "failed to get comment on post id %d"
	ERROR_GET_POST_FAILED                   = "failed to get detail post"
	ERROR_POST_NOT_FOUND                    = "post not found with given id"
	ERROR_DELETE_POST_FAILED                = "failed to delete post"
	ERROR_DELETE_COMMENT_BY_POST_ID_FAILED  = "failed to delete comment on post id %d"
	ERROR_DELETE_CATEGORY_BY_POST_ID_FAILED = "failed to delete category on post id %d"
	ERROR_DELETE_LIKE_BY_POST_ID_FAILED     = "failed to delete like on post id %d"
	ERROR_GET_LIST_CATEGORY_FAILED          = "failed to get list categories"
	ERROR_UPSERT_LIKE_FAILED                = "failed to like/unlike the post id %d"
	ERROR_GET_LIST_REASON_FAILED            = "failed to get list report reasons"
	ERROR_REPORT_POST_FAILED                = "failed to report post with id %d"
	ERROR_REPORT_POST_MINIMAL_ONE_REASON    = "report must have minimal 1 reason"
)

var (
	ERROR_GET_LIST_MUSIC_PLAYLIST_FAILED = "failed to get list music playlists"
)
