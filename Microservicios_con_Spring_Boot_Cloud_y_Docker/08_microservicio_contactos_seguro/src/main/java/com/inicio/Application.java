package com.inicio;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@ComponentScan(basePackages = {"com.inicio.controllers","com.inicio.daos","com.inicio.services","com.inicio.exceptions","com.inicio"})
@EntityScan(basePackages = {"com.inicio.model"})
@EnableJpaRepositories(basePackages = {"com.inicio.daos"})
@SpringBootApplication
public class Application {

	public static void main(String[] args) {
		SpringApplication.run(Application.class, args);
	}

}
