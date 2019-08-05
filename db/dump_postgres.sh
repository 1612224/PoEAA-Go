pg_dump --file=./$1.sql --no-owner --verbose -h localhost -p 5432 -U postgres -W $1
