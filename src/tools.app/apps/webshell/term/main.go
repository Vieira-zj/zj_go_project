/**
Created at 2019-11-15, webshell terminal demo for k8s pod.

Refer: https://github.com/maoqide/kubeutil
Workflow: xterm.js => websocket (client) => websocket (server) => k8s client remotecommand exec stdin/stdout => pod

Build: ./gorun.sh webshell
Access: open static/terminal.html on local.
Verify: kc exec -it containers-pod -n mini-test-ns -c busybox sh
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"src/mock.server/common"

	k8ssvc "src/tools.app/services/k8sio"
	wssvc "src/tools.app/services/webshell"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	defaultPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	kubeConfig  = flag.String("kubeconfig", defaultPath, "abs path to the kubeconfig file")
	addr        = flag.String("addr", ":8090", "http service address")
	cmd         = []string{"/bin/sh"}
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/query/ns", getAllNamespaces)
	router.HandleFunc("/query/pods", getAllPodsByNamespace)
	router.HandleFunc("/query/containers", getAllContainersByNsAndPod)

	router.HandleFunc("/terminal", serveTerminal)
	router.HandleFunc("/ws/{namespace}/{pod}/{container_name}/webshell", serveWs)

	log.Println("http server (websocket) is started at :8090...")
	log.Fatal(http.ListenAndServe(*addr, router))
}

// ------------------------------
// K8S resources query api (json)
//
// Namespaces: curl -v "http://localhost:8090/query/ns" | jq
// Pods: curl -v "http://localhost:8090/query/pods?ns=mini-test-ns" | jq
// Containers: curl -v "http://localhost:8090/query/containers?ns=mini-test-ns&pod=containers-pod" | jq
//
// ------------------------------

type respJSONData struct {
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data"`
}

type respErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type respK8SResources struct {
	Namespaces []string `json:"namespaces,omitempty"`
	Pods       []string `json:"pods,omitempty"`
	Containers []string `json:"containers,omitempty"`
}

func getAllNamespaces(w http.ResponseWriter, r *http.Request) {
	client, err := k8ssvc.NewK8SClient(*kubeConfig)
	if err != nil {
		log.Println("Init k8s client error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	namespaces, err := client.GetAllNamespacesName()
	if err != nil {
		log.Println("Get cluster all namespaces error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Get cluster all namespaces")
	writeOkJSONResp(w, respK8SResources{Namespaces: namespaces})
}

func getAllPodsByNamespace(w http.ResponseWriter, r *http.Request) {
	namespace := "default"

	values := r.URL.Query()
	if val, ok := values["ns"]; ok {
		namespace = val[0]
	} else {
		log.Printf("Use default namespace [%s] to query pods\n", namespace)
	}

	client, err := k8ssvc.NewK8SClient(*kubeConfig)
	if err != nil {
		log.Println("Init k8s client error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pods, err := client.GetPodsNameByNamespace(namespace)
	if err != nil {
		log.Printf("Get namespace [%s] all pods error: %s\n", namespace, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Get namespace [%s] all pods\n", namespace)
	writeOkJSONResp(w, respK8SResources{Pods: pods})
}

func getAllContainersByNsAndPod(w http.ResponseWriter, r *http.Request) {
	var namespace, pod string

	values := r.URL.Query()
	if val, ok := values["ns"]; ok {
		namespace = val[0]
	} else {
		errResp := respErrorMsg{
			Status:  499,
			Message: "Namespace is not set in query of request url",
		}
		writeJSONRespWithStatus(w, http.StatusNotAcceptable, errResp)
		return
	}

	if val, ok := values["pod"]; ok {
		pod = val[0]
	} else {
		errResp := respErrorMsg{
			Status:  499,
			Message: "Pod is not set in query of request url",
		}
		writeJSONRespWithStatus(w, http.StatusNotAcceptable, errResp)
		return
	}

	client, err := k8ssvc.NewK8SClient(*kubeConfig)
	if err != nil {
		log.Println("Init k8s client error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	containers, err := client.GetContainersNameByNsAndPod(namespace, pod)
	if err != nil {
		log.Printf("Get namespace [%s] and pod [%s] all containers error: %s\n", namespace, pod, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Get namespace [%s] and pod [%s] all containers\n", namespace, pod)
	writeOkJSONResp(w, respK8SResources{Containers: containers})
}

func writeOkJSONResp(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Set(common.TextContentType, common.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&respJSONData{Data: data}); err != nil {
		log.Println("Write json encoded response error:", err.Error())
	}
}

func writeJSONRespWithStatus(w http.ResponseWriter, retcode int, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Set(common.TextContentType, common.ContentTypeJSON)
	w.WriteHeader(retcode)

	if err := json.NewEncoder(w).Encode(&respJSONData{Data: data}); err != nil {
		log.Println("Write json encoded response error:", err.Error())
	}
}

// ------------------------------
// K8S webshell
// ------------------------------

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
