package repeatrule

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	// парсим repeat
	repeatParts := strings.Fields(repeat)
	if len(repeatParts) == 0 {
		return "", fmt.Errorf("ошибка парсинга repeat: %s", repeat)
	}
	// парсим date
	dateTime, err := time.Parse("20060102", date)
	if err != nil {
		return "", err
	}
	//проверяем параметр repeat
	var newDate time.Time
	
	switch repeatParts[0] {
	case "y": // если параметр y, то добавляем год
		newDate = dateTime.AddDate(1, 0, 0)
		for newDate.Before(now) {
				newDate = newDate.AddDate(1, 0, 0) // если новая дата меньше текущей, то добавляем еще один год
		}
		
	case "d": // если параметр d, то добавляем дни
		if len(repeatParts) > 1 {
			days, err := strconv.Atoi(repeatParts[1])
			if err != nil {
				return "", err
			}
			if days > 400 {
				return "", fmt.Errorf("количество дней не может быть больше 400")
			}
			newDate = dateTime.AddDate(0, 0, days) // добавляем дни
			for newDate.Before(now) {
				newDate = newDate.AddDate(0, 0, days) // если новая дата меньше текущей, то добавляем еще раз дни
			}
		} else {
			return "", fmt.Errorf("не указан параметр количества дней")
		}
	
	case "m":
		return "", fmt.Errorf("не реализовано") // TODO: реализовать
	
	case "w":
		return "", fmt.Errorf("не реализовано") // TODO: реализовать
	
	default:
		return "", fmt.Errorf("неизвестный параметр repeat: %s", repeatParts[0])
	}

	return newDate.Format("20060102"), nil
}
