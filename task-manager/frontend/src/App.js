import React, { useState, useEffect } from 'react';
import { Container, Row, Col, ListGroup, Button, Form, Badge } from 'react-bootstrap';
import axios from 'axios';
import './App.css';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';
import ToDoList from './To-Do-List';

function App() {
  const [tasks, setTasks] = useState([]);
  const [filteredTasks, setFilteredTasks] = useState([]);
  const [taskTitle, setTaskTitle] = useState("");
  const [filter, setFilter] = useState("all");

  // Get task from server
  useEffect(() => {
    axios.get('http://localhost:8080/tasks')
      .then(response => {
        setTasks(response.data);
        setFilteredTasks(response.data);
      })
      .catch(error => console.error('Error fetching tasks:', error));
  }, []);

  // Filter tasks
  useEffect(() => {
    if (filter === "completed") {
      setFilteredTasks(tasks.filter(task => task.completed));
    } else if (filter === "pending") {
      setFilteredTasks(tasks.filter(task => !task.completed));
    } else {
      setFilteredTasks(tasks);
    }
  }, [filter, tasks]);

  // Add new task
  const addTask = () => {
    const newTask = { title: taskTitle, completed: false };
    axios.post('http://localhost:8080/tasks', newTask)
      .then(response => {
        setTasks([...tasks, response.data]);
        setTaskTitle("");
      })
      .catch(error => console.error('Error adding task:', error));
  };

  // Delete task
  const deleteTask = (id) => {
    axios.delete(`http://localhost:8080/tasks/${id}`)
      .then(() => {
        setTasks(tasks.filter(task => task.id !== id));
      })
      .catch(error => console.error('Error deleting task:', error));
  };

  // Mark Done
  const toggleComplete = (id) => {
    const task = tasks.find(task => task.id === id);
    const updatedTask = { ...task, completed: !task.completed };
    axios.put(`http://localhost:8080/tasks/${id}`, updatedTask)
      .then(response => {
        setTasks(tasks.map(task => task.id === id ? response.data : task));
      })
      .catch(error => console.error('Error updating task:', error));
  };

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