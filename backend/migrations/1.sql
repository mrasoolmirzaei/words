create table word
(
  id integer not null AUTO_INCREMENT,
  title   varchar (50) not null, 
  constraint word_pk primary key (id),
  constraint word_the_word_uq unique (title)  
);

create table synonym
(
  word_1 integer not null,
  word_2 integer not null,
  constraint synonyms_not_match_ck check (word_1 < word_2),
  constraint synonym_pk primary key (word_1, word_2),
  constraint synonym_w1_fk foreign key (word_1) references word (id),
  constraint synonym_w2_fk foreign key (word_2) references word (id)
);