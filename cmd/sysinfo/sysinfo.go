package sysinfo

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "sysinfo",
	Short:   "Returns system information",
	Long:    `Returns information about the system formatted on a list`,
	Example: "sysinfo",
	Run: func(cmd *cobra.Command, args []string) {
		systemInfo := getSysInfo()
		fmt.Println(systemInfo)
	},
}

func getSysInfo() string {
	var stat syscall.Statfs_t

	// Get system memory information
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	systemMemory := memStats.Sys

	// Get system storage information
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Sprintf("Error getting current directory: %v", err)
	}
	syscall.Statfs(wd, &stat)
	storageSize := stat.Blocks * uint64(stat.Bsize)
	storageFree := stat.Bavail * uint64(stat.Bsize)

	// Get system CPU information
	numCPUs := runtime.NumCPU()

	// Get CPU usage percentages for each core
	percentages, err := cpu.Percent(time.Duration(1)*time.Second, true)
	if err != nil {
		fmt.Printf("Failed to get CPU usage: %s", err)
	}

	output := fmt.Sprintf("System Information:\n")
	output += fmt.Sprintf("- Memory: %s bytes\n", formatBytes(float64(systemMemory), 2))
	output += fmt.Sprintf("- Storage Size: %s\n", formatBytes(float64(storageSize), 2))
	output += fmt.Sprintf("- Storage Free: %s\n", formatBytes(float64(storageFree), 2))
	output += fmt.Sprintf("- Number of CPUs: %d\n", numCPUs)
	// Print CPU usage for each core
	for i, usage := range percentages {
		output += fmt.Sprintf("- CPU Usage Core %d: %.2f%%\n", i, usage)
	}

	return output
}

func formatBytes(bytes float64, precision int) string {
	base := 1024
	suffixes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	if bytes < 0 {
		return "0 B"
	}

	index := int(math.Floor(math.Log(bytes) / math.Log(float64(base))))
	value := bytes / math.Pow(float64(base), float64(index))
	return fmt.Sprintf(fmt.Sprintf("%%.%df %%s", precision), value, suffixes[index])
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
