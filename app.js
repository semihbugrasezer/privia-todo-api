document.addEventListener("DOMContentLoaded", function () {
  const loginForm = document.getElementById("login-form");
  const todoApp = document.getElementById("todo-app");
  const loginBtn = document.getElementById("login-btn");
  const logoutBtn = document.getElementById("logout-btn");
  const addTodoBtn = document.getElementById("add-todo-btn");
  const usernameInput = document.getElementById("username");
  const passwordInput = document.getElementById("password");
  const loginError = document.getElementById("login-error");
  const newTodoInput = document.getElementById("new-todo");
  const todoList = document.getElementById("todo-list");

  const apiUrl = "http://localhost:7878/api/v1"; // Backend API base URL
  let token = null;

  loginBtn.addEventListener("click", async function () {
    const username = usernameInput.value;
    const password = passwordInput.value;
    try {
      const response = await fetch(`${apiUrl}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });
      const data = await response.json();
      if (response.ok) {
        token = data.token; // Set the token here
        loginForm.classList.add("hidden");
        todoApp.classList.remove("hidden");
        loadTodos(); // Call loadTodos() after setting the token
      } else {
        loginError.textContent = data.message;
      }
    } catch (error) {
      loginError.textContent = "An error occurred. Please try again.";
    }
  });

  logoutBtn.addEventListener("click", function () {
    token = null;
    loginForm.classList.remove("hidden");
    todoApp.classList.add("hidden");
    usernameInput.value = "";
    passwordInput.value = "";
    todoList.innerHTML = "";
  });

  addTodoBtn.addEventListener("click", async function () {
    const newTodo = newTodoInput.value;

    try {
      const response = await fetch(`${apiUrl}/todos`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ name: newTodo }),
      });

      if (response.ok) {
        newTodoInput.value = "";
        loadTodos();
      } else {
        console.error("Failed to add todo");
      }
    } catch (error) {
      console.error("An error occurred. Please try again.");
    }
  });

  async function loadTodos() {
    try {
      const response = await fetch(`${apiUrl}/todos`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const todos = await response.json();
        todoList.innerHTML = "";
        todos.forEach((todo) => {
          const li = document.createElement("li");
          li.className =
            "list-group-item d-flex justify-content-between align-items-center";
          li.textContent = todo.name;

          const deleteBtn = document.createElement("button");
          deleteBtn.className = "btn btn-danger btn-sm";
          deleteBtn.textContent = "Delete";
          deleteBtn.addEventListener("click", function () {
            deleteTodo(todo.id);
          });

          li.appendChild(deleteBtn);
          todoList.appendChild(li);
        });
      } else {
        console.error("Failed to load todos");
      }
    } catch (error) {
      console.error("An error occurred. Please try again.");
    }
  }

  async function deleteTodo(todoId) {
    try {
      const response = await fetch(`${apiUrl}/todos/${todoId}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        loadTodos();
      } else {
        console.error("Failed to delete todo");
      }
    } catch (error) {
      console.error("An error occurred. Please try again.");
    }
  }
});
