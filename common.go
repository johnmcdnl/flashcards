package flashcards

import (
	"encoding/json"
)

func jsonString(i interface{}) string {
	j, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	}
	return string(j)
}
