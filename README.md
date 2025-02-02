# Personal Book Library Manager

# 📚 Personal Book Library Manager

A full-stack web application for managing personal book collections, built with **Go (Golang) & SQLite3** for the backend and **React, Vite & TailwindCSS** for the frontend.

## 🛠 Tech Stack

### Backend:

- **Golang** – Handles API logic and data management
- **Gin Framework** – Fast and efficient web framework
- **SQLite3** – Lightweight and efficient database

### Frontend:

- **React** – Modern UI development
- **Vite** – Fast build tool for frontend development
- **TailwindCSS** – Utility-first CSS framework for styling

---

## 🚀 Features

### **1. Landing Page:**

- Introduction to the app with **sign-up** and **sign-in** options

### **2. Authentication:**

- Sign up and log in functionality (to be implemented later)

### **3. Dashboard:**

- Displays **all books** from **"My Collection"** by default
- Search books in **My Collection** by:
  - **Title**
  - **Author**
  - **ISBN**
- Search books from **OpenLibrary** using:
  - **ISBN**

### **4. Book Management:**

- Add books **manually**
- Search books from **OpenLibrary** and add them to **My Collection**
  - The form will autofill when clicking **"Add to My Collection"**
- Update book details (title, author, isbn, notes, rating, read status)
- Delete books from My Collection

---

## 📌 Prerequisites

- **Go 1.23.3** or higher
- **SQLite3**
- **Node.js & npm** (for frontend)

---

## 🏗 Installation & Setup

### **1️⃣ Clone the Repository**

```sh
git clone https://github.com/Bantamlak12/personal_book_library_manager
cd personal_book_library_manager
```

---

### **2️⃣ Backend Setup (Go + SQLite3)**

#### **Install Dependencies**

```sh
cd backend
go mod download
```

#### **Build the Backend Application**

```sh
go build -o bin/book-library cmd/api/main.go
```

#### **Run the Backend Server**

```sh
./bin/book-library
```

- The API will be available at: **`http://localhost:8080`**

---

### **3️⃣ Frontend Setup (React + Vite + TailwindCSS)**

#### **Install Dependencies**

Navigate to the frontend folder:

```sh
cd ../frontend
npm install
```

#### **Run the Frontend**

```sh
npm run dev
```

- The frontend will be available at: **`http://localhost:5173`**

---

## 🎯 Usage

1. **Visit the landing page** and navigate to the dashboard
2. **View books** from My Collection (default view)
3. **Search books**:
   - **From My Collection** (title, author, ISBN)
   - **From OpenLibrary** (ISBN only)
4. **Manage books**:
   - Add books manually
   - Add books from OpenLibrary
   - Edit details (title, author, isbn, rating, notes, status)
   - Delete books

---

## 📌 Roadmap / Future Enhancements

- ✅ Authentication (Sign in / Sign up)
- ✅ Advanced filtering & sorting integration
- ✅ Deployment setup

---

## 🤝 Contributing

Feel free to submit issues, feature requests, or pull requests.

---

## 📜 License

MIT License – Free to use and modify.
