package main

import (
	"errors"
	"flag"
	"io"
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
