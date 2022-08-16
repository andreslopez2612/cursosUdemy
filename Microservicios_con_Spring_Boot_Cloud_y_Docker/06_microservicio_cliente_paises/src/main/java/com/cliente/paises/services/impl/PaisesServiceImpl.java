package com.cliente.paises.services.impl;

import com.cliente.paises.model.Pais;
import com.cliente.paises.services.PaisesService;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Service
public class PaisesServiceImpl implements PaisesService {

    String url = "https://restcountries.com/v3.1/all";

    @Autowired
    RestTemplate template;

    @Override
    public List<Pais> obtenerPaises() {
        String resultado = template.getForObject(url, String.class);
        ObjectMapper maper = new ObjectMapper();
        List<Pais> paises = new ArrayList<>();
        ArrayNode array;
        try {
            //obtenemos ArrayJoson con los datos de la respuesta
            array = (ArrayNode) maper.readTree(resultado);
            for (Object ob : array) {
                //obtenemos el objeto Json y extraemos
                //las propiedades que nos interesan
                ObjectNode json = (ObjectNode) ob;
                paises.add(new Pais(
                        json.get("name").get("common").asText(),
                        (json.has("capital") ? json.get("capital").get(0).asText() : "N/A"),
                        json.get("population").asInt(),
                        json.get("flag").asText()));
            }
        } catch (JsonProcessingException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }

        return paises;
    }

    @Override
    public List<Pais> buscarPaises(String name) {
        return obtenerPaises()
                .stream()
                .filter(p -> p.getNombre().contains(name))
                .collect(Collectors.toList());
    }
}
