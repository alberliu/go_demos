SELECT invited_id as id, count(*) as num, max(end_time) as time
FROM `game_pk_record`
WHERE (game_id = 1 and winner_id = invited_id and end_time > '2021-05-31 00:00:00' and end_time < '2021-06-07 00:00:00')
GROUP BY invited_id
ORDER BY num desc, time, id LIMIT 50


SELECT room_id as id, count(*) as num, max(end_time) as time
FROM `game_pk_record`
WHERE (game_id = 1 and winner_id != 0 and end_time > '2021-05-31 00:00:00' and end_time < '2021-06-07 00:00:00')
GROUP BY room_id
ORDER BY num desc, time, id LIMIT 50


火影忍者究极风暴4 96
奥特曼格斗进化3 1384
奥特曼格斗进化重生 1446
WWE2K19	 1457


菜鸡车神擂台周榜：极限竞速:地平线4、极限竞速:地平线4(全解锁)、极品飞车18、极速骑行3
(497,1517,69,1642)

菜鸡格斗擂台周榜：
假面骑士:超巅峰英雄、假面骑士斗骑大战、龙珠斗士Z全DLC、龙珠Z:卡卡罗特(全DLC)、践踏2:超级机甲联盟、JOJO全明星大乱斗、Jump大乱斗、拳皇13、地狱剑斗
(1444,1510,1562,1603,1615,1587,183,467,1625,1647)

菜鸡综合擂台周榜：
双人成行、全面战争模拟器、逃出生天、茶杯头、人类一败涂地
(1569,182,1235,119,149)

菜鸡体育擂台周榜：
NBA2K21、FIFA19、NBA2K20、NBA2K19、WWE2K竞技场、实况足球2021
(1219,307,464,196,1242,1668)

菜鸡末日生存擂台周榜：往日不再、生化危机6、生化危机7、生化危机8、消逝的光芒
(1570,159,158,1500,181)
