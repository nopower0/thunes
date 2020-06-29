# Migration File

## Basics
* All migration files should follow this format: `<model_name>_<serial_no>__<forward|backward>.sql`
  * `serial_no` can only go up, starting from 0001
  * Every `serial_no` should have both forward and backward migration files
* All model changes in a commit should have corresponding migration files
* Migration files **do not** allow changes. You can only add new files to modify table schema to match models.


## Setup DB
* An Empty Schema  
  In this case, you just need to execute all forward migrations for each model from serial_no 0001
* Incremental Migration  
  In this case, I suppose you're using `git pull` to get the latest version.
  You have to find out those forward migration files which are not applied to your DB and execute them.


## Rollback Migrations  
If you want to rollback migrations for a model, you need to find out the last forward migration file.
Then, execute the corresponding backward migration file.
