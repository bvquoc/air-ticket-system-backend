package com.ziwok.airticketsystem.api;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.EnableAspectJAutoProxy;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@SpringBootApplication
@EnableJpaRepositories
@EnableAspectJAutoProxy
public class AirTicketSystemApi {

	public static void main(String[] args) {

		SpringApplication.run(AirTicketSystemApi.class, args);
	}

}