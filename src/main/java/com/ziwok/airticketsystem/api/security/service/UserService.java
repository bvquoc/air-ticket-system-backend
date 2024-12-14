package com.ziwok.airticketsystem.api.security.service;

import com.ziwok.airticketsystem.api.model.User;
import com.ziwok.airticketsystem.api.security.dto.AuthenticatedUserDto;
import com.ziwok.airticketsystem.api.security.dto.RegistrationRequest;
import com.ziwok.airticketsystem.api.security.dto.RegistrationResponse;

/**
 * Created on Nov, 2024
 *
 * @author Quoc Bui
 */
public interface UserService {

	User findByEmail(String email);

	RegistrationResponse registration(RegistrationRequest registrationRequest);

	AuthenticatedUserDto findAuthenticatedUserByEmail(String email);

}
