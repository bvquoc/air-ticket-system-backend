package com.ziwok.airticketsystem.api.repository;

import com.ziwok.airticketsystem.api.model.Ticket;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface TicketRepository extends JpaRepository<Ticket, Integer> {

    List<Ticket> findByDirectionBookingId(Integer directionBookingId);

    List<Ticket> findByFlightId(Integer flightId);

    List<Ticket> findByPassengerId(Integer passengerId);

    List<Ticket> findByTicketStatus(String ticketStatus);
}
