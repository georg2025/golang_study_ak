package main

import "fmt"

func main() {
}

type LgTV struct {
	status bool
	model  string
}

type SamsungTV struct {
	status bool
	model  string
}

type switchOn interface {
	switchOFF() error
	switchOn() error
}

type GetStatus interface {
	GetStatus() bool
}

type SamsungHub interface {
	SamsungHub()
}

type LgHub interface {
	LGHub()
}

type GetModel interface {
	GetModel() string
}

func (f LgTV) GetModel() string {
	return f.model
}

func (f SamsungTV) GetModel() string {
	return f.model
}

func (f LgTV) GetStatus() bool {
	return f.status
}

func (f SamsungTV) GetStatus() bool {
	return f.status
}

func (f *LgTV) switchOFF() error {
	if f.status {
		f.status = false
		return nil
	}
	return fmt.Errorf("tv is already turned off")
}

func (f *LgTV) switchOn() error {
	if !f.status {
		f.status = true
		return nil
	}
	return fmt.Errorf("tv is already turned on")
}

func (f *SamsungTV) switchOFF() error {
	if f.status {
		f.status = false
		return nil
	}
	return fmt.Errorf("tv is already turned off")
}

func (f *SamsungTV) switchOn() error {
	if !f.status {
		f.status = true
		return nil
	}
	return fmt.Errorf("tv is already turned on")
}

func (f LgTV) LGHub() {
	fmt.Println("Connecting to LG Hub")
}

func (f SamsungTV) SamsungHub() {
	fmt.Println("Connecting to Samsung Hub")
}
