with recursive cte as (
  select
    word_2_id as synonym_id
  from
    synonym
  where
    word_1_id = 1
  union all
  select
    s.word_2_id as synonym_id
  from
    cte
    join synonym s on cte.synonym_id = s.word_1_id
)
select id, title from cte
join word on cte.synonym_id = word.id;