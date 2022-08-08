package com.cursos.controllers;

import org.junit.jupiter.api.MethodOrderer;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestMethodOrder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;

@SpringBootTest
@AutoConfigureMockMvc
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class CursosControllerTest {

    @Autowired
    MockMvc mockMvc;

    @Test
    void initTest() {
    }

    @Test
    @Order(0)
    void getCursosTest() throws Exception {
        mockMvc.perform(get("/cursos")).andDo(print());
    }

    @Test
    void getCursoTest() throws Exception {
    }

    @Test
    void buscarCursosTest() throws Exception {
    }

    @Test
    @Order(1)
    void eliminarCursoTest() throws Exception {
        mockMvc.perform(delete("/curso/Python"));
    }

    @Test
    @Order(2)
    void altaCursoTest() throws Exception {
        mockMvc.perform(post("/curso")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content("{\"nombre\":\"Angular 10\",\"duracion\":40,\"horario\":\"tarde\"}"))
                .andDo(print());
    }

    @Test
    @Order(3)
    void actualizaCursoTest() throws Exception {
        mockMvc.perform(put("/curso")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content("{\"nombre\":\"Angular 10\",\"duracion\":80,\"horario\":\"tarde\"}"))
                .andDo(print());
    }
}