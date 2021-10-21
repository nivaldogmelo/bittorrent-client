package bencode

import (
	"log"
	"strconv"
)

func Parser(bencodeFile string) (interface{}, error) {
	var value interface{}
	index := 0
	for index < len(bencodeFile)-1 {
		kind, err := getType(bencodeFile[index : index+1])
		if err != nil {
			log.Panic(err)
		}

		if kind == "integer" {
			index, value, _ = createInt(bencodeFile[index:])
		}

		if kind == "string" {
			index, value, _ = createString(bencodeFile[index:])
		}

		index += 1
	}

	return value, nil
}

func getType(kind string) (string, error) {
	switch kind {
	case "d":
		return "dict", nil
	case "l":
		return "list", nil
	case "i":
		return "integer", nil
	default:
		return "string", nil
	}
}

func createInt(content string) (int, int, error) {
	integer := ""
	position := 1
	for position < len(content)-1 {
		if content[position:position+1] != "e" {
			integer += content[position : position+1]
			position += 1
		} else {
			position += 1
			break
		}

	}

	value, _ := strconv.Atoi(integer)
	index := position

	return index, value, nil
}

func createString(content string) (int, string, error) {
	stringLength := 0
	position := 0

	for pos, char := range content {
		if string(char) == ":" {
			position = pos + 1
			break
		} else {
			digit, _ := strconv.Atoi(string(char))
			stringLength = stringLength*10 + digit
			position = pos
		}
	}

	returnString := content[position : position+stringLength]
	index := position + stringLength

	return index, returnString, nil

}
