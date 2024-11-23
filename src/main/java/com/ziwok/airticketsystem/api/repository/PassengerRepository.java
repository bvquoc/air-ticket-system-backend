package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.Passenger;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface PassengerRepository extends JpaRepository<Passenger, Integer> {

    Passenger findByPassportNumber(String passportNumber);
}
