package main

import "fmt"

func main() {
	flag := 656
	fmt.Println(getFilePermissions(flag))
}

func getFilePermissions(flag int) string {
	OwnerRights := flag / 100
	helpBit := 1
	ownerExecute := "-"
	if OwnerRights&helpBit != 0 {
		ownerExecute = "Execute"
	}

	helpBit = helpBit << 1
	ownerWrite := "-"
	if OwnerRights&helpBit != 0 {
		ownerWrite = "Write"
	}

	helpBit = helpBit << 1
	ownerRead := "-"
	if OwnerRights&helpBit != 0 {
		ownerRead = "Read"
	}

	GroupRights := flag % 100 / 10
	groupRead := "-"
	if GroupRights&helpBit != 0 {
		groupRead = "Read"
	}

	helpBit = helpBit >> 1
	groupWrite := "-"
	if GroupRights&helpBit != 0 {
		groupWrite = "Write"
	}

	helpBit = helpBit >> 1
	groupExecute := "-"
	if GroupRights&helpBit != 0 {
		groupExecute = "Execute"
	}

	OtherRights := flag % 10
	otherExecute := "-"
	if OtherRights&helpBit != 0 {
		otherExecute = "Execute"
	}

	helpBit = helpBit << 1
	otherWrite := "-"
	if OtherRights&helpBit != 0 {
		otherWrite = "Write"
	}

	helpBit = helpBit << 1
	otherRead := "-"
	if OtherRights&helpBit != 0 {
		otherRead = "Read"
	}

	return fmt.Sprintf("Owner:%s,%s,%s Group:%s,%s,%s Other:%s,%s,%s", ownerRead, ownerWrite, ownerExecute, groupRead,
		groupWrite, groupExecute, otherRead, otherWrite, otherExecute)
}
