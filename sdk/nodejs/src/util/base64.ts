export function decodeBase64(base64String: string) {
    const base64Chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/';
    let result = '';
    let currentBitPattern = 0;
    let bitCount = 0;

    for (let i = 0; i < base64String.length; i++) {
        if (base64String[i] === '=') {
            break;
        }

        currentBitPattern = (currentBitPattern << 6) | base64Chars.indexOf(base64String[i]);
        bitCount += 6;

        if (bitCount >= 8) {
            bitCount -= 8;
            result += String.fromCharCode((currentBitPattern >> bitCount) & 0xFF);
        }
    }

    return result;
}
