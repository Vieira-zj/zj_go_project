function getQueryVariable (variable) {
  let query = window.location.search.substring(1)
  let vars = query.split('&')

  for (let i = 0; i < vars.length; i++) {
    let pair = vars[i].split('=')
    if (pair[0] == variable) {
      return pair[1]
    }
  }
  return false
}

// backend service address
const ws_host = 'localhost:8090'

function connectWS (namespace, pod, container_name) {
  // for debug: file:///local_path/to/terminal.html?namespace=mini-test-ns&pod=hello-minikube-59ddd8676b-vkl26
  // let namespace = getQueryVariable('namespace')
  // let pod = getQueryVariable('pod')
  // let container_name = getQueryVariable('container_name')

  if (!Boolean(container_name)) {
    container_name = 'null'
  }

  if (!Boolean(namespace) || !Boolean(pod)) {
    alert('namespace or pod is empty in query!')
    return
  }
  console.log(`ns: ${namespace}, pod: ${pod}, container: ${container_name}`)

  let url = `ws://${ws_host}/ws/${namespace}/${pod}/${container_name}/webshell`
  console.log(`connect to ws url: ${url}`)

  let term = new Terminal({
    'cursorBlink': true,
  })

  if (!window['WebSocket']) {
    let item = document.getElementById('terminal')
    item.innerHTML = '<h2>Your browser does not support WebSockets.</h2>'
    return
  }

  term.open(document.getElementById('terminal'))
  term.write(`connecting to pod ${pod}...`)
  term.fit()
  // term.toggleFullScreen(true)

  // send req data to backend websocket
  term.on('data', function (data) {
    msg = { operation: 'stdin', data: data }
    conn.send(JSON.stringify(msg))
  })
  term.on('resize', function (size) {
    console.log('resize:', size)
    msg = { operation: 'resize', cols: size.cols, rows: rows }
    conn.send(JSON.stringify(msg))
  })

  // init pod term env
  conn = new WebSocket(url)
  conn.onopen = function (e) {
    term.write('\r')
    msg = { operation: 'stdin', data: 'export TERM=xterm\r' }
    conn.send(JSON.stringify(msg))
    // term.clear()
  }
  // write resp data to term
  conn.onmessage = function (event) {
    msg = JSON.parse(event.data)
    if (msg.operation === 'stdout') {
      term.write(msg.data)
    } else {
      console.warn('invalid msg operation:', msg)
    }
  }
  conn.onclose = function (event) {
    if (event.wasClean) {
      console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`)
    } else {
      console.warn('[close] Connection died')
      term.writeln("")
    }
    term.write('Connection Reset By Peer! Try Refresh.')
  }
  conn.onerror = function (error) {
    console.error('[error] Connection error')
    term.write("error: " + error.message)
    term.destroy()
  }
}
