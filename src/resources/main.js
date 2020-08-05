var server = "http://localhost:8000/";

function addItemToDOM(text) {
    var item = $('#weather-result');
    item.innerText = text;
}

function getIlkleyWeather() {
    var result = null;
    $.ajax({
        type: "GET",
        url: server + "ilkley",
            async: false,
        success: function(data) {
            result = data;
        } 
    });
    addItemToDOM(result);
}


$(function () {
    getIlkleyWeather();
})