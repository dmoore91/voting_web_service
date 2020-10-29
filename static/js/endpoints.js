/**
 * Gets the data from the form and converts it into a Hashmap
 * @param formName name of the form that contains the data
 * @return HashMap that contains data from form
 */
function getFormData(formName) {
    var formdata = $(formName).serializeArray(),
        data = {};

    $(formdata ).each(function(index, obj){
        data[obj.name] = obj.value;
    });

    return data;
}


/**
 * Gets triggered when the user signs in
 */
function signin() {
    var formData = getFormData('#signin-form'),
        storage = window.localStorage;

    storage['username'] = formData['username'];

    JSON.stringify(formData);

    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "https://localhost:8880/voting/user/login",
        success:  function() {
            window.location.href = './tfa.html';
        }
    });
}


/**
 * Gets triggered when the user signs up
 */
function signup() {
    var formData = getFormData('#signup-form'),
        secret = $('#secret').data('secret'),
        storage = window.localStorage;

    storage['username'] = formData['username'];
    storage['secret_key'] = secret;
    formData['secret_key'] = secret;

    JSON.stringify(formData);

    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "https://localhost:8880/voting/user",
        success:  function() {
            window.location.href = './tfa.html';
        }
    });
}

/**
 * Validates the user input for 2FA
 */
function validate2FA() {
    var formData = getFormData('#tfa-form'),
        storage = window.localStorage;

    formData['username'] = storage['username'];

    JSON.stringify(formData);

    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "https://localhost:8880/voting/tfa_validate",
        success:  function() {
            window.location.href = './dashboard.html';
        }
    });
}