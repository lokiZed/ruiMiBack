CREATE TABLE `challengeUser`
(
    `id`          int NOT NULL AUTO_INCREMENT,
    `userId`      int NULL DEFAULT NULL,
    `challengeId` int NULL DEFAULT NULL,
    `leaveNum`    int NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;