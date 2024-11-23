package com.ziwok.airticketsystem.api.exceptions;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

/**
 * Created on Nov, 2024
 *
 * @author Quoc Bui
 */
@Getter
@RequiredArgsConstructor
public class RegistrationException extends RuntimeException {

	private final String errorMessage;

}
