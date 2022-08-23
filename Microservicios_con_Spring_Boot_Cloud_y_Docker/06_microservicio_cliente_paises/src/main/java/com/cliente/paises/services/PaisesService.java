package com.cliente.paises.services;

import com.cliente.paises.model.Pais;

import java.util.List;

public interface PaisesService {
    List<Pais> obtenerPaises();
    List<Pais> buscarPaises(String name);
}
