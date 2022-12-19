import axios from 'axios'

export async function getData(url: string) {
    const res = axios.get(url).then((response) => {
        return response
    })
    return res
}