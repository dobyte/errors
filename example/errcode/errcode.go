/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/11/15 15:18
 * @Desc: TODO
 */

package errcode

type ErrCode int

func (i ErrCode) Int() int {
    return int(i)
}

//go:generate stringer -type ErrCode -linecomment
const (
    Success             ErrCode = 200 // success
    BusinessError       ErrCode = 400 // business error
    InternalServerError ErrCode = 500 // internal server error
)
