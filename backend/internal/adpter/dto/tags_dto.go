package dto

import "backend/internal/utils"

type TagsDto struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (t TagsDto) IsValid() bool {
	return t.ID > 0 && len(t.Name) > 0 && !utils.IsSQLInjection(t.Name)
}
