package com.ziwok.airticketsystem.api.model;

import jakarta.persistence.*;
import lombok.*;

import java.time.LocalDateTime;

@Getter
@Setter
@Entity
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "DIRECTION_BOOKINGS")
public class DirectionBooking {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer directionBookingId;

    @ManyToOne
    @JoinColumn(name = "booking_id", nullable = false)
    private Booking booking;

    private String origin;

    private String destination;

    @Column(nullable = false)
    private LocalDateTime departureDate;

    private LocalDateTime arrivalDate;
}
