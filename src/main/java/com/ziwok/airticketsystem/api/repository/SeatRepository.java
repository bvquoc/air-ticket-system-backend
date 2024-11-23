package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.Seat;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface SeatRepository extends JpaRepository<Seat, Integer> {

    List<Seat> findByAircraftId(Integer aircraftId);

    List<Seat> findByIsAvailable(Boolean isAvailable);
}
