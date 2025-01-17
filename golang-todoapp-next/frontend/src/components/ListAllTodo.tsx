"use client";
import ListCard from "@/components/ListCard";
import { useEffect, useState } from "react";

const getData = async () => {
  const res = await fetch("http://localhost:8000/todos", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    throw new Error("Failed to get all todo data");
  }
  return res.json();
};

const ListAllTodo = () => {
  const [todo, setTodo] = useState([]);

  useEffect(() => {
    const getTodo = async () => {
      try {
        const todos = await getData();
        setTodo(todos);
      } catch (error) {
        console.log(error);
      }
    };

    getTodo();
  }, []);

  return (
    <div className="flex flex-col justify-center items-center">
      {todo.map((todo: any) => (
        <div className="" key={todo.ID}>
          <ListCard todo={todo} />
        </div>
      ))}
    </div>
  );
};

export default ListAllTodo;
