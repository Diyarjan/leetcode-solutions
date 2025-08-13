-- Table: Activity

-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | player_id    | int     |
-- | device_id    | int     |
-- | event_date   | date    |
-- | games_played | int     |
-- +--------------+---------+
-- (player_id, event_date) is the primary key (combination of columns with unique values) of this table.
-- This table shows the activity of players of some games.
-- Each row is a record of a player who logged in and played a number of games (possibly 0) before logging out on someday using some device.

-- Write a solution to report the fraction of players that logged in again on the day after the day they first logged in, rounded to 2 decimal places. 
-- In other words, you need to determine the number of players who logged in on the day immediately following their initial login, and divide it by the number of total players.

with result as (select player_id,
                       event_date,
                       min(event_date) over (partition by player_id) as first_date
                from Activity
                order by player_id, event_date)

select round(sum(case when event_date - first_date = 1 then 1.0 else 0.0 end) / count(distinct (player_id)), 2) as fraction
from result;