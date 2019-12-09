package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var digitFlipReplacer = strings.NewReplacer(
	"0", "0",
	"1", "Ɩ",
	"2", "ᄅ",
	"3", "Ɛ",
	"4", "ᔭ",
	"5", "ϛ",
	"6", "9",
	"7", "Ɫ",
	"8", "8",
	"9", "6",
)

var charmap = map[string]string{}

var flip = "(╯°□°)╯︵ "
var angryflip = "(ノಠ益ಠ)ノ彡 "
var putdown = " ノ( º _ ºノ)"

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func flipPid(pid string, signal int) string {
	switch signal {
	case 9:
		return angryflip + reverse(digitFlipReplacer.Replace(pid))
	default:
		return flip + reverse(digitFlipReplacer.Replace(pid))
	}
}

func usage() string {
	return "Usage: tfkill [-s signum] pids"
}

func main() {
	var signal int
	var pids []string

	if len(os.Args) == 1 {
		fmt.Println(usage())
		os.Exit(1)
	}
	flag.Usage = func() { fmt.Println(usage()) }
	flag.IntVar(&signal, "s", 15, "signal to send to pids")
	flag.Parse()
	pids = flag.Args()
	for _, pid := range pids {
		cmd := exec.Command("kill", "-s", strconv.Itoa(signal), string(pid))
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println(pid + putdown)
			continue
		}
		fmt.Println(flipPid(pid, signal))
	}
}
