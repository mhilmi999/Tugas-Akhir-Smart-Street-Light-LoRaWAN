-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 25 Okt 2022 pada 10.12
-- Versi server: 10.4.21-MariaDB
-- Versi PHP: 8.0.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `smartstreetlight`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `area`
--

CREATE TABLE `area` (
  `area_id` int(11) NOT NULL,
  `cust_id` varchar(200) NOT NULL,
  `antares_apps` varchar(200) NOT NULL,
  `area_name` varchar(200) NOT NULL,
  `area_location` varchar(200) NOT NULL,
  `electric_rates` float NOT NULL,
  `area_last` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `area`
--

INSERT INTO `area` (`area_id`, `cust_id`, `antares_apps`, `area_name`, `area_location`, `electric_rates`, `area_last`) VALUES
(1, '1280', 'SmartPJU', 'Kawasan Lion Park', 'Karawang, Jawa Barat', 1352, '2022-06-12 08:41:36'),
(2, '1290', 'SmartPJU2', 'Kawasan Industri Mojosemi', 'Mojokerto, Jawa Timur', 1352, '2022-06-12 08:41:36'),
(3, '1300', 'SmartPJU3', 'Kawasan Industri Cikarang', 'Cikarang, Jawa Barat', 1352, '2022-06-13 18:17:07');

-- --------------------------------------------------------

--
-- Struktur dari tabel `area_division`
--

CREATE TABLE `area_division` (
  `division_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `area_id` int(11) NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `area_division`
--

INSERT INTO `area_division` (`division_id`, `user_id`, `area_id`, `date_created`) VALUES
(1, 1, 1, '2022-09-08 17:32:05'),
(2, 1, 2, '2022-06-12 08:55:43'),
(3, 2, 1, '2022-08-20 11:04:31'),
(4, 2, 2, '2022-09-06 19:40:43'),
(5, 7, 1, '2022-09-06 19:44:13'),
(6, 8, 1, '2022-09-06 19:45:12'),
(7, 9, 1, '2022-09-08 17:32:05'),
(8, 10, 1, '2022-10-05 11:40:22');

-- --------------------------------------------------------

--
-- Struktur dari tabel `detail_roles`
--

CREATE TABLE `detail_roles` (
  `role_id` int(11) NOT NULL,
  `role_name` varchar(52) NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `detail_roles`
--

INSERT INTO `detail_roles` (`role_id`, `role_name`, `date_created`) VALUES
(1, 'admin', '2022-06-04 09:14:07'),
(2, 'Pengelola Kawasan', '2022-06-04 09:14:07'),
(3, 'Engineer', '2022-06-04 09:14:07'),
(4, 'Tenant', '2022-06-04 09:14:07'),
(1, 'admin', '2022-06-04 09:14:07'),
(2, 'Pengelola Kawasan', '2022-06-04 09:14:07'),
(3, 'Engineer', '2022-06-04 09:14:07'),
(4, 'Tenant', '2022-06-04 09:14:07');

-- --------------------------------------------------------

--
-- Struktur dari tabel `device`
--

CREATE TABLE `device` (
  `device_id` int(11) NOT NULL,
  `area_id` int(11) NOT NULL,
  `antares_id` varchar(200) NOT NULL,
  `device_name` varchar(200) NOT NULL,
  `device_label` varchar(500) NOT NULL,
  `device_address` varchar(200) NOT NULL,
  `device_eui` varchar(200) NOT NULL,
  `longitude` varchar(500) NOT NULL,
  `latitude` varchar(500) NOT NULL,
  `device_location` varchar(200) NOT NULL,
  `device_cons` varchar(200) NOT NULL,
  `device_sts` int(11) NOT NULL,
  `device_last` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device`
--

INSERT INTO `device` (`device_id`, `area_id`, `antares_id`, `device_name`, `device_label`, `device_address`, `device_eui`, `longitude`, `latitude`, `device_location`, `device_cons`, `device_sts`, `device_last`) VALUES
(1, 1, 'uoJ8bqJdQ-6874Gl', 'Luminer', 'Real Device', 'f55e7014', 'a9ead5aa8dc8bac8', '107.263619', '-6.358417', 'Jln Pangandaran No. 9', '2', 1, '2022-07-05 16:01:06'),
(2, 1, 'UkmPykbCRgy__h7r', 'Device_Testing_3', 'Real Device', '-', '0e7791287b505535', '107.263625', '-6.359417', 'Jln Pangandaran No. 10', '2', 1, '2022-06-12 08:51:40'),
(3, 2, 'BoJ8bqJdQ-6874Gl', 'device 2', 'Dummy Device', 'f55e7015', '1ddad5aa8dc8bac8', '107.263619', '-6.358417', 'Jln Pangandaran No. 9', '1', 1, '2022-06-12 08:51:40'),
(4, 1, 'CoJ8bqJdQ-6874Gl', 'device 3', 'Dummy Device', 'f55e7016', 'a9ead5aa8dc8bac8', '107.263619', '-6.358417', 'Jln Pangandaran No. 9', '0', 3, '2022-06-12 08:51:40'),
(5, 1, 'DoJ8bqJdQ-6874Gl', 'device 5', 'Dummy Device', 'f55e7017', 'a9ead5aa8dc8bac8', '107.263619', '-6.358417', 'Jln Pangandaran No. 9', '1', 2, '2022-06-12 08:51:40'),
(6, 1, 'BoJ8bqJdQ-6874Gl', 'device 6', 'Dummy Device', 'f55e7018', 'a9ead5aa8dc8bac8', '107.263619', '-6.358417', 'Jln Pangandaran No. 9', '2', 1, '2022-06-12 08:51:40'),
(7, 1, 'ZuoJ8bqJdQ-6874Gl', 'device 7', 'Dummy Device', 'f55e7019', 'a9ead5aa8dc8bac8', '107.263619', '-6.358417', 'Jln Pangandaran No. 9', '1', 3, '2022-06-12 08:51:40'),
(8, 1, 'uoJ8bqJdQ-6874Gl', 'Device Test Dummy', 'Unit Test Handler', 'f55e7015', '1ddad5aa8dc8bac8', '-', '-', 'Jl. Yang Lurus', '0', 1, '2022-09-08 17:30:51'),
(9, 1, 'uoJ8bqJdQ-6874Gl', 'Device Test Dummy', 'Unit Test Handler', 'f55e7015', '1ddad5aa8dc8bac8', '-', '-', 'Jl. Yang Lurus', '0', 1, '2022-07-05 16:01:06');

-- --------------------------------------------------------

--
-- Struktur dari tabel `device_cons_d`
--

CREATE TABLE `device_cons_d` (
  `id` int(11) NOT NULL,
  `cons_name` varchar(200) NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device_cons_d`
--

INSERT INTO `device_cons_d` (`id`, `cons_name`, `date_created`) VALUES
(0, 'Padam', '2022-06-12 10:48:07'),
(1, 'Menyala', '2022-06-12 10:48:07');

-- --------------------------------------------------------

--
-- Struktur dari tabel `device_controlling`
--

CREATE TABLE `device_controlling` (
  `deviceC_id` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  `manual_Control` int(11) NOT NULL,
  `restart_sts` int(11) NOT NULL,
  `alarm_type` int(11) NOT NULL,
  `hour_on` varchar(10) NOT NULL DEFAULT '00',
  `minutes_on` varchar(10) NOT NULL DEFAULT '00',
  `hour_off` varchar(10) NOT NULL DEFAULT '00',
  `minutes_off` varchar(10) NOT NULL DEFAULT '00',
  `interval_loop` int(11) NOT NULL,
  `interval_send` int(11) NOT NULL,
  `bluetooth` int(11) NOT NULL,
  `totalisator` int(11) NOT NULL,
  `lamp_treshold` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device_controlling`
--

INSERT INTO `device_controlling` (`deviceC_id`, `device_id`, `manual_Control`, `restart_sts`, `alarm_type`, `hour_on`, `minutes_on`, `hour_off`, `minutes_off`, `interval_loop`, `interval_send`, `bluetooth`, `totalisator`, `lamp_treshold`) VALUES
(1, 1, 2, 1, 9, '10', '30', '14', '30', 10, 37, 1, 3, 200),
(2, 2, 1, 0, 0, '10', '30', '14', '30', 10, 15, 1, 150, 200),
(3, 3, 1, 0, 0, '10', '30', '14', '30', 10, 15, 1, 150, 200),
(4, 4, 1, 0, 0, '10', '30', '14', '30', 10, 15, 1, 150, 200),
(5, 5, 1, 0, 0, '10', '30', '14', '30', 10, 15, 1, 150, 200),
(6, 6, 1, 0, 0, '10', '30', '14', '30', 10, 15, 1, 150, 200),
(7, 7, 1, 0, 0, '10', '30', '14', '30', 10, 15, 1, 150, 200),
(8, 8, 0, 0, 0, '', '', '', '', 0, 0, 0, 0, 0),
(9, 9, 1, 0, 0, '', '', '', '', 0, 0, 0, 0, 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `device_cumulative_history`
--

CREATE TABLE `device_cumulative_history` (
  `historyth_id` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  `month` varchar(250) NOT NULL,
  `counter` int(11) NOT NULL,
  `total_energy` float NOT NULL,
  `avg_voltage` float NOT NULL,
  `avg_ampere` float NOT NULL,
  `date_updated` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device_cumulative_history`
--

INSERT INTO `device_cumulative_history` (`historyth_id`, `device_id`, `month`, `counter`, `total_energy`, `avg_voltage`, `avg_ampere`, `date_updated`) VALUES
(1, 1, '05/2022', 3, 0.5052, 850, 0.49, '2022-05-21 11:41:08'),
(2, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-08-20 13:52:44'),
(4, 1, '06/2022', 7, 1.6177, 787.6, 0.451, '2022-06-16 17:51:53'),
(5, 1, '07/2022', 1, 1.7198, 736.545, 0.419091, '2022-07-05 16:01:06'),
(6, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-08-20 14:02:05'),
(7, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-08-20 14:06:50'),
(8, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-08-20 14:15:23'),
(9, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-08-23 18:52:29'),
(10, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-08-23 18:56:01'),
(11, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-08-31 18:56:19'),
(12, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-09-06 19:40:43'),
(13, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-09-06 19:44:13'),
(14, 1, '07/2022', 1, 0.1021, 226, 0.1, '2022-09-06 19:45:12'),
(15, 9, '07/2022', 1, 0.1021, 226, 0.1, '2022-09-08 17:32:06');

-- --------------------------------------------------------

--
-- Struktur dari tabel `device_history`
--

CREATE TABLE `device_history` (
  `history_id` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  `energy` float NOT NULL DEFAULT 0,
  `energy_total` float NOT NULL,
  `power` float NOT NULL DEFAULT 0,
  `voltage` float NOT NULL DEFAULT 0,
  `ampere` float NOT NULL DEFAULT 0,
  `powerFactor` float NOT NULL DEFAULT 0,
  `lightSensor` float NOT NULL DEFAULT 0,
  `battery` float NOT NULL DEFAULT 0,
  `brightness` float NOT NULL DEFAULT 0,
  `device_cons` varchar(200) NOT NULL,
  `history_date` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device_history`
--

INSERT INTO `device_history` (`history_id`, `device_id`, `energy`, `energy_total`, `power`, `voltage`, `ampere`, `powerFactor`, `lightSensor`, `battery`, `brightness`, `device_cons`, `history_date`) VALUES
(1, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2022-05-17 17:51:53'),
(2, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2022-06-12 17:51:53'),
(3, 3, 450, 0, 150, 350, 120, 150, 230, 90, 80, '1', '2022-06-15 17:51:53'),
(4, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '0', '2022-06-12 17:51:53'),
(5, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '0', '2022-06-12 17:51:53'),
(6, 3, 450, 0, 150, 350, 120, 150, 230, 90, 80, '1', '2022-06-15 17:51:53'),
(7, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '0', '2022-06-15 08:47:52'),
(8, 2, 550, 0, 250, 300, 100, 150, 230, 100, 75, '1', '2022-06-15 17:51:53'),
(9, 3, 650, 0, 150, 350, 120, 150, 230, 90, 80, '1', '2022-06-15 17:51:53'),
(10, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2022-06-15 17:51:53'),
(11, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '0', '2022-06-13 17:51:53'),
(12, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '0', '2022-06-13 17:51:53'),
(13, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '0', '2022-06-15 08:47:50'),
(14, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '0', '2022-06-16 17:51:53'),
(15, 2, 550, 0, 200, 320, 110, 150, 230, 85, 90, '1', '2022-06-21 11:40:15'),
(16, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2022-05-19 11:41:00'),
(17, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2022-05-21 11:41:08'),
(18, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(19, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(20, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(21, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(22, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(23, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(24, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(25, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(26, 1, 0.1684, 0, 25, 850, 0.49, 1, 18837, 10.74, 38, '1', '2646-04-05 16:01:06'),
(27, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, '1', '2022-07-03 19:43:23'),
(28, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, '1', '2022-07-03 19:43:31'),
(29, 1, 0.1021, 0.1021, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-06-05 16:01:06'),
(30, 1, 0.1021, 0.2042, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-06-05 16:01:06'),
(31, 1, 0.1021, 0.3063, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-06-05 16:01:06'),
(32, 1, 0.1021, 0.4084, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-06-05 16:01:06'),
(33, 1, 0.1021, 0.5105, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(34, 1, 0.1021, 0.6126, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(35, 1, 0.1021, 0.7147, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(36, 1, 0.1021, 0.8168, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(37, 1, 0.1021, 0.9189, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(38, 1, 0.1021, 1.021, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(39, 1, 0.1021, 1.1231, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(40, 1, 0.1021, 1.2252, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(41, 1, 0.1021, 1.3273, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(42, 1, 0.1021, 1.4294, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(43, 1, 0.1021, 1.5315, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(44, 1, 0.1021, 1.6336, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(45, 1, 0.1021, 1.7357, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(46, 1, 0.1021, 1.8378, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(47, 1, 0.1021, 1.9399, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(48, 1, 0.1021, 2.042, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(49, 1, 0.1021, 2.1441, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(50, 1, 0.1021, 2.2462, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(51, 1, 0.1021, 2.3483, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(52, 1, 0.1021, 2.4504, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(53, 1, 0.1021, 2.5525, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(54, 1, 0.1021, 2.6546, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(55, 1, 0.1021, 2.7567, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(56, 1, 0.1021, 2.8588, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(57, 1, 0.1021, 2.9609, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(58, 1, 0.1021, 3.063, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(59, 1, 0.1021, 3.1651, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(60, 1, 0.1021, 3.2672, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(61, 1, 0.1021, 3.3693, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(62, 1, 0.1021, 3.4714, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(63, 1, 0.1021, 3.5735, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(64, 1, 0.1021, 3.6756, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(65, 1, 0.1021, 3.7777, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(66, 1, 0.1021, 3.8798, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(67, 1, 0.1021, 3.9819, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(68, 1, 0.1021, 4.084, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(69, 1, 0.1021, 4.1861, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(70, 1, 0.1021, 4.2882, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(71, 1, 0.1021, 4.3903, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(72, 1, 0.1021, 4.4924, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(73, 1, 0.1021, 4.5945, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(74, 1, 0.1021, 4.6966, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(75, 1, 0.1021, 4.7987, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(76, 1, 0.1021, 4.9008, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(77, 1, 0.1021, 5.0029, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(78, 1, 0.1021, 5.105, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(79, 1, 0.1021, 5.2071, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(80, 1, 0.1021, 5.3092, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(81, 1, 0.1021, 5.4113, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(82, 1, 0.1021, 5.5134, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(83, 1, 0.1021, 5.6155, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(84, 1, 0.1021, 5.7176, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(85, 1, 0.1021, 5.8197, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(86, 1, 0.1021, 5.9218, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(87, 1, 0.1021, 6.0239, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(88, 1, 0.1021, 6.126, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(89, 1, 0.1021, 6.2281, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(90, 1, 0.1021, 6.3302, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(91, 1, 0.1021, 6.4323, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(92, 1, 0.1021, 6.5344, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(93, 1, 0.1021, 6.6365, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(94, 1, 0.1021, 6.7386, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(95, 1, 0.1021, 6.8407, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(96, 1, 0.1021, 6.9428, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(97, 1, 0.1021, 7.0449, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(98, 1, 0.1021, 7.147, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(99, 1, 0.1021, 7.2491, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(100, 1, 0.1021, 7.3512, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(101, 1, 0.1021, 7.4533, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(102, 1, 0.1021, 7.5554, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(103, 1, 0.1021, 7.6575, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(104, 1, 0.1021, 7.7596, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(105, 1, 0.1021, 7.8617, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(106, 1, 0.1021, 7.9638, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(107, 1, 0.1021, 8.0659, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(108, 0, 0.1021, 0.1021, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(109, 1, 0.1021, 8.168, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(110, 1, 0.1021, 8.2701, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(111, 1, 0.1021, 8.3722, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(112, 1, 0.1021, 8.4743, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(113, 1, 0.1021, 8.5764, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(114, 1, 0.1021, 8.6785, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(115, 1, 0.1021, 8.7806, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(116, 1, 0.1021, 8.8827, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(117, 1, 0.1021, 8.9848, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(118, 1, 0.1021, 9.0869, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(119, 1, 0.1021, 9.189, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(120, 1, 0.1021, 9.2911, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(121, 1, 0.1021, 9.3932, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(122, 1, 0.1021, 9.4953, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(123, 1, 0.1021, 9.5974, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(124, 1, 0.1021, 9.6995, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(125, 1, 0.1021, 9.8016, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(126, 1, 0.1021, 9.9037, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(127, 1, 0.1021, 10.0058, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(128, 1, 0.1021, 10.1079, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(129, 1, 0.1021, 10.21, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(130, 1, 0.1021, 10.3121, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(131, 1, 0.1021, 10.4142, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(132, 1, 0.1021, 10.5163, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(133, 1, 0.1021, 10.6184, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(134, 1, 0.1021, 10.7205, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(135, 1, 0.1021, 10.8226, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(136, 1, 0.1021, 10.9247, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(137, 1, 0.1021, 11.0268, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(138, 1, 0.1021, 11.1289, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(139, 1, 0.1021, 11.231, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(140, 1, 0.1021, 11.3331, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(141, 1, 0.1021, 11.4352, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(142, 1, 0.1021, 11.5373, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(143, 1, 0.1021, 11.6394, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(144, 1, 0.1021, 11.7415, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(145, 1, 0.1021, 11.8436, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(146, 1, 0.1021, 11.9457, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(147, 1, 0.1021, 12.0478, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(148, 1, 0.1021, 12.1499, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(149, 1, 0.1021, 12.252, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(150, 1, 0.1021, 12.3541, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(151, 1, 0.1021, 12.4562, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(152, 1, 0.1021, 12.5583, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(153, 1, 0.1021, 12.6604, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(154, 1, 0.1021, 12.7625, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(155, 1, 0.1021, 12.8646, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06'),
(156, 9, 0.1021, 0.1021, 25, 226, 0.1, 1, 8190, 4.5, 38, '1', '2022-07-05 16:01:06');

-- --------------------------------------------------------

--
-- Struktur dari tabel `device_monitoring`
--

CREATE TABLE `device_monitoring` (
  `deviceM_id` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  `energy` float NOT NULL,
  `energy_total` float NOT NULL,
  `power` float NOT NULL,
  `voltage` float NOT NULL,
  `ampere` float NOT NULL,
  `powerFactor` float NOT NULL,
  `lightSensor` float NOT NULL,
  `battery` float NOT NULL,
  `brightness` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device_monitoring`
--

INSERT INTO `device_monitoring` (`deviceM_id`, `device_id`, `energy`, `energy_total`, `power`, `voltage`, `ampere`, `powerFactor`, `lightSensor`, `battery`, `brightness`) VALUES
(1, 1, 0.1021, 12.8646, 25, 226, 0.1, 1, 8190, 4.5, 38),
(2, 2, 550, 0, 250, 300, 100, 150, 230, 100, 75),
(3, 3, 550, 0, 250, 300, 100, 150, 230, 100, 75),
(4, 4, 550, 0, 250, 300, 100, 150, 230, 100, 75),
(5, 5, 550, 0, 250, 300, 100, 150, 230, 100, 75),
(6, 6, 550, 0, 250, 300, 100, 150, 230, 100, 75),
(7, 7, 550, 0, 250, 300, 100, 150, 230, 100, 75),
(8, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(9, 9, 0.1021, 0.1021, 25, 226, 0.1, 1, 8190, 4.5, 38),
(10, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `device_month`
--

CREATE TABLE `device_month` (
  `deviceth_id` int(11) NOT NULL,
  `device_id` int(11) NOT NULL,
  `date` varchar(250) NOT NULL,
  `energy_m` float NOT NULL,
  `energy_price` float NOT NULL,
  `date_updated` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device_month`
--

INSERT INTO `device_month` (`deviceth_id`, `device_id`, `date`, `energy_m`, `energy_price`, `date_updated`) VALUES
(1, 1, '05/2022', 0.2042, 276.078, '2022-07-13 18:28:05'),
(4, 1, '06/2022', 0.2042, 276.078, '2022-07-15 19:12:33'),
(5, 1, '07/2022', 12.4562, 16840.8, '2022-07-15 19:13:41'),
(6, 0, '07/2022', 0.1021, 0, '2022-07-18 15:28:00'),
(7, 9, '07/2022', 0.1021, 138.039, '2022-09-08 17:32:06');

-- --------------------------------------------------------

--
-- Struktur dari tabel `device_sts_d`
--

CREATE TABLE `device_sts_d` (
  `id` int(11) NOT NULL,
  `sts_name` varchar(200) NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `device_sts_d`
--

INSERT INTO `device_sts_d` (`id`, `sts_name`, `date_created`) VALUES
(1, 'Aktif', '2022-06-12 08:37:45'),
(2, 'Offline', '2022-06-12 08:37:45'),
(3, 'Error', '2022-06-12 08:37:45');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
  `name` varchar(200) NOT NULL,
  `email` varchar(200) NOT NULL,
  `password` varchar(300) NOT NULL,
  `avatar` varchar(200) NOT NULL,
  `date_updated` datetime NOT NULL DEFAULT current_timestamp(),
  `date_created` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`user_id`, `name`, `email`, `password`, `avatar`, `date_updated`, `date_created`) VALUES
(1, 'Rizqi', 'rizqi@gmail.com', '$2a$04$4l9IdV8WgKkSSgvD3UHotOWC3c5Qcg5di7jEPVgFH6KptW2Pon8YG', 'Images/avatar.png', '2022-04-26 19:14:47', '2022-04-26 19:14:47'),
(2, 'Hilmi', 'Hilmi@gmail.com', '$2a$04$2/beL7PtXpWbpN1/y7VLU..Oq5fCjy9x5fASphnpSLnmxuWUMlArm', 'Images/avatar.png', '2022-06-12 08:57:24', '2022-06-12 08:57:28'),
(3, 'bangjago', 'bangjago@gmail.com', '$2a$04$2uyprOcf1suSYf9Qg98o4u44UXpnX19AYhZGHUwmU9RmuL9nxxa4.', 'Images/avatar.png', '2022-06-05 08:57:42', '2022-06-05 08:57:46'),
(4, 'wijaya', 'wijaya@gmail.com', '$2a$04$HOJMyx8BEp/4EvbiB6WVQOXzXT7vys7wTGkNbNLQTIp8XjcJvEfDC', 'Images/avatar.png', '2022-06-07 08:57:49', '2022-06-07 08:57:53'),
(5, 'UnitTest', 'unittest@gmail.com', '$2a$04$0IfhCYHp0oDJ.XTPIX906.GaovrZqoVXIsyIJVivOSwakSf0UjMES', 'Images/avatar.png', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(6, 'UnitTest', 'unittest@gmail.com', '$2a$04$jvcEdadSpkOrGAKnwfhv3uzOsLirgb4BqypY7PupDFAkKkaqGNX6m', 'Images/avatar.png', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(7, 'UnitTest', 'unittest@gmail.com', '$2a$04$UbA.rtVl4v.ZmobcS.ZXNOKUQtYaC8aNq3auXMSIUkFhrrHa64g0K', 'Images/avatar.png', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(8, 'UnitTest', 'unittest@gmail.com', '$2a$04$zbzXwF0vnEYm5kVGkq.sduXF17gd/3Tt.xyAMFhKRtel8vLJOG9ta', 'Images/avatar.png', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(9, 'UnitTest', 'unittest@gmail.com', '$2a$04$ns.gFDgsR1vS/u/p4OtPnOpmJWqYEIjmZ3hx1IoWoYJ.vn1G2TabK', 'Images/avatar.png', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(10, 'Rama', 'mhilmi999@gmail.com', '$2a$04$qwHZcPbuZ5VPhy/Bwx/vveSrQOQEKPK5l/rgv3R8sOdHJia3VYKzW', 'Images/avatar.png', '0000-00-00 00:00:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user_roles`
--

CREATE TABLE `user_roles` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `user_roles`
--

INSERT INTO `user_roles` (`id`, `user_id`, `role_id`) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 4),
(4, 5, 1),
(5, 6, 1),
(6, 7, 1),
(7, 8, 1),
(8, 9, 1),
(9, 10, 1);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `area`
--
ALTER TABLE `area`
  ADD PRIMARY KEY (`area_id`);

--
-- Indeks untuk tabel `area_division`
--
ALTER TABLE `area_division`
  ADD PRIMARY KEY (`division_id`);

--
-- Indeks untuk tabel `device`
--
ALTER TABLE `device`
  ADD PRIMARY KEY (`device_id`);

--
-- Indeks untuk tabel `device_cons_d`
--
ALTER TABLE `device_cons_d`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `device_controlling`
--
ALTER TABLE `device_controlling`
  ADD PRIMARY KEY (`deviceC_id`);

--
-- Indeks untuk tabel `device_cumulative_history`
--
ALTER TABLE `device_cumulative_history`
  ADD PRIMARY KEY (`historyth_id`);

--
-- Indeks untuk tabel `device_history`
--
ALTER TABLE `device_history`
  ADD PRIMARY KEY (`history_id`);

--
-- Indeks untuk tabel `device_monitoring`
--
ALTER TABLE `device_monitoring`
  ADD PRIMARY KEY (`deviceM_id`);

--
-- Indeks untuk tabel `device_month`
--
ALTER TABLE `device_month`
  ADD PRIMARY KEY (`deviceth_id`);

--
-- Indeks untuk tabel `device_sts_d`
--
ALTER TABLE `device_sts_d`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- Indeks untuk tabel `user_roles`
--
ALTER TABLE `user_roles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `area`
--
ALTER TABLE `area`
  MODIFY `area_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `area_division`
--
ALTER TABLE `area_division`
  MODIFY `division_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `device`
--
ALTER TABLE `device`
  MODIFY `device_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `device_cons_d`
--
ALTER TABLE `device_cons_d`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `device_controlling`
--
ALTER TABLE `device_controlling`
  MODIFY `deviceC_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `device_cumulative_history`
--
ALTER TABLE `device_cumulative_history`
  MODIFY `historyth_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT untuk tabel `device_history`
--
ALTER TABLE `device_history`
  MODIFY `history_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=157;

--
-- AUTO_INCREMENT untuk tabel `device_monitoring`
--
ALTER TABLE `device_monitoring`
  MODIFY `deviceM_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `device_month`
--
ALTER TABLE `device_month`
  MODIFY `deviceth_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `device_sts_d`
--
ALTER TABLE `device_sts_d`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `user_roles`
--
ALTER TABLE `user_roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
