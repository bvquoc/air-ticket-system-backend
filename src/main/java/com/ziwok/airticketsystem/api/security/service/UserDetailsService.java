package com.ziwok.airticketsystem.api.security.service;

import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails;

public interface UserDetailsService {
    UserDetails loadUserByEmail(String username) throws UserDetailsNotFoundException;

}
