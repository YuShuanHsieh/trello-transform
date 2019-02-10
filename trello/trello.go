package trello

import (
	"encoding/json"

	"github.com/YuShuanHsieh/trello-transform/models"
)

func UnmarshalJson(jsonData []byte, t *models.TrelloTable) error {
	err := json.Unmarshal(jsonData, t)
	if err != nil {
		return err
	}
	return nil
}
