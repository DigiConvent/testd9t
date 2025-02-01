package sys_repository

import (
	"encoding/json"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/file_repo"
)

func (r *SysRepository) ListFlavoursForVersion() ([]string, core.Status) {
	var flavours []string

	rawData, err := file_repo.NewRepoRemote().ReadRawFile(".meta/flavours.json")

	if err != nil {
		return flavours, *core.InternalError("Could not read the flavours file")
	}

	json.Unmarshal(rawData, &flavours)

	return flavours, *core.StatusSuccess()
}
