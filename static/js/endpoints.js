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

/**
 * Gets the candidates
 */
function getCandidates() {
    $.ajax({
        type: 'GET',
        url: "https://localhost:8880/voting/candidate",
        statusCode: {
            200: function(data) {
                for (var i = 0; i < data['candidates'].length; i++) {
                    console.log('candidates', data['candidates'][i])
                    var elem = data['candidates'][i];
                    console.log(elem)
                    $('#candidates').append(
                        $('<input>').prop({
                            class: 'mr-2',
                            type: 'radio',
                            id: elem['username'],
                            name: elem['party'],
                            party: elem['party']
                        })
                    ).append(
                        $('<label>').prop({
                            for: elem['party']
                        }).html('Candidate: ' + elem['first_name'] + ' ' +  elem['last_name'] +' Party: ' + elem['party'])
                    ).append('<br>');
                }
                $('#candidates').append('<button type="submit" value="Submit" class="signupbtn">Submit Candidate</button>')
            }
        }
    });
}

function postCandidate() {
    var radios = document.getElementsByClassName('cand');

    for (var i = 0, length = radios.length; i < length; i++) {
        if (radios[i].checked) {
            var username = $(radios[i]).attr('id');

            $.ajax({
                type: 'POST',
                dataType: "text",
                url: "https://localhost:8880/voting/" + username
            });
        }
    }

    // TODO verify this
    $.ajax({
        type: 'GET',
        dataType: "text",
        url: "https://localhost:8880/voting",
        statusCode: {
            200: function (data) {
                debugger;
                console.log('daata', data);
            }
        }
    });
}
