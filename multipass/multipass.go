package multipass

import "strings"

const (
	Name        = "Name:"
	State       = "State:"
	IPv4        = "IPv4:"
	Release     = "Release:"
	ImageHash   = "Image hash:"
	Load        = "Load:"
	DiskUsage   = "Disk usage:"
	MemoryUsage = "Memory usage:"
)

type Instance struct {
	Name        string
	IP          string
	State       string
	Image       string
	ImageHash   string
	DiskUsage   string
	TotalDisk   string
	MemoryUsage string
	MemoryTotal string
	Load        string
}

func parseInstance(out string) *Instance {

	var instance Instance

	for _, line := range strings.Split(out, "\n") {

		if strings.Contains(line, Name) {
			instance.Name = strings.TrimSpace(strings.ReplaceAll(line, Name, ""))
		}

		if strings.Contains(line, State) {
			instance.State = strings.TrimSpace(strings.ReplaceAll(line, State, ""))
		}

		if strings.Contains(line, IPv4) {
			instance.IP = strings.TrimSpace(strings.ReplaceAll(line, IPv4, ""))
		}

		if strings.Contains(line, Release) {
			instance.Image = strings.TrimSpace(strings.ReplaceAll(line, Release, ""))
		}

		if strings.Contains(line, ImageHash) {
			instance.ImageHash = strings.TrimSpace(strings.ReplaceAll(line, ImageHash, ""))
		}

		if strings.Contains(line, Load) {
			instance.Load = strings.TrimSpace(strings.ReplaceAll(line, Load, ""))
		}

		if strings.Contains(line, DiskUsage) {
			diskUsage := strings.TrimSpace(strings.ReplaceAll(line, DiskUsage, ""))
			diskUsageOut := strings.Split(diskUsage, "out of")
			instance.DiskUsage = strings.TrimSpace(diskUsageOut[0])
			instance.TotalDisk = strings.TrimSpace(diskUsageOut[1])
		}

		if strings.Contains(line, MemoryUsage) {
			memoryUsage := strings.TrimSpace(strings.ReplaceAll(line, MemoryUsage, ""))
			memoryUsageOut := strings.Split(memoryUsage, "out of")
			instance.MemoryUsage = strings.TrimSpace(memoryUsageOut[0])
			instance.MemoryTotal = strings.TrimSpace(memoryUsageOut[1])
		}
	}

	return &instance
}
