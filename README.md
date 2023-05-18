# Z-B-Validator
Repository for the zerobounce email validator service. Invention and beyond

## ASSUMPTIONS
This are assumptions made during testing of zbv some might be changed or made configurable for a runtime server.
1. Mysql default password is root
2. Project runs on port 3000
3. User/admin has no database known as zbv
4. No default admin set, use cli to create your aown administrator account. Creators assume only the admin has access to cli binary.
*implementing cli in a few :)*
5. A well configured golang installation
6. Mysql maria db installed correctly.
7. Maria db is running locally at 127.0.0.1 and not remotely.

## Project Structure
db-zbv.sql Data description of all db entities. *import it into your mysql db*
All code files are in the zbv.
  zbv/assets
    Static files for loading the server admin panel UI.
  zbv/templates
    Contains template pages gor the server admin UI
  Go Files
    api.go Function definitions for creating the api key
    helpers.go Helper functions to assist various functionalities
    reqlog.go A logger for logging all incoming requests

## DEPLOY
*For a clean install/runtime, import your db-zbv.sql into your maria db server before running*
As per your operating system, cd into the zbv directory /my_project_directory/sbv and go build to create a binary specific to your operating system.
> go build -ldflags="-s -w"  --race  .

You can then proceed to run the binary, *by default it runs on port 3000* future implementation will be on port 80 and 443 for ssl.

## Create a custom directory
To create a working directory for a production server:
1. Create a working dir (dir)
2. Copy assets to dir
3. Copy templates to dir
4. Create .data inside dir
5. Run your server
