package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast		 // supports boradcast access capability
	FlagLoopback		// is a loopback interface
	FlagPointToPoint	// belongs to a point-to-point link
	FlagMulticast		// supports multicast access capability
)

func IsUp(v Flags) bool	{ return v&FlagUp == FlagUp}
func TurnDown(v *Flags)	{ *v &^= FlagUp}
func SetBroadcast(v *Flags) { *v |= FlagBroadcast}
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\
	
