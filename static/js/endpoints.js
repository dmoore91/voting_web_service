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
    return data
}

function signin() {
    var formData = getFormData('#signin-form');
}


function signup(event) {
    var formData = getFormData('#signup-form');
    JSON.stringify(formData);

    $.ajax({
        type: 'POST',
        data: JSON.stringify(formData),
        dataType: "text",
        url: "http://localhost:8880/voting/user"
    }).then(function(e) {
        console.log(e)
        debugger;
    });
}
