package com.inicio.daos;

import com.inicio.model.Contacto;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;

import javax.transaction.Transactional;

public interface AgendaJpaSpring extends JpaRepository<Contacto, Integer> {
    Contacto findByEmail(String email);

    @Transactional
    @Modifying
    @Query("Delete from Contacto cWhere c.email=?1")
    void eliminarPorEmail(String email);
}
