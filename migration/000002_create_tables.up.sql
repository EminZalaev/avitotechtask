CREATE TABLE IF NOT EXISTS user1
(
    id           bigint not null
    name         varchar not null,
    cardNumber varchar not null,
    balance      bigint default 0
);

INSERT INTO public.user1 (name, cardNumber)
VALUES ('Emin', '1111 1111'),
       ('Ilia', '2222 2222'),
       ('Nurbek', '3333 3333'),
       ('Artem', '4444 4444'),
       ('Lenya', '5555 5555'),
       ('Slava', '6666 6666');