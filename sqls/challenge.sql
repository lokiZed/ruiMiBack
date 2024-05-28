CREATE TABLE `challenge`
(
    `id`        int NOT NULL AUTO_INCREMENT,
    `gameId`    int NULL DEFAULT NULL,
    `gameLevel` int NULL DEFAULT NULL,
    `entTime`   int NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;