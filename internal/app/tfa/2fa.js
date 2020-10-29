const speakeasy = require('speakeasy');
const qrcode = require('qrcode');

/**
 * Generates 2FA secret for a specific user when user account is being created
 *
 * @retuns dictionary with the secret and the qrcode
 */

// generates 2FA secret for the user to have so the user can use 2FA
// this will be stored in the User table
// every user will have unique secret to distinctly allow that user to 2FA
var secret = speakeasy.generateSecret({
    name: 'voting_two_factor_authenticator'
});

// console.log(secret);

qrcode.toDataURL(secret.otpauth_url, function(err, data) {
    // console.log(data);

    console.log( JSON.stringify({
        'secret': secret['ascii'],
        'data': data
    }))
});



