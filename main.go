package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "kubeswitches",
		Short: "A utility to manage Kubernetes config files",
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(setCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Kubernetes config files",
	Run: func(cmd *cobra.Command, args []string) {
		kubeConfigs := os.Getenv("KUBECONFIGS")
		if kubeConfigs == "" {
			fmt.Println("Environment variable KUBECONFIGS is not set.")
			return
		}

		files, err := ioutil.ReadDir(kubeConfigs)
		if err != nil {
			fmt.Printf("Error reading KUBECONFIGS directory: %v\n", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("No Kubernetes config files found.")
			return
		}

		fmt.Println("Available Kubernetes config files:")
		for _, file := range files {
			if !file.IsDir() {
				fmt.Println("-", file.Name())
			}
		}
	},
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current Kubernetes config",
	Run: func(cmd *cobra.Command, args []string) {
		kubeConfigPath, err := getKubeConfigPath()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		content, err := ioutil.ReadFile(kubeConfigPath)
		if err != nil {
			fmt.Printf("Error reading current kube config: %v\n", err)
			return
		}

		c := exec.Command("kubectl", "config", "view", "--minify")
		output, err := c.Output()
		if err != nil {
			fmt.Printf("Error running kubectl config view: %v\n", err)
			fmt.Printf("Raw kubeconfig file: %v\n", string(content))
			return
		}

		lastConfig := ""
		configFile := filepath.Join(os.Getenv("HOME"), ".kube", ".last_config")
		data, err := os.ReadFile(configFile)
		if err == nil {
			lastConfig = string(data)
		}

		fmt.Printf("Current kubeconfig: %v\n", lastConfig)
		fmt.Println(string(output))
	},
}

var setCmd = &cobra.Command{
	Use:   "set <configfilename>",
	Short: "Set the specified Kubernetes config as current",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configName := args[0]
		kubeConfigs := os.Getenv("KUBECONFIGS")
		if kubeConfigs == "" {
			fmt.Println("Environment variable KUBECONFIGS is not set.")
			return
		}

		sourcePath := filepath.Join(kubeConfigs, configName)
		if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
			fmt.Printf("Config file %s does not exist in %s\n", configName, kubeConfigs)
			return
		}

		kubeConfigPath, err := getKubeConfigPath()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		err = os.Remove(kubeConfigPath)
		if err != nil && !os.IsNotExist(err) {
			fmt.Printf("Error removing existing kube config: %v\n", err)
			return
		}

		err = copyFile(sourcePath, kubeConfigPath)
		if err != nil {
			fmt.Printf("Error setting new kube config: %v\n", err)
			return
		}

		configFile := filepath.Join(os.Getenv("HOME"), ".kube", ".last_config")
		err = os.WriteFile(configFile, []byte(configName), 0644)
		if err != nil {
			fmt.Printf("Error saving config name: %v\n", err)
			return
		}

		fmt.Printf("Switched to config %s successfully.\n", configName)
	},
}

func getKubeConfigPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, ".kube", "config"), nil
}

func copyFile(source, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destDir := filepath.Dir(destination)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
