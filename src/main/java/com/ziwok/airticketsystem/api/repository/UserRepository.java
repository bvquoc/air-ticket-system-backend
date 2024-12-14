package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.User;
import com.ziwok.airticketsystem.api.model.UserRole;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {

	User findByEmail(String email);

	boolean existsByEmail(String email);

	// Find users by role
	List<User> findByUserRole(UserRole userRole);

	// Find users by email containing a string (useful for search/filter)
	List<User> findByEmailContainingIgnoreCase(String email);

	// Custom query to find a user by email and role
	Optional<User> findByEmailAndUserRole(String email, UserRole userRole);

	// Delete user by username
	void deleteByUsername(String username);
}
