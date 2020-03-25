package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/SimplyVC/oasis_api_server/src/config"
	lgr "github.com/SimplyVC/oasis_api_server/src/logger"
)

// Setting data to test with, valid and invalid path locations
const (
	mainConfigFile = "testdata/test_user_config_main.ini"
	nodesFile      = "testdata/test_config_nodes.ini"
	mainFileFail   = "testdata/test_config_main_fail.ini"
	nodesFileFail  = "testdata/test_config_nodes_fail.ini"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	os.Chdir("../")

	// Set Logger that will be used by API through all packages
	lgr.SetLogger(os.Stdout, os.Stdout, os.Stderr)
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

// Testing loading of default port configuration file
func TestLoadMainConfiguration_Success_1(t *testing.T) {
	mainConf := config.LoadMainConfiguration()
	if mainConf == nil {
		t.Errorf("Failed to load port file from path.")
	}
}

// Testing loading of default socket configuration file
func TestLoadNodesConfig_Success_1(t *testing.T) {
	nodesConf := config.LoadNodesConfiguration()
	if nodesConf == nil {
		t.Errorf("Failed to load socket file from path.")
	}
}

// Testing successful retrieval of ports once configuration is loaded
func TestLoadMainConfiguration_Success_2(t *testing.T) {
	mainConf := config.LoadNodesConfiguration()
	if mainConf == nil {
		t.Errorf("Failed to load port file from path.")
	}
	mainConf_1 := config.GetMain()
	if mainConf_1 == nil {
		t.Errorf("Failed to load port file from path.")
	}
}

// Testing successful retrieval of sockets once configuration is loaded
func TestLoadNodesConfig_Success_2(t *testing.T) {
	nodesConf := config.LoadNodesConfiguration()
	if nodesConf == nil {
		t.Errorf("Failed to load socket file from path.")
	}
	nodesConf_1 := config.GetNodes()
	if nodesConf_1 == nil {
		t.Errorf("Failed to load socket file from path.")
	}
}

// Testing failed loading of another port configuration file
func TestLoadMainConfiguration_Failure_1(t *testing.T) {
	config.SetMainFile(mainFileFail)
	mainConf := config.LoadMainConfiguration()
	if mainConf != nil {
		t.Errorf("Failed to not load port file from path.")
	}
}

// Testing failed loading of another socket configuration file
func TestLoadNodesConfig_Failure_1(t *testing.T) {
	config.SetNodesFile(nodesFileFail)
	nodesConf := config.LoadNodesConfiguration()
	if nodesConf != nil {
		t.Errorf("Failed to not load socket file from path.")
	}
}
