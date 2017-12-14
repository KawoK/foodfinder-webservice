-- phpMyAdmin SQL Dump
-- version 4.6.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Dec 14, 2017 at 09:05 AM
-- Server version: 5.6.33
-- PHP Version: 7.0.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `foodfinder`
--

-- --------------------------------------------------------

--
-- Table structure for table `Canteen`
--

CREATE TABLE `Canteen` (
  `id` int(11) NOT NULL,
  `nama_canteen` text NOT NULL,
  `location` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `Canteen`
--

INSERT INTO `Canteen` (`id`, `nama_canteen`, `location`) VALUES
(0, 'kantinSaraga', 'Saraga'),
(1, 'kantinBorju', 'GedungKimia'),
(2, 'kantinGKUBarat', 'GKUBarat'),
(3, 'kantinGKUTimur', 'GKUTimur'),
(4, 'kantinEtitu', 'CCTimur'),
(5, 'kantinCCBarat', 'CCBarat'),
(6, 'kantinLabtekBiru', 'SITH'),
(7, 'kantinEastCorner', 'FTTM'),
(8, 'kantinSalman', 'Salman'),
(9, 'TokemaSBM', 'SBM');

-- --------------------------------------------------------

--
-- Table structure for table `drink`
--

CREATE TABLE `drink` (
  `id` int(11) NOT NULL,
  `nama_minuman` text NOT NULL,
  `harga` text NOT NULL,
  `nama_canteen` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `drink`
--

INSERT INTO `drink` (`id`, `nama_minuman`, `harga`, `nama_canteen`) VALUES
(1, 'Bread - Crumbs, Bulk', '6135.06', 'Salman'),
(2, 'Turnip - Mini', '3108.35', 'Saraga'),
(3, 'Russian Prince', '2770.75', 'Borju'),
(4, 'Cheese - Wine', '1082.96', 'GKUBarat'),
(5, 'Cheese - Cheddar, Old White', '8819.33', 'GKUTimur'),
(6, 'Snapple - Mango Maddness', '2404.81', 'Etitu'),
(7, 'Pastry - Chocolate Marble Tea', '3318.92', 'CCBarat'),
(8, 'Beer - Heinekin', '4270.07', 'LabtekBiru'),
(9, 'Energy Drink', '3915.25', 'EastCorner'),
(10, 'Sesame Seed Black', '8959.31', 'TokemaSBM');

-- --------------------------------------------------------

--
-- Table structure for table `Food`
--

CREATE TABLE `Food` (
  `id` int(11) NOT NULL,
  `nama_makanan` text NOT NULL,
  `harga` text NOT NULL,
  `nama_canteen` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `Food`
--

INSERT INTO `Food` (`id`, `nama_makanan`, `harga`, `nama_canteen`) VALUES
(1, 'Hipnotiq Liquor', '4991.55', 'Salman'),
(2, 'Country Roll', '9929.42', 'Saraga'),
(3, 'Pail For Lid 1537', '7648.00', 'Borju'),
(4, 'Brandy Apricot', '3947.02', 'GKUBarat'),
(5, 'Pepper - Orange', '1576.56', 'GKUTimur'),
(6, 'Pastry - Banana Tea Loaf', '6408.18', 'Etitu'),
(7, 'Artichoke - Hearts, Canned', '295.83', 'CCBarat'),
(8, 'Foam Cup 6 Oz', '7924.40', 'LabtekBiru'),
(9, 'Cinnamon Buns Sticky', '8709.27', 'EastCorner'),
(10, 'Cabbage Roll', '5355.60', 'TokemaSBM');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Canteen`
--
ALTER TABLE `Canteen`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `drink`
--
ALTER TABLE `drink`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `Food`
--
ALTER TABLE `Food`
  ADD PRIMARY KEY (`id`);
