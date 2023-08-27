-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Waktu pembuatan: 08 Sep 2022 pada 06.13
-- Versi server: 5.7.33
-- Versi PHP: 7.4.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `belajar-golang-database`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `comment`
--

CREATE TABLE `comment` (
  `id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `comment` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `comment`
--

INSERT INTO `comment` (`id`, `email`, `comment`) VALUES
(1, 'wandirayandra@gmail.com', 'Hallo Gaiz'),
(2, 'wandirayandra@gmail.com', 'Hallo Gaiz'),
(3, 'wandi0@gmail.com', 'Ini Komen Ke-0'),
(4, 'wandi1@gmail.com', 'Ini Komen Ke-1'),
(5, 'wandi2@gmail.com', 'Ini Komen Ke-2'),
(6, 'wandi3@gmail.com', 'Ini Komen Ke-3'),
(7, 'wandi4@gmail.com', 'Ini Komen Ke-4'),
(8, 'wandi5@gmail.com', 'Ini Komen Ke-5'),
(9, 'wandi6@gmail.com', 'Ini Komen Ke-6'),
(10, 'wandi7@gmail.com', 'Ini Komen Ke-7'),
(11, 'wandi8@gmail.com', 'Ini Komen Ke-8'),
(12, 'wandi9@gmail.com', 'Ini Komen Ke-9'),
(13, 'repository@test.com', 'Test Repository');

-- --------------------------------------------------------

--
-- Struktur dari tabel `costumer`
--

CREATE TABLE `costumer` (
  `id` varchar(225) NOT NULL,
  `name` varchar(225) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `costumer`
--

INSERT INTO `costumer` (`id`, `name`) VALUES
('1218015', 'Adrian Nugraha'),
('1218022', 'Rayandra Wandi Marselana');

-- --------------------------------------------------------

--
-- Struktur dari tabel `costummer`
--

CREATE TABLE `costummer` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(100) NOT NULL,
  `balance` int(11) NOT NULL,
  `rating` double NOT NULL,
  `birth_date` date NOT NULL,
  `maried` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `costummer`
--

INSERT INTO `costummer` (`id`, `name`, `email`, `balance`, `rating`, `birth_date`, `maried`) VALUES
('1218001', 'yayu', '', 123213, 2, '2022-09-01', 0),
('1218022', 'Rayandra Wandi Marselana', 'wandirayandra@gmail.com', 4000000, 5, '1999-01-01', 1),
('1218023', 'Raya Lil Ayunda', 'lilraya@gmail.com', 2000000, 3, '2010-10-10', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`username`, `password`) VALUES
('admin', 'admin');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `comment`
--
ALTER TABLE `comment`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `costumer`
--
ALTER TABLE `costumer`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `costummer`
--
ALTER TABLE `costummer`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`username`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `comment`
--
ALTER TABLE `comment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
