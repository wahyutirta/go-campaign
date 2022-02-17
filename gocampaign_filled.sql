-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 17, 2022 at 10:03 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.0.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gocampaign`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `short_description` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `perks` text DEFAULT NULL,
  `backer_count` int(11) DEFAULT NULL,
  `goal_amount` int(11) DEFAULT NULL,
  `current_amount` int(11) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'campaign baru satu habis update', 'ini campaign baru banget', 'ini campaign baru banget fix sih ya gitu kata admin', 'satu perks, dua perks, tiga perks, empat perks', 0, 10000, 0, 'test-campaign', '2022-01-12 13:48:14', '2022-01-25 21:03:23'),
(2, 1, 'Test Penggalangan Dana', 'short test', 'long test', 'satu, dua, tiga', 0, 10000, 0, 'test-penggalangan-dana-1', '2022-01-22 23:54:51', '2022-01-22 23:54:51'),
(3, 1, 'campaign baru satu', 'ini campaign baru banget', 'ini campaign baru banget fix sih ya gitu kata admin', 'satu perks, dua perks, tiga perks', 0, 10000, 0, 'campaign-baru-satu-1', '2022-01-23 00:30:07', '2022-01-23 00:30:07'),
(4, 1, 'campaign baru satu habis update', 'ini campaign baru banget', 'ini campaign baru banget fix sih ya gitu kata admin', 'satu perks, dua perks, tiga perks, empat perks', 0, 10000, 0, 'campaign-baru-satu-habis-update-1', '2022-01-25 21:02:24', '2022-01-25 21:02:24'),
(5, 1, 'campaign baru satu habis update', 'ini campaign baru banget', 'ini campaign baru banget fix sih ya gitu kata admin', 'satu perks, dua perks, tiga perks, empat perks', 0, 10000, 0, 'campaign-baru-satu-habis-update-1', '2022-01-25 21:03:03', '2022-01-25 21:03:03'),
(6, 1, 'post new campaign', 'ini campaign baru banget', 'ini campaign baru banget fix sih ya gitu kata admin', 'satu perks, dua perks, tiga perks, empat perks', 0, 10000, 0, 'post-new-campaign-1', '2022-01-25 21:06:07', '2022-01-25 21:06:07');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `is_primary` tinyint(4) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'images/1-32089-NYH31X.jpg', 1, '2022-01-26 23:03:36', '2022-01-26 23:03:36'),
(2, 1, 'images/1-badut.jpg', 0, '2022-01-26 23:05:40', '2022-01-26 23:05:40');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 10000, 'paid', NULL, '2022-02-16 07:39:57', '2022-02-16 07:39:57'),
(2, 1, 2, 15000, 'paid', NULL, '2022-02-16 08:24:34', '2022-02-16 08:24:34');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(1, 'admin', 'student', 'admin@admin.com', '$2a$04$HwQDcayugfcT5Xjh2UVO4e9wir/gvJIMrrzznZP86j5MVKPFp93yS', 'images/1-1.PNG', 'user', NULL, '2022-01-12 20:45:54', '2022-01-12 20:47:28'),
(2, 'admin2', 'student', 'admin2@admin.com', '$2a$04$g/rkf5YLn3zB/XubwJ0I6epyNOYTqHQ8xbYPPCfSwN9KgGtTnTo3i', '', 'user', NULL, '2022-01-25 21:21:30', '2022-01-25 21:21:30');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
