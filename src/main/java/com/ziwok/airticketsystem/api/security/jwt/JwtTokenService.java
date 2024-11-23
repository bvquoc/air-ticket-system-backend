package com.ziwok.airticketsystem.api.security.jwt;

import com.ziwok.airticketsystem.api.security.mapper.UserMapper;
import com.ziwok.airticketsystem.api.security.service.UserService;
import com.ziwok.airticketsystem.api.model.User;
import com.ziwok.airticketsystem.api.security.dto.AuthenticatedUserDto;
import com.ziwok.airticketsystem.api.security.dto.LoginRequest;
import com.ziwok.airticketsystem.api.security.dto.LoginResponse;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.stereotype.Service;

/**
 * Created on Nov, 2024
 *
 * @author Quoc Bui
 */
@Slf4j
@Service
@RequiredArgsConstructor
public class JwtTokenService {

	private final UserService userService;

	private final JwtTokenManager jwtTokenManager;

	private final AuthenticationManager authenticationManager;

	public LoginResponse getLoginResponse(LoginRequest loginRequest) {

		final String username = loginRequest.getUsername();
		final String password = loginRequest.getPassword();

		final UsernamePasswordAuthenticationToken usernamePasswordAuthenticationToken = new UsernamePasswordAuthenticationToken(username, password);

		authenticationManager.authenticate(usernamePasswordAuthenticationToken);

		final AuthenticatedUserDto authenticatedUserDto = userService.findAuthenticatedUserByUsername(username);

		final User user = UserMapper.INSTANCE.convertToUser(authenticatedUserDto);
		final String token = jwtTokenManager.generateToken(user);

		log.info("{} has successfully logged in!", user.getUsername());

		return new LoginResponse(token);
	}

}
