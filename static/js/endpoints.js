/**
 * Gets the data from the form and converts it into a Hashmap
 * @param formName name of the form that contains the data
 * @return HashMap that contains data from form
 */
function getFormData(formName) {
    var formdata = $(formName).serializeArray();
    var data = {};
    $(formdata ).each(function(index, obj){
        data[obj.name] = obj.value;
    });
    return data;
}


/**
 * Gets triggered when the user signs in
 */
function signin() {
    var formData = getFormData('#signin-form');
    JSON.stringify(formData);
    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "http://localhost:8880/voting/login",
        statusCode: {
            200: function(responseObject, textStatus, jqXHR) {
                window.location.href = './dashboard.html';
            }
        }
    });
}


/**
 * Gets triggered when the user signs up
 */
function signup() {
    var formData = getFormData('#signup-form');
    JSON.stringify(formData);
    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "http://localhost:8880/voting/user",
        statusCode: {
            200: function(responseObject, textStatus, jqXHR) {
                window.location.href = './dashboard.html';
            }
        }
    });
}
