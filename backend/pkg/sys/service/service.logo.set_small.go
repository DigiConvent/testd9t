// exempt from testing

package sys_service

import "github.com/DigiConvent/testd9t/core"

func (s *SysService) SetSmallLogo(data []byte) *core.Status {
	return setLogo("small", data)
}
