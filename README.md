# school-management-system
School-Management-System is a simple REST API written in Go using Docker, Gin and MYSql that allows school administrators to manage some activities and processes of the school.

This api provides:
- Student Management with CRUD endpoints (Enroll student to school, Update student information, list all students, and Disenroll student).
- Courses Management CRUD endpoints (Add new course to database, update course information, list all courses, and delete course).

Other endpoints include:
- List courses taken by a student 
- List students taking a particular course
- Update list of courses taken by a student (add and remove).

## How To Install
- Fork this repository
- Run`git clone https://github.com/ogbenioye/school-management-system.git` in terminal
- Run `docker compose up`
- Connect to MYSQL database

## Running/Testing API
Using postman or any API testing platforms, create all requests in router directory
### Courses
- Create Course: 

    First, add course(s) to the database. Only the courses in the database can be registered to a student. This looks like this in postman:
    ```
    {
      "code": "MTS101",
      "courseTitle": "Algebra"
    }
    ```
- Update Course:

    Attach the course-id(course-code) you want to update in the request url. Course-code cannot be updated/edited because it also serves as the primary key, you can only update the course-title. If there is any need to update course code,       you need to delete and re-add course.
    
- List Courses:
  
  Returns the list of courses from the database.
    
- Delete Course:

    Similar to Updateing, attach course-id that is to be deleted.
### Students
- Enroll Student:

    Input the new student information and register student for one or more courses. For example:
    ```
    {
      "firstName": "John",
      "lastName": "Doe",
      "gender": "Male",
      "dept": "Computer Science",
      "courses": [
        {"code": "MTS106"},
        {"code": "COM202"},
        {"code": "COM103"}
      ]
    }
    ```
    *NB:* Courses must exist in database first to be added.
- Update Student Information:

    Attach valid student-id in the request url. You can only update/edit the first-name, last-name, gender, and department fields 
- List Students:

    Returns all the students enrolled in the school.
- Disenroll Student:

    To disenroll student from school, attach a valid student-id in the request url.
### Other features
- List Student's courses:
    
    Attach a valid student-id to get the list of courses the student registered.
- List Students registered to Course:
  
    Attach a valid course-code(course-id) to get the list of students taking the course.
- Update List of Courses Taken by a Student:

    This is split into two:
      
1. Adding one or more courses to the list of courses a student is taking:
      
      To add more courses to student's course list, attach the student's id to the request url. Only the valid coursecode is to be be sent to the request body. For example:
          
```
{
  "courses": [
    {"code": "MTS102"},
    {"code": "MTS101"},
    {"code": "MTS104"}
  ]
}
```
2. Deleting one or more courses from the list of courses a student is taking:

  To delete, attach the student's id to the request url. Only the valid coursecode is to be be sent to the request body. For example:
```
{
  "courses": [
    {"code": "MTS102"}
   ]
}
```
This will delete the course from the joint table in the database.
