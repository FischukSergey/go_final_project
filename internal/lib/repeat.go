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
		// weekdays:=strings.Split(repeatParts[1],",") // получаем дни недели
		// if len(weekdays) == 0 {
		// 	return "", fmt.Errorf("не указан параметр дней недели")
		// }
		// weekdaysInt:=make([]int,0,len(weekdays)) // преобразуем дни недели в слайс int
		// for _, day := range weekdays {  // валидируем дни недели
		// 	dayInt, err := strconv.Atoi(day)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	if dayInt <= 0 || dayInt > 7 {
		// 		return "", fmt.Errorf("недопустимый день недели: %s", day)
		// 	}
		// 	weekdaysInt = append(weekdaysInt, dayInt)
		// }
		// sort.Ints(weekdaysInt) // сортируем дни недели

		// for newDate.Before(now) {
		// 	for _, weekday := range weekdaysInt { // проверяем каждый день недели из списка
		// 		for newDate.Weekday() != time.Weekday(weekday) { // если день недели не совпадает с текущим, то добавляем день
		// 			newDate = newDate.AddDate(0, 0, 1) // Move to the next day
		// 		}
		// 		// If the found date is before now, we need to add a week
		// 		if newDate.Before(now) {
		// 			newDate = newDate.AddDate(0, 0, 7)
		// 		}
		// 	}
		// 	newDate = newDate.AddDate(0, 0, 7) // если новая дата меньше текущей, то добавляем еще раз дни
		// }
		return "", fmt.Errorf("не реализовано") // TODO: реализовать
	
	default:
		return "", fmt.Errorf("неизвестный параметр repeat: %s", repeatParts[0])
	}

	return newDate.Format("20060102"), nil
}
