package constant

var ErrorInternalServer = "internal server error"

// Error package auth
var (
	ErrorUserAlreadyExists      = "alamat email sudah terdaftar"
	ErrorCreateUserFailed       = "gagal mendaftarkan user"
	ErrorGenerateToken          = "gagal untuk membuat token"
	ErrorGenerateOtpCode        = "gagal mendapatkan kode otp"
	ErrorUserNotFound           = "user tidak ditemukan"
	ErrorGetUserFailed          = "gagal mendapatkan detail user"
	ErrorPasswordNotMatch       = "password tidak sesuai"
	ErrorAccountNotVerified     = "akun belum diverifikasi"
	ErrorVerifyUserFailed       = "gagal verifikasi user"
	ErrorOtpCodeExpired         = "kode verifikasi anda sudah expired"
	ErrorAccountAlreadyVerified = "akun sudah diverifikasi sebelumnya"
	ErrorResendLimitReached     = "mencapai batas untuk kirim ulang kode, mohon coba lagi beberapa saat lagi"
	ErrorChangePasswordFailed   = "gagal mengganti password"
	ErrorUpdateUserFailed       = "gagal memperbarui user"
	ErrorSaveProfilePhotoFailed = "gagal menyimpan foto profil"
	//ERROR_BLOCK_USER_FAILED         = "failed to block user"
	//ERROR_GET_BLOCKED_USER_FAILED   = "failed to get blocked user"
)

// var (
//
//	ERROR_FAILED_TO_CONVERT_HOUR = "failed to convert hour"
//	ERROR_MOOD_NOT_LISTED        = "mood given is not listed in our database"
//	ERROR_CREATE_EMOTION_FAILED  = "failed to create new emotion"
//	ERROR_EMOTION_ALREADY_EXISTS = "this patient emotion for this time frame today is already exists"
//	ERROR_GET_ALL_EMOTION_FAILED = "failed to get all emotions"
//
// )
//
// var (
//
//	ERROR_INVALID_EMAIL      = "invalid email given"
//	ERROR_SUBSCRIBE_FAILED   = "failed to subscribe"
//	ERROR_UNSUBSCRIBE_FAILED = "failed to unsubscribe"
//
// )
//
// var (
//
//	ERROR_MOVIE_MINIMAL_ONE_MOOD    = "movie must be associated with minimal 1 mood"
//	ERROR_MOVIE_ALREADY_EXISTS      = "movie already exists"
//	ERROR_CREATE_MOVIE_FAILED       = "failed to create new movie"
//	ERROR_STORE_MOVIE_DETAIL_FAILED = "failed to store new movie %s detail"
//	ERROR_GET_LIST_MOVIE_FAILED     = "failed to get list movies"
//	ERROR_GET_MOVIE_FAILED          = "failed to get detail movie"
//	ERROR_MOVIE_NOT_FOUND           = "movie not found with given id"
//
// )
//
// var (
//
//	ERROR_GET_LIST_CATEGORIES_FAILED    = "failed to get list categories"
//	ERROR_GET_TOTAL_PER_CATEGORY_FAILED = "failed to get total per category"
//
// )
//
// var (
//
//	ERROR_GET_LIST_BREATHING_EXERCISES_FAILED = "failed to get list breathing exercises"
//	ERROR_SET_LAST_PLAYED_FAILED              = "failed to set last played %s"
//	ERROR_GET_LAST_PLAYED_FAILED              = "failed to get last played %s"
//
// )
//
// var (
//
//	ERROR_GET_LIST_SELF_HEALING_AUDIO_THEMES_FAILED = "failed to get list self healing audio themes"
//	ERROR_GET_SELF_HEALING_AUDIO_THEME_FAILED       = "failed to get self healing audio theme"
//	ERROR_SELF_HEALING_AUDIO_THEME_NOT_FOUND        = "self healing audio theme not found with given id"
//	ERROR_GET_SELF_HEALING_AUDIO_FAILED             = "failed to get self healing audio"
//	ERROR_SELF_HEALING_AUDIO_NOT_FOUND              = "self healing audio not found with given id"
//	ERROR_GET_LIST_SELF_HEALING_AUDIO_FAILED        = "failed to get list self healing audios"
//
// )
//
// var (
//
//	ERROR_POST_MINIMAL_ONE_CATEGORY         = "post must be associated with minimal 1 category"
//	ERROR_POST_ALREADY_EXISTS               = "post already exists"
//	ERROR_COMMENT_ALREADY_EXISTS            = "comment already exists"
//	ERROR_CREATE_POST_FAILED                = "failed to create new post"
//	ERROR_CREATE_COMMENT_FAILED             = "failed to create new comment on post id %d"
//	ERROR_CREATE_CATEGORY_POST_FAILED       = "failed to associate category on post id %d"
//	ERROR_GET_LIST_POST_FAILED              = "failed to get list posts"
//	ERROR_GET_CATEGORY_BY_POST_ID_FAILED    = "failed to get category on post id %d"
//	ERROR_GET_COMMENT_BY_POST_ID_FAILED     = "failed to get comment on post id %d"
//	ERROR_GET_POST_FAILED                   = "failed to get detail post"
//	ERROR_POST_NOT_FOUND                    = "post not found with given id"
//	ERROR_DELETE_POST_FAILED                = "failed to delete post"
//	ERROR_DELETE_COMMENT_BY_POST_ID_FAILED  = "failed to delete comment on post id %d"
//	ERROR_DELETE_CATEGORY_BY_POST_ID_FAILED = "failed to delete category on post id %d"
//	ERROR_DELETE_LIKE_BY_POST_ID_FAILED     = "failed to delete like on post id %d"
//	ERROR_GET_LIST_CATEGORY_FAILED          = "failed to get list categories"
//	ERROR_UPSERT_LIKE_FAILED                = "failed to like/unlike the post id %d"
//	ERROR_GET_LIST_REASON_FAILED            = "failed to get list report reasons"
//	ERROR_REPORT_FAILED                     = "failed to report %s with id %d"
//	ERROR_REPORT_POST_MINIMAL_ONE_REASON    = "report must have minimal 1 reason"
//
// )
//
// var (
//
//	ERROR_GET_LIST_MUSIC_PLAYLIST_FAILED = "failed to get list music playlists"
//
// )
var (
	//ErrorGetListTopicsFailed     = "gagal mendapatkan daftar topik"
	ErrorExpertMinimalOneTopic = "ahli harus memiliki minimal satu topik"
	ErrorExpertAlreadyExists   = "ahli ini sudah terdaftar"
	ErrorCreateExpertFailed    = "gagal membuat ahli baru"
	//ErrorGetListExpertsFailed    = "gagal mendapatkan daftar ahli"
	ErrorExpertNotFound  = "ahli dengan id tersebut tidak ditemukan"
	ErrorGetExpertFailed = "gagal mendapatkan detail ahli"
	//ErrorDeleteExpertFailed      = "gagal menghapus ahli"
	ErrorGetExpertTopicIdsFailed = "gagal mendapatkan topik yang dimiliki ahli"
	ErrorUpdateExpertFailed      = "gagal mengupdate data ahli"
)

var (
	//
	//	ErrorUpcomingScheduleNotFound        = "jadwal sesi berikutnya dari ahli dengan id tersebut tidak ditemukan"
	//	ErrorGetExpertUpcomingScheduleFailed = "gagal mendapatkan jadwal ahli sesi berikutnya"
	ErrorGetExpertScheduleFailed = "gagal mendapatkan jadwal ahli"
	//	ErrorBookExpertScheduleFailed        = "gagal melakukan pemesanan jadwal ahli"
	//	ErrorExpertScheduleAlreadyExists = "jadwal untuk ahli ini sudah ada"
	ErrorCreateExpertScheduleFailed = "gagal membuat jadwal ahli"
	ErrorUpdateExpertScheduleFailed = "gagal memperbarui jadwal ahli"
)

var (
	ErrorGetConsultationFailed          = "Gagal mendapatkan daftar konsultasi"
	ErrorGetDetailConsultationFailed    = "Gagal mendapatkan detail konsultasi"
	ErrorUpdateStatusConsultationFailed = "Gagal memperbarui status konsultasi"
	ErrorInvalidStatusConsultation      = "Status yang diisi salah"
)

var (
	ErrorGetExpertReviewFailed = "gagal mendapatkan penilaian dan ulasan ahli"
)
