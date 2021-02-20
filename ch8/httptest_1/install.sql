drop database if exists gwp;
create database gwp;
drop user if exists gwp;
create user gwp with password 'gwp'
    superuser
    createdb
    createrole
    replication
    bypassrls;
grant all privileges on database gwp to gwp;
