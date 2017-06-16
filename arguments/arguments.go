package arguments

import (
  "fmt"
  "os"
  "strconv"
)

type Arguments struct  {
	Array []string
	Help string
}

func (args *Arguments) ShowHelp () {
	fmt.Println()
	os.Exit(1)
}

func (args *Arguments) ParseArguments () bool {
	if args.GetLenght() == 0 {
		fmt.Println("Arguments error. Run with 'help' argument to learn about using.")
		return false
	}

	if args.Array[1] == "help" {
		args.ShowHelp()
	}

	return true
}

func (args *Arguments) GetLenght () int {
	return len(args.Array) - 1
}

func (args *Arguments) ToInteger (i int) int {
	s, _ := strconv.Atoi(args.Array[i])
	return s
}
