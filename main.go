package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/michalpodeszwa/gowatcher/entities"
	"github.com/michalpodeszwa/gowatcher/outputers"
	"github.com/michalpodeszwa/gowatcher/readers"
)

func handleFileCreated(people chan entities.Person, fileName string) {
	go readers.ReadFile(fileName, people)
}

func handlePersonCreated(people chan entities.Person, outputDir string, outputType string) {
	for person := range people {
		go outputers.CreateFile(person, outputDir, outputType)
	}
}

func createDirectoriesIfNotExist(directories []string) error {
	for _, directory := range directories {
		if err := os.MkdirAll(directory, 0755); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	parsedArguments := handleCliArguments()

	createDirectoriesIfNotExist(parsedArguments.GetDirectories())

	log.Println("Starting watching")
	watcher, _ := fsnotify.NewWatcher()

	defer watcher.Close()
	people := make(chan entities.Person)

	go handlePersonCreated(people, parsedArguments.OutputDir, parsedArguments.OutputType)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					go handleFileCreated(people, event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err := watcher.Add(parsedArguments.InputDir)
	if err != nil {
		log.Fatal(err, ". Make sure you've got input directory")
	}
	<-done
}
