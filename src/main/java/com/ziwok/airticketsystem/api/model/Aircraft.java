package com.ziwok.airticketsystem.api.model;

import jakarta.persistence.*;
import lombok.*;

@Getter
@Setter
@Entity
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "AIRCRAFTS")
public class Aircraft {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer aircraftId;

    private String aircraftModel;

    private String manufacturer;

    private Integer seatingCapacity;
}
