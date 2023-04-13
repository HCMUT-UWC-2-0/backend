-- name: CreateBackOfficers :one
INSERT INTO
  "BackOfficers" (
    "email",
    "ssn",
    "hashed_password",
    "name",
    "phone",
    "age",
    "gender" ,
    "date_of_birth",
    "place_of_birth"
  )
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;


-- name: GetBackOfficer :one
SELECT * FROM "BackOfficers" 
WHERE email = $1 LIMIT 1;




