# Monkeytype Frontend

This is the frontend application for the Monkeytype clone, built using React and TypeScript. It interacts with the Go backend to provide a typing test experience.

## Project Structure

- **public/**: Contains static files.
  - **index.html**: The main HTML file that serves as the entry point for the application.
  - **favicon.ico**: The favicon for the application.

- **src/**: Contains the source code for the application.
  - **components/**: Contains reusable React components.
    - **TypingTest.tsx**: The component that handles the typing test functionality.
  - **pages/**: Contains the main pages of the application.
    - **HomePage.tsx**: The main page that renders the TypingTest component.
  - **services/**: Contains API service functions.
    - **api.ts**: Functions for making API calls to the Go backend.
  - **App.tsx**: The main App component that sets up routing and structure.
  - **index.tsx**: The entry point for the React application.
  - **styles/**: Contains CSS styles.
    - **App.css**: Global styles for the application.

## Getting Started

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd monkeytype-frontend
   ```

2. **Install dependencies**:
   ```
   npm install
   ```

3. **Run the application**:
   ```
   npm start
   ```

4. **Open your browser** and navigate to `http://localhost:3000` to view the application.

## Usage

The application provides a typing test interface where users can practice their typing skills. The TypingTest component manages the test functionality and communicates with the backend to handle typing requests.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or features you'd like to add.

## License

This project is licensed under the MIT License. See the LICENSE file for details.