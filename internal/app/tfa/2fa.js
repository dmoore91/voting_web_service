const speakeasy = require('speakeasy');
const qrcode = require('qrcode');

// generates 2FA secret for the user to have so the user can use 2FA
// this will be stored in the User table
// every user will have unique secret to distinctly allow that user to 2FA
var secret = speakeasy.generateSecret({
    name: 'voting_two_factor_authenticator'
});

// displays the secret and the data
qrcode.toDataURL(secret.otpauth_url, function(err, data) {
    console.log( JSON.stringify({
        'secret': secret['ascii'],
        'data': data
    }))
});
