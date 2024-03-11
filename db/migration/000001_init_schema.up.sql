
CREATE TABLE teachers (
    teacher varchar,
    student varchar,
    PRIMARY KEY (teacher, student) 
);

CREATE TABLE suspensions (
    student varchar PRIMARY KEY 
);