SHOW DATABASES;
SHOW CREATE TABLE todolist;
use gorestfulapi_exercise;
select * from todolist;
desc todolist;

rename table `todolistcontent` to `todolist`;
alter table `todolist` rename column `namatodolist` to `todolisttitle`;
alter table `todolist` add column `todolistcontent` text after `todolisttitle`;
alter table `todolist` add column `todolisttabledate` DATETIME after `id`;
alter table `todolist` add column `todolistcontentchecked` bool after `todolistcontent`;
alter table `todolist` rename column `todolistcontentchecked` to `checked`;
alter table `todolist` add column `todolistsubcontent` text after `todolistcontent`;
alter table `todolist` MODIFY column `checked` boolean; 
alter table `todolist` modify column `datetime` timestamp not null;

delete from todolist where id="";
insert into `todolist` (`todolisttitle`, `todolistcontent`, `todolistsubcontent`, `checked`)
values
('to do list 1', 'contoh isi main to do list 1', 'sub kategori to do list', true);
