# Answer

### no.1 
```
SELECT employee_id, name, job_title, salary, department, joined_date 
FROM itb_db.employees;

employee_id|name       |job_title|salary  |department|joined_date|
-----------+-----------+---------+--------+----------+-----------+
          1|Jhon Smith |Manager  |60000.00|Sales     | 2022-01-15|
          2|Jane Doe   |Analyst  |45000.00|Marketing | 2022-02-01|
          3|Mike Brown |Developer|55000.00|IT        | 2022-03-10|
          4|Anna Lee   |Manager  |65000.00|Sales     | 2021-12-05|
          5|Mark Wong  |Developer|50000.00|IT        | 2023-05-20|
          6|Emilly Chen|Analyst  |48000.00|Marketing | 2023-06-02|
```

### no.2. 
```
SELECT COUNT(*)
FROM itb_db.employees WHERE job_title = 'Manager';

COUNT(*)|
--------+
       2|
```

### no.3.
```
SELECT name, salary
FROM itb_db.employees WHERE department  = 'Sales' OR department = 'Marketing';

name       |salary  |
-----------+--------+
Jhon Smith |60000.00|
Jane Doe   |45000.00|
Anna Lee   |65000.00|
Emilly Chen|48000.00|
```

### no.4
```
SELECT AVG(salary) AS salary_average
FROM itb_db.employees
WHERE joined_date >= DATE_ADD(CURRENT_DATE(), INTERVAL -5 YEAR)

salary_average|
--------------+
  53833.333333|
```

### no.5
```
SELECT e.* , SUM(sd.sales) AS total_sales FROM employees e 
INNER JOIN sales_data sd ON sd.employee_id  = e.employee_id 
GROUP BY sd.employee_id 
ORDER BY total_sales DESC
LIMIT 5

employee_id|name      |job_title|salary  |department|joined_date|total_sales|
-----------+----------+---------+--------+----------+-----------+-----------+
          1|Jhon Smith|Manager  |60000.00|Sales     | 2022-01-15|   35000.00|
          2|Jane Doe  |Analyst  |45000.00|Marketing | 2022-02-01|   26000.00|
          4|Anna Lee  |Manager  |65000.00|Sales     | 2021-12-05|   22000.00|
          5|Mark Wong |Developer|50000.00|IT        | 2023-05-20|   19000.00|
          3|Mike Brown|Developer|55000.00|IT        | 2022-03-10|   18000.00|
```


### no.6
```
SELECT name, salary FROM  employees 
WHERE salary > (
	SELECT AVG(salary) FROM  employees 
	GROUP BY department 
	ORDER BY salary DESC
	LIMIT 1	
)

name    |salary  |
--------+--------+
Anna Lee|65000.00|
```



### no.7
```
SELECT  x.name, x.salary, x.total_sales, 
CASE
    WHEN total_sales > 30000 THEN "Ranking 1"
    WHEN total_sales <= 30000 AND total_sales >= 20000 THEN "Ranking 2"
    ELSE "Ranking 3"
END AS ranking
FROM (
	SELECT e.name , e.salary, SUM(sd.sales) AS total_sales 
	FROM  employees e INNER JOIN sales_data sd ON sd.employee_id = e.employee_id 
	GROUP BY sd.employee_id 
	ORDER BY total_sales DESC
) x


name       |salary  |total_sales|ranking  |
-----------+--------+-----------+---------+
Jhon Smith |60000.00|   35000.00|Ranking 1|
Jane Doe   |45000.00|   26000.00|Ranking 2|
Anna Lee   |65000.00|   22000.00|Ranking 2|
Mark Wong  |50000.00|   19000.00|Ranking 3|
Mike Brown |55000.00|   18000.00|Ranking 3|
Emilly Chen|48000.00|   13000.00|Ranking 3|
```


### no.8

```
DELIMITER $$
CREATE PROCEDURE employeelist (department VARCHAR(255))
   BEGIN
     SELECT name, salary FROM employees
     WHERE department = department;
   END$$
DELIMITER ;

CALL employeelist("Sales")

name      |salary  |
----------+--------+
Jhon Smith|60000.00|
Anna Lee  |65000.00|
```