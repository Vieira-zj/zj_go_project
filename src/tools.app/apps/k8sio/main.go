package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	mysvc "tools.app/services/k8sio"
)

var (
	kubeConfig, namespace, podName *string
)

func main() {
	// Refer: https://github.com/kubernetes/kubernetes/tree/v1.6.1/test/e2e/framework
	log.Println("K8S Client app start.")

	defaultPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	kubeConfig = flag.String("kubeconfig", defaultPath, "abs path to the kubeconfig file")
	namespace = flag.String("nspace", "mini-test-ns", "specified namespace")
	podName = flag.String("podname", "hello-minikube-59ddd8676b-vkl26", "specified pod name")
	flag.Parse()

	// for debug
	if true {
		printClusterInfo()
		printK8SResources()
	}

	if false {
		// cli := &ExecLocalCommand{}
		cli := &ExecRemoteCommand2{}
		execCliCommand(cli)
	}
}

func printClusterInfo() {
	client, err := mysvc.NewK8SClient(*kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Cluster pods info:")
	if err := client.PrintNumberOfAllPods(); err != nil {
		panic(err.Error())
	}

	log.Printf("Namespace [%s] pod [%s] info:\n", *namespace, *podName)
	if err := client.PrintPodInfo(*namespace, *podName); err != nil {
		panic(err.Error())
	}
}

func printK8SResources() {
	var testNs, testPod string

	client, err := mysvc.NewK8SClient(*kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Cluster all namespaces name:")
	if ns, err := client.GetAllNamespacesName(); err != nil {
		panic(err.Error())
	} else {
		log.Println(strings.Join(ns, ","))
		for _, name := range ns {
			if strings.Contains(name, "test") {
				testNs = name
			}
		}
	}

	log.Printf("Namespace [%s] all pods name:\n", testNs)
	pods, err := client.GetPodsNameByNamespace(testNs)
	if err != nil {
		panic(err.Error())
	}
	log.Println(strings.Join(pods, ","))

	testPod = pods[0]
	log.Printf("Namespace [%s] pod [%s] all containers name:\n", testNs, testPod)
	containers, err := client.GetContainersNameByNsAndPod(testNs, testPod)
	if err != nil {
		panic(err.Error())
	}
	log.Println(strings.Join(containers, ","))
}

func execCliCommand(cli ExecCommand) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("go-cli$ ")
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}

		cmdStr = strings.TrimSuffix(cmdStr, "\n")
		if len(cmdStr) == 0 {
			continue
		}
		cmdSlice := strings.Fields(cmdStr)
		if cmdSlice[0] == "exit" {
			log.Println("go-cli shell exit.")
			os.Exit(0)
		}

		if err = cli.run(ExecOptions{
			Configs:   *kubeConfig,
			Namespace: *namespace,
			PodName:   *podName,
			Command:   cmdStr,
		}); err != nil {
			panic(err.Error())
		}
	}
}

// ExecOptions contains exec options to run a command.
type ExecOptions struct {
	Configs   string
	Namespace string
	PodName   string
	Command   string
}

// ExecCommand an interface to run a command.
type ExecCommand interface {
	run(ExecOptions) error
}

// ExecLocalCommand runs local command.
type ExecLocalCommand struct {
}

func (cli *ExecLocalCommand) run(options ExecOptions) error {
	cmdSlice := strings.Fields(options.Command)
	cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExecRemoteCommand runs remote command on given pod in namespace with output in stdout.
type ExecRemoteCommand struct {
}

func (cli *ExecRemoteCommand) run(options ExecOptions) error {
	re, err := mysvc.NewRemoteExec(options.Configs)
	if err != nil {
		return err
	}
	return re.RunCommandWithStdout(options.Namespace, options.PodName, options.Command)
}

// ExecRemoteCommand2 runs remote command on given pod in namespace, and returns stdout string.
type ExecRemoteCommand2 struct {
}

func (cli *ExecRemoteCommand2) run(options ExecOptions) error {
	rc2, err := mysvc.NewRemoteExec2(options.Configs)
	if err != nil {
		return err
	}

	cmds := strings.Split(options.Command, " ")
	stdout, stderr, err := rc2.RunCommandInPod(options.Namespace, options.PodName, cmds...)
	if err != nil {
		if len(stderr) > 0 {
			log.Println("Error:", stderr)
		}
		return err
	}
	fmt.Print(stdout)
	return nil
}
