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

}
