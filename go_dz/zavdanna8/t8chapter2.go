// Гриценко Ірина Валеріївна завдання 8 частина 2 (big hard task)
package main

import "fmt"

type data struct {
	number    uint
	Firstname []string
	Lastname  []string
	marks     []uint
}

//<----- функція виведення даних про студента
func printData(student *data) {
	fmt.Println("Номер студентського посвідчення: ", student.number)

	fmt.Print("Cтудент: ")
	l := len(student.Firstname)
	var firstname string
	for i := 0; i < l; i++ {
		firstname += student.Firstname[i]
	}
	l = len(student.Lastname)
	var lastname string
	for i := 0; i < l; i++ {
		lastname += student.Lastname[i]
	}
	fmt.Println(lastname, firstname)

	fmt.Print("Оцінки за останньою сесією: ")
	l = len(student.marks)
	for i := 0; i < l; i++ {
		fmt.Print(student.marks[i], " ")
	}
	fmt.Print("\n")
}

//<----- допоміжна функція(сер. арифм.) для функції-сортувальника
func avrg(arr []uint) uint {
	lenth := len(arr)
	var sum uint
	for i := 0; i < lenth; i++ {
		sum += arr[i]
	}
	return sum / uint(lenth)
}

//<----- функція-сортувальник
func sort(student [9]*data, lenth int) [9]*data {
	for i := 1; i < lenth; i++ {
		for k := 1; k < lenth; k++ {
			if (avrg(student[k].marks)) > (avrg(student[k-1].marks)) {
				tmp := student[k]
				student[k] = student[k-1]
				student[k-1] = tmp
			}
		}
	}
	return student
}

//<----- функція-фільтр
func findLastnamesEndedA(student [9]*data, lenth int) {
	fmt.Println("\nДані студентів прізвища, яких закінчуються літерою 'а' :\n")
	for i := 0; i < lenth; i++ {
		lenOfLastname := len(student[i].Lastname)
		if student[i].Lastname[lenOfLastname-1] == "а" {
			printData(student[i])
			fmt.Print("\n")
		}
	}
}

//<----- функція виведення масиву
func printElsOfArr(student [9]*data, lenth int) {
	fmt.Print("\n")
	for i := 0; i < lenth; i++ {
		fmt.Println(*student[i])
	}
	fmt.Print("\n")
}

func main() {
	students := [9]data{
		{1001, []string{"М", "а", "р", "г", "а", "р", "и", "т", "а"}, []string{"A", "в", "е", "р", "к", "і", "н", "а"}, []uint{78, 99, 95, 87, 89, 93, 90}},
		{1002, []string{"В", "і", "к", "т", "о", "р"}, []string{"Б", "о", "н", "д", "а", "р"}, []uint{78, 69, 95, 87, 83, 93, 88}},
		{1003, []string{"А", "н", "д", "р", "і", "й"}, []string{"В", "е", "р", "б", "и", "ц", "ь", "к", "и", "й"}, []uint{78, 66, 95, 63, 89, 95, 84}},
		{1004, []string{"О", "л", "е", "к", "с", "а", "н", "д", "р"}, []string{"Г", "о", "р", "д", "і", "є", "н", "к", "о"}, []uint{95, 97, 95, 97, 89, 79, 91}},
		{1005, []string{"І", "р", "и", "н", "а"}, []string{"Г", "р", "и", "ц", "е", "н", "к", "о"}, []uint{99, 100, 90, 95, 97, 93, 98}},
		{1006, []string{"Б", "о", "г", "д", "а", "н"}, []string{"З", "а", "я", "ц"}, []uint{78, 93, 95, 87, 62, 78, 94}},
		{1007, []string{"А", "н", "н", "а"}, []string{"Л", "у", "ц", "е", "н", "к", "о"}, []uint{78, 99, 95, 87, 60, 93, 90}},
		{1008, []string{"Ю", "р", "і", "й"}, []string{"О", "н", "и", "щ", "е", "н", "к", "о"}, []uint{87, 100, 70, 85, 89, 91, 92}},
		{1009, []string{"О", "л", "е", "г"}, []string{"Р", "і", "д", "ч", "е", "н", "к", "о"}, []uint{78, 99, 95, 86, 75, 93, 91}},
	}
	var studentsptr [9]*data
	for i := 0; i < len(students); i++ {
		studentsptr[i] = &students[i]
	}

	l := len(students)
	// виведення початкового масиву
	fmt.Println("\nМасив студентів:")
	for i := 0; i < l; i++ {
		fmt.Println(students[i])
	}

	// сортування масиву за зменшенням оцінок
	studentsptr = sort(studentsptr, l)
	// демонструємо зміни в масиві
	fmt.Println("\nМасив студентів після сортування:")
	for i := 0; i < l; i++ {
		fmt.Println(students[i])
	}
	fmt.Print("\n")
	// демонструємо зміни в масиві вказівників
	fmt.Println("\nМасив вказівників(виведуться значення на які, вони вказують) студентів після сортування:")
	printElsOfArr(studentsptr, l)

	//фільтрація з одночасним виведенням даних студентів, якщо прізвище закінчуються літерою а
	findLastnamesEndedA(studentsptr, l)
}
