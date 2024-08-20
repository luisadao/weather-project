DROP TABLE IF EXISTS weather_data;
CREATE TABLE weather_data (
  id         INT AUTO_INCREMENT NOT NULL,
  city      VARCHAR(128) NOT NULL,
  temperature DECIMAL(5,2) NOT NULL,
  description VARCHAR(128) NOT NULL,
  timestamp  DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO weather_data
  (city, temperature, description, timestamp)
VALUES
  ('Coimbra', 35.21,"clear sky", "2024-02-21 14:24:36"),
  ('Faro', 40.21,"cloudy", "2024-02-21 22:24:36"),
  ('Leiria', 12.21,"rain", "2024-02-21 4:24:36"),
  ('Lisboa', 15.21,"thunder", "2024-02-21 10:24:36"),
  ('Porto', 5.21, "clear sky","2024-02-21 19:24:36");