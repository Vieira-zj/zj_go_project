/**
Created at 2019-11-15, webshell terminal demo for k8s pod.
Refer: https://github.com/maoqide/kubeutil

Workflow: xterm.js => websocket (client) => websocket (server) => k8s client remotecommand exec stdin/stdout => pod
Test URL: file:///local_path/to/terminal.html?namespace=mini-test-ns&pod=hello-minikube-59ddd8676b-vkl26
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	k8ssvc "tools.app/services/k8sio"
	wssvc "tools.app/services/webshell"
)

var (
	defaultPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	kubeConfig  = flag.String("kubeconfig", defaultPath, "abs path to the kubeconfig file")
	addr        = flag.String("addr", ":8090", "http service address")
	cmd         = []string{"/bin/sh"}
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/terminal", serveTerminal)
	router.HandleFunc("/ws/{namespace}/{pod}/{container_name}/webshell", serveWs)

	log.Println("http server (websocket) is started at :8090...")
	log.Fatal(http.ListenAndServe(*addr, router))
}

func internalError(conn *websocket.Conn, msg string, err error) {
	log.Printf("message: %s, error: %v\n", msg, err)
	conn.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
}

func serveTerminal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./static/terminal.html")
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	namespace := pathParams["namespace"]
	pod := pathParams["pod"]
	containerName := pathParams["container_name"]
	log.Printf("ws request: exec pod:%s, container:%s, namespace:%s", pod, containerName, namespace)

	term, err := wssvc.NewTerminalSession(w, r, nil)
	if err != nil {
		log.Printf("get terminal session failed: %v", err)
		return
	}
	defer func() {
		log.Println("close session")
		term.Close()
	}()

	k8sClient, err := k8ssvc.NewK8SClient(*kubeConfig)
	if err != nil {
		log.Printf("init k8s client failed: %v", err)
		return
	}

	// check and set container name
	if containerName != "null" {
		ok, err := k8sClient.CheckPod(namespace, pod, containerName)
		if err != nil {
			log.Printf("check pod failed: pod:%s, container:%s, namespace:%s\n", pod, containerName, namespace)
			return
		}
		if !ok {
			msg := fmt.Sprintf("Validate pod error: %v", err)
			log.Println(msg)
			term.Write([]byte(msg))
			term.Done()
			return
		}
	} else {
		pod, err := k8sClient.GetPod(namespace, pod)
		if err != nil {
			log.Printf("get pod failed: pod:%s, namespace:%s\n", pod, namespace)
			return
		}
		containerName = pod.Spec.Containers[0].Name
	}

	// term is pty handler for exec pod stdin and stdout
	if err := wssvc.ExecPod(k8sClient.KubeClient, k8sClient.KubeConfig, cmd, term, namespace, pod, containerName); err != nil {
		msg := fmt.Sprintf("Exec pod error: %v", err)
		log.Println(msg)
		term.Write([]byte(msg))
		term.Done()
	}
}
