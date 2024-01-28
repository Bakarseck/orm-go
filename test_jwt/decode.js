const token = "eyJhbGciOiJhbGciLCJ0eXAiOiJqd3QifQ.eyJOYW1lIjoiQmFrYXIiLCJJZCI6MTJ9.FZSUE8G6uAilMKM-_viokDdcrH8Lmz92EDLXUVcefns"

const parts = token.split('.');

function base64UrlDecode(str) {
    str = str.replace(/-/g, '+').replace(/_/g, '/');
    switch (str.length % 4) {
        case 0:
            // Pas de padding nécessaire
            // La longueur de la chaîne est déjà un multiple de 4
            break;
        case 2:
            // Ajouter deux caractères de padding
            // La chaîne a besoin de 2 caractères supplémentaires pour atteindre un multiple de 4
            str += '==';
            break;
        case 3:
            // Ajouter un caractère de padding
            // La chaîne a besoin de 1 caractère supplémentaire pour atteindre un multiple de 4
            str += '=';
            break;
        default:
            // Si aucun des cas ci-dessus n'est vrai, alors la chaîne est invalide pour du Base64
            throw 'Chaîne Base64URL invalide';
    }
    return atob(str);
}

const payload = base64UrlDecode(parts[1]);
const payloadObj = JSON.parse(payload);

console.log(payloadObj);
