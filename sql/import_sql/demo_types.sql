--
-- Database: `playground`
--
-- --------------------------------------------------------
--
-- Table structure for table `demo_types`
--
CREATE TABLE IF NOT EXISTS `demo_types` (
  `tinyint` TINYINT(10) NOT NULL AUTO_INCREMENT,
  `updated` TIMESTAMP,
  `created` TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE = MyISAM DEFAULT CHARSET = latin1 AUTO_INCREMENT = 10001;
--
-- Dumping data for table `user_details`
--
INSERT INTO `demo_types` (`tinyint`)
VALUES (1),
  (2),
  (3);