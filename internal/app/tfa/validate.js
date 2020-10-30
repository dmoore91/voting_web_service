const speakeasy = require('speakeasy');

var arguments = process.argv;


// validates 2FA Token that user inputs
var verified = speakeasy.totp.verify({
    secret: arguments[2],
    encoding: 'ascii',
    token: arguments[3]
});

console.log(verified);
