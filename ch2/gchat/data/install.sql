drop database if exists chat;
create database chat;
drop user if exists chat;
create user chat with password 'chat'
    superuser
    createdb
    createrole
    replication
    bypassrls;
grant all privileges on database chat to chat;
