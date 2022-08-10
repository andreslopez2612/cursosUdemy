package com.inicio.daos.impl;

import com.inicio.daos.AgendaDao;
import com.inicio.daos.AgendaJpaSpring;
import com.inicio.model.Contacto;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class AgendaDaoImpl implements AgendaDao {

    @Autowired
    AgendaJpaSpring agendaJpaSpring;

    @Override
    public void agregarContacto(Contacto contacto) {
        agendaJpaSpring.save(contacto);
    }

    @Override
    public Contacto recuperarContacto(String email) {
        return agendaJpaSpring.findByEmail(email);
    }

    @Override
    public void eliminarContacto(String email) {
        agendaJpaSpring.eliminarPorEmail(email);
    }

    @Override
    public List<Contacto> devolverContactoList() {
        return agendaJpaSpring.findAll();
    }

    @Override
    public Contacto recuperarContacto(int idContacto) {
        return agendaJpaSpring.findById(idContacto).orElse(null);
    }

    @Override
    public void actualizarContacto(Contacto contacto) {
        agendaJpaSpring.save(contacto);
    }
}
