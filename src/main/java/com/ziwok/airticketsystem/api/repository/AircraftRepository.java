package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.Aircraft;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AircraftRepository extends JpaRepository<Aircraft, Integer> {

    Aircraft findByAircraftModel(String aircraftModel);
}
