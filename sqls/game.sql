CREATE TABLE `game`
(
    `id`           int                                                           NOT NULL AUTO_INCREMENT,
    `gameId`       int                                                           NOT NULL,
    `gameName`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `gameMaxLevel` int NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;