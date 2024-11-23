package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.DirectionBooking;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface DirectionBookingRepository extends JpaRepository<DirectionBooking, Integer> {

    List<DirectionBooking> findByBookingId(Integer bookingId);
}
