insert into users (name, nick, email, password)
VALUES
("user1", "user_1", "user_1@gmail.com", "$2a$10$s82MqrkS//ItF1IyzDD6..pg096HbWaLEhr/358C5y2zn4iSumJwG"),
("user2", "user_2", "user_2@gmail.com", "$2a$10$s82MqrkS//ItF1IyzDD6..pg096HbWaLEhr/358C5y2zn4iSumJwG"),
("user3", "user_3", "user_3@gmail.com", "$2a$10$s82MqrkS//ItF1IyzDD6..pg096HbWaLEhr/358C5y2zn4iSumJwG"),
("user4", "user_4", "user_4@gmail.com", "$2a$10$s82MqrkS//ItF1IyzDD6..pg096HbWaLEhr/358C5y2zn4iSumJwG");

insert into followers(user_id, follower_id)
VALUES
(1,2),
(3,2),
(4,1);

insert into publications(title, content, author_id)
values
("Publication user 1", "This is a simple post from user 1", 1),
("Publication user 2", "This is a simple post from user 2", 2),
("Publication user 3", "This is a simple post from user 3", 3),
("Publication user 4", "This is a simple post from user 4", 4);