SELECT invited_id as id, count(*) as num, max(end_time) as time
FROM `game_pk_record`
WHERE (game_id = '96' and winner_id = invited_id and end_time > '2021-05-17 00:00:00')
GROUP BY invited_id
ORDER BY num desc, time desc LIMIT 20