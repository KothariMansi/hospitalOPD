CREATE TABLE `Client` (
  `id` int PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `age` int NOT NULL
);

CREATE TABLE `User` (
  `id` int PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `gender` ENUM ('MALE', 'FEMALE', 'OTHER') NOT NULL,
  `age` int
);

CREATE TABLE `Hospital` (
  `id` int PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `photo` varchar(255),
  `state` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `phone` varchar(255)
);

CREATE TABLE `Speciality` (
  `id` int PRIMARY KEY,
  `speciality_name` varchar(255) NOT NULL
);

CREATE TABLE `DoctorSpeciality` (
  `id` int  PRIMARY KEY,
  `speciality_id` int NOT NULL,
  `docter_id` int NOT NULL
);

CREATE TABLE `OPDLine` (
  `id` int PRIMARY KEY,
  `reg_time` datetime,
  `token_number` int NOT NULL,
  `client_id` int NOT NULL,
  `doctor_id` int NOT NULL,
  `isChecked` bool NOT NULL,
  `checked_time` datetime
);

CREATE TABLE `EmergencyLine` (
  `id` int PRIMARY KEY,
  `reg_time` datetime,
  `token_number` int NOT NULL,
  `client_id` int NOT NULL,
  `doctor_id` int NOT NULL,
  `isChecked` bool NOT NULL,
  `checked_time` datetime
);

CREATE TABLE `CheckUpTime` (
  `id` int PRIMARY KEY,
  `morning` datetime,
  `evening` datetime,
  `night` datetime
);

CREATE TABLE `Doctor` (
  `id` int PRIMARY KEY,
  `Name` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `hospital_id` int NOT NULL,
  `resident_address` varchar(255),
  `isMorning` bool,
  `isEvening` bool,
  `isNight` bool,
  `checkup_time_id` int NOT NULL,
  `isOnLeave` bool
);

ALTER TABLE `DoctorSpeciality` ADD FOREIGN KEY (`speciality_id`) REFERENCES `Speciality` (`id`);

ALTER TABLE `DoctorSpeciality` ADD FOREIGN KEY (`docter_id`) REFERENCES `Doctor` (`id`);

ALTER TABLE `OPDLine` ADD FOREIGN KEY (`client_id`) REFERENCES `Client` (`id`);

ALTER TABLE `OPDLine` ADD FOREIGN KEY (`doctor_id`) REFERENCES `Doctor` (`id`);

ALTER TABLE `EmergencyLine` ADD FOREIGN KEY (`client_id`) REFERENCES `Client` (`id`);

ALTER TABLE `EmergencyLine` ADD FOREIGN KEY (`doctor_id`) REFERENCES `Doctor` (`id`);

ALTER TABLE `Doctor` ADD FOREIGN KEY (`hospital_id`) REFERENCES `Hospital` (`id`);

ALTER TABLE `Doctor` ADD FOREIGN KEY (`checkup_time_id`) REFERENCES `CheckUpTime` (`id`);
