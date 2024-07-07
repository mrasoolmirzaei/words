with recursive cte as (
  select
    word_2_id as synonym_id
  from
    synonym
  where
    word_2_id = 4
  union all
  select
    s.word_1_id as synonym_id
  from
    cte
    join synonym s on cte.synonym_id = s.word_2_id
)
select id, title from cte
join word on cte.synonym_id = word.id;