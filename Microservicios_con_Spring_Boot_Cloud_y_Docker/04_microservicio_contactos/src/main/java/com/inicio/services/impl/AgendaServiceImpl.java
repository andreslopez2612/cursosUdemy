package com.inicio.services.impl;

import com.inicio.daos.AgendaDao;
import com.inicio.model.Contacto;
import com.inicio.services.AgendaService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AgendaServiceImpl implements AgendaService {
    AgendaDao dao;

    @Override
    public boolean agregarContacto(Contacto contacto) {
        if (dao.recuperarContacto(contacto.getIdContacto()) == null) {
            dao.agregarContacto((contacto));
            return true;
        }
        return false;
    }

    @Override
    public List<Contacto> recuperarContactos() {
        return dao.devolverContactoList();
    }

    @Override
    public void actualizarContacto(Contacto contacto) {
        if(dao.recuperarContacto(contacto.getIdContacto())!=null){
            dao.actualizarContacto(contacto);
        }
    }

    @Override
    public boolean eliminarContacto(int idContacto) {
        if(dao.recuperarContacto(idContacto)!=null){
            dao.recuperarContacto(idContacto);
            return true;
        }
        return false;
    }

    @Override
    public Contacto buscarContacto(int idContacto) {
        return dao.recuperarContacto(idContacto);
    }
}
