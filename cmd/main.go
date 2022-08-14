package main

import (
	animalizer "civilrights3.com/animalese/internal"
	_ "embed"
	"log"
)

func main() {
	conf, err := animalizer.ProcessParams()
	if err != nil {
		log.Fatalf("Error(s) processing command line parameters: %v", err)
	}

	err = animalizer.Convert(conf)
	if err != nil {
		log.Fatalf("Error converting text to audio: %v", err)
	}

	log.Println("Operation completed successfully")
}
