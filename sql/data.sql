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