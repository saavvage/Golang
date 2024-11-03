// frontend/App.js
import React, { useState, useEffect } from 'react';
import { Container, Row, Col, ListGroup, Button, Form, Badge } from 'react-bootstrap';
import axios from 'axios';
import './App.css';
import Login from './Login';

function App() {
  const [tasks, setTasks] = useState([]);
  const [filteredTasks, setFilteredTasks] = useState([]);
  const [taskTitle, setTaskTitle] = useState("");
  const [filter, setFilter] = useState("all");
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  // authorize user
  const handleLoginSuccess = () => {
    setIsAuthenticated(true);
    fetchTasks();
  };

  // get tasks from server
  const fetchTasks = () => {
    axios.get('http://localhost:8080/tasks', { withCredentials: true })
      .then(response => {
        setTasks(response.data);
        setFilteredTasks(response.data);
      })
      .catch(error => console.error('Error fetching tasks:', error));
  };

  useEffect(() => {
    if (isAuthenticated) {
      fetchTasks();
    }
  }, [isAuthenticated]);

  // filtr tasks
  useEffect(() => {
    if (filter === "completed") {
      setFilteredTasks(tasks.filter(task => task.completed));
    } else if (filter === "pending") {
      setFilteredTasks(tasks.filter(task => !task.completed));
    } else {
      setFilteredTasks(tasks);
    }
  }, [filter, tasks]);

  // add task
  const addTask = () => {
    const newTask = { title: taskTitle, completed: false };
    axios.post('http://localhost:8080/tasks', newTask, { withCredentials: true })
      .then(response => {
        setTasks([...tasks, response.data]);
        setTaskTitle("");
      })
      .catch(error => console.error('Error adding task:', error));
  };

  // delete task
  const deleteTask = (id) => {
    axios.delete(`http://localhost:8080/tasks/${id}`, { withCredentials: true })
      .then(() => {
        setTasks(tasks.filter(task => task.id !== id));
      })
      .catch(error => console.error('Error deleting task:', error));
  };

  // change status of task
  const toggleComplete = (id) => {
    const task = tasks.find(task => task.id === id);
    const updatedTask = { ...task, completed: !task.completed };
    axios.put(`http://localhost:8080/tasks/${id}`, updatedTask, { withCredentials: true })
      .then(response => {
        setTasks(tasks.map(task => task.id === id ? response.data : task));
      })
      .catch(error => console.error('Error updating task:', error));
  };

  if (!isAuthenticated) {
    return <Login onLoginSuccess={handleLoginSuccess} />;
  }

  return (
    <Container fluid>
      <Row>
        <Col xs={2} className="sidebar bg-light">
          <h4 className="mt-4">Task Manager</h4>
          <ListGroup>
            <ListGroup.Item action onClick={() => setFilter('all')}>All Tasks</ListGroup.Item>
            <ListGroup.Item action onClick={() => setFilter('completed')}>Completed</ListGroup.Item>
            <ListGroup.Item action onClick={() => setFilter('pending')}>Pending</ListGroup.Item>
          </ListGroup>
          <Form.Group className="mt-4">
            <Form.Control 
              type="text" 
              placeholder="New Task" 
              value={taskTitle} 
              onChange={(e) => setTaskTitle(e.target.value)} 
            />
            <Button variant="primary" className="mt-2" onClick={addTask}>Add Task</Button>
          </Form.Group>
        </Col>

        <Col xs={10}>
          <h2 className="mt-4">Tasks</h2>
          <ListGroup>
            {filteredTasks.map(task => (
              <ListGroup.Item key={task.id} className="d-flex justify-content-between align-items-center">
                <div>
                  <Form.Check 
                    type="checkbox" 
                    checked={task.completed} 
                    onChange={() => toggleComplete(task.id)} 
                    label={task.title} 
                  />
                  <Badge bg={task.completed ? 'success' : 'warning'} className="ms-3">
                    {task.completed ? 'Completed' : 'Pending'}
                  </Badge>
                </div>
                <Button variant="danger" onClick={() => deleteTask(task.id)}>Delete</Button>
              </ListGroup.Item>
            ))}
          </ListGroup>
        </Col>
      </Row>
    </Container>
  );
}

export default App;
