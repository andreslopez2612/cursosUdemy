package com.inicio.controllers;

import com.inicio.model.Contacto;
import com.inicio.services.AgendaService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@CrossOrigin(origins = "*")
@RestController
public class ContactosController {
    @Autowired
    AgendaService agendaService;

    @GetMapping(value = "contactos", produces = MediaType.APPLICATION_JSON_VALUE)
    public List<Contacto> recuperarContactos() {
        return agendaService.recuperarContactos();
    }

    @GetMapping(value = "contactos/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
    public Contacto recuperarContactos(@PathVariable("id") int id) {
        return agendaService.buscarContacto(id);
    }

    @PostMapping(value = "contactos", consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public String guardarContacto(@RequestBody Contacto contacto) {
        return String.valueOf(agendaService.agregarContacto(contacto));
    }

    @PutMapping(value = "contactos", consumes = MediaType.APPLICATION_JSON_VALUE)
    public void actualizarContacto(@RequestBody Contacto contacto) {
        agendaService.actualizarContacto(contacto);
    }

    @DeleteMapping(value = "contactos/{id}")
    public void eliminarPorId(@PathVariable("id") int idContacto) {
        agendaService.eliminarContacto(idContacto);
    }
}
