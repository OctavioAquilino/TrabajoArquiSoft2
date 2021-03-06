-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost
-- Tiempo de generación: 25-06-2022 a las 05:41:52
-- Versión del servidor: 10.4.24-MariaDB
-- Versión de PHP: 8.1.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Base de datos: `trabajoarqui`
--
CREATE DATABASE IF NOT EXISTS `trabajoarqui` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `trabajoarqui`;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `categories`:
--

--
-- Volcado de datos para la tabla `categories`
--

INSERT INTO `categories` VALUES
(1, 'Celulares'),
(2, 'Deporte y Fitness'),
(3, 'Televisores'),
(4, 'Indumentaria'),
(5, 'Videojuegos'),
(6, 'Electrodomesticos'),
(7, 'Electronicos y Audio'),
(8, 'Computacion'),
(9, 'Deco'),
(10, 'Vehiculos'),
(11, 'Libros'),
(12, 'Perfumes');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `monto_final` decimal(60,4) NOT NULL,
  `fecha` datetime NOT NULL,
  `id_user` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `orders`:
--   `id_user`
--       `users` -> `id`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `order_details`
--

CREATE TABLE `order_details` (
  `id` int(11) NOT NULL,
  `precio_unitario` decimal(60,4) NOT NULL,
  `cantidad` decimal(60,4) NOT NULL,
  `total` decimal(60,4) NOT NULL,
  `nombre` varchar(550) NOT NULL,
  `id_product` int(150) NOT NULL,
  `id_order` int(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `order_details`:
--   `id_order`
--       `orders` -> `id`
--   `id_product`
--       `products` -> `id`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(350) NOT NULL,
  `price` decimal(60,4) NOT NULL,
  `picture` varchar(350) NOT NULL,
  `stock` int(150) NOT NULL,
  `id_category` int(150) NOT NULL,
  `description` varchar(350) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `products`:
--   `id_category`
--       `categories` -> `id`
--

--
-- Volcado de datos para la tabla `products`
--

INSERT INTO `products` VALUES
(1, 'Samsung Galaxy S22', '800.0000', 'https://images.samsung.com/is/image/samsung/p6pim/es/2202/gallery/es-galaxy-s22-plus-s906-sm-s906blbgeub-530863207?$650_519_PNG$', 8, 1, '128gb,color celeste'),
(2, 'Samsung Galaxy S21', '700.0000', 'https://media.flixcar.com/f360cdn/Samsung-81181525-ar-galaxy-s21-5g-g991-sm-g991bzwlaro-368338861Download-Source-zoom.png', 5, 1, '64gb,color blanco'),
(3, 'Samsung Galaxy S20', '650.0000', 'https://www.cetrogar.com.ar/media/catalog/product/t/e/te2602_2.jpg?width=500&height=500&canvas=500,500&quality=80&bg-color=255,255,255&fit=bounds', 6, 1, '128gb,color gris'),
(4, 'Samsung Galaxy S10', '550.0000', 'https://images.fravega.com/f300/bad99152b8ae689a850661c347500ded.jpg.webp', 3, 1, '256gb,color negro'),
(5, 'Samsung Galaxy S9', '350.0000', 'https://http2.mlstatic.com/D_NQ_NP_613196-MLA31593018897_072019-O.jpg', 2, 1, '64gb,color violeta'),
(6, 'Iphone 13', '850.0000', 'https://newphonemx.com/wp-content/uploads/2022/04/Mzalx1Vqjqrfaa137ID_13v1.png', 11, 1, '128gb,color verde'),
(7, 'Iphone 12', '780.0000', 'https://www.macstation.com.ar/img/productos/small/2155-1111.jpg', 8, 1, '128gb,color azul'),
(8, 'Iphone 11', '600.0000', 'https://http2.mlstatic.com/D_NQ_NP_865864-MLA46114990464_052021-O.jpg', 5, 1, '128gb,color negro'),
(9, 'Iphone SE 2', '450.0000', 'https://http2.mlstatic.com/D_NQ_NP_2X_745945-MLA46552310508_062021-V.webp', 15, 1, '128gb,color blanco'),
(10, 'Iphone XR', '500.0000', 'https://http2.mlstatic.com/D_NQ_NP_873479-MLA46544714698_062021-O.jpg', 4, 1, '128gb,color coral'),
(11, 'Samsung Galaxy S22', '1000.0000', 'https://personalpy.vteximg.com.br/arquivos/ids/156245-1000-1000/s22-black-main.png?v=637826043219870000', 18, 1, '256gb,color negro'),
(12, 'Samsung Galaxy S21', '800.0000', 'https://images.samsung.com/is/image/samsung/p6pim/latin/galaxy-s21/gallery/latin-galaxy-s21-5g-g991-sm-g991bzajtpa-368316714?$650_519_PNG$', 3, 1, '128gb,color gris'),
(13, 'Samsung Galaxy S20', '570.0000', 'https://www.losdistribuidores.com/wp-content/uploads/2021/04/Samsung_s20_plus_4.jpg', 6, 1, '64gb,color celeste'),
(14, 'Samsung Galaxy S10', '500.0000', 'https://exitocol.vtexassets.com/arquivos/ids/11627860/GALAXY-S10-BLANCO-128-GB-1475882_a.jpg?v=637783932584630000', 7, 1, '128gb,color blanco'),
(15, 'Samsung Galaxy S9', '480.0000', 'https://resources.claroshop.com/medios-plazavip/mkt/5c6db9af06fc7_samsung_galaxy_s9_azuljpg.jpg', 5, 1, '128gb,color azul'),
(16, 'Iphone 13', '990.0000', 'https://www.macstation.com.ar/img/productos/2583-5.jpg', 25, 1, '256gb,color rojo'),
(17, 'Iphone 12', '620.0000', 'https://www.macstation.com.ar/img/productos/small/2317-1.jpg', 11, 1, '64gb,color violeta'),
(18, 'Iphone 11', '720.0000', 'https://http2.mlstatic.com/D_NQ_NP_699194-MLA46115014420_052021-O.jpg', 5, 1, '256gb,color amarillo'),
(19, 'Iphone SE 2', '390.0000', 'https://http2.mlstatic.com/D_NQ_NP_658260-MLA46552695787_062021-O.jpg', 7, 1, '64gb,color negro'),
(20, 'Iphone XR', '430.0000', 'https://http2.mlstatic.com/D_NQ_NP_718561-MLA46545234812_062021-O.jpg', 6, 1, '64gb,color celeste'),
(21, 'Paleta de pádel Wilson Bela Pro', '499.2500', 'https://wilsonstore.com.ar/media/catalog/product/cache/fac7280e0d0ff2a98b5e31fe1a2d4a3a/w/r/wr065511u_1_bela_pro_paddle_bl_rd.jpg', 25, 2, 'color negro,Peso 350g'),
(22, 'Raqueta Tenis Wilson Pro Open Aro 100', '753.2500', 'https://http2.mlstatic.com/D_NQ_NP_747861-MLA49204195621_022022-O.jpg', 12, 2, 'color amarillo,peso 300g'),
(23, 'Raquetero Wilson Team', '959.3300', 'https://d2r9epyceweg5n.cloudfront.net/stores/080/906/products/maletin-wilson-federer-team-iii-triple-red-tenis-d_nq_np_411725-mco25485196654_042017-f1-99ae8625e2fc252ad015220203994561-1024-1024.jpg', 8, 2, 'Con espacio para 2 raquetas'),
(24, 'Tubo Pelotas Tenis Wilson', '4.2500', 'https://http2.mlstatic.com/D_NQ_NP_924679-MLA44525562956_012021-O.jpg', 16, 2, '6 pelotas por tubo'),
(25, 'Paleta de pádel adidas Zukur', '442.7500', 'https://d3ugyf2ht6aenh.cloudfront.net/stores/069/005/products/paleta-padel-adidas-adipower1-eff525940d83f9ca0a15713199713314-1024-1024.jpg', 20, 2, 'color azul,Peso 350g'),
(26, 'Mochila Padel adidas', '99.8500', 'https://d3ugyf2ht6aenh.cloudfront.net/stores/069/005/products/mochila-adidas-adipower-19-gris1-71a17bf4498bb6fa6e15784268559145-1024-1024.jpg', 7, 2, 'color gris,capacidad 50lt'),
(27, 'Camiseta de la Selección Argentina', '110.4800', 'https://assets.adidas.com/images/w_600,f_auto,q_auto/10c04e97f68143d6ba3eae6101429157_9366/Camiseta_Titular_Seleccion_Argentina_Messi_10_Blanco_GA7550_01_laydown.jpg', 35, 2, 'Remera oficial de la Seleccion Argentina - 10'),
(28, 'Camiseta de PSG', '123.2500', 'https://www.dexter.com.ar/on/demandware.static/-/Sites-dabra-catalog/default/dw170585e1/products/NI_CV7903-411/NI_CV7903-411-1.JPG', 28, 2, 'Remera oficial del PSG - 30'),
(29, 'Camiseta de Boca Junior', '150.0000', 'https://assets.adidas.com/images/w_600,f_auto,q_auto/e245a16bc8384eb98bfbab3c013f08f4_9366/Camiseta_Local_Boca_Juniors_20-21_Azul_GL4174_01_laydown.jpg', 30, 2, 'Remera oficial de Boca Juniors - 9'),
(30, 'Camiseta de River Plate', '150.0000', 'https://tiendariver.vteximg.com.br/arquivos/ids/168463-1000-1000/GU9603_FR_Torso_Ecom.jpg?v=637572254211130000', 28, 2, 'Remera oficial de River Plate - 9'),
(31, 'Camiseta de Manchester United', '125.7800', 'https://assets.adidas.com/images/w_600,f_auto,q_auto/48dbcd9454ca4735af53ab7400cf1bc4_9366/Camiseta_Local_Manchester_United_20-21_Rojo_FM4292_01_laydown.jpg', 14, 2, 'Remera oficial del Manchester United -7'),
(32, 'Camiseta de Juventus', '100.0000', 'https://assets.adidas.com/images/w_600,f_auto,q_auto/31149ec3b1b84784bce2acda00f28ffa_9366/Camiseta_Local_Juventus_21-22_Blanco_GR0604_01_laydown.jpg', 7, 2, 'Remera oficial del Juventus - 10'),
(33, 'Camiseta de Bayer Munich', '150.0000', 'https://todosobrecamisetas.com/wp-content/uploads/bayern-munich-2022-23-adidas-home-kit-9.jpg', 19, 2, 'Remera oficial del Bayer Munich - 9'),
(34, 'Kaddygolf Set Completo Hombre Palos Golf Wilson Profile', '1000.0000', 'https://http2.mlstatic.com/D_NQ_NP_851948-MLA46345761423_062021-O.webp', 3, 2, 'Set completo con los 10 palos de golf'),
(35, 'Smart TV Noblex', '125.4800', 'https://images.fravega.com/f1000/d4e7ff5ed447687612e508f0cb5391d2.jpg', 18, 3, 'LED HD 32 pulgadas 220V'),
(36, 'Smart TV Samsung', '195.0000', 'https://images.samsung.com/is/image/samsung/latin-uhdtv-nu7095-un55nu7095gxzs-frontblack-229131451?$2160_1728_PNG$', 25, 3, 'LED HD 65 pulgadas 220V'),
(37, 'Smart TV Lg', '184.0000', 'https://www.lg.com/ar/images/televisores/md07529642/gallery/UN7310-1-1100x730.jpg', 10, 3, 'LED HD 40 pulgadas 220V'),
(38, 'Smart TV Lg', '128.0000', 'https://www.lg.com/ar/images/televisores/md07529642/gallery/UN7310-1-1100x730.jpg', 10, 3, 'LED HD 32 pulgadas 220V'),
(39, 'Smart TV Samsung', '180.0000', 'https://images.samsung.com/is/image/samsung/latin-uhdtv-nu7095-un55nu7095gxzs-frontblack-229131451?$2160_1728_PNG$', 28, 3, 'LED HD 32 pulgadas 220V'),
(40, 'Smart TV Philips', '178.2500', 'https://www.philips.es/c-dam/b2c/tv/categorypage/gb/smarttv-redesign/pus8506-8536-hero-banner.png', 16, 3, 'LED HD 32 pulgadas 220V'),
(41, 'Smart TV Philips', '148.0000', 'https://www.philips.es/c-dam/b2c/tv/categorypage/gb/smarttv-redesign/pus8506-8536-hero-banner.png', 18, 3, 'LED HD 55 pulgadas 220V'),
(42, 'Smart TV Sony', '101.0500', 'https://http2.mlstatic.com/D_NQ_NP_745471-MLA43656886736_102020-O.jpg', 7, 3, 'LED HD 32 pulgadas 220V'),
(43, 'Smart TV Sony', '115.7200', 'https://http2.mlstatic.com/D_NQ_NP_745471-MLA43656886736_102020-O.jpg', 10, 3, 'LED HD 50 pulgadas 220V'),
(44, 'Campera Pluma Hombre', '120.0000', 'http://d3ugyf2ht6aenh.cloudfront.net/stores/807/624/products/07-hombre-frente-azul_marino1-29841e0ebef170d26015974533467238-640-0.jpg', 8, 4, 'Campera Pluma Hombre Columbia Snow Country C/capucha'),
(45, 'Campera Pluma Mujer', '122.0000', 'https://www.cordonandino.com/img/articulos/campera_ansilta_de_mujer_piuquen_3_pluma_800_fp_allied_imagen2.jpg', 15, 4, 'Columbia Sister Brook Omni Heat'),
(46, 'Polo con estampado con corte regular Hombre', '100.0000', 'https://tommyargentina.vteximg.com.br/arquivos/ids/196297-753-1004/MW0MW15498_YBR_1.jpg?v=637768283012330000', 18, 4, 'Tommy Hilfiger - Hombre'),
(47, 'Polo con estampado con corte regular', '80.0000', 'https://www.ornellafashion.com/wp-content/uploads/2020/10/POLO-DE-ALGODON-negro-con-letras-doradas-MARCA-Guess-Atras.jpg', 22, 4, 'Guess - Hombre'),
(48, 'Polo lisa con corte regular', '48.2500', 'https://tommyargentina.vteximg.com.br/arquivos/ids/182353-753-1004/MW0MW10765_C5M_1.jpg?v=637587205804370000', 11, 4, 'Tommy Hilfiger - Hombre'),
(49, 'Remera de corte cropped con logo', '20.9000', 'https://tommyargentina.vteximg.com.br/arquivos/ids/192133-753-1004/DW0DW09070_C87_4.jpg?v=637768211168500000', 12, 4, 'Tommy Hilfiger - Mujer'),
(50, 'Remera de corte cropped con logo', '88.0000', 'https://img01.ztat.net/article/spp-media-p1/36db116b7dbf39e4831b9484b3ad0ce5/c70fc9e2f2f440eb8b7989780580b667.jpg?imwidth=762', 27, 4, 'Gucci - Mujer'),
(51, 'Chomba de Algodon con logo', '65.0000', 'https://tommyargentina.vteximg.com.br/arquivos/ids/183615-753-1004/DW0DW06813_P01_5.jpg?v=637589008792830000', 14, 4, 'Tommy Hilfiger - Mujer'),
(52, 'Buzo con logo', '79.9900', 'https://tommyargentina.vteximg.com.br/arquivos/ids/186529-500-667/WW0WW29236PKH.jpg?v=637689635940900000', 9, 4, 'Tommy Hilfiger - Mujer'),
(53, 'Chomba de Algodon con logo', '100.0000', 'https://http2.mlstatic.com/D_NQ_NP_722433-MLA48905681187_012022-W.jpg', 5, 4, 'Gucci - Hombre'),
(54, 'FIFA 20', '60.0000', 'https://i.3djuegos.com/juegos/16637/fifa_20/fotos/ficha/fifa_20-4971679.jpg', 25, 5, 'EA - 2019 - PS4'),
(55, 'FIFA 21', '65.0000', 'https://media.contentapi.ea.com/content/dam/ea/fifa/fifa-21/buy/common/champions/fifa21-ce-pre-ps4-front-norating-rgb.png', 22, 5, 'EA - 2020 - PS4'),
(56, 'FIFA 22', '70.0000', 'https://arsonyb2c.vtexassets.com/arquivos/ids/357835/PS5-FIFA-22-Cover.jpg?v=637695585994970000', 18, 5, 'EA - 2021 - PS5'),
(57, 'Pro Evolution Soccer 2020', '60.0000', 'https://www.portalgames.com.ar/wp-content/uploads/2019/11/Pes-2020.jpg', 28, 5, 'Konami - 2019 - PS4'),
(58, 'Pro Evolution Soccer 2021', '64.0000', 'https://images.fravega.com/f1000/e92d4654f4f0b8300c717e6ed92206d2.jpg', 15, 5, 'Konami - 2020- PS4'),
(59, 'Pro Evolution Soccer 2022', '70.0000', 'https://as01.epimg.net/meristation/imagenes/2021/09/02/noticias/1630576478_746285_1630576536_portada_normal.jpg', 30, 5, 'Konami - 2021 - PS5'),
(60, 'Pokemon Sword', '45.0000', 'https://m.media-amazon.com/images/I/81-QH2Lyy4L._SL1500_.jpg', 23, 5, 'Nintendo - 2019 - Nintendo Switch'),
(61, 'Pokemon Shield', '45.0000', 'https://worldfantasy.com.ar/wp-content/uploads/2020/09/71lz62-F84L._AC_SL1500_.jpg', 21, 5, 'Nintendo - 2019- Nintendo Switch'),
(62, 'Pokemon Sun', '34.0000', 'https://m.media-amazon.com/images/I/81mWKPCNRPL._SX425_.jpg', 8, 5, 'Nintendo - 2017- Nintendo 3DS'),
(63, 'Pokemon Moon', '34.0000', 'https://images.fravega.com/f500/fc8978f6cca3b045f0850b421c2b74f7.jpg', 12, 5, 'Nintendo - 2017- Nintendo 3DS'),
(64, 'Heladera French Door SpaceMax 629L', '750.0000', 'https://images.samsung.com/is/image/samsung/ar-french-door-rf28k9380sr-rf28k9380sr-bg-frontsilver-61646139?$2160_1728_PNG$', 5, 6, 'color negro - Samsung'),
(65, 'Heladera French Door con Beverage Center™, 713L', '850.0000', 'https://samsungpy.vtexassets.com/arquivos/ids/159557/Samsung-94445744-py-beverage-center-rf71a9671sg-pe-502077882Download-Source.png?v=637686053288770000', 8, 6, 'color plateada - Samsung'),
(66, 'Heladera French Door con Tecnología Triple Cooling, 684 L', '620.0000', 'https://http2.mlstatic.com/D_NQ_NP_697535-MLA31644801479_072019-O.jpg', 10, 6, 'color plateada - Samsung'),
(67, 'Lavarropas 7Kg Inverter con EcoBubble', '157.2200', 'https://images.samsung.com/is/image/samsung/ar-washer-ww70j4463gs-ww70j4463gsubg-frontsilver-148671992?$2160_1728_PNG$', 25, 6, 'color blanco - Samsung'),
(68, 'Lavarropas 6,5kg con control de voltaje', '100.0000', 'https://samsungar.vtexassets.com/arquivos/ids/156903/Samsung-44801097-ar-ww65m0nhwu-ww65m0nhwu-xbg-kg-149275648PD_GALLERY_PNG.png?v=636995947631600000', 15, 6, 'color plateado - Samsung'),
(69, 'Microondas con Grill de Interior Cerámico, 28L', '95.8800', 'https://images.samsung.com/is/image/samsung/ar-microwave-oven-grill-mg28f3k3tas-mg28f3k3tas-bg-001-front-silver?$2160_1728_PNG$', 10, 6, 'color negro - Samsung'),
(70, 'Microondas con Grill de Interior Cerámico, 23L', '101.9900', 'https://images.samsung.com/is/image/samsung/ar-microwave-oven-grill-mg23f3k3tak-mg23f3k3tak-bg-004-front-black?$650_519_PNG$', 8, 6, 'color negro - Samsung'),
(71, 'JBL Flip 5 Altavoz inalámbrico portátil con Bluetooth', '55.0000', 'https://audioimport.com.ar/wp-content/uploads/2020/06/jbl-flip-5-parlante-bluetooth-portatil-inalambrico-D_NQ_NP_640868-MLA40809778396_022020-F-1.jpg', 15, 7, 'Hasta 12h de reproducción con sonido de calidad, negro'),
(72, 'JBL Charge 5 Altavoz inalámbrico portátil con Bluetooth', '60.0000', 'https://m.media-amazon.com/images/I/71TRom4y1+L._AC_SY450_.jpg', 25, 7, 'Hasta 20h de reproducción con sonido de calidad, turquesa'),
(73, 'JBL Xtreme 3 Altavoz Bluetooth', '58.5500', 'https://www.culturasonora.es/wp-content/uploads/2021/03/Xtreme-3-cuerpo-cultura.jpg', 18, 7, 'Hasta 15h de reproducción con sonido de calidad, negro'),
(74, 'Sony MHC-V73D Altavoz de Alta Potencia', '76.9900', 'https://www.sony.com.ar/image/3b29c6c5bca35ecb4846d7a71dbfd6e9?fmt=pjpeg&wid=330&bgcolor=FFFFFF&bgc=FFFFFF', 21, 7, 'Con el Sonido y Luces de Fiesta, negro'),
(75, 'JBL GO 3 Altavoz inalámbrico portátil con Bluetooth', '21.5500', 'https://cf.shopee.com.ar/file/1d020940c8dee1842b0914fb35ed3b4b_tn', 30, 7, 'Hasta 5h de reproducción con sonido de calidad, negro'),
(76, 'Dell XPS 13 OLED', '500.0000', 'https://m.media-amazon.com/images/I/81BHuVIOtaL._AC_SS450_.jpg', 15, 8, 'Core i3-1115G4, 8 GB de RAM, un 256 GB SSD'),
(77, 'Apple MacBook Air M1', '1000.0000', 'https://cdn-ipoint.waugi.com.ar/20564-large_default/macbook-air-13-m1-2020-8-core-cpu-256-gb-gold.jpg', 25, 8, 'Chip M1, puedes configurarla hasta 16 GB de RAM y hasta 2 TB SSD'),
(78, 'Acer Swift 3 (Ryzen)', '700.0000', 'https://m.media-amazon.com/images/I/81nN5u1MEuL._AC_SL1500_.jpg', 5, 8, ' 8 GB de RAM y un SSD de 512 GB'),
(79, 'Apple MacBook Pro 16', '790.0000', 'https://cdn-ipoint.waugi.com.ar/22343-large_default/macbook-pro-16-apple-m1-max-chip-10core-cpu-32core-gpu-1tb-ssd-silver.jpg', 35, 8, 'Chip M1 Max, puedes configurarla hasta 32 GB de RAM y hasta 2 TB SSD'),
(80, 'Dell XPS 17', '800.0000', 'https://www.notebookcheck.org/uploads/tx_nbc2/DellXPS17-9700__1__02.jpg', 20, 8, 'Core i9-1115G4, 8 GB de RAM, un 256 GB SSD'),
(81, 'Razer Blade 14', '520.0000', 'https://assets3.razerzone.com/qtzQaOk8i7eFFPUVrSDDszOQa80=/1500x1000/https%3A%2F%2Fhybrismediaprod.blob.core.windows.net%2Fsys-master-phoenix-images-container%2Fhbd%2Fh06%2F9286402375710%2Fblade14-p8-fhd-500x500.png', 10, 8, 'Gran pantalla de 1440p de 250 Hz'),
(82, 'Google Pixelbook Go', '760.0000', 'https://lh3.googleusercontent.com/cewixHQrBsI-iviE4qPNPLppaYuNTccxIBTi9v2XusjhRvp-UdBpOAYr78exyrJPM5lyFjWHnEQFBSUyJuSSCd3sI-UGN67G8Nbi=s2048', 4, 8, 'Cuenta con sistema operativo propio de google'),
(83, 'Microsoft Surface Pro 8', '880.0000', 'https://p.turbosquid.com/ts-thumb/Ie/w27PWC/EB/microsoft_surface_pro8_gray_01/jpg/1634900784/600x600/fit_q87/3f1c0d309e5eede7e508cc4ab8f935b2403db24d/microsoft_surface_pro8_gray_01.jpg', 5, 8, 'Es la mejor representación de una 2-en-1 que Windows'),
(84, 'HP Elite Dragonfly G2', '1101.0000', 'https://ssl-product-images.www8-hp.com/digmedialib/prodimg/lowres/c07024384.png', 11, 8, 'Intel Core de 11 generación, con la opción vPro amigable a los negocios, vienen con hasta 32 GB de RAM y hasta 2 TB'),
(85, 'Lenovo Chromebook Flex 5', '520.0000', 'https://m.media-amazon.com/images/I/61qIQqN3blS._AC_SL1100_.jpg', 5, 8, 'ntel Core i3, una buena pantalla de 1080p y una batería impresionante'),
(86, 'Iphone SE 2', '550.0000', 'https://http2.mlstatic.com/D_NQ_NP_2X_951846-MLA46552056642_062021-V.webp', 15, 1, '256gb,color rojo'),
(87, 'Samsung Galaxy S22', '900.5800', 'https://unionstore.com/image/cache/catalog/2022/Samsung/Samsung%20-%20S22%20%208%20128GB%20Violet-600x800.jpg', 12, 1, '128gb,color violeta'),
(88, 'Puff de pelo largo natural', '45.2400', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/3216152?fmt=jpg&fit=fit,1&wid=200&hei=200', 36, 9, 'Ancho 40cm, alto 40cm, profundidad 40cm, peso 2kg, color natural'),
(89, 'Sillón Wald 3 cuerpos gris claro', '216.9800', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/2413000?fmt=jpg&fit=fit,1&wid=200&hei=200', 36, 9, 'Ancho 86cm, alto 188cm, profundidad 85cm, peso 32kg, color gris'),
(90, 'Futón Niza 2 cuerpos 1 1/2 plaza de eco cuero negro', '289.3700', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/2179504?fmt=jpg&fit=fit,1&wid=200&hei=200', 36, 9, 'Ancho 76cm, alto 180cm, profundidad 91cm, peso 36kg, color natural'),
(91, 'Juego de comedor Ciudades mesa + 4 sillas diseño', '312.0700', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/2179660?fmt=jpg&fit=fit,1&wid=200&hei=200', 12, 9, 'Set de comedor, mesa de vidrio y metal, sillas de cuero'),
(92, 'Mesa de exterior redonda 60 cm', '73.8000', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/2482886?fmt=jpg&fit=fit,1&wid=200&hei=200', 47, 9, 'Mesa de aluminio, color metal, con un diametro de 60cm'),
(93, 'Mesa de exterior cuadrada 70 cm', '72.3700', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/2482894?fmt=jpg&fit=fit,1&wid=200&hei=200', 26, 9, 'Mesa de aluminio, color metal, con un diametro de 70cm'),
(94, 'Canvas con bastidor', '35.8900', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/222223X?fmt=jpg&fit=fit,1&wid=200&hei=200', 85, 9, 'Decoración de pared, con una medida de 50x100'),
(95, 'Canvas con bastidor', '32.8000', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/2222248?fmt=jpg&fit=fit,1&wid=200&hei=200', 63, 9, 'Decoración de pared, con una medida de 40x120'),
(96, 'Reloj de pared Wooden', '11.2500', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/2660342?fmt=jpg&fit=fit,1&wid=200&hei=200', 26, 9, 'Medida del reloj 29x29'),
(97, 'Reloj de pared New-Mod', '9.9900', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/3218090?fmt=jpg&fit=fit,1&wid=200&hei=200', 15, 9, 'Medida del reloj 30x30'),
(98, 'Puff Jazz negro', '28.8900', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/3201325?fmt=jpg&fit=fit,1&wid=200&hei=200', 12, 9, 'Ancho 40cm, alto 40cm, profundidad 40cm, peso 2kg, color negro'),
(99, 'Puff Jazz beige', '26.3000', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/3201317?fmt=jpg&fit=fit,1&wid=200&hei=200', 25, 9, 'Ancho 40cm, alto 40cm, profundidad 40cm, peso 2kg, color beige'),
(100, 'Puff de pana Jazz Capitone negro', '22.7300', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/3216136?fmt=jpg&fit=fit,1&wid=200&hei=200', 10, 9, 'Ancho 40cm, alto 40cm, profundidad 40cm, peso 2kg, color negro'),
(101, 'Puff Jazz Marrón', '18.2100', 'https://cote-montagne.fr/1842-large_default/pouf-rond-microfibre-cuir-vieilli-leonies-marron-composition-tissu-microfibre-leonies.jpg', 10, 9, 'Ancho 40cm, alto 40cm, profundidad 40cm, peso 2kg, color marron'),
(102, 'Puff Jazz gris', '22.7300', 'https://sodimacar.scene7.com/is/image/SodimacArgentina/3201309?fmt=jpg&fit=fit,1&wid=200&hei=200', 16, 9, 'Ancho 40cm, alto 40cm, profundidad 40cm, peso 2kg, color gris'),
(103, 'Jeep Renegade Sport', '30000.0000', 'https://centraljeep.divit.com.ar/wp-content/uploads/2022/06/2019_4_renegade_Sport_Anti_Verme-removebg-preview-1.png', 18, 10, '4x2, automatico, rojo'),
(104, 'Jeep Renegade Sport', '29500.0000', 'https://centraljeep.divit.com.ar/wp-content/uploads/2022/06/2019_5_renegade_Sport_Anti_Bille-removebg-preview.png', 15, 10, '4x2, manual, gris'),
(105, 'Jeep Renegade Sport', '30000.0000', 'https://centraljeep.divit.com.ar/wp-content/uploads/2022/06/2019_1_renegade_Sport_Anti_Branc-removebg-preview.png', 25, 10, '4x2, automatico, blanco'),
(106, 'Jeep Renegade Sport', '29500.0000', 'https://www.jeep-latam.com/wp-content/uploads/2020/04/SportPLUS-y-Connect_Negro-Carbon.png', 5, 10, '4x2, manual, negro'),
(107, 'Mini Cooper', '35200.0000', 'https://www.autodato.com/wp-content/uploads/2014/12/MINI-Cooper-5-Puertas-2.png', 12, 10, '5 puertas, automatico, celeste con blanco'),
(108, 'Mini Cooper', '35000.0000', 'https://www.autodato.com/wp-content/uploads/2015/05/MINI-Cooper-F56-2.png', 21, 10, '3 puertas, automatico, rojo'),
(109, 'Mini Cooper', '35250.0000', 'https://www.pngplay.com/wp-content/uploads/13/Mini-Cooper-S-No-Background.png', 5, 10, '3 puertas, manual, mostaza'),
(110, 'Peugeot 208 Feline', '24159.0000', 'https://www.martinautos.com.ar/web/img/files/versionesautos/photo/009c67d5-7ffa-4e13-9251-ecb1243d45d7/Nuevo-208-AGSM-transp%20%281%29.png', 8, 10, 'Full, automatico, azul'),
(111, 'Peugeot 208 Feline', '20250.0000', 'https://www.peugeotindiana.com.ar/assets/images/frente-gris-966x500.png', 21, 10, 'Base, manual, gris'),
(112, 'Peugeot 208 Feline', '25500.0000', 'https://store.peugeot.cl/digital/nuevo208/images/AMARILLO.PNG', 17, 10, 'Full, automatico, amarillo'),
(113, 'Peugeot 208 Feline', '22750.0000', 'https://store.peugeot.cl/digital/nuevo208/images/ROJO_ELIXIR.PNG', 27, 10, 'Full, manual, rojo'),
(114, 'El Principito', '10.9800', 'https://librosetiqueta.com/wp-content/uploads/2022/06/El_Principito.png', 20, 11, 'Editorial Salamandra - 1943 - Saint Exupéry'),
(115, 'Harry Potter y la piedra filosofal', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/harry-1.png', 20, 11, 'Editorial Salamandra - 1997 - J.K. Rowling - 1er Libro'),
(116, 'Harry Potter y la cámara secreta', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/harry-2.png', 20, 11, 'Editorial Salamandra - 1998 - J.K. Rowling - 2do Libro'),
(117, 'Harry Potter y el prisionero de Azkaban', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/harry-3.png', 20, 11, 'Editorial Salamandra - 1999 - J.K. Rowling - 3er Libro'),
(118, 'Harry Potter y el cáliz de fuego', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/HARRY-4-EDIT-2.png', 20, 11, 'Editorial Salamandra - 2000 - J.K. Rowling - 4to Libro'),
(119, 'Harry Potter y la Orden del Fénix', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/harry-5.png', 20, 11, 'Editorial Salamandra - 2003 - J.K. Rowling - 5to Libro'),
(120, 'Harry Potter y el misterio del príncipe', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/harry-6.png', 20, 11, 'Editorial Salamandra - 2005 - J.K. Rowling - 6to Libro'),
(121, 'Harry Potter y las reliquias de la muerte', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/harry-7.png', 20, 11, 'Editorial Salamandra - 2007 - J.K. Rowling - 7mo Libro'),
(122, 'Harry Potter y el legado maldito', '20.9900', 'https://librosetiqueta.com/wp-content/uploads/2022/05/harry-8-1.png', 20, 11, 'Editorial Salamandra - 2016 - J.K. Rowling - 8vo Libro'),
(123, 'Rusia', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/1622.jpg', 15, 11, 'Roca Editorial - 1991 - Edward Rutherfurd'),
(124, 'China', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/5537.jpg', 15, 11, 'Roca Editorial - 2021 - Edward Rutherfurd'),
(125, 'Londres', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/1540.jpg', 15, 11, 'Roca Editorial - 1997 - Edward Rutherfurd'),
(126, 'Paris', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/1585.jpg', 15, 11, 'Roca Editorial - 2012 - Edward Rutherfurd'),
(127, 'Nueva York', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/1571.jpg', 15, 11, 'Roca Editorial - 2009 - Edward Rutherfurd'),
(128, 'Principes de Irlanda', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/1113.jpg', 15, 11, 'Roca Editorial - 2004 - Edward Rutherfurd'),
(129, 'Rebeldes de Irlanda', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/1114.jpg', 15, 11, 'Roca Editorial - 2006 - Edward Rutherfurd'),
(130, 'El Bosque', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/2549.jpg', 15, 11, 'Roca Editorial - 2000 - Edward Rutherfurd'),
(131, 'Sarum', '24.9000', 'https://www.rocalibros.com/archivos/imagenes/mayores/1713.jpg', 15, 11, 'Roca Editorial - 1987 - Edward Rutherfurd'),
(132, 'Chevrolet Nuevo Onix', '25000.0000', 'https://assets.static-gm.com/Assets/4094fc05-4ebc-4a4e-8773-5114c565dc62/34d98625-0fad-4589-a74a-d888b0c75fb8/v-1641945314/Desktop.webp', 13, 10, 'Full, automatico, rojo'),
(133, 'Chevrolet Nuevo Onix', '22750.0000', 'https://www.chevrolet.com.ar/content/dam/chevrolet/mercosur/argentina/espanol/index/cars/2020-onix-premier/colorizer/01-images/julio-20/2-onix-premier-seeker-met-5.png?imwidth=960', 22, 10, 'Base, manual, azul'),
(134, 'Chevrolet Nuevo Onix', '24500.0000', 'https://secure-developments.com/commonwealth/argentina/gm_forms/assets/front/images/jellys/6137a6206c5a8.png', 11, 10, 'Full, manual, gris'),
(135, 'Chevrolet Nuevo Onix', '25000.0000', 'https://www.chevrolet.com.ar/content/dam/chevrolet/mercosur/argentina/espanol/index/cars/2020-onix-premier/colorizer/01-images/julio-20/4-onix-premier-black-meet-kettle-met.png?imwidth=960', 15, 10, 'Full, automatico, negro'),
(136, 'Toyota Yaris', '25000.0000', 'https://media.a24.com/p/ef4ae0b2872aabef8e546a2ccc6e0103/adjuntos/296/imagenes/008/019/0008019945/1200x1200/smart/el-toyota-yaris-nuevas-versiones-caja-automatica-cvt.png', 13, 10, 'Full, automatico, rojo'),
(137, 'Toyota Yaris', '22750.0000', 'https://www.toyotafederico.com/images/yaris/c_h_5-2022.png', 22, 10, 'Base, manual, azul'),
(138, 'Toyota Yaris', '24500.0000', 'https://www.toyotafederico.com/images/yaris/c_h_4-2022.png', 11, 10, 'Full, manual, gris'),
(139, 'Toyota Yaris', '25000.0000', 'https://www.toyotapanamericana.com/uploads/vehiculos/nuevos/colores/color_20-05-2019_Toyota_Yaris_Blanco.png', 15, 10, 'Full, automatico, blanco'),
(140, 'Nissan Kicks', '28000.0000', 'https://iqcarsgermany.azureedge.net/img/BrandNewCarThumb/1628430587555.22392021-Kicks.png', 13, 10, 'Full, automatico, rojo'),
(141, 'Nissan Kicks', '27750.0000', 'https://nissancapital.com/wp-content/uploads/2021/12/Catalogo_Nuevo_Nissan_Kicks_21-AZUL-BITONO.png', 22, 10, 'Base, manual, azul'),
(142, 'Nissan Kicks', '26500.0000', 'https://nissancapital.com/wp-content/uploads/2021/12/Catalogo_Nuevo_Nissan_Kicks_21-GRI-PLATA.png', 11, 10, 'Full, manual, gris'),
(143, 'Nissan Kicks', '28000.0000', 'https://invygo-car-images.s3.ap-south-1.amazonaws.com/c5451552-9d54-4cc4-9833-e81f9bed5552_Nissan%20-%20Kicks-white.webp', 15, 10, 'Full, automatico, blanco'),
(144, 'Volkswagen T-Cross', '35000.0000', 'https://www.perumotor.com.pe/wp-content/uploads/2021/03/6-19.png', 13, 10, 'Full, automatico, bronce'),
(145, 'Volkswagen T-Cross', '34750.0000', 'https://secplanvw.com.ar/wp-content/uploads/2021/08/tcross.png', 22, 10, 'Base, manual, azul'),
(146, 'Volkswagen T-Cross', '33500.0000', 'https://baronstokai.co.za/media/3681/barons-t-cross-highline-range-image.png', 11, 10, 'Full, manual, naranja'),
(147, 'Volkswagen T-Cross', '37000.0000', 'https://immagini.alvolante.it/sites/default/files/styles/image_gallery_big/public/serie_auto_galleria/2019/02/volkswagen_t-cross_ant.png?itok=hzrjSu_B', 15, 10, 'Full, automatico, celeste'),
(148, 'Camiseta del Real Madrid', '110.4800', 'https://cf.camisetasfutbol.com.cn/upload/ttmall/img/20210806/d1ac1254d993df54a145783e59a0bf6e.png', 35, 2, 'Remera oficial de Real Madrid - 10'),
(149, 'Camiseta de Barcelona', '123.2500', 'https://static.nike.com/a/images/t_default/3712faa8-d7c6-4078-ac06-adff9fe6508d/primera-equipacion-stadium-fc-barcelona-2022-23-camiseta-de-futbol-dri-fit-Hn9LCv.png', 28, 2, 'Remera oficial del Barcelona - 10'),
(150, 'Camiseta de Manchester City', '122.7800', 'https://shop.mancity.com/dw/image/v2/BDWJ_PRD/on/demandware.static/-/Sites-master-catalog-MAN/default/dwb6c90258/images/large/701221537002_pp_01_mcfc.png?sw=1600&sh=1600&sm=fit', 30, 2, 'Remera oficial del Manchester City - 9'),
(151, 'Camiseta del Atletico del Madrid', '111.5400', 'https://cf.camisetasfutbol.com.cn/upload/ttmall/img/20210730/c2396ce2b3f726270a7c969aa005a677.png', 28, 2, 'Remera oficial del Atletico del Madrid - 9'),
(152, 'Camiseta de Instituto', '125.7800', 'https://encrypted-tbn2.gstatic.com/shopping?q=tbn:ANd9GcTWkkhc9c3gtF3gMYP37M5U-7DhgQtrP9rLVXnBOtYOYeHwrfxNVl6IkPHObDeMOELGWx4JRI5H0jsacFQ5FZhRNU7Tn1dJUOl3cZ7d2rm88PiCnd3GuqNA&usqp=CAE', 14, 2, 'Remera oficial de Instituto - 10'),
(153, 'Camiseta de Belgrano', '100.9900', 'https://d3ugyf2ht6aenh.cloudfront.net/stores/001/672/177/products/11-390d1d3db640131a9f16444336011611-480-0.png', 7, 2, 'Remera oficial de Belgrano - 10'),
(154, 'Camiseta de Talleres', '109.9900', 'https://i0.wp.com/tienda.clubtalleres.com.ar/wp-content/uploads/2022/02/titular-frente-copia-2.jpg?fit=1200%2C1200&ssl=1', 19, 2, 'Remera oficial de Talleres - 9'),
(155, 'Camiseta de Estudiantes de Rio Cuarto', '109.9900', 'https://http2.mlstatic.com/D_NQ_NP_717938-MLA47548433058_092021-O.jpg', 19, 2, 'Remera oficial de Estudiantes de Rio Cuarto - 9'),
(156, 'iPad Pro 12.9‑inch (5th generation)', '1099.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-pro-12-select-cell-spacegray-202104?wid=470&hei=556&fmt=p-jpg&qlt=95&.v=1617126613000', 12, 8, 'M1 chip with next-generation Neural Engine, color Space Gray, 256gb'),
(157, 'iPad Pro 12.9‑inch (5th generation)', '1299.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-pro-12-select-cell-silver-202104?wid=470&hei=556&fmt=p-jpg&qlt=95&.v=1617126601000', 15, 8, 'M1 chip with next-generation Neural Engine, color Silver, 512gb'),
(158, 'iPad Air (5th generation)', '825.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-air-select-wifi-spacegray-202203?wid=470&hei=556&fmt=png-alpha&.v=1645066742664', 10, 8, 'M1 chip with next-generation Neural Engine, color Space Gray, 256gb'),
(159, 'iPad Air (5th generation)', '678.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-air-select-wifi-pink-202203?wid=470&hei=556&fmt=png-alpha&.v=1645066399526', 11, 8, 'M1 chip with next-generation Neural Engine, color Pink, 64gb'),
(160, 'iPad Air (5th generation)', '700.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-air-select-wifi-purple-202203?wid=470&hei=556&fmt=png-alpha&.v=1645066730601', 8, 8, 'M1 chip with next-generation Neural Engine, color Purple, 128gb'),
(161, 'iPad Air (5th generation)', '850.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-air-select-wifi-blue-202203?wid=470&hei=556&fmt=png-alpha&.v=1645065732688', 16, 8, 'M1 chip with next-generation Neural Engine, color Blue, 256gb'),
(162, 'iPad Air (5th generation)', '700.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-air-select-wifi-starlight-202203?wid=470&hei=556&fmt=png-alpha&.v=1645895139236', 9, 8, 'M1 chip with next-generation Neural Engine, color Starlight, 128gb'),
(163, 'iPad Mini (6th generation)', '450.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-mini-select-wifi-space-gray-202109?wid=470&hei=556&fmt=png-alpha&.v=1629840743000', 15, 8, 'A15 Bionic chip with Neural Engine, color Space Gray, 64gb'),
(164, 'iPad Mini (6th generation)', '500.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-mini-select-wifi-pink-202109?wid=470&hei=556&fmt=png-alpha&.v=1629840714000', 25, 8, 'A15 Bionic chip with Neural Engine, color Pink, 256gb'),
(165, 'iPad Mini (6th generation)', '620.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-mini-select-wifi-purple-202109?wid=470&hei=556&fmt=png-alpha&.v=1629840735000', 33, 8, 'A15 Bionic chip with Neural Engine, color Purple, 256gb'),
(166, 'iPad Mini (6th generation)', '480.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/ipad-mini-select-wifi-starlight-202109?wid=470&hei=556&fmt=png-alpha&.v=1629840745000', 21, 8, 'A15 Bionic chip with Neural Engine, color Starlight, 128gb'),
(167, 'iMac 24 pulgadas', '1700.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/imac-24-blue-selection-hero-202104?wid=452&hei=420&fmt=jpeg&qlt=95&.v=1617492405000', 21, 8, 'Apple M1 Chip, 8-Core CPU, 8-Core GPU, color Blue, 1T'),
(168, 'iMac 24 pulgadas', '2100.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/imac-24-green-selection-hero-202104?wid=452&hei=420&fmt=jpeg&qlt=95&.v=1617492405000', 21, 8, 'Apple M1 Chip, 8-Core CPU, 8-Core GPU, color Green, 2T'),
(169, 'iMac 24 pulgadas', '1300.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/imac-24-pink-selection-hero-202104?wid=452&hei=420&fmt=jpeg&qlt=95&.v=1617492407000', 21, 8, 'Apple M1 Chip, 8-Core CPU, 8-Core GPU, color Pink, 512gb'),
(170, 'iMac 24 pulgadas', '1099.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/imac-24-purple-selection-hero-202104?wid=452&hei=420&fmt=jpeg&qlt=95&.v=1617492406000', 21, 8, 'Apple M1 Chip, 8-Core CPU, 8-Core GPU, color Purple, 256gb'),
(171, 'iMac 24 pulgadas', '2100.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/imac-24-silver-selection-hero-202104?wid=452&hei=420&fmt=jpeg&qlt=95&.v=1617492408000', 21, 8, 'Apple M1 Chip, 8-Core CPU, 8-Core GPU, color Silver, 2T'),
(172, 'iMac 24 pulgadas', '1700.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/imac-24-yellow-selection-hero-202104?wid=452&hei=420&fmt=jpeg&qlt=95&.v=1617492407000', 21, 8, 'Apple M1 Chip, 8-Core CPU, 8-Core GPU, color Yellow, 1T'),
(173, 'iMac 24 pulgadas', '1300.0000', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/imac-24-orange-selection-hero-202104?wid=452&hei=420&fmt=jpeg&qlt=95&.v=1617492406000', 21, 8, 'Apple M1 Chip, 8-Core CPU, 8-Core GPU, color Orange, 512gb'),
(174, 'HomePod mini', '99.9900', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/homepod-mini-select-spacegray-202110_FV1?wid=934&hei=440&fmt=jpeg&qlt=95&.v=1633086020000', 12, 7, 'Sonido que llena la sala. El asistente perfecto. Control de toda tu casa. Se lleva de maravilla con el iPhone. Privacidad y seguridad, color Space Gray'),
(175, 'HomePod mini', '99.9900', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/homepod-mini-select-blue-202110_FV1?wid=934&hei=440&fmt=jpeg&qlt=95&.v=1633086025000', 18, 7, 'Sonido que llena la sala. El asistente perfecto. Control de toda tu casa. Se lleva de maravilla con el iPhone. Privacidad y seguridad, color Blue'),
(176, 'HomePod mini', '99.9900', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/homepod-mini-select-white-202110_FV1?wid=934&hei=440&fmt=jpeg&qlt=95&.v=1633086025000', 25, 7, 'Sonido que llena la sala. El asistente perfecto. Control de toda tu casa. Se lleva de maravilla con el iPhone. Privacidad y seguridad, color White'),
(177, 'HomePod mini', '99.9900', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/homepod-mini-select-yellow-202110_FV1?wid=934&hei=440&fmt=jpeg&qlt=95&.v=1633086026000', 5, 7, 'Sonido que llena la sala. El asistente perfecto. Control de toda tu casa. Se lleva de maravilla con el iPhone. Privacidad y seguridad, color Yellow'),
(178, 'HomePod mini', '99.9900', 'https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/homepod-mini-select-orange-202110_FV1?wid=934&hei=440&fmt=jpeg&qlt=95&.v=1633086020000', 9, 7, 'Sonido que llena la sala. El asistente perfecto. Control de toda tu casa. Se lleva de maravilla con el iPhone. Privacidad y seguridad, color Orange'),
(179, 'Cocina Gourmet Inox multigas 4 hornallas puerta con visor 91L', '642.0000', 'https://images.precialo.com/products/cocina-ormay-gourmet-inox-multigas-4-hornallas-negra-mate-puerta-con-visor-91l/834ab085-5bdd-4011-8ff7-bf97d0eda73b.jpeg', 15, 6, 'negra mate - Ormay'),
(180, 'Cocina Gourmet Inox multigas 4 hornallas puerta con visor 91L', '650.0000', 'https://http2.mlstatic.com/D_NQ_NP_626181-MLA50161119835_062022-O.jpg', 12, 6, 'gris acero - Ormay');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(350) NOT NULL,
  `last_name` varchar(250) NOT NULL,
  `user_name` varchar(150) NOT NULL,
  `password` varchar(150) NOT NULL,
  `address` varchar(350) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `users`:
--

--
-- Volcado de datos para la tabla `users`
--

INSERT INTO `users` VALUES
(1, 'Nicolas', 'Rojo', 'NicoRed', '67a24c14b19206fec22a318c8b6e7616', 'Italia 2139 Villa Cabrera, Cordoba'),
(2, 'Octavio', 'Aquilino', 'TayiAqui', '699b79976a6310009b7ec649fd654a16', 'Santiago del Estero 288 Centro, La Rioja'),
(3, 'Santiago', 'Riveros', 'SantiRiv', '0e13e726197867bd8cfc7265c2e89ad8', 'Chacabuco 53 Centro, Cordoba');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_order_us` (`id_user`);

--
-- Indices de la tabla `order_details`
--
ALTER TABLE `order_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_order_det_prod` (`id_product`),
  ADD KEY `fk_order_det_ord` (`id_order`);

--
-- Indices de la tabla `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_prod_cat` (`id_category`);

--
-- Indices de la tabla `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_name` (`user_name`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT de la tabla `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `order_details`
--
ALTER TABLE `order_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=283;

--
-- AUTO_INCREMENT de la tabla `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `fk_order_us` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`);

--
-- Filtros para la tabla `order_details`
--
ALTER TABLE `order_details`
  ADD CONSTRAINT `fk_order_det_ord` FOREIGN KEY (`id_order`) REFERENCES `orders` (`id`),
  ADD CONSTRAINT `fk_order_det_prod` FOREIGN KEY (`id_product`) REFERENCES `products` (`id`);

--
-- Filtros para la tabla `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `fk_prod_cat` FOREIGN KEY (`id_category`) REFERENCES `categories` (`id`);
