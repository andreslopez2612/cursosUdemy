package com.inicio.daos;

import com.inicio.model.Contacto;

import java.util.List;

public interface AgendaDao {
    void agregarContacto(Contacto contacto);

    Contacto recuperarContacto(String email);

    void eliminarContacto(String email);

    List<Contacto> devolverContactoList();

    void eliminarContacto(int idContacto);

    Contacto recuperarContacto(int idContacto);

    void actualizarContacto(Contacto contacto);
}
