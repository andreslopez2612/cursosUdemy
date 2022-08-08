package com.cursos.controllers;

import com.cursos.model.Curso;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.PostConstruct;
import java.util.ArrayList;
import java.util.List;

@RestController
public class CursosController {

    private List<Curso> cursos = new ArrayList<>();
    @PostConstruct
    public void init(){
        cursos.add(new Curso("Spring",25,"tarde"));
        cursos.add(new Curso("Spring Boot",20,"tarde"));
        cursos.add(new Curso("Python",30,"tarde"));
        cursos.add(new Curso("Java EE",50,"Fin de semana"));
        cursos.add(new Curso("Java Basico",30,"mañana"));
    }

    @GetMapping(value="cursos", produces= MediaType.APPLICATION_JSON_VALUE)
    public List<Curso> getCursos(){
        return cursos;
    }

    @GetMapping(value="curso", produces= MediaType.APPLICATION_JSON_VALUE)
    public Curso getCurso(){
        return new Curso("Java",100,"mañana");
    }

    @GetMapping(value="cursos/{name}", produces= MediaType.APPLICATION_JSON_VALUE)
    public List<Curso> buscarCursos(@PathVariable("name") String nombre){
        List<Curso> aux=new ArrayList<>();
        for(Curso c:cursos){
            if(c.getNombre().contains(nombre)){
                aux.add(c);
            }
        }
        return aux;
    }
}
