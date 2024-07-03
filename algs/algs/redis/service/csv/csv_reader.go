package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/username/myAwesomeProject/models"
)

// LoadCSV loads user scores from a CSV file
func LoadCSV(filePath string) ([]models.UserScore, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var userScores []models.UserScore

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		score, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		userScores = append(userScores, models.UserScore{
			UserID: record[0],
			Score:  score,
		})
	}

	return userScores, nil
}
