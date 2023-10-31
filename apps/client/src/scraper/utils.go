package scraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
)

func DefaultInt(input string, defaultValue int) int {
	if input == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(input)

	if err != nil {
		return defaultValue
	}

	return value
}

func DefaultBytes(input string, defaultValue int) string {
	value, err := humanize.ParseBytes(input)

	if err != nil {
		return humanize.Bytes(0)
	}

	return humanize.Bytes(value)
}

func TimestampToString(input string) string {
	intValue, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		panic(err)
	}

	timeValue := time.Unix(intValue, 0)
	return TimeDifference(timeValue)
}

func TimeDifference(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	days := int(diff.Hours() / 24)
	months := int(days / 30)

	if months > 0 {
		return fmt.Sprintf("%d month(s) ago", months)
	}

	return fmt.Sprintf("%d day(s) ago", days)
}

func GetJson(url string, response interface{}) error {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Failed to parse JSON", err, string(body))
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Println("Failed to GET JSON:", resp.StatusCode, "for", url)
		return errors.New(string(resp.StatusCode))
	}

	return nil
}
