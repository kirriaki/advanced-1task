# Medical Appointment

## Project Description
"Medical Appointment" is a web application that allows users to book appointments with doctors. The main goal of the project is to simplify the scheduling and management of medical consultations for both patients and healthcare providers. The application targets patients, doctors, and clinic administrators.

## Team Members
- Kamila Nurlybayeva
- Kamila Mukhitdinova
- Aruzhan Zhuanysh

## Homepage Screenshot
![Homepage](./screenshots/homepage.png)

## How to Run the Project
### 1. Clone the Repository
```bash
git clone https://github.com/your-team/medical-appointment.git
cd medical-appointment
```

### 2. Install Dependencies
#### Server (Backend)
Navigate to the server directory and run:
```bash
cd server
npm install
```

#### Client (Frontend)
Navigate to the client directory and run:
```bash
cd client
npm install
```

### 3. Run the Project
#### Start the Server
```bash
npm start
```
The server will run at `http://localhost:5000`.

#### Start the Client
In another terminal window, run:
```bash
npm start
```
The client will run at `http://localhost:3000`.

### 4. Database Setup
Set up a MongoDB database and configure the `.env` file in the `server` directory with the appropriate connection settings:
```
DB_URI=mongodb://localhost:27017/medical_appointment
```

Run the necessary database initialization scripts:
```bash
npm run init-db
```

### 5. Open the Application
Open your browser and go to: `http://localhost:3000`.

## Tools and Resources
- **Frontend**: CSS, HTML, Bootstrap.
- **Backend**: Node.js, Express.
- **Database**: MongoDB.
- **Development Tools**: VS Code, Git

## Repository Checklist
- Code: All source code is included in the repository.
- Screenshots: The `screenshots` folder contains interface images.
- Public Access: The repository is open for viewing [here](https://github.com/your-team/medical-appointment).
