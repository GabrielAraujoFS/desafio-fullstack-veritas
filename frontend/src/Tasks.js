import { useEffect, useState } from "react";
import api from "./services/api";
import "./Tasks.css";

function Tasks() {
  const [tasks, setTasks] = useState([]);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [draggedTask, setDraggedTask] = useState(null);

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    const response = await api.get("/tasks");
    setTasks(response.data);
  };

  const addTask = async () => {
    if (!title) return alert("Informe um título!");
    await api.post("/tasks", { title, description, status: "todo" });
    setTitle("");
    setDescription("");
    fetchTasks();
  };

  const deleteTask = async (id) => {
    await api.delete(`/tasks/${id}`);
    fetchTasks();
  };

  const moveTask = async (id, newStatus) => {
    await api.put(`/tasks/${id}`, { status: newStatus });
    fetchTasks();
  };

  const handleDragStart = (task) => {
    setDraggedTask(task);
  };

  const handleDrop = async (status) => {
    if (draggedTask && draggedTask.status !== status) {
      await moveTask(draggedTask.id, status);
      setDraggedTask(null);
    }
  };

  const columns = [
    { key: "todo", title: "A Fazer", color: "#ffb3b3" },
    { key: "inprogress", title: "Em Progresso", color: "#fff2a8" },
    { key: "done", title: "Concluídas", color: "#b3ffb3" },
  ];

  return (
    <div className="kanban-container">
      <h1>Lista de tarefa Kanban</h1>

      <div className="form">
        <input
          type="text"
          placeholder="Título"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <input
          type="text"
          placeholder="Descrição"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button onClick={addTask}>Adicionar</button>
      </div>

      <div className="kanban-board">
        {columns.map((col) => (
          <div
            key={col.key}
            className="kanban-column"
            style={{ backgroundColor: col.color }}
            onDragOver={(e) => e.preventDefault()}
            onDrop={() => handleDrop(col.key)}
          >
            <h2>{col.title}</h2>
            {tasks
              .filter((t) => t.status === col.key)
              .map((t) => (
                <div
                  key={t.id}
                  className="task-card"
                  draggable
                  onDragStart={() => handleDragStart(t)}
                >
                  <h3>{t.title}</h3>
                  <p>{t.description}</p>
                  <button
                    className="delete-btn"
                    onClick={() => deleteTask(t.id)}
                  >
                    🗑
                  </button>
                </div>
              ))}
          </div>
        ))}
      </div>
    </div>
  );
}

export default Tasks;
