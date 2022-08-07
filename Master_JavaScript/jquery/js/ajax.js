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

    $('#formulario').submit(function(e){
        e.preventDefault();
        var usuario = {
            "name": $('input[name="name"]').val(),
            "web" : $('input[name="web"]').val(),
        }
    
        console.log(usuario);
    
        // $.post($(this).attr("action"), usuario, function(response){
        //     console.log(response);
        // }).done(function(){
        //     alert("Usuario añadido correctamente");
        // });

        $.ajax({
            type: 'POST',
            // dataType: 'json',
            // contectType: 'application/json',
            url: $(this).attr("action"),
            data: usuario,
            beforeSend: function(){
                console.log("enviando usuario...");
            },
            success: function(response){
                console.log(response);
            },
            error: function(){
                console.log("Ha ocurrido un error");
            },
            timeout: 5000
        });

        return false;
    });
});