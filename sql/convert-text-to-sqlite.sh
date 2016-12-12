rm -f words2.sqlite
awk  -F \|  'BEGIN {print "create table words(id INTEGER PRIMARY KEY ASC, word varchar(10)); begin transaction;"}
    {if(length($2)<=9  && length($2)>2)print "insert into words(word) values(\""$2"\");"}
    END{print "commit;"}' words.txt | sqlite3 words2.sqlite
