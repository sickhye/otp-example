// decode.js

// Retrieve Base64 data from command line arguments
const base64Data = process.argv[2]; // First argument (0 is node execution, 1 is file name)

// Base64 decoding function
function base64Decode(base64) {
    const decodedString = Buffer.from(base64, 'base64').toString('utf-8');
    return decodedString;
}

// Function to convert JSON string to object
function parseJson(jsonString) {
    return JSON.parse(jsonString); // Convert JSON string to object
}

// Decode Base64 data and parse JSON
const decodedJsonData = base64Decode(base64Data);
const jsonData = parseJson(decodedJsonData);

// Output GUID and OTP
console.log("Clinet (Javascript)");
console.log("GUID:", jsonData.guid);
console.log("OTP:", jsonData.otp);
