package actioninfo

import "log"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, datastring := range dataset {
		err := dp.Parse(datastring)
		if err != nil {
			log.Println("Ошибка парсинга:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка формирования:", err)
			continue
		}

		log.Println(info)
	}
}
