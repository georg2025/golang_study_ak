package main

import "fmt"

func main() {
	flag := 656
	fmt.Println(getFilePermissions(flag))
}

func getFilePermissions(flag int) string {
	Ownerrights := flag / 100
	helpbit := 1
	ownerexecute := "-"
	if Ownerrights&helpbit != 0 {
		ownerexecute = "Execute"
	}
	helpbit = helpbit << 1
	ownerwrite := "-"
	if Ownerrights&helpbit != 0 {
		ownerwrite = "Write"
	}
	helpbit = helpbit << 1
	ownerread := "-"
	if Ownerrights&helpbit != 0 {
		ownerread = "Read"
	}
	Grouprights := flag % 100 / 10
	groupread := "-"
	if Grouprights&helpbit != 0 {
		groupread = "Read"
	}
	helpbit = helpbit >> 1
	groupwrite := "-"
	if Grouprights&helpbit != 0 {
		groupwrite = "Write"
	}
	helpbit = helpbit >> 1
	groupexecute := "-"
	if Grouprights&helpbit != 0 {
		groupexecute = "Execute"
	}
	Otherrights := flag % 10
	otherexecute := "-"
	if Otherrights&helpbit != 0 {
		otherexecute = "Execute"
	}
	helpbit = helpbit << 1
	otherwrite := "-"
	if Otherrights&helpbit != 0 {
		otherwrite = "Write"
	}
	helpbit = helpbit << 1
	otherread := "-"
	if Otherrights&helpbit != 0 {
		otherread = "Read"
	}

	return fmt.Sprintf("Owner:%s,%s,%s Group:%s,%s,%s Other:%s,%s,%s", ownerread, ownerwrite, ownerexecute, groupread,
		groupwrite, groupexecute, otherread, otherwrite, otherexecute)
}
