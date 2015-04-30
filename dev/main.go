package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func checkPath(cmd string) (ret string, err error) {
	ret, err = exec.LookPath(cmd)
	if err != nil {
		err = fmt.Errorf("%s: executable file not found in $PATH", cmd)
		return
	}

	return ret, nil
}

func runCommand(data string, args []string) (ret string, err error) {
	bin, err := checkPath(args[0])
	if err != nil {
		return
	}

	cmd := exec.Command(bin)
	cmd.Args = args

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}

	io.WriteString(stdin, data)
	stdin.Close()
	out, err := cmd.Output()
	if err != nil {
		return
	}
	ret = string(out)

	return
}

func runLikePipe(args []string) (ret string, err error) {
	ret = args[0]

	for i := 1; i < len(args); i++ {
		if args[i] == "" {
			continue
		}

		ret, err = runCommand(ret, strings.Split(args[i], " "))
		if err != nil {
			return
		}
		ret = strings.TrimRight(ret, "\n")
	}

	return
}

func runCommandDirectly(arg string) (ret string, err error) {
	args := strings.Split(arg, " ")
	bin, err := checkPath(args[0])
	if err != nil {
		return
	}
	cmd := exec.Command(bin)
	cmd.Args = args[0:]

	out, err := cmd.Output()
	if err != nil {
		return
	}
	ret = string(out)

	return
}

func main() {
	file := "/Users/b4b4r07/a.go"
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	//p, _ := pygmentize(file)
	d, err := runLikePipe([]string{
		string(contents),
		"nkf",
		"pygmentize -O style=solarized -f console256 -g",
		//func() string {
		//	//if false {
		//	//	return "pygmentize -l " + lexerExtension(file)
		//	//}
		//	//return "pygmentize -O style=solarized -f console256 -g" // + " -l " + lexerExtension(file)
		//}(),
		//func() (ret string) {
		//	if false {
		//		ret = "cat -n"
		//	}
		//	return
		//}(),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	//d, _ = runCommandDirectly("pygmentize -O style=solarized -f console256 -g" + file)
	//d, _ = pygmentize(file)
	fmt.Println(d)
}

func pygmentize(f string) (ret string, err error) {
	cmd, err := checkPath("pygmentize")
	if err != nil {
		return
	}

	out, err := exec.Command(
		cmd,
		"-O",
		"style=solarized",
		"-f",
		"console256",
		"-g",
		f,
	).Output()
	if err != nil {
		return
	}

	ret = string(out)
	return
}
