# 统计

# 统计离职人数
select sum(if(leave_time>0,1,0)) leave_nums from table_name;

# 统计当前月离职人数
select sum(if(from_unixtime(leave_time,'%c')=$month,1,0)) leave_nums from table_name;