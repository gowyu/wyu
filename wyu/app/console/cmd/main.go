package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/**
 * Todo: BUILD: go build -o Yu -mod=vendor app/console/cmd/main.go | RUN: ./Yu dev --env=dev
 * Todo: RUN: go run -mod=vendor app/console/cmd/main.go dev --env=dev
**/

var (
	contents string = `
0 › RUN		» run 	 « 	
1 › STOP	» stop 	 « 
2 › RELOAD	» reload « 
3 › PID		» pid 	 « 
4 › CTRL	» -c	 «
8 › HELP	» -h 	 « » -h:run | -h:stop | -h:reload | -h:pid «
9 › EXIT	» exit 	 « 
	`
	numbers []interface{} = []interface{}{
		"0", "run",
		"1", "stop",
		"2", "reload",
		"3", "pid",
		"4", "-c",
		"8", "-h", "-h=run", "-h=stop", "-h=reload",
		"9", "exit",
	}
)

func main() {
	flag.Parse()
	args := flag.Args()

	if args == nil || len(args) != 1 {
		fmt.Println("» args error! (eg: dev)")
		return
	}

	console := New(args)
	if err := console.checkEnvironment(); err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf(contents)

	for{
		fmt.Printf("\n» Command Line › ")
		data, _ := console.getReadLine()
		fmt.Printf("» ››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››››› ‹\n")

		if data == "0" || data == "run" {
			if err := console.start(); err != nil {
				fmt.Println(err.Error())
				break
			}
		}

		if data == "1" || data == "stop" {
			if err := console.stop(true); err != nil {
				fmt.Println(err.Error())
				break
			}
		}

		if data == "2" || data == "reload" {
			if err := console.reload(); err != nil {
				fmt.Println(err.Error())
				break
			}
		}

		if data == "3" || data == "pid" {
			d, err := console.showPid()
			if err != nil {
				fmt.Printf("» PID » Error Process ID: %v\n", err.Error())
			} else {
				fmt.Printf("» PID » Shown Process ID: %s\n", d)
			}
		}

		if strings.ContainsAny(data, "-c") || strings.ContainsAny(data, "4") {
			fmt.Printf("» CTRL » Add Controller!\n")
		}

		if strings.ContainsAny(data, "-h") || strings.ContainsAny(data, "8") {
			help := strings.Split(data, ":")
			if len(help) != 2 {
				fmt.Printf("» HELP » Plz Add Helper Content!\n")
			} else {
				switch help[1] {
				case "run":
					fmt.Printf(`» 启动应用 start
  启动前请确保.env中环境变量配置正确
					`)
					break

				case "stop":
					fmt.Printf("» 停止应用\n")
					break

				case "reload":
					fmt.Printf("» 重启应用\n")
					break

				case "pid":
					fmt.Printf("» 查看正在运行中的进程号\n")
					break

				default:
					fmt.Printf("» Help » Plz Add Correct Helper Content[run/stop/reload]!\n")
					break
				}
			}
		}

		if console.noCommand(data, numbers ...) == false {
			fmt.Printf("» Unknown » command error « %v \n", data)
		}

		if len(data) == 0 || data == "9" || data == "exit" {
			fmt.Printf("» EXIT » exit command!\n\n")
			break
		}
	}
}

type cmd struct {
	args []string
	reader *bufio.Reader
}

func New(args []string) *cmd {
	return &cmd{
		args: args,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (console *cmd) start() (err error) {
	exelsof := exec.Command("lsof", "-t", "-i", ":8887")
	_, err = exelsof.Output()
	if err == nil {
		fmt.Printf("» running ... [%v]\n", err)
		_, err = console.isQuit()
		return
	}

	if err = console.ready(); err != nil {
		return
	}

	argsEnv := "--env="+console.args[0]
	command := exec.Command("./start", argsEnv)

	err = command.Start()
	if err != nil {
		return
	}

	fmt.Printf("» RUN  » application start, [PID] %d running ...\n", command.Process.Pid)
	_, err = console.isQuit()

	return
}

func (console *cmd) stop(Q bool) (err error) {
	exelsof := exec.Command("lsof", "-t", "-i", ":8887")
	o, err := exelsof.Output()
	if err != nil {
		if Q {
			fmt.Printf("» stopped!\n")
			_, err = console.isQuit()
		}
		return
	}

	oFields := strings.Fields(string(o))
	command := exec.Command("kill", oFields[0])

	err = command.Start()
	if err != nil {
		return
	}

	fmt.Printf("» STOP » application ended, [PID] %v stop ...\n", oFields[0])

	if Q {
		_, err = console.isQuit()
	}

	return
}

func (console *cmd) reload() (err error) {
	err = console.stop(false)
	if err != nil {
		err = errors.New(fmt.Sprintf("» stopped! error:%v\n", err.Error()))
		return
	}

	err = console.start()
	return
}

func (console *cmd) ready() (err error) {
	create := exec.Command("go", "mod", "tidy")
	if err = create.Start(); err != nil {
		return
	}

	vendor := exec.Command("go", "mod", "vendor")
	if err = vendor.Start(); err != nil {
		return
	}

	builds := exec.Command("go", "build", "-o", "start", "-mod=vendor", "app/main.go")
	if err = builds.Start(); err != nil {
		return
	}

	return
}

func (console *cmd) isQuit() (d string, err error) {
	fmt.Printf("\n» do you wanna quit the command?[y/n]: ")

	d, err = console.getReadLine()
	if err != nil {
		d = ""
		return
	}

	if strings.ToLower(d) == "y" {
		err = errors.New(fmt.Sprintf("» quit command!\n\n"))
		return
	}

	return
}

func (console *cmd) getReadLine() (d string, err error) {
	var src []byte

	src, _, err = console.reader.ReadLine()
	if err != nil {
		d = ""
		return
	}

	d = string(src)
	return
}

func (console *cmd) showPid() (d string, err error) {
	command := exec.Command("lsof", "-t", "-i", ":8887")
	o, err := command.Output()
	if err != nil {
		err = errors.New(fmt.Sprintf("» stopped!, No Process ID: %v", err.Error()))
		return
	}

	d = strings.Fields(string(o))[0]
	return
}

func (console *cmd) noCommand(data string, nums ...interface{}) (b bool) {
	for _, val := range nums {
		if strings.Contains(data, val.(string)) {
			b = true
			return
		}
	}

	return
}

func (console *cmd) checkEnvironment() (err error) {
	golang := exec.Command("go", "version")
	o, err := golang.Output()
	if err != nil {
		err = errors.New(fmt.Sprintf("\n» Environment Error » Plz install golang in the first place\n\n"))
		return
	}

	if strings.Contains(string(o), "1.13") == false {
		err = errors.New(fmt.Sprintf("\n» Environment Error » Plz use go1.13!\n\n"))
		return
	}

	goEnvs := exec.Command("go", "env")
	o, err  = goEnvs.Output()
	if err != nil {
		err = errors.New(fmt.Sprintf("\n» Environment Error » unknown error\n\n"))
		return
	}

	goproxy := ""
	res := strings.Split(string(o), "\n")
	for _, val := range res {
		src := strings.Split(val, "=")
		if src[0] == "GOPROXY" {
			goproxy = src[1]
		}
	}

	if goproxy == "" {
		err = errors.New(fmt.Sprintf("\n» Environment Error » GOPROXY is not existed!\n\n"))
		return
	}

	if strings.Contains(goproxy, "goproxy.cn") == false {
		err = errors.New(fmt.Sprintf(`
» Environment Error » Plz change GOPROXY: https://goproxy.cn,direct
                    » command » go env -w GOPROXY=https://goproxy.cn,direct
		`)+"\n")
		return
	}

	return
}
