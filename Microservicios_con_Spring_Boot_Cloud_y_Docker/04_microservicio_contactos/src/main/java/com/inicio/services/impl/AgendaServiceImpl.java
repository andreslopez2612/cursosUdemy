package com.inicio.services.impl;

import com.inicio.daos.AgendaDao;
import com.inicio.model.Contacto;
import com.inicio.services.AgendaService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AgendaServiceImpl implements AgendaService {
    @Autowired
    AgendaDao dao;

    @Override
    public void agregarContacto(Contacto contacto) throws Exception {
        if (dao.recuperarContacto(contacto.getEmail()) == null) {
            dao.agregarContacto((contacto));
            return;
        }
        throw new Exception("Contacto repetido");
    }

    @Override
    public List<Contacto> recuperarContactos() {
    	try {
    		Thread.sleep(8000);
    	}catch (InterruptedException e){
    		e.printStackTrace();
    	}    	
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
            dao.eliminarContacto(idContacto);
            return true;
        }
        return false;
    }

    @Override
    public Contacto buscarContacto(int idContacto) {
        return dao.recuperarContacto(idContacto);
    }
}
