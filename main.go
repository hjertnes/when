package main

import (
	"fmt"
	"git.sr.ht/~hjertnes/when/data"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

func help(){
	fmt.Println("when is a small utility to keep track of when you did something")
	fmt.Println("Usage:")
	fmt.Println("\t when list")
	fmt.Println("\t when set \"name\"")
	fmt.Println("\t when remove \"name\"")
	fmt.Println("\t when help")
	fmt.Println("Configuration")
	fmt.Println("\t the data is stored to a yaml file at ~/txt/when.yml or to the env var WHEN_HOME if set")
}

func list(){
	db, err := data.Read()
	if err != nil{
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "When"})

	for title, date := range db{
		table.Append([]string{
			title,
			date.Format("2006-01-02"),
		})
	}

	table.Render()
}

func set(name string){
	db, err := data.Read()
	if err != nil{
		panic(err)
	}
	db[name] = time.Now()

	err = data.Write(db)
	if err != nil{
		panic(err)
	}
}

func remove(name string){
	db, err := data.Read()
	if err != nil{
		panic(err)
	}

	delete(db, name)

	err = data.Write(db)
	if err != nil{
		panic(err)
	}
}


func main(){
	numberOfArguments := len(os.Args)

	if numberOfArguments == 1{
		help()
		return
	}

	if os.Args[1] == "list"{
		list()
		return
	} else if os.Args[1] == "set" {
		if numberOfArguments == 3 {
			set(os.Args[2])
			return
		}
	} else if os.Args[1] == "remove" {
		if numberOfArguments == 3 {
			remove(os.Args[2])
			return
		}
	}

	help()
	return
}
