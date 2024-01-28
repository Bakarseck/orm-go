const token = "eyJhbGciOiJhbGciLCJ0eXAiOiJqd3QifQ.eyJOYW1lIjoiQmFrYXIiLCJJZCI6MTJ9.FZSUE8G6uAilMKM-_viokDdcrH8Lmz92EDLXUVcefns"

const parts = token.split('.');

function base64UrlDecode(str) {
    str = str.replace(/-/g, '+').replace(/_/g, '/');
    switch (str.length % 4) {
        case 0:
            break;
        case 2:
            str += '==';
            break;
        case 3:
            str += '=';
            break;
        default:
            throw 'Chaîne Base64URL invalide';
    }
    return atob(str);
}

const payload = base64UrlDecode(parts[1]);
const payloadObj = JSON.parse(payload);

console.log(payloadObj);
