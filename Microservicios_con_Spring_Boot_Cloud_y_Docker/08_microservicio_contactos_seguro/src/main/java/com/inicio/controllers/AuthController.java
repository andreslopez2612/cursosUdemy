/**
 * 
 */
package com.inicio.controllers;

import java.util.Date;
import java.util.stream.Collectors;

import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

import static com.inicio.utils.Constans.*;

/**
 * 
 * @author andre
 *
 */
@RestController
public class AuthController {
	AuthenticationManager authManager;

	/**
	 * 
	 * @param authManager
	 */
	public AuthController(AuthenticationManager authManager) {
		this.authManager = authManager;
	}

	/**
	 * 
	 * @param user
	 * @param pwd
	 * @return
	 */
	@PostMapping("login")
	public String login(@RequestParam("user") String user, @RequestParam("pws") String pwd) {
		Authentication authentication = authManager.authenticate(new UsernamePasswordAuthenticationToken(user, pwd));
		// Si el usuario esta autenticado, se genera el token
		if (authentication.isAuthenticated()) {
			return getToken(authentication);
		} else {
			return "No autenticado";
		}
	}

	private String getToken(Authentication authentication) {
		/**
		 * En el body del token se incluye el usuario y los roles a los que pertenece,
		 * ademas de la fecha de caducidad y los datos de la firma
		 */
		String token = Jwts.builder().setIssuedAt(new Date()) // FECHA CREACION
				.setSubject(authentication.getName()) // USUARIO
				.claim("authorities", authentication.getAuthorities().stream()// ROLES
						.map(GrantedAuthority::getAuthority).collect(Collectors.toList()))
				.setExpiration(new Date(System.currentTimeMillis() + TIEMPO_VIDA)) // FECHA CADUCIDAD
				.signWith(SignatureAlgorithm.HS512, CLAVE)// CLAVE Y ALGORITMO PARA FIRMA
				.compact(); //GENERACION DEL TOKEN
		return token;
	}
}
