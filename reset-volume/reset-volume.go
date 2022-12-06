package main

import (
	"github.com/onyx-and-iris/voicemeeter"
	"log"
)

func main() {
	vm, err := vmConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	vm.Bus[0].SetGain(0)
}

func vmConnect() (*voicemeeter.Remote, error) {
	vm, err := voicemeeter.NewRemote("banana", 20)
	if err != nil {
		return nil, err
	}

	err = vm.Login()
	if err != nil {
		return nil, err
	}

	return vm, nil
}
