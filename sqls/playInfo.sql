CREATE TABLE `playInfo`
(
    `id`          int NOT NULL AUTO_INCREMENT,
    `challengeId` int NOT NULL,
    `userId`      int NOT NULL,
    `score`       int NOT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;