package repository

import (
    "fmt"
)

func formatLength(minutes int) string {
	hours := minutes / 60
	minutes = minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func calculateActorAge(actorId int, releaseYear int) int {
	return releaseYear - wholeData.Persons[actorId].BirthYear
}
