package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.Booking;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface BookingRepository extends JpaRepository<Booking, Integer> {

    List<Booking> findByUserId(Long userId);

    List<Booking> findByStatus(String status);
}
