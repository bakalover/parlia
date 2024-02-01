package idgen

import "math/rand"

var initPSeed, initASeed, initLSeed = 61, 17, 29

func GenerateInitConfig() (uint8, uint8, uint8) {
	kProposers := rand.Intn(initPSeed)*2 + 3
	kAcceptors := rand.Intn(initASeed)*2 + 1 // Prefferable to be odd (Fault-tolerance points)
	kLearners := rand.Intn(initLSeed)*2 + 3
	return uint8(kProposers), uint8(kAcceptors), uint8(kLearners)
}
