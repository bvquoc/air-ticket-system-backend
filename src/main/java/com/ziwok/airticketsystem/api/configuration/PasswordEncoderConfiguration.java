package com.ziwok.airticketsystem.api.configuration;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

/**
 * Created on March, 2023
 *
 * @author Quoc Bui
 */

@Configuration
public class PasswordEncoderConfiguration {

	@Bean
	public BCryptPasswordEncoder encoder() {
		return new BCryptPasswordEncoder();
	}

}
