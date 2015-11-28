drop database gwp;
create database gwp;
drop user gwp;
create user gwp with password 'gwp';
grant all privileges on database gwp to gwp;