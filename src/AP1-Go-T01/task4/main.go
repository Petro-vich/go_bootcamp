package main

import (
	"fmt"
	"time"
)

const (
	Save         = "Save"
	GetHistory   = "GetHistory"
	GetLastVisit = "GetLastVisit"
	Exit         = "Exit"
)

type Visit struct {
	Specialization string
	Date           string
}

func getUserCommand() string {
	fmt.Println("Введите комманду: ")
	fmt.Println("1. Save -> позволяет сохранять в картотеку ФИО посетителя, специализацию врача, к которому он приходил,\nи время визита.")
	fmt.Println("2. GetHistory -> позволяет просмотреть историю посещения пациентом больницы.")
	fmt.Println("3. GetLastVisit -> позволяет получить последнее посещение пациентом определенного специалиста в больнице.")

	var command string
	for {
		_, err := fmt.Scanln(&command)
		if err != nil ||
			(command != Save &&
				command != GetHistory &&
				command != GetLastVisit &&
				command != Exit) {
			fmt.Println("Unknown command. Please, try again")
			continue
		}

		break
	}
	return command
}

func implCommand(command *string, pacientVisit map[string][]Visit) bool {

	switch *command {
	case Save:
		implSave(pacientVisit)
	case GetHistory:
		implGetHistory(pacientVisit)
	case GetLastVisit:
		implGetLastVisit(pacientVisit)
	case Exit:
		fmt.Println("Exiting...")
		return false
	default:
		fmt.Println("unknown command")
		return true
	}
	return true
}

func implSave(pacientVisit map[string][]Visit) {
	var fullName string
	fmt.Println("Введите ФИО пациента:")
	fmt.Scanln(&fullName)

	var doctorSpecialization string
	for {
		fmt.Println("Введите специализацию врача")
		_, err := fmt.Scanln(&doctorSpecialization)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	var dateVisit string
	for {
		fmt.Println("Введите время визита в формате YYYY-MM-DD")
		_, err := fmt.Scanln(&dateVisit)
		if err != nil {
			fmt.Println("err")
			continue
		}

		if _, err = time.Parse("2006-01-02", dateVisit); err != nil {
			fmt.Println("Неверный формат даты. Ожидается YYYY-MM-DD")
			continue
		}
		break
	}

	visit := Visit{
		Specialization: doctorSpecialization,
		Date:           dateVisit,
	}

	pacientVisit[fullName] = append(pacientVisit[fullName], visit)
}

func implGetHistory(pacientVisit map[string][]Visit) {
	var fullName string
	fmt.Println("Введи ФИО пациента")

	for {
		_, err := fmt.Scanln(&fullName)
		if err != nil {
			fmt.Println("Тут будет реализация вывода собственного типа ошибки")
			continue
		}
		break
	}

	history, ok := pacientVisit[fullName]
	_ = history
	_ = ok
	fmt.Println(history)
}

func implGetLastVisit(pacientVisit map[string][]Visit) {
	fmt.Println("Введите ФИО пациента")
	var fullName string
	for {
		_, err := fmt.Scanln(&fullName)
		if err != nil {
			fmt.Println("invalid input. Try again")
			continue
		} else if _, ok := pacientVisit[fullName]; !ok {
			fmt.Println("Пациент не найден")
			continue
		}
		break
	}

	var doctorSpecialization string
	fmt.Println("Введите специализацию врача")
	for {
		_, err := fmt.Scanln(&doctorSpecialization)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	var history string
	for _, visit := range pacientVisit[fullName] {
		if visit.Specialization == doctorSpecialization {
			history = visit.Date
		}
	}
	fmt.Println("Последний визит:", history)
}

func main() {
	pacientVisit := make(map[string][]Visit)

	isWorking := true

	for isWorking {
		command := getUserCommand()
		isWorking = implCommand(&command, pacientVisit)
	}
}
