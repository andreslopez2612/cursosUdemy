package com.inicio.services;

import com.inicio.model.Contacto;

import java.util.List;

public interface AgendaService {
    void agregarContacto(Contacto contacto) throws  Exception;
    List<Contacto> recuperarContactos();
    void actualizarContacto(Contacto contacto);
    boolean eliminarContacto(int idContacto);
    Contacto buscarContacto(int idContacto);
}
