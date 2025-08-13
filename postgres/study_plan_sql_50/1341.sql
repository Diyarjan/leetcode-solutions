-- 1341. Movie Rating

-- The first method
with max_user as (select m.user_id, u.name as results, count(m.rating) as cnt
                  from MovieRating m
                           join users u on m.user_id = u.user_id
                  group by m.user_id, u.name
                  order by cnt desc, u.name
                  limit 1),
     max_movie as (select r.movie_id,
                          m.title       as results,
                          AVG(r.rating) as rat
                   from MovieRating r
                            join Movies m on m.movie_id = r.movie_id
                   where r.created_at >= '2020-02-01'::date
                     and r.created_at < '2020-03-01'::date
                   group by r.movie_id, m.title
                   order by rat desc, m.title
                   limit 1)

select results
from max_user
union all
select results
from max_movie;


-- The second method
(select u.name as results
 from MovieRating m
          join users u on m.user_id = u.user_id
 group by m.user_id, u.name
 order by count(m.rating) desc, u.name
 limit 1)

union all

(select m.title as results
 from MovieRating r
          join Movies m on m.movie_id = r.movie_id
 where r.created_at >= '2020-02-01'::date and r.created_at < '2020-03-01'::date
 group by r.movie_id, m.title
 order by AVG(r.rating) desc, m.title
 limit 1);