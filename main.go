package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Arguments map[string]string

type Users struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func Perform(args Arguments, writer io.Writer) error {
	if args["fileName"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	if args["operation"] == "" {
		return errors.New("-operation flag has to be specified")
	}
	switch args["operation"] {
	case "add":
		return add(args, writer)
	case "list":
		return list(args, writer)
	case "findById":
		return findById(args, writer)
	case "remove":
		return remove(args, writer)
	default:
		return fmt.Errorf("Operation %s not allowed!", args["operation"])
	}
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	id := flag.String("id", "", "")
	item := flag.String("item", "", "")
	operation := flag.String("operation", "", "")
	fileName := flag.String("fileName", "", "")
	flag.Parse()
	return Arguments{
		"id":        *id,
		"item":      *item,
		"operation": *operation,
		"fileName":  *fileName,
	}
}

func add(args Arguments, writer io.Writer) error {
	if args["item"] == "" {
		return errors.New("-item flag has to be specified")
	}
	file, err := os.OpenFile(args["fileName"], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	db, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var person Users
	err = json.Unmarshal([]byte(args["item"]), &person)
	if err != nil {
		return err
	}

	var people []Users
	err = json.Unmarshal(db, &people)

	if len(db) == 0 {
		people = append(people, person)
		jsonPeople, err := json.Marshal(people)
		_, err = file.Write(jsonPeople)
		if err != nil {
			return err
		}
		return nil
	}

	for _, value := range people {
		if value.Id == person.Id {
			TextError := fmt.Sprintf("Item with id %s already exists", person.Id)
			_, err = writer.Write([]byte(TextError))
			if err != nil {
				return err
			}
			people = append(people, person)
		}
		jsonPeople, err := json.Marshal(people)

		_, err = file.Write(jsonPeople)
		if err != nil {
			return err
		}
	}
	return nil
}

func list(args Arguments, writer io.Writer) error {
	file, err := os.OpenFile(args["fileName"], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	db, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = writer.Write(db)
	if err != nil {
		return err
	}
	return nil
}

func findById(args Arguments, writer io.Writer) error {
	if args["id"] == "" {
		return errors.New("-id flag has to be specified")
	}
	file, err := os.OpenFile(args["fileName"], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	db, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	var person Users
	err = json.Unmarshal([]byte(args["item"]), &person)
	if err != nil {
		return err
	}
	var people []Users
	if len(db) > 0 {
		err = json.Unmarshal(db, &people)
	}
	for _, value := range people {
		if value.Id == person.Id {
			TextError := fmt.Sprintf("Item with id %s not found", args["id"])
			_, err = writer.Write([]byte(TextError))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func remove(args Arguments, writer io.Writer) error {
	if args["id"] == "" {
		return errors.New("-id flag has to be specified")
	}
	file, err := os.OpenFile(args["fileName"], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	db, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	var person Users
	err = json.Unmarshal([]byte(args["item"]), &person)
	if err != nil {
		return err
	}
	var people []Users
	if len(db) > 0 {
		err = json.Unmarshal(db, &people)
	}
	for _, value := range people {
		if value.Id == person.Id {
			TextError := fmt.Sprintf("Item with id %s not found", args["id"])
			_, err = writer.Write([]byte(TextError))
			if err != nil {
				return err
			}
		}
	}
	for key, value := range people {
		if value.Id == args["id"] {
			people = append(people[:key], people[key+1:]...)
			break
		}
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	jsonPeople, err := json.Marshal(people)

	_, err = file.Write(jsonPeople)
	if err != nil {
		return err
	}
	return nil
}
