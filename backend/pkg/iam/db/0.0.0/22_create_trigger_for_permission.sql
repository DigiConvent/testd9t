create trigger insert_parent_permissions
after insert on permissions
for each row
begin
    with recursive split_permission as (
        select 
            new.permission as full_permission,
            substr(new.permission, 1, instr(new.permission || '.', '.') - 1) as parent_permission,
            substr(new.permission, instr(new.permission || '.', '.') + 1) as remaining_permission
        union all
        select 
            full_permission,
            substr(remaining_permission, 1, instr(remaining_permission || '.', '.') - 1) as parent_permission,
            substr(remaining_permission, instr(remaining_permission || '.', '.') + 1) as remaining_permission
        from split_permission
        where remaining_permission != ''
    )
    insert or ignore into permissions (permission)
    select parent_permission
    from split_permission
    where parent_permission != '';
end;