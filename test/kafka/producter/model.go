package main

import "log"

var fakeDB string

func saveMessage(text string) {
	log.Printf("msg: [%s]", text)
	fakeDB = text
}

func getMessage() string { return fakeDB }