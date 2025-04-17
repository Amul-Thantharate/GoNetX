package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error excuting command: %w output: %s", err, output)
	}
	return strings.TrimSpace(string(output)), nil
}

func GetOS() string {
	return runtime.GOOS
}

func DisplayIPAddresses(c *cli.Context) error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	color.Blue("IP Addresses:\n")
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			color.Yellow("  %s: %v\n", iface.Name, err)
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				color.Green("  Interface: ")
				color.Cyan("%s, ", iface.Name)
				color.Green("Address: ")
				color.Magenta("%s\n", ipNet.IP.String())
			}
		}
	}
	return nil
}

func DisplayHostName(c *cli.Context) error {
	hostName, err := os.Hostname()
	if err != nil {
		return err
	}
	color.Blue("Host Name: ")
	color.Cyan("%s\n", hostName)
	return nil
}

func RunPingCommand(c *cli.Context) error {
	host := c.Args().First()
	if host == "" {
		return cli.NewExitError("Please provide a host to ping", 1)
	}

	var cmd string
	var args []string

	osName := GetOS()
	switch osName {
	case "windows":
		cmd = "ping"
		args = []string{"-n", "4", host}
	case "darwin":
		cmd = "ping"
		args = []string{"-c", "4", host}
	case "linux":
		cmd = "ping"
		args = []string{"-c", "4", host}
	default:
		return cli.NewExitError(fmt.Sprintf("OS '%s' not supported", osName), 1)
	}

	output, err := ExecuteCommand(cmd, args...)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}

func DisplayOSInformation(c *cli.Context) error {
	osName := GetOS()
	color.Blue("Operating System: ")
	color.Magenta("%s\n", osName)

	switch osName {
	case "linux":
		output, err := ExecuteCommand("lsb_release", "-a")
		if err != nil {
			color.Red("Error getting Linux distribution info: %v\n", err) // Red
		} else {
			fmt.Println(output)
		}
	case "darwin":
		output, err := ExecuteCommand("sw_vers")
		if err != nil {
			color.Red("Error getting macOS version info: %v\n", err) // Red
		} else {
			fmt.Println(output)
		}
	case "windows":
		cmd := "systeminfo"
		args := []string{"|", "findstr", "/B", "/C:\"OS Name\"", "/C:\"OS Version\""}
		output, err := ExecuteCommand(cmd, args...)
		if err != nil {
			color.Red("Error getting Windows version info: %v\n", err)
		} else {
			fmt.Println(output)
		}

	default:
		color.Red("Operating system information not available for %s\n", osName) // Red
	}
	return nil
}

func DisplayRAMInformation(c *cli.Context) error {
	osName := GetOS()
	color.Blue("RAM Information:\n")
	switch osName {
	case "linux":
		output, err := ExecuteCommand("free", "-h")
		if err != nil {
			return fmt.Errorf("%s: %w", color.RedString("error getting RAM info"), err)
		}
		fmt.Println(output)
	case "darwin":
		output, err := ExecuteCommand("top", "-l", "1", "-s", "0", "-n", "0", "-stats", "rsize,vsize,pgrp,mem")
		if err != nil {
			return fmt.Errorf("%s: %w", color.RedString("error getting RAM info"), err)
		}
		fmt.Println(output)
	case "windows":
		cmd := "systeminfo"
		args := []string{"|", "findstr", "/C:\"Total Physical Memory\""}
		output, err := ExecuteCommand(cmd, args...)
		if err != nil {
			return fmt.Errorf("%s: %w", color.RedString("error getting RAM info"), err)
		}
		fmt.Println(output)
		fmt.Println(output)
	default:
		color.Red("RAM information not available for %s\n", osName)
	}
	return nil
}
func DisplayHardDiskInformation(c *cli.Context) error {
	osName := GetOS()
	color.Blue("Hard Disk Information:\n")
	switch osName {
	case "linux":
		output, err := ExecuteCommand("df", "-h")
		if err != nil {
			return fmt.Errorf("%s: %w", color.RedString("error getting disk info"), err)
		}
		fmt.Println(output)
	case "darwin":
		output, err := ExecuteCommand("df", "-h")
		if err != nil {
			return fmt.Errorf("%s: %w", color.RedString("error getting disk info"), err)
		}
		fmt.Println(output)
	case "windows":
		cmd := "powershell"
		args := []string{"Get-PSDrive", "-PSProvider", "FileSystem"}
		output, err := ExecuteCommand(cmd, args...)
		if err != nil {
			return fmt.Errorf("%s: %w", color.RedString("error getting disk info"), err)
		}
		fmt.Println(output)
	default:
		color.Red("Hard disk information not available for %s\n", osName)
	}
	return nil
}

func main() {
	app := &cli.App{
		Name:                 "net-Utils",
		Usage:                "A simple network utility CLI",
		Version:              "1.0.0",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "Ip",
				Usage:   "Display IP addresses",
				Aliases: []string{"address"},
				Action:  DisplayIPAddresses,
			},
			{
				Name:    "Display Hostname",
				Usage:   "Display the Host name",
				Aliases: []string{"hostname"},
				Action:  DisplayHostName,
			},
			{
				Name:    "Ping",
				Usage:   "Ping a host to check connected to the internet or not",
				Aliases: []string{"ping"},
				Action:  RunPingCommand,
			},
			{
				Name:    "Os Information",
				Usage:   "Display OS information",
				Action:  DisplayOSInformation,
				Aliases: []string{"os"},
			},
			{
				Name:    "Ram Information",
				Usage:   "Display RAM information",
				Action:  DisplayRAMInformation,
				Aliases: []string{"ram"},
			},
			{
				Name:    "Disk Information",
				Usage:   "Display Hard Disk information",
				Action:  DisplayHardDiskInformation,
				Aliases: []string{"disk"},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
