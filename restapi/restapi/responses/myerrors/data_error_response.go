package myerrors

import (
	"errors"
)

var (
	Success           = errors.New("Thành công")
	UnSuccess         = errors.New("Không thành công")
	ErrUnknown        = errors.New("Lỗi không xác định")
	BadRequest        = errors.New("Yêu cầu không hợp lệ")
	RoleExisted       = errors.New("Quyền đã tồn tại")
	EmptyData         = errors.New("Dữ liệu rỗng")
	UserNotExisted    = errors.New("Không tìm thấy người dùng")
	ErrNotInt         = errors.New("Err Int")
	SessionExpired    = errors.New("Session Expired")
	VerifyCodeExpired = errors.New("Verify Code Expired")
	Existed           = errors.New("Dữ liệu đã tồn tại")
	AdminExisted      = errors.New("Quản trị viên đã tồn tại")
	AddressNotExist   = errors.New("Địa chỉ không tồn tại")
	NotAdmin          = errors.New("Không phải Quản trị viên!")
	NotFound          = errors.New("Không tìm thấy dữ liệu")
	NotPermission     = errors.New("Không được phép thực hiện thao tác")
	NotOwner          = errors.New("Dữ liệu không thuộc quyền sở hữu của bạn")
	UnAuthorized      = errors.New("UnAuthorized")
	ErrLogin          = errors.New("Login Error")
	ErrSystem         = errors.New("Lỗi hệ thống")
	NotExisted        = errors.New("Dữ liệu không tồn tại")
	UserIdNotExisted  = errors.New("Mã định danh người dùng không tồn tại")
	ErrChangePass     = errors.New("Change Password Error !")
	NotMore           = errors.New("No More Data")
	ExistedPosUser    = errors.New("Existed Pos User")
	LimitCharacter    = errors.New("Limit character")
	CannotEmpty       = errors.New("Dữ liệu không được để trống")
	ErrSig            = errors.New("Sig error")
	TimeNotExceed     = errors.New("Chưa đủ thời gian thực hiện thao tác")
	Banned            = errors.New("Tài khoản bị cấm")
	AmountInvalid     = errors.New("Số lượng nhập không hợp lệ")
	TimeInvalid       = errors.New("Lần nhập không hợp lệ")

	MapDescription = map[error]string{
		AmountInvalid:                       AmountInvalid.Error(),
		TimeInvalid:                         TimeInvalid.Error(),
		Success:                             Success.Error(),
		ErrUnknown:                          ErrUnknown.Error(),
		UnAuthorized:                        "UnAuthorized",
		BadRequest:                          BadRequest.Error(),
		EmptyData:                           EmptyData.Error(),
		UserNotExisted:                      UserNotExisted.Error(),
		AdminExisted:                        AdminExisted.Error(),
		AddressNotExist:                     AddressNotExist.Error(),
		ErrNotInt:                           "Param not int!",
		UnSuccess:                           UnSuccess.Error(),
		VerifyCodeExpired:                   "Verify Code Expired",
		SessionExpired:                      "SessionExpired!",
		Existed:                             Existed.Error(),
		RoleExisted:                         RoleExisted.Error(),
		NotAdmin:                            NotAdmin.Error(),
		NotFound:                            NotFound.Error(),
		NotPermission:                       NotPermission.Error(),
		ErrLogin:                            "Wrong username, password. ",
		ErrSystem:                           ErrSystem.Error(),
		NotExisted:                          NotExisted.Error(),
		UserIdNotExisted:                    UserIdNotExisted.Error(),
		ErrChangePass:                       "Wrong username, password.",
		NotMore:                             NotMore.Error(),
		ExistedPosUser:                      ExistedPosUser.Error(),
		ErrSig:                              "Sig error !",
		CannotEmpty:                         CannotEmpty.Error(),
		TimeNotExceed:                       TimeNotExceed.Error(),
		Banned:                              Banned.Error(),
		TRANSFER_BLOCKED:                    TRANSFER_BLOCKED.Error(),
		LIMIT_WITHDRAWAL_ERROR:              LIMIT_WITHDRAWAL_ERROR.Error(),
		PUBKEY_INVALID_ERROR:                PUBKEY_INVALID_ERROR.Error(),
		BLOCK_CHAIN_ADDRESS_NOT_EXIST_ERROR: BLOCK_CHAIN_ADDRESS_NOT_EXIST_ERROR.Error(),
		PRIVATE_KEY_INVALID:                 PRIVATE_KEY_INVALID.Error(),
		BALANCE_INSUFFICIENT_ERROR:          BALANCE_INSUFFICIENT_ERROR.Error(),
		LEVEL_SECURITY_ERROR:                "Level security to low for access feature",
		INVALID_ID:                          INVALID_ID.Error(),
		USER_INFO_EMPTY:                     USER_INFO_EMPTY.Error(),
		ADDRESS_POOL_EMPTY:                  ADDRESS_POOL_EMPTY.Error(),
		COIN_NOT_EXIST:                      COIN_NOT_EXIST.Error(),
		BLOCK_CHAIN_ADDRESS_INVALID:         BLOCK_CHAIN_ADDRESS_INVALID.Error(),
		PUBKEY_NOT_EXIST_ERROR:              PUBKEY_NOT_EXIST_ERROR.Error(),
		TIME_INVALID:                        TIME_INVALID.Error(),
	}
	MapErrorCode = map[error]int64{
		Success:                             200,
		UnSuccess:                           201,
		ErrNotInt:                           302,
		SessionExpired:                      303,
		NotExisted:                          401,
		UserIdNotExisted:                    415,
		Existed:                             420,
		RoleExisted:                         400,
		ErrChangePass:                       306,
		NotAdmin:                            307,
		NotFound:                            404,
		NotPermission:                       403,
		NotMore:                             309,
		ExistedPosUser:                      310,
		CannotEmpty:                         311,
		LimitCharacter:                      312,
		ErrSig:                              316,
		BadRequest:                          303,
		EmptyData:                           202,
		UserNotExisted:                      303,
		AddressNotExist:                     303,
		AdminExisted:                        303,
		ErrUnknown:                          401,
		ErrLogin:                            402,
		ErrSystem:                           500,
		UnAuthorized:                        405,
		VerifyCodeExpired:                   406,
		TimeNotExceed:                       408,
		Banned:                              400,
		LIMIT_WITHDRAWAL_ERROR:              407,
		PUBKEY_INVALID_ERROR:                407,
		BLOCK_CHAIN_ADDRESS_NOT_EXIST_ERROR: 408,
		BALANCE_INSUFFICIENT_ERROR:          409,
		INVALID_ID:                          501,
		LEVEL_SECURITY_ERROR:                410,
		USER_INFO_EMPTY:                     303,
		ADDRESS_POOL_EMPTY:                  410,
		BLOCK_CHAIN_ADDRESS_INVALID:         411,
		COIN_NOT_EXIST:                      412,
		PRIVATE_KEY_INVALID:                 413,
		PUBKEY_NOT_EXIST_ERROR:              414,
		TIME_INVALID:                        414,
		LIMIT_NOT_EXISTED:                   416,
	}
)

// Returns a error.
// swagger:response Err
type Err struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message string `json:"message"`
}

type ValidateErr struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message map[string]string `json:"message"`
}

func NewErr(err error) *Err {
	if _, exist := MapErrorCode[err]; exist {
		return &Err{
			Code:    MapErrorCode[err],
			Message: MapDescription[err],
		}
	} else {
		return &Err{
			Code:    401,
			Message: err.Error(),
		}
	}
}

func NewErrByText(err error, text string) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: text,
	}
}
