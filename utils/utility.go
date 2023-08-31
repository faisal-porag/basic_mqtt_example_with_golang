package utils

import (
	"github.com/rs/zerolog/log"
	"time"
)

func GetBDTimeZone() time.Time {
	// Load the Asia/Dhaka time zone location
	bangladeshTimeZone, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		log.Error().Err(err).Msgf("Error loading Bangladesh time zone:", err)
		return time.Time{} // Return a zero time value to indicate an error
	}

	// Get the current time in the Bangladesh time zone
	bangladeshTime := time.Now().UTC().In(bangladeshTimeZone)
	return bangladeshTime
}
