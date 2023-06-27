import axios from "axios";


export default async function (params) {
    const response = await axios.get('http://tisserv.net:9009/docs/')

    return response.data
}