select distinct id, title from word
join synonym s on word.id = s.word_1_id or word.id = s.word_2_id 
where (s.word_2_id = ($1) or s.word_1_id = ($1)) and word.id != ($1);