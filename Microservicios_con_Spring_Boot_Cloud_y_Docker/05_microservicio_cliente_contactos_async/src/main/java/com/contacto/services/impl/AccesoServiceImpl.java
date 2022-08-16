package com.contacto.services.impl;

import com.contacto.model.Persona;
import com.contacto.services.IAccesoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.Arrays;
import java.util.List;
import java.util.concurrent.CompletableFuture;

@Service
public class AccesoServiceImpl implements IAccesoService {
    @Autowired
    RestTemplate template;

    String urlBase = "http://localhost:8080";

    @Async
    public CompletableFuture<List<Persona>> llamadaServicio(Persona persona){
        template.postForLocation(urlBase + "/contactos", persona); //Se ejecuta la peticion tupo POST
        Persona[] personas = template.getForObject(urlBase + "/contactos", Persona[].class); //Se obtiene la lista de los contactos
        return CompletableFuture.completedFuture(Arrays.asList(personas));
    }
}
