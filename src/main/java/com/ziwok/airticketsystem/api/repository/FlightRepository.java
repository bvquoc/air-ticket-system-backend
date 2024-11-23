package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.Flight;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface FlightRepository extends JpaRepository<Flight, Integer> {

    Flight findByFlightNumber(String flightNumber);

    List<Flight> findByOriginAndDestination(String origin, String destination);

    List<Flight> findByAircraftId(Integer aircraftId);
}
