package com.ziwok.airticketsystem.api.security.dto;

import com.ziwok.airticketsystem.api.model.UserRole;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

/**
 * Created on Nov, 2024
 *
 * @author Quoc Bui
 */
@Getter
@Setter
@NoArgsConstructor
public class AuthenticatedUserDto {

	private String name;

	private String username;

	private String password;

	private UserRole userRole;

}
