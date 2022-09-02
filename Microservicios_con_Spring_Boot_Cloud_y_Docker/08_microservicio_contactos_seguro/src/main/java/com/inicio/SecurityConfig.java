package com.inicio;

import com.inicio.utils.JWTAuthorizationFilter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;


@EnableWebSecurity
@Configuration
public class SecurityConfig extends WebSecurityConfigurerAdapter {
    /**
     * Definicion de roles y usuarios
     */
    @Bean
    @Override
    protected void configure(AuthenticationManagerBuilder auth) throws Exception {
        auth
                .inMemoryAuthentication()
                .withUser("user1")
                .password("{noop}user1") //"noop" se pone para no obligar a usar mecanismo de encriptacion
                .roles("USER")
                .and()
                .withUser("admin")
                .password("{noop}admin")
                .roles("USER", "ADMIN");

        /**
         * Lo siguiente es para el caso de que se encripte la password
         * auth
         *                 .inMemoryAuthentication()
         *                 .withUser("user1")
         *                 .password(new BCryptPasswordEncoder().encode("user1")) //"noop" se pone para no obligar a usar mecanismo de encriptacion
         *                 .roles("USER")
         *                 .and()
         *                 .withUser("admin")
         *                 .password(new BCryptPasswordEncoder().encode("admin"))
         *                 .roles("USER", "ADMIN");
         */
        /**
         * La siguiente configuracion es para usuarios en una base de datos
         *  auth
         *                 .jdbcAuthentication()
         *                 .dataSource(dataSource)
         *                 .usersByUsernameQuery("select username, password, enable" + " from users where username=?")
         *                 .authoritiesByUsernameQuery("select username, authority " + "from authorities where username=?");
         */
    }

    /**
     * Definicion de politicas de seguridad
     * @param http
     * @throws Exception
     */
    @Override
    protected void configure(HttpSecurity http) throws Exception{
        http.csrf().disable()
                .authorizeRequests()
                .antMatchers(HttpMethod.POST,"/contactos").hasRole("ADMIN")
                //.antMatchers("/**").authenticated()
                //.antMatchers("/contactos/**").authenticated()
                .and()
                .addFilter(new JWTAuthorizationFilter(authenticationManager()));
    }
}
