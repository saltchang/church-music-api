package helper

import "strconv"

// SidGenerator struct
type SidGenerators struct{}

// GenSID method
func (SidGenerator *SidGenerators) GenSID(lang string, numC string, numI string) (success bool, newSID string) {
	success = true
	if lang == "Chinese" {
		newSID = "1"
	} else if lang == "Taiwanese" {
		newSID = "2"
	} else {
		success = false
	}

	if len(numC) < 1 || len(numI) < 1 {
		return false, ""
	}

	numCInt, err := strconv.Atoi(numC)
	if err != nil || numCInt < 1 || numCInt > 999 {
		return false, ""
	}
	numIInt, err := strconv.Atoi(numI)
	if err != nil || numIInt < 1 || numIInt > 999 {
		return false, ""
	}

	if len(numC) > 3 || len(numI) > 3 {
		return false, ""
	}

	if len(numC) < 2 {
		newSID = newSID + "00" + numC
	} else if len(numC) < 3 {
		newSID = newSID + "0" + numC
	} else {
		newSID = newSID + numC
	}

	if len(numI) < 2 {
		newSID = newSID + "00" + numI
	} else if len(numI) < 3 {
		newSID = newSID + "0" + numI
	} else {
		newSID = newSID + numI
	}

	return true, newSID
}

var (
	// SidGenerator var
	SidGenerator = new(SidGenerators)
)
