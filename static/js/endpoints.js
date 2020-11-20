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

    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "https://localhost:8880/voting/user/login",
        success: function() {
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

    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "https://localhost:8880/voting/user",
        statusCode: {
            200: function() {
                window.location.href = './tfa.html';
            }
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

    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "https://localhost:8880/voting/tfa_validate",
        statusCode: {
            200: function() {
                storage['session_id'] = formData['token'];
                window.location.href = './dashboard.html';
            }
        }
    });
}

/**
 * Gets the candidates so users can see it
 */
function getCandidates() {
    $.ajax({
        type: 'GET',
        url: "https://localhost:8880/voting/candidate",
        statusCode: {
            200: function(data) {
                for (var i = 0; i < data['candidates'].length; i++) {
                    var elem = data['candidates'][i];
                    console.log(elem)
                    $('#candidates').append(
                        $('<input>').prop({
                            class: 'cand mr-2',
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

/**
 * Posts the user input into the database so it is recorded
 */
function postCandidate() {
    var radios = document.getElementsByClassName('cand');

    for (var i = 0, length = radios.length; i < length; i++) {
        if (radios[i].checked) {
            var username = $(radios[i]).attr('id');

            $.ajax({
                type: 'POST',
                dataType: "text",
                url: "https://localhost:8880/voting/vote/" + username
            });
        }
    }
}

/***
 * Gets the candidate votes, so it can be displayed in the results
 */
function getCandidateVotes() {
    $.ajax({
        type: 'GET',
        url: "https://localhost:8880/voting/vote",
        statusCode: {
            200: function (data) {
                console.log(data['candidates'])
                for (var i = 0; i < data['candidates'].length; i++) {
                    $('.result').append(
                        data['candidates'][i].first_name + ' '
                    ).append(data['candidates'][i].last_name + ' Count: '
                    ).append(data['candidates'][i].votes
                    ).append(' Party: ' + data['candidates'][i].party
                    ).append('<br>')
                }
            }
        }
    });
}

/**
 * Authenticates the user that enters the dashboard for security purposes
 */
function authenticateUser() {
    var localStorage = window.localStorage;

    $.ajax({
        type: 'GET',
        data: 'text',
        url: "https://localhost:8880/voting/session/" + localStorage.getItem("username") + "/" + localStorage.getItem("session_id"),
        complete: function(xhr, textStatus) {
            if (xhr.status !== 200) {
                window.location.href = './index.html';
            }
        }
    });
}
