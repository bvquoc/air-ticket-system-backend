create table airport
(
    airport_code varchar(10)  not null
        primary key,
    name         varchar(255) not null
);

alter table airport
    owner to postgres;

create table aircraft
(
    aircraft_id serial
        primary key,
    status      varchar(50),
    name        varchar(255) not null
);

alter table aircraft
    owner to postgres;

create table seat
(
    seat_code   varchar(20) not null
        primary key,
    aircraft_id integer     not null
        references aircraft,
    position    varchar(50),
    class       varchar(50)
);

alter table seat
    owner to postgres;

create table flight
(
    flight_id           serial
        primary key,
    aircraft_id         integer     not null
        references aircraft,
    departure_airport   varchar(10) not null
        references airport,
    destination_airport varchar(10) not null
        references airport,
    departure_time      timestamp   not null,
    destination_time    timestamp   not null,
    status              varchar(50)
);

alter table flight
    owner to postgres;

create table seat_availability
(
    flight_id integer     not null
        references flight,
    seat_code varchar(20) not null
        references seat,
    status    varchar(50),
    price     numeric(10, 2),
    primary key (flight_id, seat_code)
);

alter table seat_availability
    owner to postgres;

create table passenger
(
    passenger_id serial
        primary key,
    name         varchar(255) not null,
    phone_number varchar(15),
    email        varchar(255)
);

alter table passenger
    owner to postgres;

create table booking
(
    booking_id      serial
        primary key,
    status          varchar(20) default 'init'::character varying,
    ticket_amount   numeric(10, 2),
    fee_amount      numeric(10, 2),
    discount_amount numeric(10, 2),
    total_amount    numeric(10, 2),
    checkout_at     timestamp
);

alter table booking
    owner to postgres;

create table ticket
(
    ticket_id    serial
        primary key,
    passenger_id integer     not null
        references passenger,
    flight_id    integer     not null
        references flight,
    seat_code    varchar(20) not null
        references seat,
    status       varchar(50),
    amount       numeric(10, 2),
    booking_id   integer     not null
        references booking,
    issued_at    timestamp   not null
);

alter table ticket
    owner to postgres;

create table transaction
(
    txn_id         serial
        primary key,
    status         varchar(50),
    type           varchar(20)
        constraint transaction_type_check
            check ((type)::text = ANY ((ARRAY ['PAYMENT'::character varying, 'REFUND'::character varying])::text[])),
    debit_acc      varchar(50),
    credit_acc     varchar(50),
    payment_method varchar(50),
    fee_charge     numeric(10, 2),
    amount         numeric(10, 2),
    booking_id     integer
        references booking
);

alter table transaction
    owner to postgres;

create table booking_passenger
(
    booking_id   integer not null
        references booking,
    passenger_id integer not null
        references passenger,
    primary key (booking_id, passenger_id)
);

alter table booking_passenger
    owner to postgres;

create table pricing_config
(
    key   varchar(50) not null
        primary key,
    value numeric(10, 2)
);

alter table pricing_config
    owner to postgres;

create table booking_direction
(
    bd_id      serial
        primary key,
    booking_id integer not null
        references booking,
    direction  varchar(20)
        constraint booking_direction_direction_check
            check ((direction)::text = ANY ((ARRAY ['AWAY'::character varying, 'HOME'::character varying])::text[])),
    flight_id  integer not null
        references flight,
    amount     numeric(10, 2)
);

alter table booking_direction
    owner to postgres;

create table bd_ticket
(
    bd_id     integer not null
        references booking_direction,
    ticket_id integer not null
        references ticket,
    primary key (bd_id, ticket_id)
);

alter table bd_ticket
    owner to postgres;

create table booking_line
(
    bl_id           serial
        primary key,
    bd_id           integer not null
        references booking_direction,
    type            varchar(20)
        constraint booking_line_type_check
            check ((type)::text = ANY
        ((ARRAY ['TICKET'::character varying, 'VAT'::character varying, 'FEE'::character varying])::text[])),
    unit_amount     numeric(10, 2),
    quantity        integer,
    subtotal_amount numeric(10, 2)
);

alter table booking_line
    owner to postgres;

create table entry
(
    entry_id       serial
        primary key,
    sign           char
        constraint entry_sign_check
            check (sign = ANY (ARRAY ['+'::bpchar, '-'::bpchar])),
    acc            varchar(50),
    amount         numeric(10, 2),
    status         varchar(50),
    transaction_id integer not null
        references transaction
);

alter table entry
    owner to postgres;

