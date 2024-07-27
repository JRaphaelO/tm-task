# Task Service

Task service is microsservice with responsability the management for task. This service, is called by a RESTFULL api.

## Task Fields
### Task Entity
- id: 
    - Unique identify for task
    - Type uuid v4
    - Primary key
- name:
    - Text to identify task
    - Type text
    - Max 120 characters
- description:
    - Text to describe task
    - Type text
    - Default `NULL`
- status:
    - Progress status by a task
    - Type ENUM <CREATED, STOPED, WORKING, DONE >
    - Default `CREATED`
- prevision_date:
    - Date for prevision finished task
    - Type timestamp
    - Defaul `NULL`
- finished_date:
    - Date for finished task
    - Type timestamp
    - Default `NULL`
- created_at:
    - Date for created task
    - Type timestamp
    - Defaul `NOW`
- updated_at:
    - Date for update task
    - Type timestamp
    - Defaul `NOW`


## Task Requirements

- Task is field for include presion date for coclusion
- Task when created is status equals CREATED
- Task name is not empty and 120 characters max
