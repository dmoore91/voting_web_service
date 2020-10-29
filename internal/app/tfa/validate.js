const speakeasy = require('speakeasy');

var arguments = process.argv;

var verified = speakeasy.totp.verify({
    secret: arguments[2],
    encoding: 'ascii',
    token: arguments[3]
});

console.log(verified);
