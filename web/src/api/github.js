import axios from 'axios'

const service = axios.create()

export function Commits(page) {
  return service({
    url: 'https://api.github.com/repos/defeng-hub/ByOfficeAutomatic/commits?page=' + page,
    method: 'get'
  })
}


