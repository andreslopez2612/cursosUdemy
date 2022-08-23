package com.contactos.controllers;

import com.contactos.model.Persona;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.HttpStatusCodeException;
import org.springframework.web.client.RestTemplate;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

@RestController
public class PersonaController {

    @Autowired
    RestTemplate template;

    String urlBase = "http://localhost:8080";

    @GetMapping(value = "/personas/{nombre}/{email}/{edad}", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<List<Persona>> altaPersona(@PathVariable("nombre") String nombre,
                                                     @PathVariable("email") String email,
                                                     @PathVariable("edad") int edad) {
        Persona persona = new Persona(nombre, email, edad);//Se crea la variable "persona" la cual arma el objeto JSON

        try {
            /**
             * Al lanzar la peticion post, si la respuesa contiene un codigo de error
             * diferente a 200, se producira una excepcion
             */
            template.postForLocation(urlBase + "/contactos", persona); //Se ejecuta la peticion tupo POST
            /**
             * Si todo ha ido bien, devolvemos como respuesta la lsita de persona
             * e indicamos OK como codigo de estado de respuesta
             */
            Persona[] personas = template.getForObject(urlBase + "/contactos", Persona[].class); //Se obtiene la lista de los contactos
            return new ResponseEntity<List<Persona>>(Arrays.asList(personas), HttpStatus.OK); //Se retorna la lista en un array JSON
        } catch (HttpStatusCodeException exception) {
            /**
             * Si hubo un error en la llamada al microservicio, enviamos a nuestro lcienta final
             * una cabecera con el mensaje de error, una lista vacia de personas en el cuerpo
             */
            HttpHeaders headers = new HttpHeaders();
            headers.add("Error", exception.getResponseBodyAsString());
            return new ResponseEntity<List<Persona>>(new ArrayList<Persona>(), headers, exception.getStatusCode());
        }


    }

    @GetMapping(value = "/personas/{edad1}/{edad2}", produces = MediaType.APPLICATION_JSON_VALUE)
    public List<Persona> buscarEdades(@PathVariable("edad1") int edad1, @PathVariable("edad2") int edad2) {
        Persona[] personas = template.getForObject(urlBase + "/contactos", Persona[].class);
        return Arrays.stream(personas).filter(p -> p.getEdad() >= edad1 && p.getEdad() <= edad2)
                .collect(Collectors.toList());
    }
}
