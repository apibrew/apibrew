class crypto {
    static getRandomValues(arr) {
        for (let i = 0; i < arr.length; i++) {
            arr[i] = Math.floor(Math.random() * 256);
        }
        return arr;
    }
}

class ReadableStream {
    constructor() {
        this.buffer = [];
    }

    getReader() {
        return {
            read: () => {
                return new Promise((resolve, reject) => {
                    if (this.buffer.length > 0) {
                        resolve({value: this.buffer.shift(), done: false});
                    } else {
                        resolve({done: true});
                    }
                });
            }
        }
    }

    put(data) {
        this.buffer.push(data);
    }
}

class TextEncoder {
    encode(str) {
        return new Uint8Array([...str].map(c => c.charCodeAt(0)));
    }
}

class Headers {
    constructor() {
        this.headers = {};
    }

    append(key, value) {
        this.headers[key] = value;
    }
}

class Request {
    constructor(url, options) {
        this.url = url;
        this.options = options;
    }
}

class URL {
    constructor(url) {
        this.url = url;
    }
}
