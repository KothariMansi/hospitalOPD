CREATE TABLE `Client` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `state` VARCHAR(255) NOT NULL,
  `city` VARCHAR(255) NOT NULL,
  `age` INT NOT NULL
);

CREATE TABLE `User` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `state` VARCHAR(255) NOT NULL,
  `city` VARCHAR(255) NOT NULL,
  `gender` ENUM ('MALE', 'FEMALE', 'OTHER') NOT NULL,
  `age` INT
);

CREATE TABLE `Hospital` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `photo` VARCHAR(255),
  `state` VARCHAR(255) NOT NULL,
  `city` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(255)
);

CREATE TABLE `Speciality` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `speciality_name` VARCHAR(255) NOT NULL
);

CREATE TABLE `CheckUpTime` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `morning` TIMESTAMP ,
  `evening` TIMESTAMP ,
  `night` TIMESTAMP 
);

CREATE TABLE Doctor (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  username VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  hospital_id BIGINT NOT NULL,
  resident_address VARCHAR(255),
  checkup_time_id BIGINT NOT NULL,
  is_on_leave BOOLEAN DEFAULT FALSE,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (hospital_id) REFERENCES Hospital(id),
  FOREIGN KEY (checkup_time_id) REFERENCES CheckUpTime(id)
);

CREATE TABLE `DoctorSpeciality` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `speciality_id` BIGINT NOT NULL,
  `docter_id` BIGINT NOT NULL,
  FOREIGN KEY (`speciality_id`) REFERENCES `Speciality` (`id`),
  FOREIGN KEY (`docter_id`) REFERENCES `Doctor` (`id`)
);

CREATE TABLE `OPDLine` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `reg_time` DATETIME,
  `token_number` INT NOT NULL,
  `client_id` BIGINT NOT NULL,
  `doctor_id` BIGINT NOT NULL,
  `isChecked` BOOLEAN NOT NULL,
  `checked_time` DATETIME,
  FOREIGN KEY (`client_id`) REFERENCES `Client` (`id`),
  FOREIGN KEY (`doctor_id`) REFERENCES `Doctor` (`id`)
);

CREATE TABLE `EmergencyLine` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `reg_time` DATETIME,
  `token_number` BIGINT NOT NULL,
  `client_id` BIGINT NOT NULL,
  `doctor_id` BIGINT NOT NULL,
  `isChecked` BOOLEAN NOT NULL,
  `checked_time` DATETIME,
  FOREIGN KEY (`client_id`) REFERENCES `Client` (`id`),
  FOREIGN KEY (`doctor_id`) REFERENCES `Doctor` (`id`)
);
