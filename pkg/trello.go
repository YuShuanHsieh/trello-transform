package trello

import (
	"encoding/json"
	"fmt"

	"github.com/YuShuanHsieh/trello-transform/models"
)

func UnmarshalJson(jsonData []byte, t *models.TrelloTable) error {
	err := json.Unmarshal(jsonData, t)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal json: %s", err.Error())
	}
	return nil
}
