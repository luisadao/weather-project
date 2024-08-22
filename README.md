# Project Overview
This is my first project using GO, it consists on using GO for a simple backend that serves as an API, that uses OpenWeather API.
The frontend is a very basic Vue.js just to consume the API.

# WIP:
- [x] Authentication
- [x] Backend Folder Structure
- [ ] Middleware to the frontend

# Getting it to work:
## Backend:
- Ensure that you have mysql installed on your system
- Use the .env provided for the backend
- Source the create-tables.sql file to your database example: 
    - 1:`mysql -u root -p`
    - 2:`create database weather;`
    - 3:`use weather;`
    - 4:`source path/create-tables.sql`
- Inside the backend directory run `make run` to run the project

## Frontend:
- Inside the frontend directory run `npm install` to install the project and `npm run dev` to run it
