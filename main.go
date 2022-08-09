package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var startNode *node = new(node)

	for {

		printData(startNode)

		fmt.Println("1 - Добавить запись в книжку")
		fmt.Println("2 - Удалить запись из книжки")
		fmt.Println("3 - Сохранить в файл")
		fmt.Println("4 - Загрузить из файла")
		fmt.Println("5 - Выход")

		var readKey int
		fmt.Fscanln(os.Stdin, &readKey)

		switch readKey {
		case 1:

			add(startNode)

		case 2:

			var deleteNum int

			fmt.Println("Введите номер записи для удаления")
			fmt.Fscanln(os.Stdin, &deleteNum)

			startNode = delete(startNode, deleteNum)

		case 3:

			save(startNode)

		case 4:

			read(startNode)

		case 5:

			os.Exit(0)

		default:

			fmt.Print("Вы ввели неверное значение\n\n")

		}
	}
}

type node struct {
	number      int
	name        string
	lastname    string
	phoneNumber string
	nextNode    *node
} //Создание структуры 1 ноды

func (n node) setData(Node *node, number int, name string, lastname string, phoneNumber string) {

	Node.number = number
	Node.name = name
	Node.lastname = lastname
	Node.phoneNumber = phoneNumber

}

func printData(startNode *node) {

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("|Номер|                 Имя|             Фамилия|Номер телефона|")
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("|%5d|%20s|%20s|%14s|\n", startNode.number, startNode.name, startNode.lastname, startNode.phoneNumber)
	fmt.Println("---------------------------------------------------------------")

	if startNode.nextNode != nil {

		var newNode = startNode.nextNode

		for newNode.nextNode != nil {

			fmt.Printf("|%5d|%20s|%20s|%14s|\n", newNode.number, newNode.name, newNode.lastname, newNode.phoneNumber)
			fmt.Println("---------------------------------------------------------------")
			newNode = newNode.nextNode

		}
	}

}

func add(startNode *node) {

	var name string
	var lastname string
	var phone string

	fmt.Println("Введите имя")
	fmt.Fscanln(os.Stdin, &name)

	fmt.Println("Введите фамилию")
	fmt.Fscanln(os.Stdin, &lastname)

	fmt.Println("Введите номер телефона")
	fmt.Fscanln(os.Stdin, &phone)

	if startNode.nextNode == nil {

		startNode.setData(startNode, 0, name, lastname, phone)
		startNode.nextNode = new(node)

	} else {

		lastNode, number := findLastNode(startNode)
		lastNode.nextNode = new(node)
		lastNode.setData(lastNode, number, name, lastname, phone)

	}

}

func save(startNode *node) {

	var fileName string
	var Node *node = startNode

	fmt.Println("Введите имя файла")
	fmt.Fscanln(os.Stdin, &fileName)

	fileName += ".txt"
	file, _ := os.Create(fileName)

	for Node.nextNode != nil {

		file.WriteString(fmt.Sprintf("%s %s %s\n", Node.name, Node.lastname, Node.phoneNumber))
		Node = Node.nextNode

	}

	file.Close()

}

func delete(startNode *node, number int) *node {

	if startNode.number != number {

		var findNode = startNode.nextNode
		var previousNode = startNode

		for findNode.number != number {
			previousNode = findNode
			findNode = findNode.nextNode
		}
		previousNode.nextNode = findNode.nextNode
		for previousNode.nextNode != nil {

			previousNode.nextNode.number -= 1
			previousNode = previousNode.nextNode

		}

		return startNode

	} else if startNode.nextNode != nil {

		startNode.nextNode.number = 0

		return startNode.nextNode

	} else {
		startNode.number = 0
		startNode.name = ""
		startNode.lastname = ""
		startNode.phoneNumber = ""

		return startNode
	}
}

func findLastNode(startNode *node) (lastNode *node, num int) {

	lastNode = startNode
	num = 0

	for lastNode.nextNode != nil {

		lastNode = lastNode.nextNode
		num++
	}
	return
}

func saveThisData(thisString string, lastNode *node, numNode int) {

	s := strings.Split(thisString, " ")

	lastNode.setData(lastNode, numNode, s[0], s[1], s[2])
	lastNode.nextNode = new(node)

} //функция разделяет строку из файла на фамилию имя и телефон и помещает в ноду

func read(startNode *node) {

	//lastNode := startNode
	var fileName string

	fmt.Println("Введите имя файла из которого считать: ")
	fmt.Fscanln(os.Stdin, &fileName)

	file, err := os.Open(fileName)

	if err != nil {

		fmt.Println("Ошибка при открытии файла ", fileName)
		os.Exit(1)

	}

	data := make([]byte, 1)
	var thisString = ""

	for {

		_, err := file.Read(data)

		if err == io.EOF { // если конец файла

			lastNode, numNode := findLastNode(startNode)
			saveThisData(thisString, lastNode, numNode)

			thisString = ""

			break // выходим из цикла
		}

		if rune(data[0]) != '\n' {

			thisString += string(data[0])

		} else {

			lastNode, numNode := findLastNode(startNode)
			saveThisData(thisString, lastNode, numNode)

			thisString = ""

		}

	}

	file.Close()

}
