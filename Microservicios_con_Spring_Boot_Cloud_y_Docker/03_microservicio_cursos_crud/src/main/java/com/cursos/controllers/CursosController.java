package com.cursos.controllers;

import com.cursos.model.Curso;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.*;

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

    @DeleteMapping(value="curso/{name}")
    public void eliminarCurso(@PathVariable("name") String nombre){
        cursos.removeIf(c->c.getNombre().equals(nombre));
    }

    @PostMapping(value="curso",consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public List<Curso> altaCurso(@RequestBody Curso curso){
        cursos.add(curso);
        return cursos;
    }

    @PutMapping(value="curso",consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public List<Curso> actualizaCurso(@RequestBody Curso curso){
        for(int i = 0;i<cursos.size();i++){
            if(cursos.get(i).getNombre().equals(curso.getNombre())){
                cursos.set(i,curso);
            }
        }
        return cursos;
    }
}
