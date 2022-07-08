package main

import (
	"encoding/json"
	"errors"
	"flag"
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
	if args["item"] == "" {
		return errors.New("-item flag has to be specified")
	}
	if args["id"] == "" {
		return errors.New("-id flag has to be specified")
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
	}
	return nil
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
	file, err := os.OpenFile("fileName", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var item Users
	err = json.Unmarshal([]byte(args["item"]), &item)
	if err != nil {
		return err
	}
	var people []Users
	for _, value := range people {
		if value.Id == item.Id {
			errors.New("Item with this Id is already existed")
		}
		people = append(people, item)
	}
	jsonPeople, err := json.Marshal(people)

	_, err = file.Write(jsonPeople)
	if err != nil {
		return err
	}
	return nil
}

func list(args Arguments, writer io.Writer) error {

	return nil
}

func findById(args Arguments, writer io.Writer) error {

	return nil
}

func remove(args Arguments, writer io.Writer) error {

	return nil
}
