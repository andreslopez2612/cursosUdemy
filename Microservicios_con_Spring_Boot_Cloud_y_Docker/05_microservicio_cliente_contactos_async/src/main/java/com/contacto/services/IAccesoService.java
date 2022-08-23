package com.contacto.services;

import com.contacto.model.Persona;

import java.util.List;
import java.util.concurrent.CompletableFuture;

public interface IAccesoService {
    CompletableFuture<List<Persona>> llamadaServicio(Persona persona);
}
