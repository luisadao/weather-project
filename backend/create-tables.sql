DROP TABLE IF EXISTS weather_data;
CREATE TABLE weather_data (
  id         INT AUTO_INCREMENT NOT NULL,
  city      VARCHAR(128) NOT NULL,
  temperature DECIMAL(5,2) NOT NULL,
  icon  VARCHAR(128)  NOT NULL,
  description VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO weather_data
  (city, temperature, icon, description)
VALUES
  ('Coimbra', 35.21,01,"clear sky"),
  ('Faro', 40.21,02,"few clouds"),
  ('Leiria', 12.21,10,"rain"),
  ('Lisboa', 15.21,11,"thunder"),
  ('Porto', 5.21,01,"clear sky");