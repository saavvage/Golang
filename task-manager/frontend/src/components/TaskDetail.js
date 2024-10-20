import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import axios from 'axios';

const TaskDetail = () => {
  const { id } = useParams(); 
  const [task, setTask] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    
    const fetchTask = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/tasks/${id}`);
        setTask(response.data); 
      } catch (error) {
        setError('Task not found');
      } finally {
        setLoading(false); 
      }
    };

    fetchTask();
  }, [id]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h2>Task Details</h2>
      {task ? (
        <div>
          <h3>{task.title}</h3>
          <p>{task.description}</p>
          <p>Status: {task.completed ? 'Completed' : 'Not Completed'}</p>
          <Link to={`/tasks/${id}/edit`} className="btn btn-primary">
            Edit Task
          </Link>
        </div>
      ) : (
        <p>Task not found</p>
      )}
      <Link to="/" className="btn btn-secondary">
        Back to Task List
      </Link>
    </div>
  );
};

export default TaskDetail;
