package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	//"strconv"
	"strings"
	"sync"
)

const internal_default_style = "solarized"

func isExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getStyle() (ret []string, err error) {
	python_code := `from pygments.styles import get_all_styles
s = list(get_all_styles())
print ' '.join(s)`

	c, err := checkPath("python")
	if err != nil {
		return
	}
	cmd := exec.Command(c)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}

	io.WriteString(stdin, python_code)
	stdin.Close()
	out, err := cmd.Output()
	if err != nil {
		return
	}
	ret = strings.Split(strings.TrimRight(string(out), "\n"), " ")

	return
}

func setStyle(s string) (ret string, err error) {
	styles, err := getStyle()
	if err != nil {
		return
	}

	if stringInSlice(s, styles) {
		ret = s
	} else {
		ret = "default"
	}

	return
}

func checkPath(cmd string) (ret string, err error) {
	ret, err = exec.LookPath(cmd)
	if err != nil {
		err = fmt.Errorf("%s: executable file not found in $PATH", cmd)
		return
	}

	return ret, nil
}

func main() {
	var style = flag.String("s", "", "pygmentize style")
	var verbose = flag.Bool("v", false, "verbose")
	//flag.Usage = "a"
	flag.Parse()
	if *style == "" {
		*style = internal_default_style
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Input file is missing\n")
		os.Exit(1)
	}

	p, err := checkPath("pygmentize")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if false {
		var wg sync.WaitGroup
		for _, item := range args {
			wg.Add(1)
			go func(item string) {
				defer wg.Done()

				if !isExists(item) {
					fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", item)
					os.Exit(1)
				}

				style, err := setStyle(*style)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}

				out, err := exec.Command(
					p,
					"-O",
					"style="+style,
					"-f",
					"console256",
					"-g",
					item,
				).Output()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}

				fmt.Print(string(out))
			}(item)
		}
		wg.Wait()
	} else {
		var wg sync.WaitGroup
		for _, item := range args {
			wg.Add(1)
			go func(item string) {
				defer wg.Done()

				if !isExists(item) {
					fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", item)
					os.Exit(1)
				}

				style, err := setStyle(*style)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}

				cmd := exec.Command(
					p,
					"-O",
					"style="+style,
					"-f",
					"console256",
					"-g",
					//"-l",
					//lexerExtension(item),
				)
				stdin, err := cmd.StdinPipe()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}

				io.WriteString(stdin, nkf(item))
				stdin.Close()
				out, err := cmd.Output()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}

				fmt.Print(string(out))
				if *verbose {
					//max, _ := strconv.Atoi(os.Getenv("LINES"))
					//line := ""
					//for i := 0; i < max; i++ {
					//	line = line + "-"
					//}
					//fmt.Println(line)
					fmt.Println("------------------")
				}
			}(item)
		}
		wg.Wait()
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func lexerExtension(s string) (ret string) {
	extension := strings.TrimLeft(filepath.Ext(s), ".")
	if extension == "" {
		lines, _ := readLines(s)
		basename := filepath.Base(lines[0])
		if ok, _ := regexp.MatchString("^env", basename); ok {
			t := strings.Split(lines[0], " ")
			basename = t[1]
		}
		ret = basename
	} else {
		ret = extension
	}

	return
}

func outputData(s []byte) (ret string) {
	ret = string(s)

	c, err := checkPath("nkf")
	if err != nil {
		return
	}
	cmd := exec.Command(c)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}

	io.WriteString(stdin, ret)
	stdin.Close()
	out, err := cmd.Output()
	if err != nil {
		return
	}
	ret = string(out)

	return
}

func nkf(s string) (ret string) {
	cmd, err := checkPath("nkf")
	if err != nil {
		a, _ := readLines(s)
		return strings.Join(a, "\n")
	}

	out, err := exec.Command(
		cmd,
		s,
	).Output()
	if err != nil {
		return
	}

	ret = string(out)
	return
}

// にほんごであそぼ
// vim: fdm=marker
