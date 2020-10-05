-- migrate:up
DROP TABLE IF EXISTS `checkins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `checkins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uuid` varchar(36) DEFAULT NULL,
  `business_uuid` varchar(36) DEFAULT NULL,
  `business_name` varchar(50) DEFAULT NULL,
  `business_address` varchar(300) DEFAULT NULL,
  `chckr_uuid` varchar(36) DEFAULT NULL,
  `chckr_name` varchar(500) DEFAULT NULL,
  `chckr_phone` varchar(100) DEFAULT NULL,
  `chckr_email` varchar(255) DEFAULT NULL,
  `chckr_street` varchar(500) DEFAULT NULL,
  `chckr_city` varchar(100) DEFAULT NULL,
  `chckr_country` varchar(100) DEFAULT NULL,
  `chckr_registered` boolean DEFAULT false,
  `paper` boolean DEFAULT false,
  `checked_in_at` datetime DEFAULT NULL,
  `checked_out_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_checkins_uuid` (`uuid`),
  KEY `idx_checkins_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

-- migrate:down
DROP TABLE IF EXISTS `checkins`;
