const token = "eyJhbGciOiJhbGciLCJ0eXAiOiJqd3QifQ.eyJOYW1lIjoiQmFrYXIiLCJJZCI6MTJ9.FZSUE8G6uAilMKM-_viokDdcrH8Lmz92EDLXUVcefns"

const payloadObj = getPayLoad(token);

function getPayLoad(token) {
    const parts = token.split('.');
    const payload = parts[1].replace(/-/g, '+').replace(/_/g, '/');
    const paddedPayload = payload.padEnd(payload.length + (4 - payload.length % 4) % 4, '=');
    const decodedPayload = atob(paddedPayload);
    const payloadObj = JSON.parse(decodedPayload);
    return payloadObj;
}

console.log(payloadObj);

