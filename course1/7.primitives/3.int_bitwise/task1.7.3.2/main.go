package main

import "fmt"

func main() {
	flag := 656
	fmt.Println(getFilePermissions(flag))
}

func getFilePermissions(flag int) string {
	ownerRights := flag / 100
	helpBit := 1
	ownerExecute := "-"
	if ownerRights&helpBit != 0 {
		ownerExecute = "Execute"
	}

	helpBit = helpBit << 1
	ownerWrite := "-"
	if ownerRights&helpBit != 0 {
		ownerWrite = "Write"
	}

	helpBit = helpBit << 1
	ownerRead := "-"
	if ownerRights&helpBit != 0 {
		ownerRead = "Read"
	}

	groupRights := flag % 100 / 10
	groupRead := "-"
	if groupRights&helpBit != 0 {
		groupRead = "Read"
	}

	helpBit = helpBit >> 1
	groupWrite := "-"
	if groupRights&helpBit != 0 {
		groupWrite = "Write"
	}

	helpBit = helpBit >> 1
	groupExecute := "-"
	if groupRights&helpBit != 0 {
		groupExecute = "Execute"
	}

	otherRights := flag % 10
	otherExecute := "-"
	if otherRights&helpBit != 0 {
		otherExecute = "Execute"
	}

	helpBit = helpBit << 1
	otherWrite := "-"
	if otherRights&helpBit != 0 {
		otherWrite = "Write"
	}

	helpBit = helpBit << 1
	otherRead := "-"
	if otherRights&helpBit != 0 {
		otherRead = "Read"
	}

	return fmt.Sprintf("Owner:%s,%s,%s Group:%s,%s,%s Other:%s,%s,%s", ownerRead, ownerWrite, ownerExecute, groupRead,
		groupWrite, groupExecute, otherRead, otherWrite, otherExecute)
}
