package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

/**
 * Todo: BUILD: go build -o Yu -mod=vendor app/console/cmd/main.go | RUN: ./Yu dev --env=dev
 * Todo: RUN: go run -mod=vendor app/console/cmd/main.go dev
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
		fmt.Println("» args error!「 eg:dev 」")
		return
	}

	okEnv := true
	for _, val := range []string{"dev", "stg", "prd"} {
		if strings.Contains(args[0], val) {
			okEnv = false
		}
	}

	if okEnv {
		fmt.Println("» args error!「 must be around dev, stg or prd 」")
		return
	}

	console := New(args)

	env := viper.New()
	env.SetConfigType("yaml")
	env.AddConfigPath(".")
	env.SetConfigName(".env." + console.args[0])

	if err := env.ReadInConfig(); err != nil {
		fmt.Println("» config file error!「",".env."+console.args[0]+".yaml","」")
		return
	}

	ok := env.IsSet("App.Port")
	if ok {
		console.port = env.GetString("App.Port")
	} else {
		console.port = "8887"
	}

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
				console.stdout(err.Error())
				break
			}
		}

		if data == "1" || data == "stop" {
			if err := console.stop(true); err != nil {
				console.stdout(err.Error())
				break
			}
		}

		if data == "2" || data == "reload" {
			if err := console.reload(); err != nil {
				console.stdout(err.Error())
				break
			}
		}

		if data == "3" || data == "pid" {
			d, err := console.arrPids()
			if err != nil {
				console.stdout("PID", fmt.Sprintf("Error Process ID: %v", err.Error()))
			} else {
				console.stdout("PID", fmt.Sprintf("Shown Process ID: %v", d))
			}
		}

		if strings.ContainsAny(data, "-c") || strings.ContainsAny(data, "4") {
			console.stdout("CTRL", "Add Controller!")
		}

		if strings.ContainsAny(data, "-h") || strings.ContainsAny(data, "8") {
			help := strings.Split(data, ":")
			if len(help) != 2 {
				console.stdout("HELP", "Plz Add Helper Content!")
			} else {
				switch help[1] {
				case "run":
					fmt.Printf(`» 启动应用 start
  启动前请确保.env中环境变量配置正确
					`)
					break

				case "stop":
					console.stdout("停止应用")
					break

				case "reload":
					console.stdout("重启应用")
					break

				case "pid":
					console.stdout("查看正在运行中的进程号")
					break

				default:
					console.stdout("HELP", "Plz Add Correct Helper Content[run/stop/reload]!")
					break
				}
			}
		}

		if console.noCommand(data, numbers ...) == false {
			console.stdout("UNKNOWN", "command error")
		}

		if len(data) == 0 || data == "9" || data == "exit" {
			console.stdout("EXIT", "exit command!\n")
			break
		}
	}
}

type cmd struct {
	port string
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
	o, err := console.arrPids()

	if len(o) > 1 {
		err = console.stop(false)
		if err != nil {
			return
		}
	} else {
		if err == nil || o != nil {
			console.stdout(fmt.Sprintf("running ... [%v:%v]", o, err))
			_, err = console.isQuit()
			return
		}
	}

	if err = console.ready(); err != nil {
		return
	}

	argsEnv := "--env="+console.args[0]
	command := exec.Command("./start", argsEnv)

	console.stdout("go run ... ")
	if err = command.Start(); err != nil {
		return
	}

	console.stdout("RUN", fmt.Sprintf("application start, [PID] %d running ...", command.Process.Pid))
	_, err = console.isQuit()

	return
}

func (console *cmd) stop(Q bool) (err error) {
	o, err := console.arrPids()
	if err != nil {
		if Q {
			console.stdout("stopped")
			_, err = console.isQuit()
		}

		return
	}

	for _, pid := range o {
		command := exec.Command("kill", pid)
		if err = command.Run(); err != nil {
			return
		}

		console.stdout("END", fmt.Sprintf("application ended, [PID] %v stop ...", pid))
	}

	os.Remove("start")

	if Q {
		_, err = console.isQuit()
	}

	return
}

func (console *cmd) reload() (err error) {
	err = console.stop(false)
	if err != nil {
		err = errors.New(fmt.Sprintf("stopped! error: %v", err.Error()))
		return
	}

	err = console.start()
	return
}

func (console *cmd) ready() (err error) {
	console.stdout("go mod tidy")
	create := exec.Command("go", "mod", "tidy")
	if err = create.Run(); err != nil {
		return
	}

	console.stdout("go mod vendor")
	vendor := exec.Command("go", "mod", "vendor")
	if err = vendor.Run(); err != nil {
		return
	}

	console.stdout("go build ... ")
	builds := exec.Command("go", "build", "-o", "start", "-mod=vendor", "app/main.go")
	if err = builds.Run(); err != nil {
		return
	}

	return
}

func (console *cmd) isQuit() (d string, err error) {
	fmt.Printf("\n» quit the command?[y/n] › ")
	d, err = console.getReadLine()
	if err != nil {
		d = ""
		return
	}

	if strings.ToLower(d) == "y" {
		err = errors.New(fmt.Sprintf("quit command!\n"))
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

func (console *cmd) arrPids() (d []string, err error) {
	command := exec.Command("lsof", "-t", "-i", ":"+console.port)

	o, err := command.Output()
	if err != nil {
		err = errors.New(fmt.Sprintf("stopped!, No Process ID: %v", err.Error()))
		return
	}

	d = strings.Fields(string(o))
	return
}

func (console *cmd) noCommand(data string, nums ...interface{}) (b bool) {
	if len(nums) == 0 {
		return
	}

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

func (console *cmd) stdout(data ...string) {
	switch len(data) {
	case 1:
		fmt.Printf("» %v\n", data[0])
		break

	case 2:
		fmt.Printf("» %v » %v\n", data[0], data[1])
		break

	default:
		fmt.Printf("» unkown error!")
		break
	}
}
