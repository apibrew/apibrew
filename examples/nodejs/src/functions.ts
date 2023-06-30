import axios from "axios"

export async function test1(params) {
    const response = await axios.get('http://tisserv.net:9009/docs/')

    console.log('Loading axios !@!! 22222')

    return 'test1 loaded'
}

export async function test2(params) {
    const response = await axios.get('http://tisserv.net:9009/docs/')

    console.log('Loading axios !@!! 11111')

    return 'test2 loaded'
}