# Task Manager API

# Necessary Environment Variables

```
// Expects your postgres connection string.
DATABASE_URL
```

# Required DB Config

```
 CREATE TABLE public.tasks (    
     id uuid  NOT NULL,         
     title text  NULL,          
     description text  NULL,    
     status boolean  NULL);
```

After the table is created you should be able to use the API to create new tasks.
