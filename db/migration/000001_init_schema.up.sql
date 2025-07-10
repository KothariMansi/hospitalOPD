CREATE TABLE `Client` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `state` VARCHAR(255) NOT NULL,
  `city` VARCHAR(255) NOT NULL,
  `age` INT NOT NULL
);

CREATE TABLE `User` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `state` VARCHAR(255) NOT NULL,
  `city` VARCHAR(255) NOT NULL,
  `gender` ENUM ('MALE', 'FEMALE', 'OTHER') NOT NULL,
  `age` INT
);

CREATE TABLE `Hospital` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `photo` VARCHAR(255),
  `state` VARCHAR(255) NOT NULL,
  `city` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(255)
);

CREATE TABLE `Speciality` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `speciality_name` VARCHAR(255) NOT NULL
);

CREATE TABLE `CheckUpTime` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `morning` DATETIME,
  `evening` DATETIME,
  `night` DATETIME
);

CREATE TABLE `Doctor` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `Name` VARCHAR(255) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `hospital_id` INT NOT NULL,
  `resident_address` VARCHAR(255),
  `isMorning` BOOLEAN,
  `isEvening` BOOLEAN,
  `isNight` BOOLEAN,
  `checkup_time_id` INT NOT NULL,
  `isOnLeave` BOOLEAN,
  FOREIGN KEY (`hospital_id`) REFERENCES `Hospital`(`id`),
  FOREIGN KEY (`checkup_time_id`) REFERENCES `CheckUpTime`(`id`)
);

CREATE TABLE `DoctorSpeciality` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `speciality_id` INT NOT NULL,
  `docter_id` INT NOT NULL,
  FOREIGN KEY (`speciality_id`) REFERENCES `Speciality` (`id`),
  FOREIGN KEY (`docter_id`) REFERENCES `Doctor` (`id`)
);

CREATE TABLE `OPDLine` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `reg_time` DATETIME,
  `token_number` INT NOT NULL,
  `client_id` INT NOT NULL,
  `doctor_id` INT NOT NULL,
  `isChecked` BOOLEAN NOT NULL,
  `checked_time` DATETIME,
  FOREIGN KEY (`client_id`) REFERENCES `Client` (`id`),
  FOREIGN KEY (`doctor_id`) REFERENCES `Doctor` (`id`)
);

CREATE TABLE `EmergencyLine` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `reg_time` DATETIME,
  `token_number` INT NOT NULL,
  `client_id` INT NOT NULL,
  `doctor_id` INT NOT NULL,
  `isChecked` BOOLEAN NOT NULL,
  `checked_time` DATETIME,
  FOREIGN KEY (`client_id`) REFERENCES `Client` (`id`),
  FOREIGN KEY (`doctor_id`) REFERENCES `Doctor` (`id`)
);
