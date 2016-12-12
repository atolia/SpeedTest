awk -F \| 'BEGIN {print "drop table if exists `words`; create table words(id INTEGER PRIMARY KEY AUTO_INCREMENT, word varchar(10)); begin;"}
    {if(length($2)<=9  && length($2)>2)print "insert into words(word) values(\""$2"\");"}
    END{print "commit;"}' $1 |mysql test
