package searchindex

import (
	"encoding/json"
	"os"
)

// Saves the index to a file
func (searchIndex SearchIndex[Identifier]) Save(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600) //#nosec
	if err != nil {
		return err
	}
	// Write the data
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(searchIndex.tokensIndex); err != nil {
		_ = file.Close()
		return err
	}
	// Close the file
	return file.Close()
}
