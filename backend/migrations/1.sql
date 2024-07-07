drop table if exists synonym;
drop table if exists word;

create table word
(
  id serial primary key,
  title   varchar (50) not null, 
  constraint word_the_word_uq unique (title)  
);

create table synonym
(
  word_1_id integer not null,
  word_2_id integer not null,
  constraint synonyms_not_match_ck check (word_1_id < word_2_id),
  constraint synonym_pk primary key (word_1_id, word_2_id),
  constraint synonym_w1_fk foreign key (word_1_id) references word (id),
  constraint synonym_w2_fk foreign key (word_2_id) references word (id)
);