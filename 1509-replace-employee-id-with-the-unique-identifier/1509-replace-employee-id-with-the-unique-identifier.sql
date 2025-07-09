# Write your MySQL query statement below
SELECT euni.unique_id,e.name
FROM Employees e 
LEFT JOIN employeeUNI euni
ON e.id = euni.id