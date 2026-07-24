package system

import (
	"os"
	"runtime"

	"go-agent/internal/version"
)

type Info struct {
	MachineID    string `json:"machine_id"`
	Hostname     string `json:"hostname"`
	AgentVersion string `json:"agent_version"`
	OS           string `json:"os"`
	Arch         string `json:"arch"`
}

func GetInfo() (*Info, error) {
	machineID, err := MachineID()
	if err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Info{
		MachineID:    machineID,
		Hostname:     hostname,
		AgentVersion: version.Version,
		OS:           runtime.GOOS,
		Arch:         runtime.GOARCH,
	}, nil
}
