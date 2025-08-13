-- 185. Department Top Three Salaries

with result as (
    select d.name as Department,
       e.name as Employee,
       e.salary as Salary,
       dense_rank() OVER (partition by e.departmentId order by e.salary desc) as rat
    from Employee e
    join Department d on e.departmentId = d.id)

select Department, Employee, Salary from result
where rat < 4
order by Department, Salary desc;