CREATE TABLE IF NOT EXISTS users1
(
    id           serial,
    name         varchar not null,
    cardNumber varchar not null,
    balance      int default 0
);

create unique index users_id_uindex
    on users (id);

INSERT INTO public.user2 (id, name, "cardNumber", balance)
VALUES (1, 'Emin', '1111 1111', DEFAULT),
       (2, 'Ilia', '2222 2222', DEFAULT),
       (3, 'Nurbek', '3333 3333', DEFAULT),
       (4, 'Artem', '4444 4444', DEFAULT),
       (5, 'Lenya', '5555 5555', DEFAULT),
       (6, 'Slava', '6666 6666', DEFAULT);

drop table public.user