package com.inicio.utils;

import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

import static com.inicio.utils.Constans.*;
import static org.springframework.security.config.annotation.web.configuration.SecurityReactorContextConfiguration.SecurityReactorContextSubscriberRegistrar.getAuthentication;

public class JWTAuthorizationFilter extends BasicAuthenticationFilter {

    public JWTAuthorizationFilter(AuthenticationManager authenticationManager) {
        super(authenticationManager);
    }

    @Override
    protected void doFilterInternal(HttpServletRequest req, HttpServletResponse res, FilterChain chain)
            throws IOException, ServletException {
        String header = req.getHeader(ENCABEZADO);
        if (header == null || !header.startsWith(PREFIJO_TOKEN)) {
            chain.doFilter(req, res);
            return;
        }
        /**
         * Se obtienen los datos del usuario apartir del token
         */
        UsernamePasswordAuthenticationFilter authenticationFilter = getAuthentication();
        /**
         * La informacion del usuario se almacena en el contexto de seguridad
         * para que pueda ser utilizada por Spring security durante el proceso de autorizacion
         */

    }
}
