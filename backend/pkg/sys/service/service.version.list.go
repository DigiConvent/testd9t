// exempt from testing
package sys_service

import (
	"encoding/json"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/file_repo"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) ListReleaseTags() ([]sys_domain.ReleaseTag, *core.Status) {
	contents, err := file_repo.NewRepoRemote().ReadRawFile(".meta/release_tags.json")

	if err != nil {
		return nil, core.InternalError(err.Error())
	}

	versions := make([]sys_domain.ReleaseTag, 0)
	err = json.Unmarshal(contents, &versions)
	if err != nil {
		return nil, core.InternalError(err.Error())
	}

	return versions, core.StatusSuccess()
}
