// Mock Utils

const methodGet = 'get'
const methodPost = 'post'
const resultPass = 'success'
const resultFail = 'failed'
// elements id
const registerResultId = 'reg_result'
const mockResultId = 'mock_result'

function buildQuery(params) {
  items = []
  for (let param of params) {
    items.push(`${param.key}=${param.value}`)
  }
  return items.join("&")
}

function addField(fields) {
  fields.push({idx: fields.length, key: '', value: ''})
}

function removeField(fields) {
  if (fields.length === 1) {
    return
  }
  fields.pop()
}

function filterHeaders(headers) {
  let retHeaders = {'Content-Type': 'application/json;charset=utf-8'}
  for (let header of headers) {
    if (Boolean(header.key) && Boolean(header.value)) {
      retHeaders[header.key] = header.value
    }
  }
  return retHeaders
}

function updateResultElementStyle(id, result) {
  let s = document.querySelector(`#${id}`).style
  s.fontWeight = 'bold'
  s.color = result ? 'green' : 'red'
}
