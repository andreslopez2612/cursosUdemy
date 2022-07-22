$(document).ready(()=>{
    //https://reqres.in/
    // $("#datos").load()
    $.get("https://reqres.in/api/users", {page: 2}, function(response){
        console.log(response);
        response.data.forEach((element, index) => {
            console.log(element);
            $("#datos").append("<p>" + element.first_name+" "+ element.last_name +"</p>");
        });
    })

    var usuario = {
        "name":"Andres Lopez",
        "web" : "andreslopez.com"
    }

    $.post("https://reqres.in/api/users", usuario, function(response){
        console.log(response);
    })
});