package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

/**
 * Created on Nov, 2024
 *
 * @author Quoc Bui
 */
public interface UserRepository extends JpaRepository<User, Long> {

	User findByUsername(String username);

	boolean existsByEmail(String email);

	boolean existsByUsername(String username);

}
